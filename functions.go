//go has property to have more than one return type and it can have variable parameters

package main

import "fmt"

func vals() (int, int) {
    return 3, 7                   // two int returns
}



func sum(nums ...int) {          // any number of parameters
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}


func intSeq() func() int {       // closure or inline function (return type is a func itself)
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    a, _ := vals()
    fmt.Println(a)
	
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	
	nextInt := intSeq()
	fmt.Println(nextInt())
    fmt.Println(nextInt())
}