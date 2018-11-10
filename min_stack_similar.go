import (
    "fmt"
    //"errors"
)

type MinStack struct {
    arr []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{make([]int,0)}
}


func (this *MinStack) Push(val int)  {
    // 
    this.arr = append(this.arr, val)
}


func (this *MinStack) Pop() int {
    // arr len
    l := len(this.arr)

    // empty stack, return
    if l == 0 {
        return 0
    }

    // latt item, out
    res := this.arr[l-1]

    // 0 to l-2 element, inclusive  
    this.arr = this.arr[:l-1]
    //
    return res
}


func (this *MinStack) Top() int {
    // arr len
    l := len(this.arr)

    // empty stack, return
    if l == 0 {
        return 0
    }

    // latt item, out
    res := this.arr[l-1]

    return res
}


func (this *MinStack) GetMin() int {
    l := len(this.arr)
    min := this.arr[l-1] 
    for _, val := range this.arr {
        if(val < min) {
            min = val
        } else {
        
        }
    }

    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
