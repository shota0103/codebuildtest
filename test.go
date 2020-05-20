package mysum

import (
	"fmt"
)

func mysum(x, y int) int {
	return x + y * 2 
}

func main() {
	fmt.Println(mysum(5, 2))
}
