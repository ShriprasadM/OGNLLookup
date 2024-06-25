package main

import (
	// mognl "OGNLLookup/ognl/bidresponse"
	mognl "OGNLLookup/ognl1"
	"fmt"

	"github.com/prebid/openrtb"
)

func generateBidResponseOGNLMap() {
	initF(openrtb.BidResponse{}, "bidResponse", "./ognl1/bid_response_ognl.go")

}

var bidResponse = openrtb.BidResponse{
	BidID: "233e",
	SeatBid: []openrtb.SeatBid{
		{Bid: []openrtb.Bid{{AdM: "native"}}, Seat: "pubmatic"},
		{Bid: []openrtb.Bid{{AdM: "video"}, {AdM: "<VAST></VAST>"}, {AdM: "<VAST1></VAST1>"}}},
	},
}

func test() {
	bid, err := mognl.LookUpBidResponse("bidresponse.seatbid.bid", openrtb.BidResponse{}, mognl.IndexInfo{}, openrtb.Bid{})
	if err == nil {
		fmt.Println(bid)
	} else {
		fmt.Println(err)
	}
}

func admLookup() {
	bid, err := mognl.LookUpBidResponse("bidresponse.seatbid.bid.adm", bidResponse, mognl.IndexInfo{
		SeatbidIndex: 1,
		BidIndex:     44,
	}, "")

	if err == nil {
		fmt.Println(bid)
	} else {
		fmt.Println(err)
	}
}

func seatLookup() {
	seat, err1 := mognl.LookUpBidResponse("bidresponse.seatbid.seat", bidResponse, mognl.IndexInfo{}, "")
	if err1 == nil {
		fmt.Println(seat)
	} else {
		fmt.Println(err1)
	}
}
