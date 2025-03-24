new things learned from each project:
    * initialize
        * Use `testing.B` for benchmark tests
        * Use `b.N` pattern for running operations multiple times
        * Global variables to prevent compiler optimizations in benchmarks
        * run many benchmark `go test` with `-bench=.` flag
    * allocate
        * Use `-gcflags '-m -m'` to see compiler optimizations and escape analysis
        * When values escape to heap:
          - Values passed to fmt.Println escape to heap (e.g., `t escapes to heap`)
          - Pointer-created structs escape to heap (e.g., `&Item{} escapes to heap`)
        * `does not escape` => put into stack
        * Stack vs Heap allocation:
          - Stack: faster, automatically managed memory (local variables)
          - Heap: slower, requires garbage collection, but necessary for values with unknown size or that outlive their scope
        * heap is where the largest amounts of memory are usually stored