package main

import (
	"fmt"
	"gogo/customer"
	"gogo/user"
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

	// -------------------------------------------------------------------------------------------------------------------

	println(customer.Name)
	println(customer.Print_name())

	test1 := 10
	fmt.Println((&test1))

	// -------------------------------------------------------------------------------------------------------------------

	q := 10
	// var w *int
	// w = &q
	w := &q
	println(q, w)

	*w = 999
	println(q, w)

	st1 := "test jaa"
	println(st1)
	set_value_format(&st1)
	println(st1)

	println(user.Get_info(user.Person{Name: "peter", Age: 32}))

	user_1 := user.Person{Name: "u1", Age: 1}
	println(user_1.Name)
	user_1.Rename("new new")
	println(user_1.Name)

}

func sum(a int, b int) int {
	return a + b
}

func cal(fn func(int, int) int) {
	ans := fn(30, 40)
	println(ans)
}

func set_value_format(input *string) {
	*input = "I'm new format of " + *input
}
