package main

import (
	"./list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println("length: ", l.Len())

	fmt.Println("emptyness: ", l.Empty())
	l.Unshift("1")
	l.Unshift(42)
	l.Append(43)
	l.Delete(2)
	l.Append("Finish")
	fmt.Println("emptyness: ", l.Empty())
	fmt.Println("length: ", l.Len())

	item, err := l.Next()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item)

	item, err = l.Next()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item)

	item, err = l.Next()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item)
}
