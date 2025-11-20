package main

import "fmt"

func main() {

	// If statement
	gradle := 199

	if gradle <= 100 {
		fmt.Println("gradle is greater than ", gradle)
	} else {
		fmt.Println("gradle is less than ", gradle)
	}

	score := 60
	if score >= 90 {
		fmt.Println("A grade")
	} else if score >= 80 {
		fmt.Println("B grade")
	} else if score >= 70 {
		fmt.Println("C grade")
	} else if score >= 65 {
		fmt.Println("D grade")
	} else {
		fmt.Println("Failing grade")
	}

	// for loop
	for i := 0; i < 100; i++ {
		println("value is", i)
	}

	// while loop using for loop
	i2 := 1
	for i2 <= 5 {
		fmt.Println(i2)
		i2++
	}

	// break statement
	for i3 := 1; i3 <= 5; i3++ {
		if i3 == 3 {
			// terminate the loop when i3 is 3
			break
		}
		fmt.Println(i3)
	}

	// continue statement
	for i4 := 1; i4 <= 5; i4++ {
		if i4 == 3 {
			// skip the current iteration when i4 is 3
			continue
		}
		fmt.Println(i4)
	}
	// goto statement
	for i5 := 1; i5 <= 5; i5++ {
		if i5 == 3 {
			// jump to the label when i5 is 3
			goto label
		}
		fmt.Println(i5)
	}
	// label
label:
	fmt.Println("Jumped to label")

	// looping through maps
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	
	// looping through strings
	str := "Hello, !"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}
