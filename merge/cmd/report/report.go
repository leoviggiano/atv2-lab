package main

import (
	"fmt"
	"log"
	"path"
	"strings"

	"merge/entity"
	"merge/lib/csv"
)

func main() {
	repositories, err := csv.ReadRepositories("./etc/repositories.csv")
	if err != nil {
		log.Fatal(err)
	}

	metrics, err := csv.ReadMetrics("./etc/metrics.csv")
	if err != nil {
		log.Fatal(err)
	}

	mapMetrics := make(map[string]*entity.Metrics, len(metrics))
	for _, v := range metrics {
		mapMetrics[strings.ToLower(v.Repository)] = v
	}

	reports := make([]*entity.Report, 0, len(repositories))
	for _, r := range repositories {
		if metrics, ok := mapMetrics[strings.ToLower(path.Base(r.Url))]; ok {
			reports = append(reports, &entity.Report{
				Repository: r,
				Metric:     metrics,
			})
		} else {
			fmt.Println(r.Name)
		}
	}

	err = csv.SaveReports(reports)
	if err != nil {
		log.Fatal(err)
	}
}
