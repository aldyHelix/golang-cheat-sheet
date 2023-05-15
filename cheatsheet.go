package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

/**
go mod init <module-name or project-name>
go install .
*/

/**
* Hello world
* build : go build
* run : go run hello.go
 */

func iniMain() {
	message := greetMe("world")
	fmt.Println(message)
}

func greetMe(name string) string {
	return "Hello, " + name + "!"
}

//end hello world

/**
* Constant
* Constatant can be character, stringm boolean, or numberic values
 */

const Phi = 1.618      //numberic
const Ok = "OK"        //string
const isLeave = true   //boolean
const gradeLevel = "C" //character
const maxScore = 100   //also numberic

//end Constatant

/**
* Basic types
 */

//Strings
//String are of type string (text)
func mainKata() {
	str := "hello"
	str := `multiline
	string`

	//Numbers
	//Typical types
	num := 3         //int
	num := 3.        //float 64
	num := 3 + 4i    //complex128
	num := byte('a') //byte (alias for uint8)

	//other type
	var u uint = 7       //uint (unsigned)
	var p float32 = 22.7 //32-bit float

	//Arrays
	//Arrays have a fixed size
	//var numbers [5]int

	numbers := [...]int{0, 0, 0, 0, 0}

	//Slices
	//Slice have a dynamic size, unlike arrays
	slice := []int{2, 3, 4}
	slice := []byte("Hello")

	//declare empty slice
	var slice []int
	slice := []int{}

	//Type convertions
	i := 2
	f := float64(i)
	u := uint(i)

	//End basic types
}

//Pointers

func iniMainLagi() {
	b := *getPointer() //output : value of address
	fmt.PrintLn("Value is", b)

	a := new(int) // reserved address on ram
	*a = 234      // put value in address
}

func getPointer() (myPointer *int) {
	a := 234
	return &a //output : ram address
}

//Pointers point to a memory location of a variable. go is fully garbage-collected
//reference : https://tour.golang.org/moretypes/1

func mainFlowControl() {
	/**
	* Flow Control
	* for best practice time complexity
	 */

	// { conditional }
	if day == "sunday" || day == "saturday" {
		rest()
	} else if day == "moday" && isTired() {
		groan
	} else {
		work()
	}

	// Statements in if
	if _, err := doThing(); err != nil {
		fmt.Println("Uh oh")
	}

	/* IMPORTANT
	You can’t write a short one-line conditional in Go; there is no ternary conditional operator. Instead of

	res = expr ? x : y
	*/

	/**
	* A condition in a if statement can be preceded with a statement before a;. Variables declared bby the statement are only in scope until the end of the if.
	see : https://tour.golang.org/flowcontrol/6
	*/

	//Switch
	switch day {
	case "sunday":
		// cases don't "fall through" by default
		fallthrough

	case "saturday":
		rest()

	default:
		work()
	}

	//For Loop

	for count := 0; count <= 10; count++ {
		fmt.Println("My counter is at", count)
	}

	//for time complexity [O(n)]

	//For-Range Loop
	//blank identifier
	sammy := "Sammy"
	//letter as variable in scope of looping
	for _, letter := range sammy {
		fmt.Printf("the letter is %c\n", letter)
	}

	entry := []string{"Jack", "Jhon", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}

	//range time complexity [O(n)]

	//While loop
	n := 0
	x := 42
	for n != x {
		n := guess()
	}
	//while time complexity [O(log n)]

	//make sure always using less time complexity
}

/**
*	Functions
 */

func mainLamdas() {
	//Lambdas
	myfunc := func() bool {
		return x > 100000
	}
	//function are first class objects.

	//Multiple return type

	a, b := getMessage()
}

func getMessage() (a string, b string) {
	return "Hello", "World"
}

//Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// By defining the return value names in the signature, a return (no args) will return variables with those names.

/**
* Packages
 */

//importing
/*See On top of this file*/
// import "fmt"
// import "math/rand"

//or
/*
import (
	"fmt" 			// gives fmt.Println
	"math/rand"		// gives rand.Intn
)
*/

// Both are the same.

//Aliases

// import r "math/rand"
// r.Intn()

//Exporting names

func Hello() {
	fmt.Println("Hellow")
}

//Exported names begin with capital letters

// package hello /*See on top of this file*/

//every package file has to start with package.

//============================================
/**
* Concurrency
* concurrency is playing 2 games at once (look like exacly at once)
* but actualy it has different time gap between execution
 */

//Goroutines
func mainLagi() {
	// A "channel"
	ch := make(chan string)

	//start concurrent routines
	go push("Moe", ch)
	go push("Larry", ch)
	go push("Curly", ch)

	//Read 3 result
	// (Since our goroutines are concurrent,
	// the order isn't guaranteed!)
	fmt.Println(<-ch, <-ch, <-ch)
}

func push(name string, ch chan string) {
	msg := "Hey, " + name + "\n"
	ch <- msg
}

//the result of the order isn't exacly same.
//Channels are concurrency-save communication objects, used in goroutines.

func mainanConcurrency() {
	//Buffered channels
	ch := name(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3

	//fatal error:
	// all goroutines are asleep - dealock
	//Buffered channels limit the amount of messages it can keep.

	// Closing channels
	//->closes a channel
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	//iterates across a channel until its closed
	for i := range ch {
		fmt.Println(ch)
	}

	//closed if ok == false
	v, ok := <-ch
}

//reference : https://tour.golang.org/concurrency/4

//WaitGroup
//import "sync" /*Imported on top of file*/

func mainLagiDanLagi() {
	var wg sync.WaitGroup

	for _, item := range itemList {
		// Increment WaitGroup Counter
		wg.Add(1)
		go doOperation(&wg, item)
		wg.Wait()
	}
}

func doOperation(wg *sync.WaitGroup, item string) {
	defer wg.Done() //will be called Done when operation is clear
	//do operation on item
	// ...
}

/**
A Waitgroup waits for a collection of go routines to finish. The main
goroutine calls Add to set the number of goroutines to wait for. The goroutine calls wg.Done() when it finishes.
references : https://golang.org/pkg/sync/#WaitGroup
*/

/**
Error Control
*/

//Defer
func lagiMain() {
	defer fmt.Println("Done")
	fmt.Println("Working ...")
}

/**
Defer running a function until the surronding function returns.
The Arguments are evaluated immadiately, but the function call is not ran until later.
reference : https://blog.golang.org/defer-panic-and-recover
*/

//Deferring functions

func diMainin() {
	defer func() {
		fmt.Println("Done")
	}()
	fmt.Println("Working ...")
}

//Lamdas are better suited for defer blocks.

func mainYangLain() {
	var d = int64(0)
	defer func(d *int64) { //points an address value of d
		fmt.Printf("& %v Unix Sec\n", *d)
	}(&d) //output the address
	fmt.Print("Done .")
	d = time.Now.Unix()
}

/*
the defer func uses current value of d, unless we use a pointer to get final value at the end of main.
*/

/*
Structs
*/

//Defining

type VertexA struct {
	X int
	Y int
}

func diMaininLagi() {
	v := VertexA{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)

	//Literals
	v := VertexA{X: 1, Y: 2}

	// Field names can be omitted
	v := VertexA{1, 2}

	//Y is implicit
	v := VertexA{X: 1}
	//You can also put field names

	//Pointers to structs
	v := &VertexA{1, 2}
	v.X = 2
	//Doing v.X is the same as doing (*v).X, when v is a pointer
}

//reference : https://tour.golang.org/moretypes/2

/*
Methods
*/

//Receivers
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//There are no classes, but you can define functions with recevers
//references : https://tour.golang.org/methods/1

//Mutation
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func mainanVertex() {

	v := Vertex{1, 2}
	v.Abs()

	v := vertex{6, 12}
	v.Scale(0.5)
	// `v` is updated
}

//By defininf your receiver as a pointer (*Vertex), you can do mutations.
//receivers : https://tour.golang.org/methods/4

/**
Interfaces
*/

//A Basic interface

type Shape interface {
	Area() float64
	Perimeter() float64
}

//Struct
type Rectangle struct {
	Length, Width float64
}

//Struct Rectangle implicity implements interface Shape by implementing all of its methods.

// Methods
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

//The methods defined in Shape are implemented in Rectangle.

//Interface example
func mainanTerakhir() {
	var r Shape = Rectangle{Length: 3, Width: 4}
	fmt.Println("Type of r: %T, Area: %v, Perimeter: %v", r, r.Area(), r.Perimeter())
}

/*
REFERENCES
A Tour of Go
https://tour.golang.org/welcome/1
Golang Wiki
https://github.com/golang/go/wiki
Effective Go
https://goalng.org/effective_go.html

Other References
Go by Example
Https://gobyexample.com
Awesome Go
https://awesome-go.com
JustForFunc Youtube
https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw
Style Guide
https://github.com/golang/go/wiki/CodeReviewComments

EOF
*/
