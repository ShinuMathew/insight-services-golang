package CSVHandler

import (
	"bufio"
	"encoding/csv"
	"os"
)

func ReadCSV(fileLocation string) *csv.Reader {

	csvFile, _ := os.Open(fileLocation)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	return reader
}
