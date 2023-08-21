package piscine

import ("fmt"
	"strings"
)

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	if x >= 2 && y >= 2 {
		var xSidesToPrint string = strings.Repeat("B", x-2)
		var ySidesToPrint string = strings.Repeat(" ", x-2)
		fmt.Println("A" + xSidesToPrint + "C")
		for i := y - 2; i > 0; i-- {
			fmt.Println("B" + ySidesToPrint + "B")
		}
		fmt.Println("C" + xSidesToPrint + "A")
	} else if x == 1 && y >= 2 {
		fmt.Println("A")
		for i := y - 2; i > 0; i-- {
			fmt.Println("B")
		}
		fmt.Println("C")
	} else if y == 1 && x >= 2 {
		var xSidesToPrint string = strings.Repeat("B", x-2)
		fmt.Println("A" + xSidesToPrint + "C")
	} else if y == 1 && x == 1 {
		fmt.Println("A")
	}
}