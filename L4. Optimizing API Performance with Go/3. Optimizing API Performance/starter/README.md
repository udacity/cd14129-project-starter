# Exercise: Optimizing API Performance

In this exercise, you will identify and address common performance bottlenecks in a simple API built with Gin. Specifically, you'll improve the API by optimizing memory usage and reducing unnecessary computations.

First, navigate to **/L4. Optimizing API Performance with Go/3. Optimizing API Performance/starter/main.go**.

## Instructions

1. Review the starter code. The provided Gin server simulates an inefficient API that unnecessarily recomputes the same result multiple times and performs extra memory allocations. Since this API endpoint performs an inefficient task, this will simulate a common performance bottleneck.

2. Write a benchmark to measure the performance of this inefficient endpoint. This will help you identify where optimizations are needed.

3. Implement optimizations to reduce memory usage and improve efficiency.

4. Rerun your benchmark to measure the performance improvements after your optimizations.

### Running the Benchmark

In this exercise, you are running the benchmark test for both the inefficient and optimized versions of the API to measure the difference in performance. Here's a brief review of running benchmarks:

1. Run the benchmark:

```
go test -bench=. -benchmem
```

2. Output example (_before_ optimization):

```
BenchmarkInefficientCompute-8         1        1200000000 ns/op       1600000 B/op     20000 allocs/op
PASS
ok      _/path/to/your/app     1.234s
```

3. Output example (_after_ optimization):

```
BenchmarkOptimizedCompute-8         1        600000000 ns/op         800000 B/op      10000 allocs/op
PASS
ok      _/path/to/your/app     0.678s
```
