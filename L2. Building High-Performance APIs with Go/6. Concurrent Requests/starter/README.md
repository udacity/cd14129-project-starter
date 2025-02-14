# Exercise: Goroutines and Channels

In this exercise, you'll practice using goroutines and channels to handle multiple API requests concurrently in Go with the Gin framework. The goal is to simulate fetching data from multiple sources (e.g., external APIs) and combine the results into a single response.

First, navigate to **/L2. Building High-Performance APIs with Go/6. Concurrent Requests/starter/main.go**.

## Instructions

1. Implement the `handleConcurrentRequests` function to simulate concurrent API requests.

2. Inside the `handleConcurrentRequests` function, simulate two API requests by creating goroutines that wait for a few seconds (using `time.Sleep`) and then send a result back via channels.

3. Create two channels Use two channels to collect the results from your two simulated "API calls."

4. Use the channels to wait for the results from the two goroutines and combine them into a single response.

5. Once both results are received, return them as a JSON response to the client.

6. Test the endpoint. Use cURL or Postman to make a request to the endpoint and observe how the data is processed concurrently.
