package main

func main() {
	// generateBidResponseOGNLMap()
	admLookup()
	seatLookup()
}

// func main() {
// 	// fmt.Println(generateBidResponse())

// 	bidResponse := openrtb.BidResponse{
// 		BidID: "233e",
// 		SeatBid: []openrtb.SeatBid{
// 			{Bid: []openrtb.Bid{{AdM: "native"}}, Seat: "pubmatic"},
// 			{Bid: []openrtb.Bid{{AdM: "video"}, {AdM: "banner"}}},
// 		},
// 	}

// 	bid, err := lookup("bidresponse.seatbid.bid.adm", bidResponse, IndexInfo{
// 		seatbidIndex: 1,
// 		bidIndex:     1,
// 	}, "")

// 	if err == nil {
// 		fmt.Println(bid)
// 	} else {
// 		fmt.Println(err)
// 	}

// 	seat, err1 := lookup("bidresponse.seatbid.seat", bidResponse, IndexInfo{}, "")
// 	if err1 == nil {
// 		fmt.Println(seat)
// 	} else {
// 		fmt.Println(err1)
// 	}

// }

// type IndexInfo struct {
// 	seatbidIndex int
// 	bidIndex     int
// }

// func lookup[R any](key string, bidResponse openrtb.BidResponse, indexInfo IndexInfo, returnType R) (R, error) {

// 	m := map[string]func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{}{
// 		"bidresponse.id": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.ID
// 		},
// 		"bidresponse.seatbid.bid.id": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].ID
// 		},
// 		"bidresponse.seatbid.bid.impid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].ImpID
// 		},
// 		"bidresponse.seatbid.bid.price": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].Price
// 		},
// 		"bidresponse.seatbid.bid.adid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].AdID
// 		},
// 		"bidresponse.seatbid.bid.nurl": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].NURL
// 		},
// 		"bidresponse.seatbid.bid.adm": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].AdM
// 		},
// 		"bidresponse.seatbid.bid.bundle": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].Bundle
// 		},
// 		"bidresponse.seatbid.bid.iurl": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].IURL
// 		},
// 		"bidresponse.seatbid.bid.cid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].CID
// 		},
// 		"bidresponse.seatbid.bid.crid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].CrID
// 		},
// 		"bidresponse.seatbid.bid.dealid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].DealID
// 		},
// 		"bidresponse.seatbid.bid.h": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].H
// 		},
// 		"bidresponse.seatbid.bid.w": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Bid[indexInfo.bidIndex].W
// 		},
// 		"bidresponse.seatbid.seat": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Seat
// 		},
// 		"bidresponse.seatbid.group": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.SeatBid[indexInfo.seatbidIndex].Group
// 		},
// 		"bidresponse.bidid": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.BidID
// 		},
// 		"bidresponse.cur": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.Cur
// 		},
// 		"bidresponse.customdata": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.CustomData
// 		},
// 		"bidresponse.nbr": func(bidResponse openrtb.BidResponse, indexInfo IndexInfo) interface{} {

// 			return bidResponse.NBR
// 		},
// 	}

// 	fn := m[key]
// 	if fn == nil {
// 		return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
// 	}
// 	result := fn(bidResponse, indexInfo)
// 	if rTypeValue, ok := result.(R); ok {
// 		return rTypeValue, nil
// 	}
// 	return returnType, fmt.Errorf("data type mismatch.Expected [%T] but found [%T]", returnType, result)
// }
