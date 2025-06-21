package main

import "fmt" //"github.com/alem-platform/ap"

func drawRow(we int, fill []int) {
	// draw blocks
	for j := 0; j < 3; j++ { // each block - 3 lines
		for i := 0; i < we; i++ {
			fmt.Printf("|")
			// filling 1 block with Space or Xs
			for k := 0; k < 7; k++ {
				if fill[i] == 1 {
					fmt.Printf(" ")
				} else {
					fmt.Printf("X")
				}
			}
		}
		fmt.Printf("|")
		if j != 3 {
			fmt.Printf("\n")
		}

	}
	// bottom line
	fmt.Printf(" ")
	for i := 0; i < we; i++ {
		for j := 0; j < 8; j++ {
			if fill[i] == 1 {
				fmt.Printf("—")
			} else {
				fmt.Printf(" ")
			}
		}
	}
	fmt.Printf("\n")
}

func drawMap(arr [][]int) {
	var hei, wei int = len(arr), len(arr[0])
	// upper bound
	var l int = wei*8 - 1
	if hei != 0 {
		fmt.Printf(" ")
		for i := 0; i < l; i++ {
			fmt.Printf("_") // ap.PutRune('-')
		}
		fmt.Printf("\n")
	}

	// insides
	for i := 0; i < hei; i++ {
		// for j := 0; j < wei; j++ {
		// draw insides (_/X)
		drawRow(wei, arr[i])
		//fmt.Printf("\n")
		//}
	}
}

// func main() {
// 	drawMap([][]int{{1, 0, 1}, {0, 1, 1}})
// 	// fmt.Print([][]int{{1, 1, 1}, {1, 1, 1}})
// 	fmt.Printf("\n")
// }
