package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricsService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	//metrics storage to prometheus
	fmt.Println("pushing metrics prometheus")
	return s.next.FetchPrice(ctx, ticker)
}
