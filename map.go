package main
import "fmt"
func main() {

fru := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range fru {
        fmt.Printf("%s -> %s\n", k, v)  // key and value
	}
}