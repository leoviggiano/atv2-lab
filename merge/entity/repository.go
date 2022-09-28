package entity

import (
	"strconv"
	"time"
)

const NoRelease = "No Releases"

type PrimaryLanguage struct {
	Name string
}

type Releases struct {
	TotalCount    int
	LatestRelease time.Time
	Nodes         []*Release
}

type Repository struct {
	ID   string
	Name string
	Url  string

	CreatedAt time.Time
	UpdatedAt time.Time

	PrimaryLanguage PrimaryLanguage
	Releases        Releases

	Stargazers struct {
		TotalCount int
	}

	PullRequests struct {
		TotalCount int
	}

	Issues struct {
		TotalCount int
		Closed     int
		Open       int
	}
}

type Release struct {
	CreatedAt time.Time
}

func (r *Repository) CsvHeader() []string {
	return []string{"CreatedAt", "UpdatedAt", "PrimaryLanguage", "Releases", "PullRequests", "Open Issues", "Closed Issues", "LastRelease", "Stars"}
}

func (r *Repository) CsvValue() []string {
	lastRelease := NoRelease
	if len(r.Releases.Nodes) > 0 {
		lastRelease = r.Releases.Nodes[0].CreatedAt.Format(time.RFC3339)
	}

	return []string{
		r.CreatedAt.String(),
		r.UpdatedAt.String(),
		r.PrimaryLanguage.Name,
		strconv.Itoa(r.Releases.TotalCount),
		strconv.Itoa(r.PullRequests.TotalCount),
		strconv.Itoa(r.Issues.Open),
		strconv.Itoa(r.Issues.Closed),
		lastRelease,
		strconv.Itoa(r.Stargazers.TotalCount),
	}
}

func (r *Repository) FillFromCSV(row []string) {
	r.Name = row[0]
	r.CreatedAt, _ = time.Parse("2006-01-02 15:04:05 -0700 MST", row[1])
	r.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05 -0700 MST", row[2])
	r.PrimaryLanguage.Name = row[3]
	r.Releases.TotalCount, _ = strconv.Atoi(row[4])
	r.PullRequests.TotalCount, _ = strconv.Atoi(row[5])
	r.Issues.Open, _ = strconv.Atoi(row[6])
	r.Issues.Closed, _ = strconv.Atoi(row[7])
	r.Issues.TotalCount = r.Issues.Open + r.Issues.Closed
	r.Stargazers.TotalCount, _ = strconv.Atoi(row[9])
	r.Url = row[10]

	if row[8] != NoRelease {
		date, _ := time.Parse(time.RFC3339, row[8])
		r.Releases.LatestRelease = date
	}
}
