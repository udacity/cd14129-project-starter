# Exercise: Enabling pprof and Profiling a Gin Server

In this exercise, you'll focus on integrating the `pprof` profiling tool into an existing Gin server and generating profiling data. You'll explore how to enable pprof to capture profiling data like CPU, memory, and goroutines.

This exercise will guide you through enabling pprof in the Gin server, simulating some workload, and running a series of tests to generate profiling data. By the end of this exercise, you'll have hands-on experience enabling pprof in a Gin server, generating profiling data, and analyzing it to identify potential performance issues.

First, navigate to **/L4. Optimizing API Performance with Go/1. Profiling/starter/main.go**.

## Instructions

1. Enable pprof in the provided Gin server. The starter code includes a basic Gin server. Your task is to enable `pprof` profiling by importing the necessary `net/http/pprof` package and setting up the corresponding routes.

2. Simulate some work. Modify the `/api/data` endpoint to perform some CPU and memory-intensive operations that will be noticeable in the profiling data.

3. Generate profiling data. After enabling `pprof` and simulating work, run the server and use profiling tools like go tool `pprof` to collect data.

4. Analyze profiling data. After collecting profiling data, you will interpret the results and identify potential performance bottlenecks in the server.

### Testing and Profiling

1. Make sure you have already started the Gin server if you had not done so already:

```
go run main.go
```

2. Access `pprof` routes. With `pprof` enabled, you can now access the profiling endpoints (running on port `6060`):

- CPU profile: `http://localhost:6060/debug/pprof/profile`
- Heap profile: `http://localhost:6060/debug/pprof/heap`
- Goroutine profile: `http://localhost:6060/debug/pprof/goroutine`

3. Generate CPU profile by running the following command to generate a CPU profile, saving it to a file:

```
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

4. Analyze CPU profile. Open the profile using the following command:

```
go tool pprof <file-name>
```

Then, within the interactive mode, use commands like `top` or `web` to view the most CPU-intensive functions or generate flame graphs.

5. To analyze memory usage or goroutines, you can also run:

```
go tool pprof http://localhost:6060/debug/pprof/heap go tool pprof http://localhost:6060/debug/pprof/goroutine
```

6. Interpret results. In general:

- Use `top` to identify the functions consuming the most resources
- Generate flame graphs to visually inspect where time/memory is being spent
- Look for inefficient code patterns or functions that take too long
