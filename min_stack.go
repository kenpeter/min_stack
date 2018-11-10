package main
import (
	"fmt"
	"sync"
	//"errors"
)

// class
type stack struct {
	// lock
	// sync
	// mutex
	// lock is attr 
	lock sync.Mutex
	// s is attr 
    arr []int
}

// func cap
// constructor 
// return pointer
func MinStack() *stack {
	// return addr, so can point
	// init lock
	// arr len 0 
    return &stack {sync.Mutex{}, make([]int,0), }
}

// connect to class
func (mystack *stack) Push(val int) {
    // push, lock, unlock
	mystack.lock.Lock()
    defer mystack.lock.Unlock()

	// 
    mystack.arr = append(mystack.arr, val)
}

func (mystack *stack) Top() (int) {
	mystack.lock.Lock()
    defer mystack.lock.Unlock()

	// arr len
    l := len(mystack.arr)

    // empty stack, return
    if l == 0 {
        return 0
    }

    // latt item, out
    res := mystack.arr[l-1]

	return res	
}

func (mystack *stack) GetMin() (int) {
	mystack.lock.Lock()
    defer mystack.lock.Unlock()

	l := len(mystack.arr)
	min := mystack.arr[l-1] 
	for _, val := range mystack.arr {
		if(val < min) {
			min = val
		} else {

		}
	}

	return min 
}

// connect to class
// return 2 
func (mystack *stack) Pop() (int) {
	// lock
    mystack.lock.Lock()
	// unlock
    defer mystack.lock.Unlock()

	// arr len
    l := len(mystack.arr)

	// empty stack, return
    if l == 0 {
        return 0
    }

	// latt item, out
    res := mystack.arr[l-1]

	// 0 to l-2 element, inclusive	
    mystack.arr = mystack.arr[:l-1]
	//
    return res
}


func main(){
	// obj
    s := MinStack()
	// 1, 2, 3
    s.Push(5)
    s.Push(3)
    s.Push(2)
	s.Push(6)
	s.Push(4)

	fmt.Println("-- get min --")
    fmt.Println(s.GetMin())

}
