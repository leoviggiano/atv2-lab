package entity

type Report struct {
	Metric     *Metrics
	Repository *Repository
}

func (r *Report) CsvHeader() []string {
	metricsHeader := r.Metric.CsvHeader()
	repositoriesHeader := r.Repository.CsvHeader()
	return append(metricsHeader, repositoriesHeader...)
}

func (r *Report) CsvValue() []string {
	metricsValue := r.Metric.CsvValue()
	repositoriesValue := r.Repository.CsvValue()
	return append(metricsValue, repositoriesValue...)
}
