package main

import "fmt"

func showVariables() {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	var king2 = false
	king3 := 3457348534
	fmt.Println(i, j, k, c, python, java, king2, king3)
	var kalyanName = "Kalyan Varma"
	fmt.Println(kalyanName)
	j2, k2, l2 := "shark", 2.05, 15

	fmt.Println(j2)
	fmt.Println(k2)
	fmt.Println(l2)

	const shark = "Sammy"
	fmt.Println(shark)
	//iota
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
}

func main() {
	k := 43289234
	fmt.Println(k)
	s := "Helllopdgkldnfgkjdf"
	f := 3463485.3458934895
	b := 5 > 9
	array := [4]string{"Heh", "fdsfs", "dflkgjdfg", "dfgkljdflkg"}
	slice := []string{"one", "two", "three"}
	m := map[string]string{"letter": "g", "number": "seven", "symbol": "&"}

	fmt.Println(s, f, b, array, slice, m)
	showVariables()
	zeroValue()
}
func zeroValue() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n\n", d, d)
	//Variablenaming
	names := []string{"Mary", "John", "Bob", "Anna"}
	for i, n := range names {
		fmt.Printf("index: %d = %q\n", i, n)
	}
}
