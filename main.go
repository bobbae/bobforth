package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ForthInterpreter holds the stack and dictionary of words
type ForthInterpreter struct {
	stack []int
	words map[string]func(*ForthInterpreter)
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
}

// Execute parses and executes a Forth input string
func (f *ForthInterpreter) Execute(input string) {
	tokens := strings.Fields(input)
	for _, token := range tokens {
		if word, exists := f.words[token]; exists {
			word(f) // Execute Forth word
		} else if number, err := strconv.Atoi(token); err == nil {
			f.push(number) // Push number onto the stack
		} else {
			fmt.Println("Unknown word:", token)
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

	// Example Forth commands
	fmt.Println("Forth Interpreter Example:")
	interpreter.Execute("3 4 + .")  // Outputs: 7
	interpreter.Execute("10 2 / .") // Outputs: 5
	interpreter.Execute("5 dup * .") // Outputs: 25
}
