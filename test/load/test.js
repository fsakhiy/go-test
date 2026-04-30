import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    // Key configurations for load testing
    stages: [
        { duration: '10s', target: 50 },  // Ramp-up to 50 users over 10 seconds
        { duration: '30s', target: 50 },  // Stay at 50 users for 30 seconds
        { duration: '10s', target: 0 },   // Ramp-down to 0 users over 10 seconds
    ],
    // For maximum flooding, you can comment out the stages above and use:
    // vus: 100, 
    // duration: '30s',
};

export default function () {
    // Target the endpoint based on your server logs
    const url = 'http://127.0.0.1:8899/api/v1/ticket/';

    const params = {
        headers: {
            'Authorization': 'Bearer fuckugo',
            'Content-Type': 'application/json',
        },
    };

    // Perform the GET request
    const res = http.get(url, params);

    // Assert that the response is 200 OK
    check(res, {
        'is status 200': (r) => r.status === 200,
        // Optional: you can add more checks here (e.g. response time)
        // 'transaction time OK': (r) => r.timings.duration < 200,
    });

    // Short sleep between iterations to not completely overwhelm your machine
    // Remove or lower this for more aggressive flooding
    sleep(0.1);
}
