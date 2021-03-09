package main

import "fmt"

func main() {
	fmt.Println("Hi go!!")

	fmt.Println("pai"+"value"+" "+"is:", 3+.1416)

	var a string = "rifat"

	fmt.Println("my name"+" "+"is:", a)

	i := 1

	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 5; j <= 10; j++ {

		if j < 9 && j > 6 {
			fmt.Println(j)
		}
	}

	if num := 15; num%5 == 0 {
		fmt.Println(num, "is divisible by 5")
	}

}
