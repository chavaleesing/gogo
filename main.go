package main

import (
	"fmt"
)

func main() {
	var x int = 4
	y := 10
	fmt.Printf("test %v, xxx %v ,,%v ,,%v\n\n", 33, true, x, y)
	if y > 10 {
		y--
	} else {
		y++
	}

	xx := []string{"a d", "f l"}
	// xx = append(xx, "6789")
	yy := map[string]string{}
	yy["sing"] = "eiei"
	zz, ok := yy["sindfgg"]
	println(ok)
	fmt.Printf("============= \n %v  ********  %v  ********  %v, %T   \n================\n", xx, yy["sing"], zz, zz)

	for i := 0; i < len(xx); i++ {
		println(xx[i])
	}

	println(sum(7, 8))

	sub := func(a, b int) int {
		return a - b
	}

	println(sub(5, 50))

	cal(sum)
	cal(sub)

	sum2 := func(x ...int) int {
		anss := 0
		for _, val := range x {
			anss += val
		}
		return anss
	}

	sum2(2, 2, 2, 2, 2, 2, 2, 2, 2, 2)

}

func sum(a int, b int) int {
	return a + b
}

func cal(fn func(int, int) int) {
	ans := fn(30, 40)
	println(ans)
}
