package main

import (
	"fmt"
	"math"
	"os"
)

// Returns a closure that produces the next Fibonacci number.
func fibonnaciFactory() func() int {
	var last, next = 0, 1
	return func() int {
		ret := last
		last, next = next, last+next
		return ret
	}

}

// Returns a closure that adds a constant to its argument. An example of partial application.
func addFactory(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Example of a curried function. This is a function to solve quadratic equations.
func curriedQuadraticSolver(a float64) func(float64) func(float64) []float64 {
	if a == 0 {
		fmt.Print("a cannot be 0")
		os.Exit(1)
	}
	return func(b float64) func(float64) []float64 {
		return func(c float64) []float64 {
			// Use the quadratic formula to solve for x
			discriminant := math.Pow(b, 2.0) - 4*a*c
			if discriminant < 0 {
				fmt.Printf("No real solutions (discriminant = %f)\n", discriminant)
				return nil
			} else if discriminant == 0 {
				return []float64{-b / 2 * a}
			}
			return []float64{(-b + math.Sqrt(discriminant)) / 2 * a, (-b - math.Sqrt(discriminant)) / 2 * a}
		}
	}
}

// Sum a bunch of numbers.
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Let's do functions!")
	fmt.Println("Fibonacci numbers:")
	fib := fibonnaciFactory()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
	fmt.Println("Partial application to make an adder:")
	add3 := addFactory(3)
	fmt.Println("3 + 5 =", add3(5))
	fmt.Println("3 + 7 =", add3(7))

	fmt.Println("Curried function example:")
	fmt.Println("Solving a quadratic equation with a = 1, b = 2, c = 3 (x^2 + 2x + 3 = 0)")
	a, b, c := 1.0, 2.0, 3.0
	fmt.Println("curriedQuadraticSolver(a)(b)(c) =", curriedQuadraticSolver(a)(b)(c))
	fmt.Println("Solving a quadratic equation with a = 1, b = 20, c = 3 (x^2 + 20x + 3 = 0)")
	a, b, c = 1.0, 20.0, 3.0
	fmt.Println("curriedQuadraticSolver(a)(b)(c) =", curriedQuadraticSolver(a)(b)(c))

	fmt.Println("Variadic function example:")
	fmt.Println("sum(1, 2, 3, 4, 5) =", sum(1, 2, 3, 4, 5))

}
