package binarytree

import (
	"fmt"
	"math/rand"
	"time"
)

// Node is a node in a binary tree
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func init() {
	// Seed the random numbers from the clock
	rand.Seed(time.Now().UnixNano())
}

// Create returns a pointer to a new Node with rootData as its Data
func Create(rootData interface{}) *Node {
	newNode := &Node{
		Data:  rootData,
		Left:  nil,
		Right: nil}

	return newNode
}

// InsertLeft inserts a new Node with the specifed Data to the left of the
// current Node
func (node *Node) InsertLeft(data interface{}) *Node {
	newNode := Create(data)
	node.Left = newNode

	return newNode
}

// InsertRight inserts a new Node with the specifed data to the right of the
// current Node
func (node *Node) InsertRight(data interface{}) *Node {
	newNode := Create(data)
	node.Right = newNode

	return newNode
}

// GenerateRandomIntTree generates a binary tree of given height with randomInt
// ints in the range [minValue, maxValue] as the Node data
func GenerateRandomIntTree(height, minValue, maxValue int) *Node {
	root := Create(getRandomIntInRange(minValue, maxValue))
	generateRandomIntTree(root, minValue, maxValue, height-1)

	return root
}

func generateRandomIntTree(parent *Node, minValue, maxValue, levelCount int) {
	if levelCount <= 0 {
		return
	}
	randomInt := getRandomIntInRange(minValue, maxValue)
	subtreeChoice := getRandomIntInRange(0, 2)

	switch subtreeChoice {

	case 0: // Left subtree
		parent.InsertLeft(randomInt)
		generateRandomIntTree(parent.Left, minValue, maxValue, levelCount-1)

	case 1: // Right subtree
		parent.InsertRight(randomInt)
		generateRandomIntTree(parent.Right, minValue, maxValue, levelCount-1)

	case 2: // Both subtree
		parent.InsertLeft(randomInt)
		generateRandomIntTree(parent.Left, minValue, maxValue, levelCount-1)
		randomInt = getRandomIntInRange(minValue, maxValue)
		parent.InsertRight(randomInt)
		generateRandomIntTree(parent.Right, minValue, maxValue, levelCount-1)
	}
}

func getRandomIntInRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}

// PrintInOrder prints an in-order traversal of the tree with the current Node
// as its root
func (node *Node) PrintInOrder() {
	printInOrder(node, false)
	fmt.Printf("\n")
}

func (node *Node) DebugPrint() {
	printInOrder(node, true)
}

func printInOrder(parent *Node, debug bool) {
	if parent == nil {
		return
	}
	printInOrder(parent.Left, debug)
	fmt.Printf("%v ", parent.Data)
	if debug {
		fmt.Printf("\n\t%v\n\t%v\n", parent.Left, parent.Right)
	}
	printInOrder(parent.Right, debug)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Sets the value at index idx of the slice pointed to by slicePtr to
// val, growing the slice if needed
func dynamicSliceGet(slicePtr *[]string, idx int) string {
	sliceResize(slicePtr, idx+1)
	return (*slicePtr)[idx]
}

func dynamicSliceSet(slicePtr *[]string, idx int, val string) {
	sliceResize(slicePtr, idx+1)
	(*slicePtr)[idx] = val
}

func sliceResize(slicePtr *[]string, newSize int) {
	if newSize <= len(*slicePtr) {
		return
	}
	grownSlice := make([]string, newSize)
	copy(grownSlice, *slicePtr)
	*slicePtr = grownSlice
}
