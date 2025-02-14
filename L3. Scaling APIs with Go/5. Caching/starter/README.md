# Exercise: In-Memory Caching in Go with Gin

In this exercise, you'll implement a simple in-memory caching mechanism with concurrency control using `sync.RWMutex` in Go, and you'll work with the Gin framework to create an API that returns a list of todos and retrieves individual todos with caching.

First, navigate to **/L3. Scaling APIs with Go/5. Caching/starter/main.go**.

## Instructions

1. Review the starter code. This server already includes:

- A `Todo` struct that includes an `ID` and a `Title`.
- A list of `todos` that represents a simple list of tasks.
- A Go `map` that implements an in-memory cache where the key is the todo's `ID`, and the value is the `Todo`.
- A `sync.RWMutex` for thread safety, ensuring that read and write operations on the cache are safely handled in a concurrent environment.

2. Implement handlers:

- `GET /todos`: This endpoint should simply return the full list of todos.
- `GET /todos/:id`: This endpoint should retrieve a single todo by its `ID`.
  - First, check if the todo is present in the cache. If found, return the cached todo.
  - If not found in the cache, search for the todo in the list, add it to the cache, and then return it.
