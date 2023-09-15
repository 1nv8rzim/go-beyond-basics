package main

import "fmt"

func main() {
	x := 1
	fmt.Println("1: x = ", x)

	{
		fmt.Println("2: x = ", x)
		x := 2
		fmt.Println("2: x = ", x)
		x = 3
		fmt.Println("2: x = ", x)

		{
			fmt.Println("3: x = ", x)
			x := 4
			fmt.Println("3: x = ", x)
		}

		fmt.Println("2: x = ", x)
	}

	fmt.Println("1: x = ", x)
}
