package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"merge/entity"
)

func main() {
	basePath := flag.String("path", "", "metrics repository path")
	flag.Parse()

	metricsPath := fmt.Sprintf("%s/class.csv", *basePath)
	clocPath := fmt.Sprintf("%s/cloc.txt", *basePath)
	pathName := path.Base(*basePath)

	metric := &entity.Metrics{
		Repository: pathName,
	}

	err := metric.ParseLinesFromTxt(clocPath)
	if err != nil {
		log.Fatal(err)
	}

	err = metric.ParseMetricsFromCSV(metricsPath)
	if err != nil {
		log.Fatal(err)
	}

	err = metric.ToCSV(*basePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", metric)
}
