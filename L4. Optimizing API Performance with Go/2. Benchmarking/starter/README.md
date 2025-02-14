# Exercise: Benchmarking

In this exercise, you'll practice writing and running benchmarks to measure the performance of an API built with Gin. Benchmarking will help you understand how to evaluate the performance of your API under different loads and identify areas for improvement.

You will:

- Write a benchmark test for an API endpoint.
- Run the benchmark test using Go’s testing package.
- Interpret the benchmark results.

First, navigate to **/L4. Optimizing API Performance with Go/2. Benchmarking/starter/main.go**.

## Instructions

1. In **main_test.go**, implement the `BenchmarkAPIData` function. This function should call an API endpoint in a loop to simulate load. You will use Go’s testing package for this.

2. Use the following command to run your benchmark test.

```
go test -bench=. -benchmem
```

This will execute the `BenchmarkAPIData` function and measure the performance of the `/api/data` endpoint.

3. Analyze the benchmark output to see how the API performs under load. Look at metrics such as the time per operation (ns/op) and memory allocations (B/op). The output should look something like the following:

```
BenchmarkAPIData-8          10000        120000 ns/op       1024 B/op       10 allocs/op
PASS
ok      _/path/to/your/app     1.234s
```

In your analysis, recall that:

- Time per operation (ns/op) -- higher values indicate that the endpoint takes longer to respond
- Memory allocations (B/op and allocs/op) -- excessive memory usage or allocations might indicate inefficiencies in the code

4. Bonus (_optional_): Modify the API code to optimize performance and rerun the benchmark to observe improvements. Can you reduce the time taken by the API to process requests? Try modifying the `time.Sleep` duration in **main.go** or optimizing the handler function logic, then rerun the benchmark to measure improvements.
