# Exercise: API Endpoints II

In this exercise, you'll extend the solution from the previous exercise (i.e., API Endpoints I) to add the remaining CRUD operations for posts: retrieving all posts, updating a specific post's details, and deleting a post. This will help you deepen your understanding of how to work with GORM and the `net/http` package for handling more advanced CRUD functionality.

You can extend your solution from the previous exercise, or you can use the provided solution code as the starter. Either way, you will be adding the following functionality:

- `GET /posts` to retrieve all posts.
- `PATCH /posts/{id}` to update an existing post's details.
- `DELETE /posts/{id}` to delete a post by its ID.

First, navigate to **/L1. API Design Principles/6. API Endpoints II/starter/main.go**.

## Instructions

1.  Define a route for `GET /posts` and create a corresponding handler that retrieves and returns a list of all posts in the database as JSON.

- Bonus - Include error handling:
  - Return a "Method Not Allowed" error if the request method is not `GET`.
  - Return a "Could Not Retrieve Posts" error if there is an issue retrieving the list of posts from the database (e.g., if the database connection fails).

2.  Define a route for `PATCH /posts/{id}` and create a handler that allows updating certain fields of a specific post (e.g., `title`, `body`, or `user_id`). Ensure that only the fields included in the request body are updated.

- Bonus - Include error handling:

  - Return a "Method Not Allowed" error if the request method is not `PATCH`.
  - Return an "Invalid Post ID" error if the provided ID is not valid (e.g., if it cannot be converted to an integer).
  - Return an "Invalid Input" error if the request body does not contain valid data.
  - Return a "Could Not Update Post" error if there is an issue updating the post in the database.

3.  Define a route for `DELETE /posts/{id}` and create a handler that removes the specified post from the database.

- Bonus - Include error handling:

  - Return a "Method Not Allowed" error if the request method is not `DELETE`.
  - Return an "Invalid Post ID" error if the provided ID is not valid (e.g., if it cannot be converted to an integer).
  - Return a "Could Not Delete Post" error if there is an issue deleting the post from the database.

4.  Test your new endpoints. Use cURL or Postman to test retrieving all posts, updating an existing post, and deleting a post by its ID.

Again, feel free to start by extending your previous solution or using the provided solution code from the previous exercise.
