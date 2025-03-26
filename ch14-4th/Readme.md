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
    * slice-leaks
        * Slices can cause memory leaks when returning a slice of a larger slice
         * This happens because a slice is just a view into an underlying array - when you slice a large array and return a small portion, the entire backing array is still referenced and can't be garbage collected
        * The compiler shows this with: `leaking param: s to result ~r0 level=0`
        * Even when only using a small portion (e.g., 3 elements), the entire backing array remains in memory
        * Solution: create a new slice and copy only the needed elements
        * Use `copy()` function to prevent slice leaks
        * Run with `-gcflags '-m -l'` to see escape analysis and identify potential leaks
        * Memory usage difference can be significant when original slices are large (e.g., 1,000,000 elements)
    * maps-leaks
        * Maps don't automatically shrink after deleting elements
        * Even after deleting all elements, the map's memory isn't fully reclaimed
        * Running garbage collection (`runtime.GC()`) helps but doesn't fully reclaim memory
        * Setting the map to nil and running GC is necessary to fully reclaim memory
        * Use `runtime.MemStats` to track memory allocation
        * `runtime.KeepAlive(objectVar)` prevents premature garbage collection of referenced objects
        * Large maps with value types (like arrays) can consume significant memory
        * Deleting map entries doesn't immediately free the memory used by their values