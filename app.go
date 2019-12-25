package main

import (
    "fmt"
    "github.com/emirpasic/gods/maps/linkedhashmap"
)

func main() {
    m := linkedhashmap.New() // empty (keys are of type int)
    m.Put("aaa", "b")            // 2->b
    m.Put("cdc", "x")            // 2->b, 1->x (insertion-order)
    m.Put("cacs", "a")            // 2->b, 1->a (insertion-order)
    b, _ := m.ToJSON()
    fmt.Println(string(b))

}
