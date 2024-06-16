package main

import (
	"github.com/prebid/openrtb"
)

func generateBidResponse() string {
	return initF(openrtb.BidResponse{}, "bidResponse")
}
