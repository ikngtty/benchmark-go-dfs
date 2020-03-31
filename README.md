# Benchmark Go DFS

## How to run the benchmark

### Local

```shell
go test -bench=. -benchmem ./traverse
```

### Docker

```shell
docker run -i --rm -v $(pwd):/go/src/github.com/ikngtty/benchmark-go-dfs golang go test -bench=. -benchmem github.com/ikngtty/benchmark-go-dfs/traverse/
```

## Result

The go version is 1.14.

commit: 3320987

```shell
BenchmarkFindPathRec-4    	     132	   8807604 ns/op	     664 B/op	       7 allocs/op
BenchmarkFindPathLoop-4   	      12	  88368313 ns/op	33555069 B/op	 1048582 allocs/op
```

commit: d9fb48e

```shell
BenchmarkFindPathRec-4    	     140	   8297480 ns/op	     664 B/op	       7 allocs/op
BenchmarkFindPathLoop-4   	     124	   9179461 ns/op	 8389896 B/op	      47 allocs/op
```

commit: d5d45d7

```shell
BenchmarkFindPathRec-4    	     150	   7938290 ns/op	     664 B/op	       7 allocs/op
BenchmarkFindPathLoop-4   	      10	 102009637 ns/op	25166953 B/op	 1572874 allocs/op
```
