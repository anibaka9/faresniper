package flightsfrom

import (
	"errors"
	"regexp"
)

func Exctact(htmlData []byte) ([]byte, error) {
	re := regexp.MustCompile(`(?s)data\(\)\s*\{\s*return\s*(\{.*?\});?\s*\},?\s*mounted\(\)`)
	matches := re.FindSubmatch(htmlData)
	if len(matches) != 2 {
		return nil, errors.New("no hjson data found in data")
	}
	rawData := matches[1]
	badFieldRegex := regexp.MustCompile(`(?s)iframeUrl:(.*?),`)
	rawCorrectData := badFieldRegex.ReplaceAll(rawData, []byte(""))
	return rawCorrectData, nil
}
