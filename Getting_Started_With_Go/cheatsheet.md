## The Go Tool Commands
1. `go build` - compiles the program
2. `go doc` - prints documentation for a package 
3. `go fmt` - formats source code files
4. `go get` - downloads packages and installs them
5. `go list` - list all installed packages
6. `go run` - compiles .go files and runs the executable
7. `go test` - runs tests using files ending in "_test.go"

## Pointers
* A pointer is an address to data in memory 
* & operator returns the address of a variable/function
* * operator returns data at an address (dereferencing)

### New
* Alternate way to create a variable
* `new()` function creates a variable and returns a pointer to the variable
ptr := new(int)
*ptr = 3 

## Variable Scope
* Places in code where a variable can be accessed

## Block 
* A sequence of declarations and statements within matching brackets, {}
* Hierarchy is in placed 

## Lexical Scoping 
* bi >= bj where bj is defined inside bi 

## Deallocating Space
* when a variable is no longer needed, it should be deallocated
* Memory space is made available
* Otherwise, we will eventually run out of memory

## Stack 
* Stack is dedicated to function calls
* Local variables are stored here
* Deallocated after function completes 

## Heap
* Heap is persistent 

## Garbage Collection in Go
* Go is a compiled language which enables garbage collection 
* Implementation is fast 
* Compiler determines stack vs heap
* Garbage collection in the background 

## Unicode Packages 
* IsDigit(r rune)
* IsSpace(r rune)
* IsLetter(r rune)
* IsLower(r rune)
* IsPunct(r rune)
* ToUpper(r rune)
* ToLower(r rune)

## String Packages
* Compare(a, b)
* Contains(s, substr)
* HasPrefix(s, prefix)
* Index(s, substr)
* Replace(s, old, new, n)
* ToLower(s)
* ToUpper(s)
* TrimSpace(s)

## Strconv Package 
* Atoi(s) - converts string s to int 
* Itoa(i) - converts int to s 
* formatFloat( f, fmt, prec, bitSize)
* parseFloat(s, bitSize)

## Iota
* Each constant is assigned to a unique integer 
* Starts at 1 and increments 
* iota represents each constant is going to be different