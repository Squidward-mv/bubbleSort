package main

import "fmt"

func main()  {
	bubbleSort()
}

func bubbleSort()  {
	var n int
	fmt.Print("Enter the number of the array: ")
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := range arr {
		fmt.Println("Enter the elements of the array: ")
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if arr[j] < arr[j-1]{
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := range arr {
		fmt.Print(arr[i], " ")
	}
}
