package utils

func HandleRoutines(results chan interface{}, errors chan error, waitedResponse int) ([]interface{}, error) {
	var responses []interface{}
	var errs []error

	for {
		select {
		case result := <-results:
			responses = append(responses, result)
		case err := <-errors:
			errs = append(errs, err)
		}

		if len(responses)+len(errs) == waitedResponse {
			break
		}
	}

	return responses, CombineErrors(errs)
}
