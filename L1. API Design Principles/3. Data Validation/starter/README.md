# Exercise: Data Validation

In this exercise, you'll practice implementing simple data validation on incoming requests in a Go web server. You will set up a basic API that accepts user registration details and responds with a success or error message depending on the validity of the provided data.

First, navigate to **/L1. API Design Principles/3. Data Validation/starter/main.go**.

# Instructions:

1. Create a new register handler function that accepts a POST request with a JSON payload containing `name`, `email`, and `age`.
2. In your handler, implement the following validation rules:

- The `name` field cannot be empty.
- The `email` field must contain an `@` symbol.
- The `age` field must be a number greater than `21`.

3. If any of the data is invalid, return a `400 Bad Request` response with an error message.
4. If all data is valid, return a `200 OK` response with a success message.
5. Set up the server to run on port `8080`, and test your validation using tools like cURL or Postman.
