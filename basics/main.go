package main

import (
	"fmt"
	"os"
)

// Computes a digit sum iteratively.
func sumDigits(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	tmp := num
	for tmp != 0 {
		sum += tmp % 10
		tmp /= 10
	}
	return sum
}

// Uses the Luhn algorithm to validate digits of an identifier, such as a credit card.
func validateLuhn(digits []int) bool {
	// Check that we have at least 2 digits
	if len(digits) < 2 {
		fmt.Println("Invalid ID number: too short")
		os.Exit(1)
	}

	// Last digit is the check digit we're comparing against
	check_digit := digits[len(digits)-1]

	// Sum all digits, doubling every other one (and summing its digits) starting with the first
	should_double := true
	all_sum := 0
	for i := len(digits) - 2; i >= 0; i-- {
		if should_double {
			all_sum += sumDigits(2 * digits[i])
		} else {
			all_sum += digits[i]
		}
		should_double = !should_double
	}
	// Check digit formula: (10 - (sum mod 10)) mod 10
	computed_check := (10 - all_sum%10) % 10
	return check_digit == computed_check
}

func main() {
	fmt.Println("Luhn's checker")
	fmt.Print("ID number to check: ")
	var identifier string
	n, err := fmt.Scanf("%s", &identifier)
	if err != nil {
		fmt.Printf("Got error while processing ID number: %s", err.Error())
		os.Exit(1)
	}
	if n != 1 {
		fmt.Println("Expected 1 ID number, got", n)
		os.Exit(1)
	}
	// Parse digits
	digits := make([]int, len(identifier))
	for i, chr := range identifier {
		tmp := int(chr - '0')
		if tmp < 0 || tmp > 9 {
			fmt.Printf("Invalid digit: %c\n", identifier[i])
			os.Exit(1)
		}
		digits[i] = tmp
	}
	isValid := validateLuhn(digits)
	if isValid {
		fmt.Println("Valid ID number")
	} else {
		fmt.Println("Invalid ID number.")
	}
}
