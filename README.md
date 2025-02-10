# Go Race Condition in Counter Increment

This repository demonstrates a common race condition in Go programs that use goroutines to increment a shared counter. The program creates 1000 goroutines, each incrementing the counter. However, without proper synchronization, the final counter value might be less than 1000 due to race conditions.

## Bug Description
The `counter++` operation is not atomic. Multiple goroutines can try to increment the counter simultaneously, leading to incorrect results.  The `WaitGroup` ensures all goroutines complete but doesn't solve the race condition.

## Solution
The solution utilizes a mutex to protect the counter from concurrent access.  This ensures that only one goroutine can increment the counter at a time, preventing race conditions and guaranteeing the correct final count.