package service

import (
	"gitsee/client"
	"log"
	"time"
)

var rateLimit struct {
	Cost      int
	Limit     int
	Remaining int
}

type RateLimit struct {
	Cost      int
	Limit     int
	Remaining int
}

var R RateLimit

func ResetEveryHour() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				InitRateLimit()
			case <-quit:
				ticker.Stop()
			}
		}
	}()
}

func InitRateLimit() {
	err := client.GHClient.Query(client.GHContext, &rateLimit, nil)
	if err != nil {
		log.Fatal(err)
	}

	r := &RateLimit{
		Cost:      rateLimit.Cost,
		Limit:     rateLimit.Limit,
		Remaining: rateLimit.Remaining,
	}

	R = *r
}
