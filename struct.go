package main

import "fmt"

type person struct {
    name string
    age  int
}

type rect struct {
    width, height int
}

func newPerson(name string) person {   

    p := person{name: name}
    p.age = 42
    return p
}


func (r *rect) area() int {        // defined on r , can also be(func (r rect) area() int{})
    return r.width * r.height
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp.age = 51
	fmt.Println(sp.age)
	
	//_________________________________________________________________

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
}


//Interface - just like class it contains diff functions