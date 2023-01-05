package src

import (
	"fmt"
	"strings"
)

func Calculate(direction string, pagesPerSheet int, numberOfPages int, printerType string) string {
	var pagesPerPage int = pagesPerSheet * 2
	var pageRemainder int = numberOfPages % pagesPerPage
	var singleSide1 string
	var singleSide2 string
	var doubleSide string
	var result string

	if pageRemainder == 0 {
		for i := 1; i < (numberOfPages / 2); i += 2 {
			var singleSide1Temp string
			var singleSide2Temp string

			if direction == "LTR" {
				singleSide1Temp = fmt.Sprintf("%d,%d,", (numberOfPages-i)+1, i)
				singleSide2Temp = fmt.Sprintf("%d,%d,", i+1, numberOfPages-i)
			} else {
				singleSide1Temp = fmt.Sprintf("%d,%d,", i, (numberOfPages-i)+1)
				singleSide2Temp = fmt.Sprintf("%d,%d,", numberOfPages-i, i+1)
			}

			singleSide1 += singleSide1Temp
			singleSide2 += singleSide2Temp
			doubleSide += singleSide1Temp + singleSide2Temp
		}

		singleSide1 = strings.TrimSuffix(singleSide1, ",")
		singleSide2 = strings.TrimSuffix(singleSide2, ",")
		doubleSide = strings.TrimSuffix(doubleSide, ",")

		if printerType == "SS" {
			result = fmt.Sprintf(`Set your printer's "pages per side" setting to %d.
				
Print these pages first: %s
				
Flip the printed pages and put them back into your printer.
				
Now print these pages: %s

`, pagesPerSheet, singleSide1, singleSide2)
		} else {
			result = fmt.Sprintf(`Set your printer's "pages per side" setting to %d.

Set your printer's "2-sided" (duplex) setting to either short or long edge, or similar option, depending on page orientation.
				
Print pages in this order:  %s

`, pagesPerSheet, doubleSide)
		}
	} else {
		result = fmt.Sprintf("Not divisible by %d. Need %d more page(s).\n", pagesPerPage, (pagesPerPage - pageRemainder))
	}

	return result
}
