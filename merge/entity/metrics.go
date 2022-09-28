package entity

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Metrics struct {
	Repository   string
	CBO          int
	DIT          int
	LCOM         int
	CodeLines    int
	CommentLines int
}

func (m *Metrics) CsvHeader() []string {
	return []string{
		"Repository",
		"CBO",
		"DIT",
		"LCOM",
		"Code Lines",
		"Comment Lines",
	}
}

func (m *Metrics) CsvValue() []string {
	return []string{
		m.Repository,
		strconv.Itoa(m.CBO),
		strconv.Itoa(m.DIT),
		strconv.Itoa(m.LCOM),
		strconv.Itoa(m.CodeLines),
		strconv.Itoa(m.CommentLines),
	}
}

func (m *Metrics) ParseMetricsFromCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csv, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range csv[1:] {
		cbo, _ := strconv.Atoi(row[3])
		dit, _ := strconv.Atoi(row[8])
		lcom, _ := strconv.Atoi(row[11])

		m.CBO += cbo
		m.DIT += dit
		m.LCOM += lcom
	}

	m.CBO /= (len(csv) - 1)
	m.DIT /= (len(csv) - 1)
	m.LCOM /= (len(csv) - 1)

	return nil
}

func (m *Metrics) ParseLinesFromTxt(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	values := strings.Split(string(content), "\n")
	m.CommentLines, _ = strconv.Atoi(values[0])
	m.CodeLines, _ = strconv.Atoi(values[1])
	return nil
}

func (m *Metrics) FillFromMetricsCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csv, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range csv[1:] {
		m.FillFromCSV(row)
	}

	return nil
}

func (m *Metrics) FillFromCSV(row []string) {
	cbo, _ := strconv.Atoi(row[1])
	dit, _ := strconv.Atoi(row[2])
	lcom, _ := strconv.Atoi(row[3])
	codeLines, _ := strconv.Atoi(row[4])
	commentLines, _ := strconv.Atoi(row[5])

	m.Repository = row[0]
	m.CBO += cbo
	m.DIT += dit
	m.LCOM += lcom
	m.CodeLines += codeLines
	m.CommentLines += commentLines
}

func (m *Metrics) ToCSV(path string) error {
	path = fmt.Sprintf("%s/metrics.csv", path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	fmt.Printf("metrics save with success in %s\n", path)
	return w.WriteAll([][]string{m.CsvHeader(), m.CsvValue()})
}
