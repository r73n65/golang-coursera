# Go Cheatsheet
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

## Slice
Underlying array
* len() - returns the length
* cap() - returns the capacity
* make() - create a slice (and array)
* 2 argument: type, length = capacity
sli = make([]int, 10)
* 3 arguments: type, length, capacity
sli = make([]int, 10, 15)
* append() - inserts into the underlying array 

## Maps
* Implementation of a hash table
* Use make()
var idMap map[string]int 
idMap = make(map[string]int) 

### Accessing Maps
* Referencing a value with [key]
* Returns zero if key is not present
fmt.Println(idMap["joe"])
* Adding a key/value pair
idMap["jane"] = 456
* Deleting a key/value pair
delete(idMap, "joe")
* Two-value assignment tests for existence of the key
id, p := idMap["joe"]
id is value, p is the presence of key 
* len() - returns number values

### Iterating through a map
for key, value := range idMap {
	fmt.Println(key, val)
}

## Struct
* new() - to initialise struct
p1 := new(Person)
* struct literal - to initialise struct 
p1 := Person (name: "joe", 
addr: "s st", 
phone: "123")

## JSON Marshalling 
* Marshal() - return JSON representation as []byte
barr, err := json.Marshal(p1)

## JSON Unmarshalling 
* Unmarshal() - converts a JSON []byte into a Go object
err := json.Unmarshal(barr, &p2)
* Pointer to Go object is passed to Unmarshal()
* Object must fit JSON [] byte (same struct)

## Files
### Basic Operations
1. Open - get handle for access
2. Read - read bytes into []byte
3. Write - write []byte into file 
4. Close - release handle
5. Seek - move read/write head

### ioutil File Read
* "io/ioutil" package has basic functions 
dat, e := ioutil.ReadFile("test.txt")
* dat is []byte filled with contents of entire file
* Explicit open/close are not needed 
* Large files cause a problem 

### ioutil File Write
* Writes []byte to file 
* Creates a file
* Unix-style permission bytes
dat = "Hello, world"
err := ioutil.WriteFile("outfile.txt", dat, 0777)

### os Package File Access
* os.Open() - opens a file 
* os.Close() - closes a file
* os.Read() - reads from a file into a []byte
* os.Write() - writes a []byte into a file

### os File Reading
* Open and reading 
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()
* Reads and fills `barr`
* `Read` returns # of bytes read
* Maybe less than []byte length

### os File Create/Write
* WriteString() - writes a string
* Write() - writes a []byte (any unicode sequence)
f, err := os.Create("outfile.txt")
barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
