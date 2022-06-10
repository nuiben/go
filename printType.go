package main

import (
	"fmt"
)

func main() {
	t1 := "text"
	t2 := []string{"apple", "strawberry", "blueberry"}
	t3 := map[string]float64{"strawberry": 3.2, "blueberry": 1.2}
	t4 := 2
	t5 := 4.5
	t6 := true

	fmt.Printf("t1: %T\n", t1)
	fmt.Printf("t2: %T\n", t2)
	fmt.Printf("t3: %T\n", t3)
	fmt.Printf("t4: %T\n", t4)
	fmt.Printf("t5: %T\n", t5)
	fmt.Printf("t6: %T\n", t6)
}
