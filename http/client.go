package httpClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func HTTPRequest(method, url string, body interface{}, params, headers map[string]string) ([]byte, error) {
	var requestBody []byte
	if body != nil {
		var err error
		requestBody, err = encodeRequestBody(body)
		if err != nil {
			return nil, err
		}
	}

	fullURL := url
	if params != nil {
		queryString := encodeQueryParams(params)
		fullURL += "?" + queryString
	}

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// log.Info().Msg("Response: " + string(responseBody))

	if resp.StatusCode != 200 {
		return nil, errors.New("error: status code not 200")
	}

	defer resp.Body.Close()

	return responseBody, nil
}

func encodeRequestBody(body interface{}) ([]byte, error) {
	return json.Marshal(body)
}

func encodeQueryParams(params map[string]string) string {
	encoded := url.Values{}
	for key, value := range params {
		encoded.Add(key, value)
	}
	return encoded.Encode()
}
