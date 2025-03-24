package main

import "fmt"

const VAT = 24

type Item struct {
	Description string
	Value       float64
}

func Value(price float64) float64 {
	total := price + price*VAT/100
	return total
}

func main() {
	t := Item{Description: "Keyboard", Value: 100}
	t.Value = Value(t.Value)
	fmt.Println(t)

	tP := &Item{}
	*&tP.Description = "Mouse"
	*&tP.Value = 100
	fmt.Println(tP)
}

// go run -gcflags '-m' .          
// # cty.sh/allocate
// ./main.go:12:6: can inline Value with cost 13 as: func(float64) float64 { total := price + price * 24 / 100; return total }
// ./main.go:17:6: cannot inline main: function too complex: cost 199 exceeds budget 80
// ./main.go:19:17: inlining call to Value
// ./main.go:20:13: inlining call to fmt.Println
// ./main.go:25:13: inlining call to fmt.Println
// ./main.go:22:8: &Item{} escapes to heap:
// ./main.go:22:8:   flow: tP = &{storage for &Item{}}:
// ./main.go:22:8:     from &Item{} (spill) at ./main.go:22:8
// ./main.go:22:8:     from tP := &Item{} (assign) at ./main.go:22:5
// ./main.go:22:8:   flow: {storage for ... argument} = tP:
// ./main.go:22:8:     from tP (interface-converted) at ./main.go:25:14
// ./main.go:22:8:     from ... argument (slice-literal-element) at ./main.go:25:13
// ./main.go:22:8:   flow: fmt.a = &{storage for ... argument}:
// ./main.go:22:8:     from ... argument (spill) at ./main.go:25:13
// ./main.go:22:8:     from fmt.a := ... argument (assign-pair) at ./main.go:25:13
// ./main.go:22:8:   flow: {heap} = *fmt.a:
// ./main.go:22:8:     from fmt.Fprintln(os.Stdout, fmt.a...) (call parameter) at ./main.go:25:13
// ./main.go:20:14: t escapes to heap:
// ./main.go:20:14:   flow: {storage for ... argument} = &{storage for t}:
// ./main.go:20:14:     from t (spill) at ./main.go:20:14
// ./main.go:20:14:     from ... argument (slice-literal-element) at ./main.go:20:13
// ./main.go:20:14:   flow: fmt.a = &{storage for ... argument}:
// ./main.go:20:14:     from ... argument (spill) at ./main.go:20:13
// ./main.go:20:14:     from fmt.a := ... argument (assign-pair) at ./main.go:20:13
// ./main.go:20:14:   flow: {heap} = *fmt.a:
// ./main.go:20:14:     from fmt.Fprintln(os.Stdout, fmt.a...) (call parameter) at ./main.go:20:13
// ./main.go:20:13: ... argument does not escape
// ./main.go:20:14: t escapes to heap
// ./main.go:22:8: &Item{} escapes to heap
// ./main.go:25:13: ... argument does not escape