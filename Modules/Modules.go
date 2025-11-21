package main

import (
	"fmt"

	"example.com/myproject/mathutils"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("UUID:", uuid.New())
	sum := mathutils.Add(33, 44)
	fmt.Println(sum)

}
