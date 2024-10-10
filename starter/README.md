# Inventory Management (starter project)

This project is the starting point for building your Inventory Management API using [Go](https://go.dev/), the [Gin](https://github.com/gin-gonic/gin) framework, and [GORM](https://gorm.io/index.html) for database operations. You can use this starter code to help seed the database, but _its usage is completely optional_ for this project (you may wish to seed the database using your own code).

## Instructions

1. Clone this project to your local machine, then navigate to the **/starter** directory.

2. Set up a PostgreSQL database on your local machine.

3. In **main.go**, create a `dsn` variable and add your PostgreSQL connection string.

```
// TODO: Add your PostgreSQL connection string here (dsn) to connect to your database
```

4. Define the `Item` struct in **main.go**. This struct represents the inventory items stored in the database.

```
// TODO: Create a struct for Item here
```

5. Before running the server, ensure all dependencies are properly managed.

```
go mod tidy
```

6. Run the server.

```
go run main.go
```

The server will connect to your PostgreSQL database and seed it with 20 sample items. Be sure to check the database to confirm that the items have been seeded successfully.
