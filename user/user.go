package user

import "strconv"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Rename(new_name string) {
	p.Name = new_name
}

func Get_info(pp Person) string {
	// u := Person{
	// 	Name: "sing",
	// 	Age:  28,
	// }
	//println(u)
	return "Name: " + pp.Name + " Age: " + strconv.Itoa(pp.Age)
}
