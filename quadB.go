package piscine

import (
	"fmt"
	"strings"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	
	if x >= 2 && y >= 2 {
		var xSidesToPrint string = strings.Repeat("*", x-2)
		var ySidesToPrint string = strings.Repeat(" ", x-2)
		fmt.Println("/" + xSidesToPrint + "\\")
		for i := y - 2; i > 0; i-- {
			fmt.Println("*" + ySidesToPrint + "*")
		}
		fmt.Println("\\" + xSidesToPrint + "/")
	} else if x == 1 && y >= 2 {
		fmt.Println("/")
		for i := y - 2; i > 0; i-- {
			fmt.Println("*")
		}
		fmt.Println("\\")
	} else if y == 1 && x >= 2 {
		var xSidesToPrint string = strings.Repeat("*", x-2)
		fmt.Println("/" + xSidesToPrint + "\\")
	} else if y == 1 && x == 1 {
		fmt.Println("/")
	}
}
