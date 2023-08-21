package piscine

import ("fmt"
		"strings"
	)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	if x >= 2 && y >= 2 {
		var xSidesToPrint string = strings.Repeat("-", x-2)
		var ySidesToPrint string = strings.Repeat(" ", x-2)
		fmt.Println("o" + xSidesToPrint + "o")
		for i := y - 2; i > 0; i-- {
			fmt.Println("|" + ySidesToPrint + "|")
		}
		fmt.Println("o" + xSidesToPrint + "o")
	} else if x == 1 && y >= 2 {
		fmt.Println("o")
		for i := y - 2; i > 0; i-- {
			fmt.Println("|")
		}
		fmt.Println("o")
	} else if y == 1 && x >= 2 {
		var xSidesToPrint string = strings.Repeat("-", x-2)
		fmt.Println("o" + xSidesToPrint + "o")
	} else if y == 1 && x == 1 {
		fmt.Println("o")
	}
}