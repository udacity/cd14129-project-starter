# Exercise: Implementing Rate Limiting Using the Token Bucket Algorithm

In this exercise, you’ll implement a rate limiter in a Gin-based API using the Token Bucket algorithm. The rate limiter will allow a burst of 5 requests and then limit subsequent requests to 1 per second.

First, navigate to **/L3. Scaling APIs with Go/4. Rate Limiting/starter/main.go**.

## Instructions

1. Create an instance of the Token Bucket to handle rate limiting. Allow 5 requests per burst and limit further requests to 1 per second.
2. Implement the Token Bucket algorithm. In the `Allow` method, calculate the number of tokens that should be added based on the elapsed time, refill the bucket, and allow or deny requests based on the number of available tokens.
3. Add rate limiting to the API:

- Use the `GET /data` route in the Gin server. Apply the rate limiter to this route.
- If the rate limiter allows the request, respond with a success message.
- If the rate limit is exceeded, return a `429 Too Many Requests` status code.

4. Test the `/data` endpoint by sending multiple requests quickly to verify if the rate limiter works correctly.
