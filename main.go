package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func PrintError(err string) {
	for i := 0; i < len(err); i++ {
	}
}

func drawRow(fill []int) {
	var we int = len(fill)
	// draw blocks
	for j := 0; j < 3; j++ { // each block - 3 lines
		for i := 0; i < we; i++ {
			fmt.Printf("|")
			// filling 1 block with Space or Xs
			for k := 0; k < 7; k++ {
				if j == 2 && (fill[i] == 1 || fill[i] == 2 || fill[i] == 3) {
					ap.PutRune('_')
				} else if fill[i] == 1 {
					ap.PutRune(' ')
				} else if fill[i] == 0 {
					ap.PutRune('X')
				} else if fill[i] == 2 {
					if j == 1 && k == 3 {
						ap.PutRune('>')
					} else {
						ap.PutRune(' ')
					}
				} else if fill[i] == 3 {
					if j == 1 && k == 3 {
						ap.PutRune('@')
					} else {
						ap.PutRune(' ')
					}
				}
			}
		}
		ap.PutRune('|')
		if j != 3 {
			ap.PutRune(' ')
		}
	}
}

func drawMap(arr [][]int) {
	var hei, wei int = len(arr), len(arr[0])
	// upper bound
	var l int = wei*8 - 1
	if hei != 0 {
		ap.PutRune(' ')
		for i := 0; i < l; i++ {
			ap.PutRune('_')
		}
		ap.PutRune('\n')
	}

	// insides
	for i := 0; i < hei; i++ {
		// draw insides (_/X)
		drawRow(arr[i])
	}
}

func main() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w) // first line input

	if h == 0 || w == 0 {
		// ErrorMes()
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
				ErrorMes() // it will be changed to ap.PutRune()
			}
			for j := 0; j < w; j++ {
				ma[i][j] = int(nums[j] - 48)
				if ma[i][j] < 0 || ma[i][j] > 3 {
					// ErrorMes()
					return
				}
			}
		}

		drawMap(ma)

		ap.PutRune('\n')

	}
}
