package mysum

import (
	"fmt"
)

func mysum(x, y int) int {
	return x + x + y
}

func main() {
	fmt.Println(mysum(5, 2))
}
