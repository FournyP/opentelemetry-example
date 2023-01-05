package utils

import "fmt"

func CombineErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}

	var combinedError string
	for _, err := range errors {
		combinedError += err.Error() + "\n"
	}

	return fmt.Errorf(combinedError)
}
