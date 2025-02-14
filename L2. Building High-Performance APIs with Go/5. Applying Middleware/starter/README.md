# Exercise: Applying Middleware

In this exercise, you will create a custom middleware in Gin that adds a custom HTTP header (`X-Custom-Header`) to all responses. You will then apply this middleware globally to all routes using the r.Use() method.

First, navigate to **/L2. Building High-Performance APIs with Go/5. Applying Middleware/starter/main.go**.

## Instructions

1. Write a middleware function that adds the `X-Custom-Header` with a static value (e.g., "Middleware-Active") to all responses.

2. Use the `r.Use()` method to apply this middleware to all routes in the server.

3. Test the middleware. Start the server and make requests to different routes. Check the responses to confirm that the custom header has been added.

4. Reflect on middleware scope. Think about how global middleware differs from route-specific middleware and when it makes sense to apply middleware globally.
