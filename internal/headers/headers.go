package headers

import (
	"encoding/json"
	"net/http"
)

func ProcessHeaders(r *http.Request) ([]byte, error) {
	headers := r.Header

	headerMap := make(map[string][]string)
	for key, values := range headers {
		headerMap[key] = values
	}

	return json.Marshal(headerMap)
}