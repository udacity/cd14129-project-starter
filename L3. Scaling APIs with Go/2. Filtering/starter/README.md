# Exercise: Filtering

In this exercise, you'll implement a filtering mechanism in your API that allows clients to filter items by specific fields. The server will return a filtered list of items based on the query parameters provided in the request.

First, navigate to **/L3. Scaling APIs with Go/2. Filtering/starter/main.go**.

## Instructions

1. Review the provided starter code, which defines a simple Gin server and a model (`Task`).
2. Run the server. The server will respond to a `GET /tasks` route, returning tasks filtered by optional query parameters.
3. Modify the handler function to accept query parameters (`status`, `priority`, and `assignee`) for filtering tasks. If a query parameter is provided, filter the tasks based on that field. If no filters are provided, return all tasks by default.
4. Test the filtering. Use query parameters like `?status=completed`, `?priority=high`, or `?assignee=John` to test your filtering logic. You server should be able to combine multiple filters as well: `?status=completed&priority=high&assignee=John`.

### Testing the API

- Filter by status: `GET /tasks?status=completed`. This should return all tasks with the status "completed".
- Filter by priority: `GET /tasks?priority=high`. This should return all tasks with the priority "high".
- Filter by assignee: `GET /tasks?assignee=John`. This should return all tasks assigned to "John".
- Combine filters: `GET /tasks?status=pending&priority=high&assignee=John`. This should return all tasks that match all of the provided filters (status "pending", priority "high", and assignee "John").
