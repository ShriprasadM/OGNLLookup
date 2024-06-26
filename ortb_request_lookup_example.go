package main

import (
	"OGNLLookup/ognl1"
	"fmt"

	openrtb "github.com/prebid/openrtb/v20/openrtb2"
)

func generateBidRequestOGNLMap() {
	initF(openrtb.BidRequest{}, "bidRequest", "./ognl1/bid_request_ognl.go")

}

var bidRequest = openrtb.BidRequest{
	Imp: []openrtb.Imp{
		{
			Video: &openrtb.Video{
				MinDuration: 15,
			},
		},
		{
			Banner: &openrtb.Banner{
				W: new(int64),
			},
		},
		{
			Video: &openrtb.Video{
				MinDuration: 1500,
			},
		},
	},
}

func videoMinDuarationLookup() {
	width := new(int64)
	*width = 300
	bidRequest.Imp[1].Banner.W = width

	videoMinDuration, err := ognl1.LookUpBidRequest("bidrequest.imp.video.minduration", bidRequest, ognl1.BidRequestIndexInfo{
		ImpIndex: 1,
	}, int64(0))
	if err == nil {
		fmt.Println(videoMinDuration)
	} else {
		fmt.Println(err)
	}
}
