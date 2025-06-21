import (
	//"github.com/alem-platform/ap"
	"fmt"
)

func main() {
	var h, w int
	fmt.Scanf("%d %d", h, w) // first line input

	// creating 2D array
	ma := make([][]int, h)
	for i := range ma {
		ma[i] = make([]int, w)
	}

	for i := 0; i < h; i++ { // second input
		fmt.Scanf()
	}
	fmt.Print(ma)
}