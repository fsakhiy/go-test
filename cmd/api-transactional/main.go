// cmd/api-transactional/main.go
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gin-test/internal/auth"
	"gin-test/internal/shared/middleware"
	"gin-test/internal/shared/response"
	"gin-test/internal/tickets"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// 1. Global Error Handling Middleware for Gin
func globalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Let the request process first

		// If a handler attached an error using c.Error(err)
		if len(c.Errors) > 0 {
			// Grab the last error thrown
			err := c.Errors.Last().Err

			var validationErrs validator.ValidationErrors

			if errors.As(err, &validationErrs) {
				// Group the errors by rule, exactly as you did before
				groupedErrors := make(map[string][]string)

				for _, fieldError := range validationErrs {
					rule := fieldError.Tag()
					field := fieldError.Field()
					groupedErrors[rule] = append(groupedErrors[rule], field)
				}

				// Return using your shared response package
				response.Err(c, 400, "validation error", groupedErrors)
				return
			}

			// Catch-all for 500 internal server errors
			response.Err(c, 500, "internal server error", err.Error())
		}
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found or could not be loaded: %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "6700"
	}

	// 2. Configure Gin's built-in validator
	// We safely extract the validator engine and register your exact JSON tag function
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// 3. Initialize Gin and attach the global error handler
	app := gin.Default()
	app.Use(globalErrorHandler())

	// 4. Database Setup (Untouched)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}

	// 5. Setup Middleware
	jwtSecret := os.Getenv("JWT_SECRET")
	authMiddleware := middleware.ValidateAuth(jwtSecret)

	// 6. Dependency Injection
	ticketRepo := tickets.NewRepository(db)
	ticketSvc := tickets.NewService(ticketRepo)
	ticketHandler := tickets.NewHandler(ticketSvc)

	// auth injection
	authRepo := auth.NewRepository(db)
	authSvc := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authSvc)

	v1 := app.Group("/api/v1")

	// register routes
	tickets.RegisterRoutes(v1, ticketHandler, authMiddleware)
	auth.RegisterRoutes(v1, authHandler)

	// start app
	if err := app.Run(":" + port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
