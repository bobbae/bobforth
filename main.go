package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ForthInterpreter holds the stack and dictionary of words
type ForthInterpreter struct {
	stack   []int
	words   map[string]func(*ForthInterpreter)
	defMode bool          // Indicates if we're defining a new word
	newWord string        // Holds the name of the new word being defined
	wordDef []string      // Holds the commands for the new word definition
}

// NewForthInterpreter creates a new Forth interpreter
func NewForthInterpreter() *ForthInterpreter {
	interpreter := &ForthInterpreter{
		stack: []int{},
		words: make(map[string]func(*ForthInterpreter)),
	}

	// Define built-in Forth words (commands)
	interpreter.defineBasicWords()

	return interpreter
}

// defineBasicWords defines the core arithmetic and stack manipulation words
func (f *ForthInterpreter) defineBasicWords() {
	// Arithmetic operations
	f.words["+"] = func(f *ForthInterpreter) {
		b, a := f.pop(), f.pop()
		f.push(a + b)
	}
	f.words["-"] = func(f *ForthInterpreter) {
		b, a := f.pop(), f.pop()
		f.push(a - b)
	}
	f.words["*"] = func(f *ForthInterpreter) {
		b, a := f.pop(), f.pop()
		f.push(a * b)
	}
	f.words["/"] = func(f *ForthInterpreter) {
		b, a := f.pop(), f.pop()
		if b == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		f.push(a / b)
	}

	// Stack operations
	f.words["."] = func(f *ForthInterpreter) {
		fmt.Println(f.pop())
	}
	f.words["dup"] = func(f *ForthInterpreter) {
		top := f.top()
		f.push(top)
	}
	f.words["swap"] = func(f *ForthInterpreter) {
		b, a := f.pop(), f.pop()
		f.push(b)
		f.push(a)
	}

	// Word definition (start with ":")
	f.words[":"] = func(f *ForthInterpreter) {
		f.defMode = true
	}
}

// Execute parses and executes a Forth input string
func (f *ForthInterpreter) Execute(input string) {
	tokens := strings.Fields(input)
	for _, token := range tokens {
		// If we are in definition mode, handle the word creation
		if f.defMode {
			if token == ";" {
				// End of word definition
				f.defMode = false
				f.addNewWord(f.newWord, f.wordDef)
				f.newWord = ""
				f.wordDef = []string{}
			} else if f.newWord == "" {
				// First token after ":" is the new word name
				f.newWord = token
			} else {
				// Subsequent tokens are part of the word definition
				f.wordDef = append(f.wordDef, token)
			}
		} else {
			// Normal execution
			if word, exists := f.words[token]; exists {
				word(f) // Execute Forth word
			} else if number, err := strconv.Atoi(token); err == nil {
				f.push(number) // Push number onto the stack
			} else {
				fmt.Println("Unknown word:", token)
			}
		}
	}
}

// addNewWord adds a new user-defined word to the dictionary
func (f *ForthInterpreter) addNewWord(name string, definition []string) {
	f.words[name] = func(f *ForthInterpreter) {
		for _, token := range definition {
			if word, exists := f.words[token]; exists {
				word(f) // Execute Forth word
			} else if number, err := strconv.Atoi(token); err == nil {
				f.push(number) // Push number onto the stack
			} else {
				fmt.Println("Unknown word in definition:", token)
			}
		}
	}
}

// push adds a value to the stack
func (f *ForthInterpreter) push(value int) {
	f.stack = append(f.stack, value)
}

// pop removes and returns the top value from the stack
func (f *ForthInterpreter) pop() int {
	if len(f.stack) == 0 {
		fmt.Println("Error: Stack underflow")
		return 0
	}
	value := f.stack[len(f.stack)-1]
	f.stack = f.stack[:len(f.stack)-1]
	return value
}

// top returns the top value from the stack without removing it
func (f *ForthInterpreter) top() int {
	if len(f.stack) == 0 {
		fmt.Println("Error: Stack underflow")
		return 0
	}
	return f.stack[len(f.stack)-1]
}

func main() {
	interpreter := NewForthInterpreter()

	// Example Forth commands with word definition
	fmt.Println("Forth Interpreter Example:")
	interpreter.Execute(": square dup * ;") // Define a new word 'square'
	interpreter.Execute("5 square .")       // Outputs: 25
	interpreter.Execute("3 4 + square .")   // Outputs: 49
}
