# Exercise: API Endpoints I

In this exercise, you'll practice using GORM and the `net/http` package to build a simple server that interacts with a PostgreSQL database for basic CRUD operations. You will set up two endpoints: one to create a new post and one to retrieve a post by its ID.

First, navigate to **/L1. API Design Principles/5. API Endpoints I/starter/main.go**.

## Instructions

1. Create a `POST` route for retrieving all posts (see the `TODO` in the starter code).

2. Write an accompanying handler function to support the route you created above.

- Bonus - Include error handling:
  - Return a "Method Not Allowed" error if the request method is not `POST`.
  - Return an "Invalid Input" error if the request body does not contain valid data for creating a post.
  - Return a "Could Not Create Post" error if there is an issue with inserting the post into the database.

3. Create a `GET` route for retrieving a single post by its ID (see the `TODO` in the starter code).

4. Write a accompanying handler function to support the create you created above.

- Bonus - Include error handling:
  - Return a "Method Not Allowed" error if the request method is not `GET`.
  - Return an "Invalid post ID" error if the ID provided in the URL is not a valid format (e.g., if it's not a number).
  - Return an error if there is an issue retrieving the post from the database (e.g., if the post does not exist).

5. Test your endpoints. Use cURL or Postman to test the creation and retrieval of posts.
