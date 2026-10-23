package main

import (
	"fmt"
	"os"
	"strconv"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	val, ok := os.LookupEnv("OUTPUT_COUNT")
	if !ok {
		fmt.Println("OUTPUT_COUNT not set")
		return
	}
	count, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("OUTPUT_COUNT must be an integer")
	}
	fmt.Printf("OUTPUT_COUNT: %d\n", count)
	for i := 0; i < count; i++ {
		fmt.Printf("%d\t", i)
	}
}
