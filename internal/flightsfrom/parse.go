package flightsfrom

import (
	"fmt"

	"github.com/hjson/hjson-go/v4"
)

func Parse(data []byte) (*FlightData, error) {
	flightData := &FlightData{}
	if err := hjson.Unmarshal(data, flightData); err != nil {
		return nil, fmt.Errorf("couldnt parse fight data: %w", err)
	}
	return flightData, nil
}
