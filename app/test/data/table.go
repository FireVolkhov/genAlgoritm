package data

import (
	"io/ioutil"
	"strings"
	"../../../app/core"
)

const indexDate = 0;
const indexOpen = 1;
const indexHigh = 2;
const indexLow = 3;
const indexClose = 4;
const indexVolume = 5;
const indexAdjClose = 6;

type TableRow struct {
	Open float64
	High float64
	Low float64
	Close float64
	Volume int
	AdjClose float64
	Dif float64
}

var table []*TableRow

func GetTableLen() int {
	return len(table)
}

func init() {
	file, _ := ioutil.ReadFile("./app/test/data/table.csv")
	dataAsString := string(file)
	lines := strings.Split(dataAsString, "\n")

	// Убираем заголовки таблицы
	// Убираем последнюю пустую строку
	lines = lines[1:len(lines) - 1]
	table = make([]*TableRow, len(lines))

	for rowIndex, line := range lines {
		lineArray := strings.Split(line, ",")

		table[rowIndex] = &TableRow{
			Open: core.ToFloat64(lineArray[indexOpen]),
			High: core.ToFloat64(lineArray[indexHigh]),
			Low: core.ToFloat64(lineArray[indexLow]),
			Close: core.ToFloat64(lineArray[indexClose]),
			Volume: core.StringToInt(lineArray[indexVolume]),
			AdjClose: core.ToFloat64(lineArray[indexAdjClose]),
			Dif: 0,
		}
	}

	// Reverse array
	last := len(table) - 1
	for i := 0; i < len(table)/2; i++ {
		table[i], table[last - i] = table[last - i], table[i]
	}

	for rowIndex := range table {
		var dif float64

		if (rowIndex == 0) {
			dif = core.ToFloat64(0);
		} else {
			prevClose := table[rowIndex - 1].Close
			dif = table[rowIndex].Close - prevClose
		}

		table[rowIndex].Dif = dif
	}

	// Убираем первую строку с 0
//	table = table[1:]
	dataLen := core.RoundToInt(0.5 * float64(len(table)))
	table = table[dataLen:]
}

type TableIterator struct {
	step int
	position int
	IsFinished bool
}

func NewTableIterator(step int) *TableIterator {
	return &TableIterator{
		step: step,
		position: 0,
		IsFinished: false,
	}
}

func (this *TableIterator) Next() []*TableRow {
	if (!this.IsFinished) {
		position := this.position
		end := position + this.step
		result := table[position:end]

		this.position++
		this.IsFinished = (len(table) < (end + 1))
		return result

	} else {
		return nil
	}
}

func (this *TableIterator) GetLen () int {
	return len(table) - (this.step - 1)
}
