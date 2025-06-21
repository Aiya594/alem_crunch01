package main

import (
	//"github.com/alem-platform/ap"
	"fmt"
)

func ErrorMes() {
	fmt.Print("E")
	fmt.Print("R")
	fmt.Print("R")
	fmt.Print("O")
	fmt.Print("R")
}

func main() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w) // first line input

	if h == 0 || w == 0 {
		ErrorMes()
	} else {
		// initializing 2D array
		ma := make([][]int, h)

		for i := range ma {
			ma[i] = make([]int, w)
		}
		// filling 2d array from terminal
		for i := 0; i < h; i++ {
			var nums string
			fmt.Scanf("%s", &nums) // to read each line as string
			if len(nums) > w {
				fmt.Print("ERROR") // it will be changed to ap.PutRune()
			}
			for j := 0; j < w; j++ {
				ma[i][j] = int(nums[j] - 48)
			}
		}

		fmt.Print(ma)

	}
}
