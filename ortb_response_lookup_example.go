package main

import (
	"encoding/json"

	// mognl "OGNLLookup/ognl/bidresponse"
	mognl "OGNLLookup/ognl1"
	"fmt"

	openrtb "github.com/prebid/openrtb/v20/openrtb2"
)

func generateBidResponseOGNLMap() {
	initF(openrtb.BidResponse{}, "bidResponse", "./ognl1/bid_response_ognl.go")

}

/*
	unmarshal

bidresponse ------> Map -> jsonlib -> lookup
bidresponse ------> lookup
*/
func bidExtLookup() {

	bidExt, err := mognl.LookUpBidResponse("bidresponse.seatbid.bid.mtype", openrtb.BidResponse{
		SeatBid: []openrtb.SeatBid{
			{
				Bid: []openrtb.Bid{
					{
						Ext: []byte(`this is bid.ext contains`),
					},
				},
			},
		},
	}, mognl.BidResponseIndexInfo{}, json.RawMessage{})

	if err == nil {
		fmt.Println(string(bidExt))
	} else {
		fmt.Println(err)
	}
}

var bidResponse = openrtb.BidResponse{
	BidID: "233e",
	SeatBid: []openrtb.SeatBid{
		{Bid: []openrtb.Bid{{AdM: "native"}}, Seat: "pubmatic"},
		{Bid: []openrtb.Bid{{AdM: "video"}, {AdM: "<VAST></VAST>"}, {AdM: "<VAST1></VAST1>"}}},
	},
}

func bidLookup() {

	bid, err := mognl.LookUpBidResponse("bidresponse.seatbid.bid", bidResponse, mognl.BidResponseIndexInfo{}, []openrtb.Bid{})
	if err == nil {
		fmt.Println(bid)
	} else {
		fmt.Println(err)
	}
}

func admLookup() {
	adm, err := mognl.LookUpBidResponse("bidresponse.seatbid.bid.adm", bidResponse, mognl.BidResponseIndexInfo{
		SeatbidIndex: 1,
		BidIndex:     2,
	}, "")

	if err == nil {
		fmt.Println(adm)
	} else {
		fmt.Println(err)
	}
}

func seatLookup() {
	seat, err1 := mognl.LookUpBidResponse("bidresponse.seatbid.seat", bidResponse, mognl.BidResponseIndexInfo{}, "")
	if err1 == nil {
		fmt.Println(seat)
	} else {
		fmt.Println(err1)
	}
}
