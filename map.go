package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "rifat"
	m["age"] = "24"

	fmt.Println("full array is:", m)
	fmt.Println("array length:", len(m))

	delete(m, "age")
	fmt.Println("full array is:", m)

	_, prs := m["name"]

	_, check := m["age"]

	fmt.Println("key exist on array? S:", prs, check) //return bollean value

	skill := map[string]int{"a": 1, "b": 2}

	fmt.Println("deleare array is:", skill)

}
