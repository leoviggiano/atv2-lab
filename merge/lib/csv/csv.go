package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"merge/entity"
)

func SaveMetrics(metrics []*entity.Metrics) error {
	if len(metrics) == 0 {
		return nil
	}

	file, err := os.Create("etc/metrics.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	headers := metrics[0].CsvHeader()
	values := make([][]string, 0, len(metrics))

	for _, r := range metrics {
		if r == nil {
			continue
		}

		values = append(values, r.CsvValue())
	}

	if err := w.Write(headers); err != nil {
		return err
	}

	if err := w.WriteAll(values); err != nil {
		return err
	}

	fmt.Printf("saved %d rows on csv with success\n", len(metrics))
	return nil
}

func SaveReports(reports []*entity.Report) error {
	if len(reports) == 0 {
		return nil
	}

	file, err := os.Create("etc/reports.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)

	headers := reports[0].CsvHeader()
	values := make([][]string, 0, len(reports))

	for _, r := range reports {
		values = append(values, r.CsvValue())
	}

	if err := w.Write(headers); err != nil {
		return err
	}

	if err := w.WriteAll(values); err != nil {
		return err
	}

	fmt.Printf("saved %d rows on csv with success\n", len(reports))
	return nil
}

func ReadRepositories(filePath string) ([]*entity.Repository, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	repositories := make([]*entity.Repository, 0)

	for _, row := range records[1:] {
		repository := &entity.Repository{}
		repository.FillFromCSV(row)

		repositories = append(repositories, repository)
	}

	return repositories, nil
}

func ReadMetrics(filePath string) ([]*entity.Metrics, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	metrics := make([]*entity.Metrics, 0)

	for _, row := range records[1:] {
		metric := &entity.Metrics{}
		metric.FillFromCSV(row)

		metrics = append(metrics, metric)
	}

	return metrics, nil
}
