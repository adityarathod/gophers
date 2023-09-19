package main

import (
	"fmt"
	"unicode/utf8"
)

type Definition struct {
	definition string
	language   string
}

// Makes a new Definition using the given definition and language.
func makeDefinition(definition string, language string) *Definition {
	return &Definition{definition, language}
}

// Definition getter
func (d *Definition) getDefinition() string {
	return d.definition
}

// Definition setter
func (d *Definition) setDefinition(definition string) {
	d.definition = definition
}

// Type alias for a map from strings to Definitions.
type Dictionary = map[string]*Definition

func main() {
	greetings := make(Dictionary)
	greetings["hello"] = makeDefinition("A standard greeting", "English")
	greetings["hola"] = makeDefinition("A standard greeting", "Spanish")
	greetings["안녕하세요"] = makeDefinition("A standard greeting", "Korean")

	fmt.Println("Here is a dictionary of greetings.")
	for key, value := range greetings {
		fmt.Printf("%s: %s (source language: %s)\n", key, value.getDefinition(), value.language)
		fmt.Printf("%s has %d runes\n", key, utf8.RuneCountInString(key))
		fmt.Println("Rune listing:", []rune(key))
		fmt.Println()
	}

	fmt.Println("Let's change the definition of 'hello' in Korean to a more specific definition.")
	greetings["안녕하세요"].setDefinition("A respectful greeting")
	fmt.Printf("The new definition of 'hello' in Korean is: %s\n", greetings["안녕하세요"].getDefinition())

}
