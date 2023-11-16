package main

import (
	"fmt"
)

func addElementToMap(myMap map[string]int, key string, value int) {
    myMap[key] = value
}


func main() {
    fmt.Println("Hello, World!")

	m := make(map[string]int)
	
    //m["v1"] = 10
    //m["v2"] = 12
    addElementToMap(m, "first", 1)
	addElementToMap(m, "second", 2)
	
    fmt.Println("map:", m)

	//n := map[string]int{"val1": 1, "val2": 2}
    //fmt.Println("map:", n)

}