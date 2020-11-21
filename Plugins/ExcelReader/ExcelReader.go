package ExcelReader

import (
	"github.com/tealeg/xlsx"
)

func Excel_reader(file, sheet string) []*xlsx.Row {

	xlfile, err := xlsx.OpenFile(file)

	if err != nil {
		panic(err.Error())
	}

	return xlfile.Sheet[sheet].Rows
}
