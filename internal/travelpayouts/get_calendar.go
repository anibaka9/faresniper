package travelpayouts

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetUrl(origin string, destination string, month string, token string) string {
	return fmt.Sprintf("http://api.travelpayouts.com/v2/prices/month-matrix?currency=usd&origin=%s&destination=%s&month=%s&show_to_affiliates=false&token=%s", origin, destination, month, token)
}

func GetCalendar(origin string, destination string, month string) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("cant load env file: %w", err)
	}

	token := os.Getenv("TRAVELPAYOUTS_TOKEN")

	if token == "" {
		return nil, fmt.Errorf("travelpayouts token is empty: %w", err)
	}

	requestUrl := GetUrl(origin, destination, month, token)

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
