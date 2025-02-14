# Exercise: Handlers

In this exercise, you'll use the `net/http` package to set up a simple web server with two routes: one that returns a greeting message and another that provides mock user data in JSON format.

By the end, your server will handle both plain text and JSON responses, giving you practical experience with HTTP handlers and JSON encoding in Go.

First, navigate to **/L1. API Design Principles/2. Handlers/starter/main.go**.

## Instructions:

1. Uncomment the commented code.
2. Create two handler functions:

- `greet()`, which responds with a simple greeting message of your choice.
- `userInfo()`, which responds with a JSON object representing a user with fields for name, age, and email.

3. Make sure to register your handlers using the appropriate `net/http` functions.
4. Serve your application on port `8080`.
5. Test your endpoints using a browser, cURL, or Postman to ensure the proper response formats (plain text for the greeting, and JSON for the user information).
