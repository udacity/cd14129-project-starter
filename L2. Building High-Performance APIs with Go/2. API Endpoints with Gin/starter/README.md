# Exercise: API Endpoints with Gin

In this exercise, you'll practice defining and implementing basic API endpoints with Gin. You’ll create two simple API endpoints: one for retrieving a list of users (GET) and another for adding a new user (POST). This exercise will introduce you to defining routes, handling JSON requests and responses, and using Gin’s context for request and response management. We'll focus on a subset of CRUD operations to keep it simple.

First, navigate to **/L2. Building High-Performance APIs with Go/2. API Endpoints with Gin/starter/main.go**.

## Instructions

1. Use Gin to define two routes: `GET /users` to retrieve all users and `POST /users` to create a new user.
2. Implement the `GET /users` handler to return the list of users (provided in the starter code) in JSON format.
3. Use Gin’s context to send the response with the appropriate status code.
4. Implement the `POST /users` handler to accept a JSON payload for a new user.
5. Use Gin’s context to parse the request body into the `User` struct and return the created user in the response with status `201 Created`.
6. Use Postman or cURL to test both endpoints. Ensure the `GET /users` returns the list of users and `POST /users` successfully adds a new user.
