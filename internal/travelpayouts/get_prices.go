package travelpayouts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const urlString = "http://api.travelpayouts.com/aviasales/v3/prices_for_dates"

func GetUrl(origin string, token string) (string, error) {
	reqUrl, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("wrong api url: %w", err)
	}
	params := reqUrl.Query()
	params.Set("origin", origin)
	params.Set("currency", "USD")
	params.Set("direct", "true")
	params.Set("limit", "1000")
	// params.Set("unique", "true")
	params.Set("token", token)
	reqUrl.RawQuery = params.Encode()
	return reqUrl.String(), nil
}

func GetPrices(origin string) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("cant load env file: %w", err)
	}

	token := os.Getenv("TRAVELPAYOUTS_TOKEN")

	if token == "" {
		return nil, fmt.Errorf("travelpayouts token is empty: %w", err)
	}

	requestUrl, err := GetUrl(origin, token)
	if err != nil {
		return nil, fmt.Errorf("error getting request url: %w", err)
	}

	fmt.Printf("getting data from %s\n", requestUrl)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("cant make request: %w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant get response data: %w", err)
	}
	return body, nil
}
