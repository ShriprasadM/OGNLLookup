package ognl1

import (
	"fmt"

	"github.com/prebid/openrtb/v20/openrtb2"
)

type BidResponseIndexInfo struct {
	SeatbidIndex int
	BidIndex     int
	ExtIndex     int
}

var bidResponseLookup = map[string]func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error){
	"bidresponse.id": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {

		return bidResponse.ID, nil
	},
	"bidresponse.seatbid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {

		return bidResponse.SeatBid, nil
	},
	"bidresponse.seatbid.bid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid, nil
	},
	"bidresponse.seatbid.bid.id": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].ID, nil
	},
	"bidresponse.seatbid.bid.impid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].ImpID, nil
	},
	"bidresponse.seatbid.bid.price": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Price, nil
	},
	"bidresponse.seatbid.bid.nurl": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].NURL, nil
	},
	"bidresponse.seatbid.bid.burl": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].BURL, nil
	},
	"bidresponse.seatbid.bid.lurl": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].LURL, nil
	},
	"bidresponse.seatbid.bid.adm": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].AdM, nil
	},
	"bidresponse.seatbid.bid.adid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].AdID, nil
	},
	"bidresponse.seatbid.bid.adomain": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].ADomain, nil
	},
	"bidresponse.seatbid.bid.bundle": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Bundle, nil
	},
	"bidresponse.seatbid.bid.iurl": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].IURL, nil
	},
	"bidresponse.seatbid.bid.cid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].CID, nil
	},
	"bidresponse.seatbid.bid.crid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].CrID, nil
	},
	"bidresponse.seatbid.bid.tactic": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Tactic, nil
	},
	"bidresponse.seatbid.bid.cattax": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].CatTax, nil
	},
	"bidresponse.seatbid.bid.cat": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Cat, nil
	},
	"bidresponse.seatbid.bid.attr": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Attr, nil
	},
	"bidresponse.seatbid.bid.apis": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].APIs, nil
	},
	"bidresponse.seatbid.bid.api": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].API, nil
	},
	"bidresponse.seatbid.bid.protocol": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Protocol, nil
	},
	"bidresponse.seatbid.bid.qagmediarating": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].QAGMediaRating, nil
	},
	"bidresponse.seatbid.bid.language": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Language, nil
	},
	"bidresponse.seatbid.bid.langb": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].LangB, nil
	},
	"bidresponse.seatbid.bid.dealid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].DealID, nil
	},
	"bidresponse.seatbid.bid.w": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].W, nil
	},
	"bidresponse.seatbid.bid.h": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].H, nil
	},
	"bidresponse.seatbid.bid.wratio": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].WRatio, nil
	},
	"bidresponse.seatbid.bid.hratio": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].HRatio, nil
	},
	"bidresponse.seatbid.bid.exp": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Exp, nil
	},
	"bidresponse.seatbid.bid.dur": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Dur, nil
	},
	"bidresponse.seatbid.bid.mtype": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].MType, nil
	},
	"bidresponse.seatbid.bid.slotinpod": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].SlotInPod, nil
	},
	"bidresponse.seatbid.bid.ext": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.BidIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Bid[indexInfo.BidIndex].Ext, nil
	},
	"bidresponse.seatbid.seat": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Seat, nil
	},
	"bidresponse.seatbid.group": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Group, nil
	},
	"bidresponse.seatbid.ext": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.SeatbidIndex >= len(bidResponse.SeatBid) || indexInfo.ExtIndex >= len(bidResponse.SeatBid[indexInfo.SeatbidIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.SeatBid[indexInfo.SeatbidIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidresponse.bidid": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {

		return bidResponse.BidID, nil
	},
	"bidresponse.cur": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {

		return bidResponse.Cur, nil
	},
	"bidresponse.customdata": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {

		return bidResponse.CustomData, nil
	},
	"bidresponse.nbr": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {

		return bidResponse.NBR, nil
	},
	"bidresponse.ext": func(bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo) (interface{}, error) {
		if indexInfo.ExtIndex >= len(bidResponse.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidResponse.Ext[indexInfo.ExtIndex], nil
	},
}

// LookUpBidResponse obtains value of type R using key from bidResponse. It uses indexIndo if required
// In case of failure returns error
func LookUpBidResponse[R any](key string, bidResponse openrtb2.BidResponse, indexInfo BidResponseIndexInfo, returnType R) (R, error) {
	fn := bidResponseLookup[key]
	if fn == nil {
		return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
	}
	result, err := fn(bidResponse, indexInfo)
	if err != nil {
		// Note: returntype here is not actual value
		return returnType, err
	}
	if rTypeValue, ok := result.(R); ok {
		return rTypeValue, nil
	}
	return returnType, fmt.Errorf("data type mismatch.Expected [%T] but found [%T]", returnType, result)
}
