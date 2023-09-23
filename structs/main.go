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

func printFunctionMappings(mappings *map[string]string) {
	for key, value := range *mappings {
		fmt.Printf("\t- %s %s\n", key, value)
	}
}

func printStructs(structs *[]string) {
	for _, value := range *structs {
		fmt.Printf("\t- %s\n", value)
	}
}

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

	fmt.Println("\n\nLet's create examples of embedded structs and using interfaces.")
	fmt.Println("The interface is called 'entity' and has the following methods:")
	methodsAndReturnTypes := map[string]string{
		"getName()":             "string",
		"getPosition()":         "position",
		"setPosition(position)": "",
		"getDistance(entity)":   "float64",
	}
	printFunctionMappings(&methodsAndReturnTypes)

	fmt.Println("The entity interface is implemented by the following structs:")
	allStructs := []string{"player", "enemy"}
	printStructs(&allStructs)

	fmt.Println("Let's create a simple player and enemy.")
	myPlayer := makePlayer(makePosition(0, 0, 0))
	myEnemy := makeEnemy(makePosition(10, 10, 10))
	fmt.Printf("Player is: %+v and enemy is %+v\n", myPlayer, myEnemy)

	fmt.Println("Let's move the player towards the enemy.")
	moveTowards(myPlayer, myEnemy.getPosition())
	fmt.Printf("Player is now: %+v and enemy is %+v\n", myPlayer, myEnemy)
}
