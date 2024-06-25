package bidresponse

import (
	"fmt"

	"github.com/prebid/openrtb"
)

type IndexInfo struct {
	SeatbidIndex int
	BidIndex     int
}

var bidResponseLookup = map[string]func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{}{
	"bidresponse.id": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.ID
	},
	"bidresponse.seatbid.bid.id": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].ID
	},
	"bidresponse.seatbid.bid.impid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].ImpID
	},
	"bidresponse.seatbid.bid.price": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Price
	},
	"bidresponse.seatbid.bid.adid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].AdID
	},
	"bidresponse.seatbid.bid.nurl": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].NURL
	},
	"bidresponse.seatbid.bid.adm": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].AdM
	},
	"bidresponse.seatbid.bid.bundle": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Bundle
	},
	"bidresponse.seatbid.bid.iurl": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].IURL
	},
	"bidresponse.seatbid.bid.cid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].CID
	},
	"bidresponse.seatbid.bid.crid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].CrID
	},
	"bidresponse.seatbid.bid.dealid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].DealID
	},
	"bidresponse.seatbid.bid.h": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].H
	},
	"bidresponse.seatbid.bid.w": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].W
	},
	"bidresponse.seatbid.seat": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Seat
	},
	"bidresponse.seatbid.group": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Group
	},
	"bidresponse.bidid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.BidID
	},
	"bidresponse.cur": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.Cur
	},
	"bidresponse.customdata": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.CustomData
	},
	"bidresponse.nbr": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

		return bidResponse.NBR
	},
}

func LookUpBidResponse[R any](key string, bidResponse openrtb.BidResponse, indexInfo IndexInfo, returnType R) (R, error) {
	fn := bidResponseLookup[key]
	if fn == nil {
		return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
	}
	result := fn(bidResponse, indexInfo)
	if rTypeValue, ok := result.(R); ok {
		return rTypeValue, nil
	}
	return returnType, fmt.Errorf("data type mismatch.Expected [%T] but found [%T]", returnType, result)
}
