# Go Map Implementation Benchmark

This repository contains benchmarks comparing the performance of Go's map implementations across different versions:
- Go 1.23: Bucket-based map implementation
- Go 1.24: Swiss map implementation

## Benchmark Setup

The benchmarks measure various map operations using text from "King Lear" as input data:
- Map insertion
- Map lookup
- Map update
- Map deletion

Each benchmark is run in separate directories with different Go versions specified in their respective `go.mod` files.

## CPU Usage Comparison Results

Based on CPU profiling, here's how the Swiss map implementation in Go 1.24 compares to the bucket map in Go 1.23:

### Map Insertion Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Total CPU sampling time | 2.37s | 2.27s | Swiss map used ~4.2% less CPU overall |
| Map assignment (`mapassign_faststr`) flat | 0.38s (16.03%) | 0.22s (9.69%) | Swiss map used ~42.1% less direct CPU time |
| Map assignment cumulative | 0.92s (38.82%) | 0.86s (37.89%) | Swiss map used ~6.5% less total CPU time |
| Hash function (`aeshashbody`) | 0.23s (9.70%) | 0.21s (9.25%) | Swiss map used ~8.7% less CPU for hashing |

### Map Update Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Total CPU sampling time | 2.07s | 2.12s | Swiss map used ~2.4% more CPU overall |
| Map assignment (`mapassign_faststr`) flat | 0.69s (33.33%) | 0.45s (21.23%) | Swiss map used ~34.8% less direct CPU time |
| Map assignment cumulative | 1.93s (93.24%) | 2.03s (95.75%) | Swiss map used ~5.2% more total CPU time |
| Hash function (`aeshashbody`) | 0.60s (28.99%) | 1.06s (50.00%) | Swiss map used ~76.7% more CPU for hashing |
| Memory comparison (`memequal`) | 0.03s (1.45%) | 0.38s (17.92%) | Swiss map used significantly more CPU for memory comparison |

### Map Lookup Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Total CPU sampling time | 1.04s | 1.09s | Swiss map used ~4.8% more CPU overall |
| Map access (`mapaccess1_faststr`) flat | 0.49s (47.12%) | 0.21s (19.27%) | Swiss map used ~57.1% less direct CPU time |
| Map access cumulative | 1.00s (96.15%) | 1.03s (94.50%) | Swiss map used ~3.0% more total CPU time |
| Hash function (`aeshashbody`) | 0.39s (37.50%) | 0.55s (50.46%) | Swiss map used ~41.0% more CPU for hashing |
| Memory comparison (`memequal`) | 0.05s (4.81%) | 0.24s (22.02%) | Swiss map used ~380% more CPU for memory comparison |

### Map Delete Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Total CPU sampling time | 1.52s | 1.46s | Swiss map used ~3.9% less CPU overall |
| Map delete (`mapdelete_faststr`) flat | 0.93s (61.18%) | 0.86s (58.90%) | Swiss map used ~7.5% less direct CPU time |
| Map delete cumulative | 0.93s (61.18%) | 0.97s (66.44%) | Swiss map used ~4.3% more total CPU time |
| Map usage tracking | 0.00s (0.00%) | 0.11s (7.53%) | Swiss map spent additional CPU time tracking map usage |

### Key Findings

1. For map insertion, the Swiss map implementation shows significant CPU usage improvements, particularly in direct map operations with a 42.1% reduction in direct CPU time.
2. For map update operations, the results are mixed:
   - The Swiss map uses 34.8% less direct CPU time for map assignments
   - However, it uses 76.7% more CPU for hash calculations during updates
   - It also spends significantly more time on memory comparisons
3. For map lookup operations:
   - The Swiss map uses 57.1% less direct CPU time for map access operations
   - However, it uses 41.0% more CPU for hash calculations during lookups
   - It spends substantially more time on memory comparisons (380% increase)
   - Overall, it uses 4.8% more total CPU for lookup operations
4. For map delete operations:
   - The Swiss map uses 7.5% less direct CPU time for map deletions
   - It spends additional CPU time on map usage tracking
   - Overall, it uses 3.9% less total CPU for delete operations
5. Overall, the Swiss map implementation is more CPU-efficient for insertion and deletion operations and for direct map access operations, but shows different performance characteristics for update and lookup operations, with increased CPU usage for hashing and memory comparisons.

## Memory Usage Comparison Results

Based on benchstat comparisons between Go 1.23's bucket map and Go 1.24's Swiss map implementations:

### Map Insertion Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Time per operation | 785.3µs | 654.4µs | Swiss map is ~16.7% faster |
| Memory allocated per op | 476.0 KiB | 426.2 KiB | Swiss map uses ~10.5% less memory |
| Heap allocations per op | 143 | 46 | Swiss map makes ~67.8% fewer allocations |

### Map Lookup Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Time per operation | 446.0µs | 411.5µs | Swiss map is ~7.7% faster |
| Memory allocated per op | 0 | 0 | No difference (no allocations) |
| Heap allocations per op | 0 | 0 | No difference (no allocations) |

### Map Update Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Time per operation | 504.0µs | 534.4µs | Swiss map is ~6.0% slower |
| Memory allocated per op | 0 | 0 | No difference (no allocations) |
| Heap allocations per op | 0 | 0 | No difference (no allocations) |

### Map Delete Operation

| Metric | Bucket Map (Go 1.23) | Swiss Map (Go 1.24) | Difference |
|--------|----------------------|---------------------|------------|
| Time per operation | 40.03µs | 34.66µs | Swiss map is ~13.4% faster |
| Memory allocated per op | 0 | 0 | No difference (no allocations) |
| Heap allocations per op | 0 | 0 | No difference (no allocations) |

### Memory Usage Key Findings

1. For map insertion, the Swiss map implementation shows dramatic improvements:
   - 16.7% faster execution time
   - 10.5% less memory allocated
   - 67.8% fewer heap allocations

2. For lookup and delete operations, the Swiss map is faster (7.7% and 13.4% respectively) with no difference in memory usage.

3. For update operations, the Swiss map is slightly slower (6.0%) with no difference in memory usage.

4. The most significant memory improvement is in the number of heap allocations for insertions, which could lead to reduced garbage collection pressure.

## How to Run the Benchmarks

To run the benchmarks for a specific operation and see memory usage:

```bash
# For Go 1.23 (bucket map)
cd bucketsmap
go test -bench=BenchmarkMap<Operation> -benchmem > results/<operation>/mem.txt

# For Go 1.24 (Swiss map)
cd swissmap
go test -bench=BenchmarkMap<Operation> -benchmem > results/<operation>/mem.txt
```

Replace `<Operation>` with one of: `Insert`, `Lookup`, `Update`, or `Delete`.

To compare results using benchstat:

```bash
# Install benchstat if you don't have it
go install golang.org/x/perf/cmd/benchstat@latest

# Compare results
benchstat bucketsmap/results/<operation>/mem.txt swissmap/results/<operation>/mem.txt
```

For more statistically significant results, use the `-count` flag:

```bash
go test -bench=BenchmarkMap<Operation> -benchmem -count=6 > results/<operation>/mem_stat.txt
```

To generate CPU profiles:

```bash
# For Go 1.23 (bucket map)
cd bucketsmap
go test -bench=BenchmarkMap<Operation> -cpuprofile=results/<operation>/cpu.prof

# For Go 1.24 (Swiss map)
cd swissmap
go test -bench=BenchmarkMap<Operation> -cpuprofile=results/<operation>/cpu.prof
```

To analyze CPU profiles:

```bash
go tool pprof -top ./bucketsmap/results/<operation>/cpu.prof > ./bucketsmap/results/<operation>/cpu.txt
go tool pprof -top ./swissmap/results/<operation>/cpu.prof > ./swissmap/results/<operation>/cpu.txt
```

For visual analysis (requires Graphviz):

```bash
go tool pprof -http=:8080 ./bucketsmap/results/<operation>/cpu.prof
go tool pprof -http=:8080 ./swissmap/results/<operation>/cpu.prof
```
