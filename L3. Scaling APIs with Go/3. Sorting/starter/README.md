# Exercise: Sorting

In this exercise, you will implement sorting functionality in your API. The server will respond to a `GET /movies` route, returning movies sorted by a specified field. Sorting will be done based on optional query parameters.

First, navigate to **/L3. Scaling APIs with Go/3. Sorting/starter/main.go**.

## Instructions

1. Set up the project and model. Start by using the provided starter code, which defines a simple Gin server and a model (`Movie`). The server will respond to the `GET /movies` route and allow sorting by `title`, `director`, `releaseYear`, and `rating`.
2. Implement sorting. Modify the handler function to accept a query parameter (`sortBy`) for specifying the field to sort by. If a valid `sortBy` query parameter is provided (e.g., `title`, `director`, `releaseYear`, or `rating`), return the movies sorted by that field. Default to sorting by ID if no parameter is provided.
3. Test the sorting. Use query parameters like `?sortBy=title`, `?sortBy=director`, or `?sortBy=rating` to test sorting.

### Testing the API

- Sort by title: `GET /movies?sortBy=title`. This should returns movies sorted alphabetically by title.
- Sort by director: `GET /movies?sortBy=director`. This should return movies sorted alphabetically by director name.
- Sort by release year: `GET /movies?sortBy=releaseYear`. Returns movies sorted by release year (oldest to newest).
- Sort by rating: `GET /movies?sortBy=rating`. This should return movies sorted by rating (highest rating first).
- Default (sort by ID): `GET /movies`. This should return movies sorted by their ID in ascending order.
