package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

var (
	sheetName = "wukong50k_release"
)

func GetExcelData(path string) [][]string {
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	rows := f.GetRows(sheetName)
	return rows
}
