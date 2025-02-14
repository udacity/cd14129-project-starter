# Exercise: Connecting to the Database with GORM and Gin

In this exercise, you’ll review how to connect a PostgreSQL database to a Go server using the Gin web framework alongside GORM. You’ll set up a simple API that connects to PostgreSQL, auto-migrates a model to the database, and serves a basic route. This exercise will help you practice connecting Gin and GORM while exploring some of the differences compared to using `net/http` as you had in Lesson 1.

First, navigate to **/L2. Building High-Performance APIs with Go/3. Database Setup with Gin/starter/main.go**.

## Instructions

1. Set up a PostgreSQL database ensure you have it running. You can use a database previously created in the course or set up a new one.
2. Implement the necessary code to configure a connection to your PostgreSQL database using GORM. Don’t forget to handle potential errors (see the `TODO` in the starter code).
3. The `Vehicle` model is already provided in the starter code. Complete the `TODO` to auto-migrate it to your PostgreSQL database.
4. Create a `GET / route` that returns the text "Database connected!" to verify that your server and database are working together.
5. Run the server and test the `/` endpoint to ensure the connection to PostgreSQL works.
