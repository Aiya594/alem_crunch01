package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func PrintStr(err string) {
	for i := 0; i < len(err); i++ {
		ap.PutRune(rune(err[i]))
	}
}

const (
	reset     = "\033[0m"
	whiteBG   = "\033[47m" // White background
	magentaBG = "\033[45m" // Magenta background (for walls)
	greenBG   = "\033[42m" // Green background (for player)
	yellowBG  = "\033[43m" // Yellow background (for awards)
)

func drawRow(fill []int) {
	var we int = len(fill)
	// draw blocks
	for j := 0; j < 3; j++ { // each block - 3 lines
		for i := 0; i < we; i++ {
			PrintStr("|")
			// filling 1 block with Space or Xs
			for k := 0; k < 7; k++ {
				if j == 2 && (fill[i] == 1 || fill[i] == 2 || fill[i] == 3) {
					PrintStr("_")
				} else if fill[i] == 1 {
					PrintStr(whiteBG + " " + reset)
				} else if fill[i] == 0 {
					PrintStr(magentaBG + "X" + reset)
				} else if fill[i] == 2 {
					if j == 1 && k == 3 {
						PrintStr(greenBG + ">" + reset) // Player symbol centered
					} else {
						PrintStr(greenBG + " " + reset) // Entire cell green
					}
				} else if fill[i] == 3 {
					if j == 1 && k == 3 {
						PrintStr(yellowBG + "@" + reset) // Award symbol centered
					} else {
						PrintStr(yellowBG + " " + reset) // Entire cell yellow
					}
				} // This closing brace was missing
			}
		}
		ap.PutRune('|')
		if j != 3 {
			ap.PutRune('\n')
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
	fmt.Scanf("%d %d/n", &h, &w) // first line input

	if h == 0 || w == 0 {
		PrintStr("Is zero") // error
	} else if h < 0 || w < 0 {
		PrintStr("Is negative")
	} else {
		// initializing 2D array
		ma := make([][]int, h)

		for i := range ma {
			ma[i] = make([]int, w)
		}

		// filling 2d array from terminal
		for i := 0; i < h; i++ {
			var nums string
			fmt.Scanf("%s/n", &nums) // to read each line as string
			if len(nums) != w {
				PrintStr("Too many")
				return
			}
			for j := 0; j < w; j++ {
				ma[i][j] = int(nums[j] - '0')
				if ma[i][j] < 0 || ma[i][j] > 3 {
					PrintStr("Not in range")
					return
				}
			}
		}

		drawMap(ma)

		ap.PutRune('\n')
	}
}
