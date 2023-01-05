package usecases

import (
	"fmt"
	"net/http"

	"test/utils"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type PingMessage struct {
	Websites []string `query:"websites"`
}

type PingResponse struct {
	Website string
	Result  string
}

type PingUseCase struct {
	httpClient *http.Client
}

func NewPingUseCase(httpClient *http.Client) *PingUseCase {
	return &PingUseCase{
		httpClient: httpClient,
	}
}

func (pingUseCase *PingUseCase) Execute(message *PingMessage) ([]interface{}, error) {
	errors := make(chan error)
	defer close(errors)

	results := make(chan interface{})
	defer close(results)

	pingUseCase.httpClient.Transport = otelhttp.NewTransport(http.DefaultTransport)

	for _, website := range message.Websites {
		go func(url string) {
			response, err := pingUseCase.httpClient.Get(url)

			if err != nil {
				errors <- fmt.Errorf("error requesting %v", url)
			}

			results <- PingResponse{
				Website: url,
				Result:  response.Status,
			}
		}(website)
	}

	responses, err := utils.HandleRoutines(results, errors, len(message.Websites))

	if err != nil {
		return nil, err
	}

	return responses, nil
}
