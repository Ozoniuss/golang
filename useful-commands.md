**Run a single benchmark file, including memory allocations (without running tests)**

```
go test -run=xxx -bench=. -benchmem main_test.go
```

**Run a single benchmark function from a file, for an increased duration if the function takes a while to execute, to improve reporting accuracy (without running tests)**

```
go test -run=xxx -bench=Fib40 -benchtime=20s main_test.go
```

## Benchmark flags

```
go test -run=xxx -bench=Fib40 -benchtime=20s -benchmem -count=10 -cpu=1,2,4,16 main_test.go
```

```
bench - regex for running benchmarks

benchmem - memory allocations, works with b.ReportAllocs as well

benchtime - increase benchmark time for long running benchmarks. Note that if
the benchmark timer is stop, the execution outside the timer is not counted
towards the total benchtime. Reduce benchtime for fast executions but long
setups

Note that benchtime=20x tells to run the benchmark exactly 20 times.

c - save the benchmark to a binary (didn't use yet)

count - run the benchmark that many times (useful for computing statistics
like variance)

cpu - run the benchmark using the number of cpus provided

gcflags=S - show the assembly of the benchmark

run - regex for running tests (which execute during benchmarks)
```

```
benchstat - statistics on benchmarks
benchstat bench.txt - compute statistics on bench results
benchstat old.txt new.txt - compare results between different benchmarks
```

## Benchmark type functions

```
b.ResetTimer() - reset benchmark timer. Be careful to not do this inside a b.N!
b.StopTimer() - stop benchmark timer. Execution doesn't count towards bench time
b.StartTimer() - restart benchmark timer. Execution time is counted again
b.ReportAllocs() - report memory allocations for that benchmark
```

### Profiling

For now use the following package for profiling:

```
go get github.com/pkg/profile
```

The standard profiling package is `runtime/pprof`.

Start the profile with an HTTP server. Graph visualizations are typically more useful than say text visualisation of the top 10 functions by execution time, because that visualization also shows the call stack, and the expensive functions we see in the code may call underlying lower-level functions that actually take time to execute.

```
go tool pprof -http=:8080 cpu.pprof
```

## Inlining

A leaf function is a function that does not call any other function, making it an ideal canditate for inlining. An inline function can be further optimised, for example, the compiler may see it does not affect any global state. In that case, the compiler may remove the instructions completely, nullifying the benchmark results.

To avoid inlining, the recommended strategy is to:

-   define a global variable;
-   define a local variable at the start of the benchmark;
-   store the result of the benchmark inside the local variable;
-   at the end of the benchmark, store the result in the local variable.

Why not store the result directly in a global variable? That is because changing a global variable (which lives on the heap) is more expensive than changing a local varaible living on the stack.

Why not store the result to \_?
