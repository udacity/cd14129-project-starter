# Exercise: Documenting a Gin-based API with Swagger

In this exercise, you'll document a simple Gin-based API using Swagger. This will help you understand the importance of API documentation and how to generate interactive documentation using Swagger UI for your Go APIs.

First, navigate to **/L4. Optimizing API Performance with Go/4. Documentation/starter/main.go**.

## Instructions

1. Set up Swagger for your project. You’ll install the necessary tools and libraries to enable Swagger in your Gin application.

If you haven't done so already, run the following command to tidy up your dependencies and install any missing ones:

```
go mod tidy
```

2. Review the starter code, which is basic Gin server with CRUD endpoints. Your task is to add Swagger annotations to the code. That is, add Swagger annotations to document your API’s behavior and endpoints.

You’ll do this for both `getAllPhotos` and `createPhoto`.

3. Use Swagger tools to generate the API documentation based on the annotations:

```
swag init
```

4. Start your server and navigate to the Swagger UI page to view your documented API. Once the documentation is generated, start the server and visit `http://localhost:8080/swagger/index.html` to view your Swagger UI.
