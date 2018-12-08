package main

import (
	"log"
	"testing"

	"github.com/tealeg/xlsx"
)

func Test_A(t *testing.T) {
	RootPath := "/Users/yanrui/project/PM/运营分析/"
	ExcelFileName := RootPath + "2018年项目验收计划副本.xlsx"
	log.Println("ExcelFileName:%s", ExcelFileName)
	xlFile, err := xlsx.OpenFile(ExcelFileName)
	if err != nil {
		log.Println("-------Error:Read %s", ExcelFileName)
		return
	}

	row_num := 0

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if row_num < 5 {
				row_data := ""
				for _, cell := range row.Cells {
					text := cell.String()
					row_data += text + ","
				}
				row_data += ";"
				log.Printf("row_data:%s \n", row_data)

			}
			row_num++
		}
	}
}
