package txt

import (
	"blaus/pkg/entity"
	"fmt"
	"os"
)

func Save(repositories []*entity.Repository) error {
	file, err := os.Create("etc/clone.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	urls := ""
	for _, r := range repositories {
		urls += fmt.Sprintf("%s\n", r.Url)
	}

	_, err = file.WriteString(urls)
	if err != nil {
		return err
	}

	return nil
}
