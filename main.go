package main

import (
	"fmt"

	"github.com/prebid/openrtb"
)

func main() {
	// fmt.Println(generateBidResponse())

	bidResponse := openrtb.BidResponse{
		BidID: "233e",
		SeatBid: []openrtb.SeatBid{
			{Bid: []openrtb.Bid{{AdM: "native"}}},
			{Bid: []openrtb.Bid{{AdM: "video"}, {AdM: "banner"}}},
		},
	}

	// adm, _ := lookup("bidresponse.bidid", bidResponse, 0, "")
	// fmt.Println(adm)

	bid, err := lookup("bidresponse.seatbid.bid.adm", bidResponse, 0, "")
	if err == nil {
		fmt.Println(bid)
	} else {
		fmt.Println(err)
	}
}

func lookup[R any](key string, bidResponse openrtb.BidResponse, i int, returnType R) (R, error) {
	m := map[string]func(bidResponse openrtb.BidResponse, i int) interface{}{
		"bidresponse.id": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.ID
		},
		"bidresponse.seatbid.bid.id": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].ID
		},
		"bidresponse.seatbid.bid.impid": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].ImpID
		},
		"bidresponse.seatbid.bid.price": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].Price
		},
		"bidresponse.seatbid.bid.adid": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].AdID
		},
		"bidresponse.seatbid.bid.nurl": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].NURL
		},
		"bidresponse.seatbid.bid.adm": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].AdM
		},
		"bidresponse.seatbid.bid.bundle": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].Bundle
		},
		"bidresponse.seatbid.bid.iurl": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].IURL
		},
		"bidresponse.seatbid.bid.cid": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].CID
		},
		"bidresponse.seatbid.bid.crid": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].CrID
		},
		"bidresponse.seatbid.bid.dealid": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].DealID
		},
		"bidresponse.seatbid.bid.h": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].H
		},
		"bidresponse.seatbid.bid.w": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Bid[i].W
		},
		"bidresponse.seatbid.seat": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Seat
		},
		"bidresponse.seatbid.group": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.SeatBid[i].Group
		},
		"bidresponse.bidid": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.BidID
		},
		"bidresponse.cur": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.Cur
		},
		"bidresponse.customdata": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.CustomData
		},
		"bidresponse.nbr": func(bidResponse openrtb.BidResponse, i int) interface{} {

			return bidResponse.NBR
		},
	}
	fn := m[key]
	if fn == nil {
		return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
	}
	result := fn(bidResponse, i)
	if rTypeValue, ok := result.(R); ok {
		return rTypeValue, nil
	}
	return returnType, fmt.Errorf("data type mismatch.Expected [%T] but found [%T]", returnType, result)
}
