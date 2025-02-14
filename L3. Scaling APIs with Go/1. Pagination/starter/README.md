# Exercise: Implementing Pagination in Go

In this exercise, you'll implement offset-based pagination for a list of items in a simple HTTP server using the Gin framework. Instead of querying a database, you'll use a static slice of items to demonstrate pagination.

First, navigate to **/L3. Scaling APIs with Go/1. Pagination/starter/main.go**.

## Instructions

1. Define a `GET /items` route.

2. Implement the handler `itemsHandler` for the endpoint `GET /items`. This handler accepts two query parameters: `limit` and `offset`. The handler should return a subset of items based on the provided `limit` and `offset` values.

3. Test the endpoint. Use a tool like Postman or cURL to test your pagination implementation.
