package main

import (
	"fmt"
	"net"
	"sync"
)

// const declared at the top of the file and grouped
const (
	Scalar  = 0.1 // all lowercase for local use, capitalised for exportable
	Version = 1.0
)

// function grouping: exported functions should be on top, then sort by function importance
func ExportedFunc()      {}
func veryImportantFunc() {}
func someUtilFunc()      {}

func Foo() int {
	// group declared variables on top of the function block
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)

	fmt.Println(foo)
	return x + y
}

// variable names use single characters
// function that panics (doesn't capture error has the prefix Must)
func MustParseIntFromString(s string) int {
	panic("oops")

	return 10
}

type Vector struct {
	x int
	y int
}

type Server struct {
	listenAddr string
	isRunning  bool

	// put the mutex closest to the field it wants to protect
	peersMu sync.RWMutex
	peers   map[string]net.Conn
}

// interface name should ends with 'er'
// interface segregation: interface mustn't do too many things
type RepositoryHandler interface {
	Getter
	Setter
}

type Getter interface {
	Get()
}

type Setter interface {
	Put()
	Delete()
}

// HTTP handler functions start with handle prefix
func handleGetUserById() {}
func handleResizeImage() {}

// writing enums
type Suit byte

const (
	// starts enum with the type it implements
	SuitHearts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

type Order struct {
	Size float64
}

// construct with order.New(), with order as the package name
func New() *Order {
	return &Order{
		Size: 69.0,
	}
}

func main() {
	vector := Vector{x: 10, y: 20} // declare struct with named variables
	fmt.Println(vector)
}
