## Exercise: Database CRUD with Gin

In this exercise, you will practice creating basic CRUD (Create, Read, Update, Delete) operations using Gin and GORM. You'll implement API routes to perform each of the CRUD operations on a Task resource, query the PostgreSQL database using GORM, and return JSON responses. The exercise will also guide you through basic error handling during each operation.

First, navigate to **/L2. Building High-Performance APIs with Go/4. Database CRUD with Gin/starter/main.go**.

## Instructions

1. Implement the necessary handler functions to:

   - Create a new task in the database.
   - Retrieve a task by its ID.
   - Update a task’s fields (e.g., title, description, is_completed).
   - Delete a task by its ID.

2. For each handler, you must:

- Use GORM's methods to create, retrieve, update, and delete tasks from the PostgreSQL database.
- Ensure your handlers return appropriate JSON responses for each operation (success messages or the created/updated/deleted task object).
- Include error handling for scenarios such as:
  - Invalid input data during creation or update.
  - Task not found when trying to retrieve, update, or delete a task.
  - Any issues interacting with the database (e.g., connection problems).

3. Test the API. Use Postman, cURL, or another tool to test each endpoint and ensure that the CRUD operations are functioning properly.
