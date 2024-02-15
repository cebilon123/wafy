package response

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ValidateStatusCode(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf(
			"status code wasn't positive, was '%d', error reading response body: %w",
			res.StatusCode,
			err,
		)
	}

	return errors.New(fmt.Sprintf("status code wasn't positive, was '%d', body: '%d'", res.StatusCode, body))
}
