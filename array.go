package main

import "fmt"

func main() {
	var arr [5]int

	arr[4] = 99

	fmt.Println("array is:", arr)

	b := []string{"a", "b", "ab"}

	fmt.Println("declare array is:", b)

	s := make([]string, 2) //2 is the length

	fmt.Println(len(s))

	s = append(s, "abc")
	s = append(s, "abcd")

	fmt.Println("array s is:", s)

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("array c is:", c)

}
