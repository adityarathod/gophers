package main

import "fmt"

func main() {
	// Simple int linked list
	fmt.Println("Creating a new linked list of type int")
	myList := makeLinkedList[int]()
	fmt.Println("Adding 1, 2, 3, 4, 5 to the linked list")
	myList.addMany(1, 2, 3, 4, 5)
	fmt.Printf("List: %s\n", myList.toString())

	fmt.Println("Finding the value at index 1")
	num, _ := myList.get(1)
	fmt.Printf("myList[1] = %d\n", num)

	fmt.Println("Finding the value at index 999 (should error out)")
	_, err := myList.get(999)
	fmt.Printf("Error: %s\n", err)
}
