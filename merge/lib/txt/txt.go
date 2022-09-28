package txt

import (
	"fmt"
	"os"
)

func SaveErrors(errors []string) error {
	file, err := os.Create("etc/errors.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	errs := ""
	for _, err := range errors {
		errs += fmt.Sprintf("%s\n", err)
	}

	_, err = file.WriteString(errs)
	if err != nil {
		return err
	}

	return nil
}
