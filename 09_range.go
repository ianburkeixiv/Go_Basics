package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n, i", id)
	}

	// Not using index..just put underscore instead of i
	for _, id := range ids {
		fmt.Printf("%d - ID: %d\n", id)
	}

	// Add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Joe": "Joe@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
