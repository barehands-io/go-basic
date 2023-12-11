package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ReadXlsx() {

	f, err := excelize.OpenFile("./uploads/sponsor.xlsx")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Get all sheet names
	for _, sheetName := range f.GetSheetList() {
		fmt.Println("Sheet Name:", sheetName)
	}

	// Get all the rows in the sheet 1

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Println("Error reading rows:", err)
		return
	}

	fmt.Println("rows", rows)
}
