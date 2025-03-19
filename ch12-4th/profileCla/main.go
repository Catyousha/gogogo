package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"time"
)

func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fibo2(n-1)) + int64(fibo2(n-2))
}

func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

func N1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func N2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func cpuProfiling()  {
	// go tool pprof /var/folders/p9/qybzqs3s1cn4bsqc2mf309_w0000gn/T/cpuProfileCla.out
	// File: profile-cla
	// Type: cpu
	// Time: 2025-03-19 10:54:13 WIB
	// Duration: 14.22s, Total samples = 530ms ( 3.73%)
	// Entering interactive mode (type "help" for commands, "o" for options)
	// (pprof) top
	// Showing nodes accounting for 530ms, 100% of 530ms total
	// Showing top 10 nodes out of 22
	// 	flat  flat%   sum%        cum   cum%
	// 	350ms 66.04% 66.04%      350ms 66.04%  main.N2 (inline)
	// 	150ms 28.30% 94.34%      150ms 28.30%  main.N1 (inline)
	// 	10ms  1.89% 96.23%       10ms  1.89%  runtime.pMask.clear (inline)
	// 	10ms  1.89% 98.11%       10ms  1.89%  runtime.pthread_cond_signal
	// 	10ms  1.89%   100%       10ms  1.89%  runtime.usleep
	cpuFilename := path.Join(os.TempDir(), "cpuProfileCla.out")
	fmt.Println(cpuFilename)
	cpuFile, err := os.Create(cpuFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	total := 0
	for i := 2; i < 100000; i++ {
		n := N1(i)
		if n {
			total = total + 1
		}
	}
	fmt.Println("Total primes:", total)

	total = 0
	for i := 2; i < 100000; i++ {
		n := N2(i)
		if n {
			total = total + 1
		}
	}
	fmt.Println("Total primes:", total)

	for i := 1; i < 90; i++ {
		n := fibo1(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	for i := 1; i < 90; i++ {
		n := fibo2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	runtime.GC()
}

func memoryProfiling() {
	// go tool pprof /var/folders/p9/qybzqs3s1cn4bsqc2mf309_w0000gn/T/memoryProfileCla.out
	// File: profile-cla
	// Type: inuse_space
	// Time: 2025-03-19 10:54:28 WIB
	// Entering interactive mode (type "help" for commands, "o" for options)
	// (pprof) top
	// Showing nodes accounting for 50371kB, 100% of 50371kB total
	// Showing top 10 nodes out of 14
	// 	flat  flat%   sum%        cum   cum%
	// 48832kB 96.94% 96.94%    48832kB 96.94%  main.memoryProfiling
	// 	1539kB  3.06%   100%     1539kB  3.06%  runtime.allocm
	// 		0     0%   100%    48832kB 96.94%  main.main
	// 		0     0%   100%    48832kB 96.94%  runtime.main
	memoryFilename := path.Join(os.TempDir(), "memoryProfileCla.out")
	
	memory, err := os.Create(memoryFilename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer memory.Close()

	for range 10 {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(50 * time.Millisecond)
	}

	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main()  {
	fmt.Println(os.TempDir())
	
	cpuProfiling()
	memoryProfiling()
}