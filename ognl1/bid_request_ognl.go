package ognl1

import (
	"fmt"

	"github.com/prebid/openrtb/v20/openrtb2"
)

type BidRequestIndexInfo struct {
	ImpIndex           int
	MetricIndex        int
	FormatIndex        int
	ExtIndex           int
	MimesIndex         int
	BattrIndex         int
	CompanionadIndex   int
	BtypeIndex         int
	ExpdirIndex        int
	ApiIndex           int
	DurfloorsIndex     int
	ProtocolsIndex     int
	RqddursIndex       int
	DeliveryIndex      int
	CompaniontypeIndex int
	DealsIndex         int
	RefsettingsIndex   int
	CatIndex           int
	DataIndex          int
	SegmentIndex       int
	KwarrayIndex       int
	SectioncatIndex    int
	PagecatIndex       int
	BrowsersIndex      int
	VersionIndex       int
	EidsIndex          int
	UidsIndex          int
	WseatIndex         int
	NodesIndex         int
}

var bidRequestLookup = map[string]func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error){
	"bidrequest.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.ID, nil
	},
	"bidrequest.imp": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Imp, nil
	},
	"bidrequest.imp.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].ID, nil
	},
	"bidrequest.imp.metric": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Metric, nil
	},
	"bidrequest.imp.metric.type": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || indexInfo.MetricIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Metric) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Metric[indexInfo.MetricIndex].Type, nil
	},
	"bidrequest.imp.metric.value": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || indexInfo.MetricIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Metric) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Metric[indexInfo.MetricIndex].Value, nil
	},
	"bidrequest.imp.metric.vendor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || indexInfo.MetricIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Metric) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Metric[indexInfo.MetricIndex].Vendor, nil
	},
	"bidrequest.imp.metric.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || indexInfo.MetricIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Metric) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Metric[indexInfo.MetricIndex].Ext, nil
	},
	"bidrequest.imp.banner": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner, nil
	},
	"bidrequest.imp.banner.format": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format, nil
	},
	"bidrequest.imp.banner.format.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].W, nil
	},
	"bidrequest.imp.banner.format.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].H, nil
	},
	"bidrequest.imp.banner.format.wratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].WRatio, nil
	},
	"bidrequest.imp.banner.format.hratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].HRatio, nil
	},
	"bidrequest.imp.banner.format.wmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].WMin, nil
	},
	"bidrequest.imp.banner.format.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Format[indexInfo.FormatIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.banner.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.W, nil
	},
	"bidrequest.imp.banner.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.H, nil
	},
	"bidrequest.imp.banner.wmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.WMax, nil
	},
	"bidrequest.imp.banner.hmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.HMax, nil
	},
	"bidrequest.imp.banner.wmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.WMin, nil
	},
	"bidrequest.imp.banner.hmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.HMin, nil
	},
	"bidrequest.imp.banner.btype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.BType, nil
	},
	"bidrequest.imp.banner.battr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.BAttr, nil
	},
	"bidrequest.imp.banner.pos": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Pos, nil
	},
	"bidrequest.imp.banner.mimes": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.MIMEs, nil
	},
	"bidrequest.imp.banner.topframe": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.TopFrame, nil
	},
	"bidrequest.imp.banner.expdir": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.ExpDir, nil
	},
	"bidrequest.imp.banner.api": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.API, nil
	},
	"bidrequest.imp.banner.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.ID, nil
	},
	"bidrequest.imp.banner.vcm": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Vcm, nil
	},
	"bidrequest.imp.banner.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Banner == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Banner.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Banner.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.video": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video, nil
	},
	"bidrequest.imp.video.mimes": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.MimesIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.MIMEs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MIMEs[indexInfo.MimesIndex], nil
	},
	"bidrequest.imp.video.minduration": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MinDuration, nil
	},
	"bidrequest.imp.video.maxduration": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MaxDuration, nil
	},
	"bidrequest.imp.video.startdelay": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.StartDelay, nil
	},
	"bidrequest.imp.video.maxseq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MaxSeq, nil
	},
	"bidrequest.imp.video.poddur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.PodDur, nil
	},
	"bidrequest.imp.video.protocols": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Protocols, nil
	},
	"bidrequest.imp.video.protocol": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Protocol, nil
	},
	"bidrequest.imp.video.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || bidRequest.Imp[indexInfo.ImpIndex].Video.W == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.W, nil
	},
	"bidrequest.imp.video.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || bidRequest.Imp[indexInfo.ImpIndex].Video.H == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.H, nil
	},
	"bidrequest.imp.video.podid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.PodID, nil
	},
	"bidrequest.imp.video.podseq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.PodSeq, nil
	},
	"bidrequest.imp.video.rqddurs": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.RqdDurs, nil
	},
	"bidrequest.imp.video.placement": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Placement, nil
	},
	"bidrequest.imp.video.plcmt": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Plcmt, nil
	},
	"bidrequest.imp.video.linearity": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Linearity, nil
	},
	"bidrequest.imp.video.skip": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Skip, nil
	},
	"bidrequest.imp.video.skipmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.SkipMin, nil
	},
	"bidrequest.imp.video.skipafter": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.SkipAfter, nil
	},
	"bidrequest.imp.video.sequence": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Sequence, nil
	},
	"bidrequest.imp.video.slotinpod": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.SlotInPod, nil
	},
	"bidrequest.imp.video.mincpmpersec": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MinCPMPerSec, nil
	},
	"bidrequest.imp.video.battr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.BattrIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.BAttr) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.BAttr[indexInfo.BattrIndex], nil
	},
	"bidrequest.imp.video.maxextended": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MaxExtended, nil
	},
	"bidrequest.imp.video.minbitrate": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MinBitRate, nil
	},
	"bidrequest.imp.video.maxbitrate": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.MaxBitRate, nil
	},
	"bidrequest.imp.video.boxingallowed": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.BoxingAllowed, nil
	},
	"bidrequest.imp.video.playbackmethod": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.PlaybackMethod, nil
	},
	"bidrequest.imp.video.playbackend": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.PlaybackEnd, nil
	},
	"bidrequest.imp.video.delivery": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Delivery, nil
	},
	"bidrequest.imp.video.pos": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || bidRequest.Imp[indexInfo.ImpIndex].Video.Pos == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Pos, nil
	},
	"bidrequest.imp.video.companionad": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd, nil
	},
	"bidrequest.imp.video.companionad.format": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex], nil
	},
	"bidrequest.imp.video.companionad.format.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].W, nil
	},
	"bidrequest.imp.video.companionad.format.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].H, nil
	},
	"bidrequest.imp.video.companionad.format.wratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].WRatio, nil
	},
	"bidrequest.imp.video.companionad.format.hratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].HRatio, nil
	},
	"bidrequest.imp.video.companionad.format.wmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].WMin, nil
	},
	"bidrequest.imp.video.companionad.format.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.video.companionad.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].W, nil
	},
	"bidrequest.imp.video.companionad.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].H, nil
	},
	"bidrequest.imp.video.companionad.wmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].WMax, nil
	},
	"bidrequest.imp.video.companionad.hmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].HMax, nil
	},
	"bidrequest.imp.video.companionad.wmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].WMin, nil
	},
	"bidrequest.imp.video.companionad.hmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].HMin, nil
	},
	"bidrequest.imp.video.companionad.btype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.BtypeIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].BType) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].BType[indexInfo.BtypeIndex], nil
	},
	"bidrequest.imp.video.companionad.battr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.BattrIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].BAttr) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].BAttr[indexInfo.BattrIndex], nil
	},
	"bidrequest.imp.video.companionad.pos": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Pos == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Pos, nil
	},
	"bidrequest.imp.video.companionad.mimes": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.MimesIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].MIMEs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].MIMEs[indexInfo.MimesIndex], nil
	},
	"bidrequest.imp.video.companionad.topframe": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].TopFrame, nil
	},
	"bidrequest.imp.video.companionad.expdir": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.ExpdirIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].ExpDir) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].ExpDir[indexInfo.ExpdirIndex], nil
	},
	"bidrequest.imp.video.companionad.api": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.ApiIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].API) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].API[indexInfo.ApiIndex], nil
	},
	"bidrequest.imp.video.companionad.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].ID, nil
	},
	"bidrequest.imp.video.companionad.vcm": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Vcm == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Vcm, nil
	},
	"bidrequest.imp.video.companionad.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionAd[indexInfo.CompanionadIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.video.api": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.ApiIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.API) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.API[indexInfo.ApiIndex], nil
	},
	"bidrequest.imp.video.companiontype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.CompanionType, nil
	},
	"bidrequest.imp.video.poddedupe": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.PodDedupe, nil
	},
	"bidrequest.imp.video.durfloors": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors, nil
	},
	"bidrequest.imp.video.durfloors.mindur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors[indexInfo.DurfloorsIndex].MinDur, nil
	},
	"bidrequest.imp.video.durfloors.maxdur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors[indexInfo.DurfloorsIndex].MaxDur, nil
	},
	"bidrequest.imp.video.durfloors.bidfloor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors[indexInfo.DurfloorsIndex].BidFloor, nil
	},
	"bidrequest.imp.video.durfloors.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors[indexInfo.DurfloorsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.DurFloors[indexInfo.DurfloorsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.video.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Video == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Video.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Video.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.audio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio, nil
	},
	"bidrequest.imp.audio.mimes": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.MimesIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.MIMEs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MIMEs[indexInfo.MimesIndex], nil
	},
	"bidrequest.imp.audio.minduration": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MinDuration, nil
	},
	"bidrequest.imp.audio.maxduration": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MaxDuration, nil
	},
	"bidrequest.imp.audio.poddur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.PodDur, nil
	},
	"bidrequest.imp.audio.protocols": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.ProtocolsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.Protocols) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.Protocols[indexInfo.ProtocolsIndex], nil
	},
	"bidrequest.imp.audio.startdelay": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || bidRequest.Imp[indexInfo.ImpIndex].Audio.StartDelay == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.StartDelay, nil
	},
	"bidrequest.imp.audio.rqddurs": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.RqddursIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.RqdDurs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.RqdDurs[indexInfo.RqddursIndex], nil
	},
	"bidrequest.imp.audio.podid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.PodID, nil
	},
	"bidrequest.imp.audio.podseq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.PodSeq, nil
	},
	"bidrequest.imp.audio.sequence": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.Sequence, nil
	},
	"bidrequest.imp.audio.slotinpod": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.SlotInPod, nil
	},
	"bidrequest.imp.audio.mincpmpersec": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MinCPMPerSec, nil
	},
	"bidrequest.imp.audio.battr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.BattrIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.BAttr) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.BAttr[indexInfo.BattrIndex], nil
	},
	"bidrequest.imp.audio.maxextended": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MaxExtended, nil
	},
	"bidrequest.imp.audio.minbitrate": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MinBitrate, nil
	},
	"bidrequest.imp.audio.maxbitrate": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MaxBitrate, nil
	},
	"bidrequest.imp.audio.delivery": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.DeliveryIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.Delivery) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.Delivery[indexInfo.DeliveryIndex], nil
	},
	"bidrequest.imp.audio.companionad": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex], nil
	},
	"bidrequest.imp.audio.companionad.format": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex], nil
	},
	"bidrequest.imp.audio.companionad.format.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].W, nil
	},
	"bidrequest.imp.audio.companionad.format.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].H, nil
	},
	"bidrequest.imp.audio.companionad.format.wratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].WRatio, nil
	},
	"bidrequest.imp.audio.companionad.format.hratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].HRatio, nil
	},
	"bidrequest.imp.audio.companionad.format.wmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].WMin, nil
	},
	"bidrequest.imp.audio.companionad.format.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.FormatIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Format[indexInfo.FormatIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.audio.companionad.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].W, nil
	},
	"bidrequest.imp.audio.companionad.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].H, nil
	},
	"bidrequest.imp.audio.companionad.wmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].WMax, nil
	},
	"bidrequest.imp.audio.companionad.hmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].HMax, nil
	},
	"bidrequest.imp.audio.companionad.wmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].WMin, nil
	},
	"bidrequest.imp.audio.companionad.hmin": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].HMin, nil
	},
	"bidrequest.imp.audio.companionad.btype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.BtypeIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].BType) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].BType[indexInfo.BtypeIndex], nil
	},
	"bidrequest.imp.audio.companionad.battr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.BattrIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].BAttr) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].BAttr[indexInfo.BattrIndex], nil
	},
	"bidrequest.imp.audio.companionad.pos": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Pos == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Pos, nil
	},
	"bidrequest.imp.audio.companionad.mimes": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.MimesIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].MIMEs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].MIMEs[indexInfo.MimesIndex], nil
	},
	"bidrequest.imp.audio.companionad.topframe": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].TopFrame, nil
	},
	"bidrequest.imp.audio.companionad.expdir": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.ExpdirIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].ExpDir) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].ExpDir[indexInfo.ExpdirIndex], nil
	},
	"bidrequest.imp.audio.companionad.api": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.ApiIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].API) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].API[indexInfo.ApiIndex], nil
	},
	"bidrequest.imp.audio.companionad.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].ID, nil
	},
	"bidrequest.imp.audio.companionad.vcm": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Vcm == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Vcm, nil
	},
	"bidrequest.imp.audio.companionad.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompanionadIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionAd[indexInfo.CompanionadIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.audio.api": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.ApiIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.API) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.API[indexInfo.ApiIndex], nil
	},
	"bidrequest.imp.audio.companiontype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.CompaniontypeIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionType) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.CompanionType[indexInfo.CompaniontypeIndex], nil
	},
	"bidrequest.imp.audio.maxseq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.MaxSeq, nil
	},
	"bidrequest.imp.audio.feed": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.Feed, nil
	},
	"bidrequest.imp.audio.stitched": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.Stitched, nil
	},
	"bidrequest.imp.audio.nvol": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.NVol, nil
	},
	"bidrequest.imp.audio.durfloors": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors[indexInfo.DurfloorsIndex], nil
	},
	"bidrequest.imp.audio.durfloors.mindur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors[indexInfo.DurfloorsIndex].MinDur, nil
	},
	"bidrequest.imp.audio.durfloors.maxdur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors[indexInfo.DurfloorsIndex].MaxDur, nil
	},
	"bidrequest.imp.audio.durfloors.bidfloor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors[indexInfo.DurfloorsIndex].BidFloor, nil
	},
	"bidrequest.imp.audio.durfloors.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors[indexInfo.DurfloorsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.DurFloors[indexInfo.DurfloorsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.audio.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Audio == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Audio.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Audio.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.native": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Native, nil
	},
	"bidrequest.imp.native.request": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Native == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Native.Request, nil
	},
	"bidrequest.imp.native.ver": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Native == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Native.Ver, nil
	},
	"bidrequest.imp.native.api": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Native == nil || indexInfo.ApiIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Native.API) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Native.API[indexInfo.ApiIndex], nil
	},
	"bidrequest.imp.native.battr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Native == nil || indexInfo.BattrIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Native.BAttr) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Native.BAttr[indexInfo.BattrIndex], nil
	},
	"bidrequest.imp.native.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Native == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Native.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Native.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.pmp": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP, nil
	},
	"bidrequest.imp.pmp.privateauction": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.PrivateAuction, nil
	},
	"bidrequest.imp.pmp.deals": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals, nil
	},
	"bidrequest.imp.pmp.deals.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].ID, nil
	},
	"bidrequest.imp.pmp.deals.bidfloor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].BidFloor, nil
	},
	"bidrequest.imp.pmp.deals.bidfloorcur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].BidFloorCur, nil
	},
	"bidrequest.imp.pmp.deals.at": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].AT, nil
	},
	"bidrequest.imp.pmp.deals.wseat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].WSeat, nil
	},
	"bidrequest.imp.pmp.deals.wadomain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].WADomain, nil
	},
	"bidrequest.imp.pmp.deals.guar": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].Guar, nil
	},
	"bidrequest.imp.pmp.deals.mincpmpersec": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].MinCPMPerSec, nil
	},
	"bidrequest.imp.pmp.deals.durfloors": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors[indexInfo.DurfloorsIndex], nil
	},
	"bidrequest.imp.pmp.deals.durfloors.mindur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors[indexInfo.DurfloorsIndex].MinDur, nil
	},
	"bidrequest.imp.pmp.deals.durfloors.maxdur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors[indexInfo.DurfloorsIndex].MaxDur, nil
	},
	"bidrequest.imp.pmp.deals.durfloors.bidfloor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors[indexInfo.DurfloorsIndex].BidFloor, nil
	},
	"bidrequest.imp.pmp.deals.durfloors.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) || indexInfo.DurfloorsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors[indexInfo.DurfloorsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].DurFloors[indexInfo.DurfloorsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.pmp.deals.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.DealsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Deals[indexInfo.DealsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.pmp.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].PMP == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].PMP.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].PMP.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.displaymanager": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].DisplayManager, nil
	},
	"bidrequest.imp.displaymanagerver": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].DisplayManagerVer, nil
	},
	"bidrequest.imp.instl": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Instl, nil
	},
	"bidrequest.imp.tagid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].TagID, nil
	},
	"bidrequest.imp.bidfloor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].BidFloor, nil
	},
	"bidrequest.imp.bidfloorcur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].BidFloorCur, nil
	},
	"bidrequest.imp.clickbrowser": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].ClickBrowser, nil
	},
	"bidrequest.imp.secure": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Secure, nil
	},
	"bidrequest.imp.iframebuster": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].IframeBuster, nil
	},
	"bidrequest.imp.rwdd": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Rwdd, nil
	},
	"bidrequest.imp.ssai": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].SSAI, nil
	},
	"bidrequest.imp.exp": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Exp, nil
	},
	"bidrequest.imp.qty": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Qty, nil
	},
	"bidrequest.imp.qty.multiplier": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Qty == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Qty.Multiplier, nil
	},
	"bidrequest.imp.qty.sourcetype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Qty == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Qty.SourceType, nil
	},
	"bidrequest.imp.qty.vendor": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Qty == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Qty.Vendor, nil
	},
	"bidrequest.imp.qty.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Qty == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Qty.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Qty.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.dt": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].DT, nil
	},
	"bidrequest.imp.refresh": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh, nil
	},
	"bidrequest.imp.refresh.refsettings": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Refresh == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings, nil
	},
	"bidrequest.imp.refresh.refsettings.reftype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Refresh == nil || indexInfo.RefsettingsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings[indexInfo.RefsettingsIndex].RefType, nil
	},
	"bidrequest.imp.refresh.refsettings.minint": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Refresh == nil || indexInfo.RefsettingsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings[indexInfo.RefsettingsIndex].MinInt, nil
	},
	"bidrequest.imp.refresh.refsettings.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Refresh == nil || indexInfo.RefsettingsIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings[indexInfo.RefsettingsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh.RefSettings[indexInfo.RefsettingsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.refresh.count": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Refresh == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh.Count, nil
	},
	"bidrequest.imp.refresh.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || bidRequest.Imp[indexInfo.ImpIndex].Refresh == nil || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Refresh.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Refresh.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.imp.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ImpIndex >= len(bidRequest.Imp) || indexInfo.ExtIndex >= len(bidRequest.Imp[indexInfo.ImpIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Imp[indexInfo.ImpIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Site, nil
	},
	"bidrequest.site.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.ID, nil
	},
	"bidrequest.site.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Name, nil
	},
	"bidrequest.site.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Domain, nil
	},
	"bidrequest.site.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.CatTax, nil
	},
	"bidrequest.site.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Cat, nil
	},
	"bidrequest.site.sectioncat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.SectionCat, nil
	},
	"bidrequest.site.pagecat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.PageCat, nil
	},
	"bidrequest.site.page": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Page, nil
	},
	"bidrequest.site.ref": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Ref, nil
	},
	"bidrequest.site.search": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Search, nil
	},
	"bidrequest.site.mobile": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Mobile, nil
	},
	"bidrequest.site.privacypolicy": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.PrivacyPolicy, nil
	},
	"bidrequest.site.publisher": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher, nil
	},
	"bidrequest.site.publisher.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher.ID, nil
	},
	"bidrequest.site.publisher.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher.Name, nil
	},
	"bidrequest.site.publisher.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher.CatTax, nil
	},
	"bidrequest.site.publisher.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Publisher == nil || indexInfo.CatIndex >= len(bidRequest.Site.Publisher.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.site.publisher.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher.Domain, nil
	},
	"bidrequest.site.publisher.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Publisher == nil || indexInfo.ExtIndex >= len(bidRequest.Site.Publisher.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Publisher.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.content": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content, nil
	},
	"bidrequest.site.content.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.ID, nil
	},
	"bidrequest.site.content.episode": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Episode, nil
	},
	"bidrequest.site.content.title": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Title, nil
	},
	"bidrequest.site.content.series": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Series, nil
	},
	"bidrequest.site.content.season": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Season, nil
	},
	"bidrequest.site.content.artist": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Artist, nil
	},
	"bidrequest.site.content.genre": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Genre, nil
	},
	"bidrequest.site.content.album": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Album, nil
	},
	"bidrequest.site.content.isrc": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.ISRC, nil
	},
	"bidrequest.site.content.producer": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer, nil
	},
	"bidrequest.site.content.producer.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer.ID, nil
	},
	"bidrequest.site.content.producer.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer.Name, nil
	},
	"bidrequest.site.content.producer.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer.CatTax, nil
	},
	"bidrequest.site.content.producer.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Producer == nil || indexInfo.CatIndex >= len(bidRequest.Site.Content.Producer.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.site.content.producer.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer.Domain, nil
	},
	"bidrequest.site.content.producer.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Producer == nil || indexInfo.ExtIndex >= len(bidRequest.Site.Content.Producer.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Producer.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.content.url": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.URL, nil
	},
	"bidrequest.site.content.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.CatTax, nil
	},
	"bidrequest.site.content.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.CatIndex >= len(bidRequest.Site.Content.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.site.content.prodq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.ProdQ, nil
	},
	"bidrequest.site.content.videoquality": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.VideoQuality, nil
	},
	"bidrequest.site.content.context": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Context, nil
	},
	"bidrequest.site.content.contentrating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.ContentRating, nil
	},
	"bidrequest.site.content.userrating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.UserRating, nil
	},
	"bidrequest.site.content.qagmediarating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.QAGMediaRating, nil
	},
	"bidrequest.site.content.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Keywords, nil
	},
	"bidrequest.site.content.kwarray": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.KwArray, nil
	},
	"bidrequest.site.content.livestream": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.LiveStream, nil
	},
	"bidrequest.site.content.sourcerelationship": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.SourceRelationship, nil
	},
	"bidrequest.site.content.len": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Len, nil
	},
	"bidrequest.site.content.language": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Language, nil
	},
	"bidrequest.site.content.langb": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.LangB, nil
	},
	"bidrequest.site.content.embeddable": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Embeddable, nil
	},
	"bidrequest.site.content.data": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data, nil
	},
	"bidrequest.site.content.data.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].ID, nil
	},
	"bidrequest.site.content.data.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Name, nil
	},
	"bidrequest.site.content.data.segment": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment, nil
	},
	"bidrequest.site.content.data.segment.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].ID, nil
	},
	"bidrequest.site.content.data.segment.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Name, nil
	},
	"bidrequest.site.content.data.segment.value": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Value, nil
	},
	"bidrequest.site.content.data.segment.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment) || indexInfo.ExtIndex >= len(bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.content.data.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.DataIndex >= len(bidRequest.Site.Content.Data) || indexInfo.ExtIndex >= len(bidRequest.Site.Content.Data[indexInfo.DataIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Data[indexInfo.DataIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.content.network": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Network, nil
	},
	"bidrequest.site.content.network.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Network.ID, nil
	},
	"bidrequest.site.content.network.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Network.Name, nil
	},
	"bidrequest.site.content.network.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Network.Domain, nil
	},
	"bidrequest.site.content.network.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Network == nil || indexInfo.ExtIndex >= len(bidRequest.Site.Content.Network.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Network.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.content.channel": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Channel, nil
	},
	"bidrequest.site.content.channel.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Channel.ID, nil
	},
	"bidrequest.site.content.channel.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Channel.Name, nil
	},
	"bidrequest.site.content.channel.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Channel.Domain, nil
	},
	"bidrequest.site.content.channel.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || bidRequest.Site.Content.Channel == nil || indexInfo.ExtIndex >= len(bidRequest.Site.Content.Channel.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Channel.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.content.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || bidRequest.Site.Content == nil || indexInfo.ExtIndex >= len(bidRequest.Site.Content.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Content.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.site.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Keywords, nil
	},
	"bidrequest.site.kwarray": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || indexInfo.KwarrayIndex >= len(bidRequest.Site.KwArray) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.KwArray[indexInfo.KwarrayIndex], nil
	},
	"bidrequest.site.inventorypartnerdomain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.InventoryPartnerDomain, nil
	},
	"bidrequest.site.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Site == nil || indexInfo.ExtIndex >= len(bidRequest.Site.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Site.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.App, nil
	},
	"bidrequest.app.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.ID, nil
	},
	"bidrequest.app.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Name, nil
	},
	"bidrequest.app.bundle": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Bundle, nil
	},
	"bidrequest.app.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Domain, nil
	},
	"bidrequest.app.storeurl": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.StoreURL, nil
	},
	"bidrequest.app.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.CatTax, nil
	},
	"bidrequest.app.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || indexInfo.CatIndex >= len(bidRequest.App.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.app.sectioncat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || indexInfo.SectioncatIndex >= len(bidRequest.App.SectionCat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.SectionCat[indexInfo.SectioncatIndex], nil
	},
	"bidrequest.app.pagecat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || indexInfo.PagecatIndex >= len(bidRequest.App.PageCat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.PageCat[indexInfo.PagecatIndex], nil
	},
	"bidrequest.app.ver": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Ver, nil
	},
	"bidrequest.app.privacypolicy": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.PrivacyPolicy == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.PrivacyPolicy, nil
	},
	"bidrequest.app.paid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Paid, nil
	},
	"bidrequest.app.publisher": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher, nil
	},
	"bidrequest.app.publisher.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher.ID, nil
	},
	"bidrequest.app.publisher.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher.Name, nil
	},
	"bidrequest.app.publisher.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher.CatTax, nil
	},
	"bidrequest.app.publisher.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil || indexInfo.CatIndex >= len(bidRequest.App.Publisher.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.app.publisher.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher.Domain, nil
	},
	"bidrequest.app.publisher.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Publisher == nil || indexInfo.ExtIndex >= len(bidRequest.App.Publisher.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Publisher.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.content": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content, nil
	},
	"bidrequest.app.content.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.ID, nil
	},
	"bidrequest.app.content.episode": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Episode, nil
	},
	"bidrequest.app.content.title": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Title, nil
	},
	"bidrequest.app.content.series": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Series, nil
	},
	"bidrequest.app.content.season": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Season, nil
	},
	"bidrequest.app.content.artist": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Artist, nil
	},
	"bidrequest.app.content.genre": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Genre, nil
	},
	"bidrequest.app.content.album": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Album, nil
	},
	"bidrequest.app.content.isrc": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.ISRC, nil
	},
	"bidrequest.app.content.producer": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer, nil
	},
	"bidrequest.app.content.producer.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer.ID, nil
	},
	"bidrequest.app.content.producer.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer.Name, nil
	},
	"bidrequest.app.content.producer.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer.CatTax, nil
	},
	"bidrequest.app.content.producer.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil || indexInfo.CatIndex >= len(bidRequest.App.Content.Producer.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.app.content.producer.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer.Domain, nil
	},
	"bidrequest.app.content.producer.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Producer == nil || indexInfo.ExtIndex >= len(bidRequest.App.Content.Producer.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Producer.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.content.url": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.URL, nil
	},
	"bidrequest.app.content.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.CatTax, nil
	},
	"bidrequest.app.content.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.CatIndex >= len(bidRequest.App.Content.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.app.content.prodq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.ProdQ == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.ProdQ, nil
	},
	"bidrequest.app.content.videoquality": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.VideoQuality == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.VideoQuality, nil
	},
	"bidrequest.app.content.context": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Context, nil
	},
	"bidrequest.app.content.contentrating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.ContentRating, nil
	},
	"bidrequest.app.content.userrating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.UserRating, nil
	},
	"bidrequest.app.content.qagmediarating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.QAGMediaRating, nil
	},
	"bidrequest.app.content.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Keywords, nil
	},
	"bidrequest.app.content.kwarray": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.KwarrayIndex >= len(bidRequest.App.Content.KwArray) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.KwArray[indexInfo.KwarrayIndex], nil
	},
	"bidrequest.app.content.livestream": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.LiveStream == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.LiveStream, nil
	},
	"bidrequest.app.content.sourcerelationship": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.SourceRelationship == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.SourceRelationship, nil
	},
	"bidrequest.app.content.len": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Len, nil
	},
	"bidrequest.app.content.language": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Language, nil
	},
	"bidrequest.app.content.langb": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.LangB, nil
	},
	"bidrequest.app.content.embeddable": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Embeddable == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Embeddable, nil
	},
	"bidrequest.app.content.data": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex], nil
	},
	"bidrequest.app.content.data.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].ID, nil
	},
	"bidrequest.app.content.data.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Name, nil
	},
	"bidrequest.app.content.data.segment": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex], nil
	},
	"bidrequest.app.content.data.segment.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].ID, nil
	},
	"bidrequest.app.content.data.segment.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Name, nil
	},
	"bidrequest.app.content.data.segment.value": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Value, nil
	},
	"bidrequest.app.content.data.segment.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Segment) || indexInfo.ExtIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.content.data.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.DataIndex >= len(bidRequest.App.Content.Data) || indexInfo.ExtIndex >= len(bidRequest.App.Content.Data[indexInfo.DataIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Data[indexInfo.DataIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.content.network": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Network, nil
	},
	"bidrequest.app.content.network.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Network.ID, nil
	},
	"bidrequest.app.content.network.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Network.Name, nil
	},
	"bidrequest.app.content.network.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Network.Domain, nil
	},
	"bidrequest.app.content.network.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Network == nil || indexInfo.ExtIndex >= len(bidRequest.App.Content.Network.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Network.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.content.channel": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Channel, nil
	},
	"bidrequest.app.content.channel.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Channel.ID, nil
	},
	"bidrequest.app.content.channel.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Channel.Name, nil
	},
	"bidrequest.app.content.channel.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Channel.Domain, nil
	},
	"bidrequest.app.content.channel.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || bidRequest.App.Content.Channel == nil || indexInfo.ExtIndex >= len(bidRequest.App.Content.Channel.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Channel.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.content.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || bidRequest.App.Content == nil || indexInfo.ExtIndex >= len(bidRequest.App.Content.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Content.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.app.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Keywords, nil
	},
	"bidrequest.app.kwarray": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || indexInfo.KwarrayIndex >= len(bidRequest.App.KwArray) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.KwArray[indexInfo.KwarrayIndex], nil
	},
	"bidrequest.app.inventorypartnerdomain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.InventoryPartnerDomain, nil
	},
	"bidrequest.app.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.App == nil || indexInfo.ExtIndex >= len(bidRequest.App.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.App.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.DOOH, nil
	},
	"bidrequest.dooh.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.ID, nil
	},
	"bidrequest.dooh.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Name, nil
	},
	"bidrequest.dooh.venuetype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.VenueType, nil
	},
	"bidrequest.dooh.venuetypetax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.VenueTypeTax, nil
	},
	"bidrequest.dooh.publisher": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher, nil
	},
	"bidrequest.dooh.publisher.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher.ID, nil
	},
	"bidrequest.dooh.publisher.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher.Name, nil
	},
	"bidrequest.dooh.publisher.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher.CatTax, nil
	},
	"bidrequest.dooh.publisher.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil || indexInfo.CatIndex >= len(bidRequest.DOOH.Publisher.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.dooh.publisher.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher.Domain, nil
	},
	"bidrequest.dooh.publisher.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Publisher == nil || indexInfo.ExtIndex >= len(bidRequest.DOOH.Publisher.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Publisher.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Domain, nil
	},
	"bidrequest.dooh.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Keywords, nil
	},
	"bidrequest.dooh.content": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content, nil
	},
	"bidrequest.dooh.content.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.ID, nil
	},
	"bidrequest.dooh.content.episode": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Episode, nil
	},
	"bidrequest.dooh.content.title": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Title, nil
	},
	"bidrequest.dooh.content.series": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Series, nil
	},
	"bidrequest.dooh.content.season": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Season, nil
	},
	"bidrequest.dooh.content.artist": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Artist, nil
	},
	"bidrequest.dooh.content.genre": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Genre, nil
	},
	"bidrequest.dooh.content.album": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Album, nil
	},
	"bidrequest.dooh.content.isrc": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.ISRC, nil
	},
	"bidrequest.dooh.content.producer": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer, nil
	},
	"bidrequest.dooh.content.producer.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer.ID, nil
	},
	"bidrequest.dooh.content.producer.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer.Name, nil
	},
	"bidrequest.dooh.content.producer.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer.CatTax, nil
	},
	"bidrequest.dooh.content.producer.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil || indexInfo.CatIndex >= len(bidRequest.DOOH.Content.Producer.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.dooh.content.producer.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer.Domain, nil
	},
	"bidrequest.dooh.content.producer.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Producer == nil || indexInfo.ExtIndex >= len(bidRequest.DOOH.Content.Producer.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Producer.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.content.url": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.URL, nil
	},
	"bidrequest.dooh.content.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.CatTax, nil
	},
	"bidrequest.dooh.content.cat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.CatIndex >= len(bidRequest.DOOH.Content.Cat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Cat[indexInfo.CatIndex], nil
	},
	"bidrequest.dooh.content.prodq": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.ProdQ == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.ProdQ, nil
	},
	"bidrequest.dooh.content.videoquality": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.VideoQuality == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.VideoQuality, nil
	},
	"bidrequest.dooh.content.context": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Context, nil
	},
	"bidrequest.dooh.content.contentrating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.ContentRating, nil
	},
	"bidrequest.dooh.content.userrating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.UserRating, nil
	},
	"bidrequest.dooh.content.qagmediarating": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.QAGMediaRating, nil
	},
	"bidrequest.dooh.content.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Keywords, nil
	},
	"bidrequest.dooh.content.kwarray": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.KwarrayIndex >= len(bidRequest.DOOH.Content.KwArray) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.KwArray[indexInfo.KwarrayIndex], nil
	},
	"bidrequest.dooh.content.livestream": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.LiveStream == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.LiveStream, nil
	},
	"bidrequest.dooh.content.sourcerelationship": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.SourceRelationship == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.SourceRelationship, nil
	},
	"bidrequest.dooh.content.len": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Len, nil
	},
	"bidrequest.dooh.content.language": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Language, nil
	},
	"bidrequest.dooh.content.langb": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.LangB, nil
	},
	"bidrequest.dooh.content.embeddable": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Embeddable == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Embeddable, nil
	},
	"bidrequest.dooh.content.data": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex], nil
	},
	"bidrequest.dooh.content.data.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].ID, nil
	},
	"bidrequest.dooh.content.data.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Name, nil
	},
	"bidrequest.dooh.content.data.segment": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex], nil
	},
	"bidrequest.dooh.content.data.segment.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].ID, nil
	},
	"bidrequest.dooh.content.data.segment.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Name, nil
	},
	"bidrequest.dooh.content.data.segment.value": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Value, nil
	},
	"bidrequest.dooh.content.data.segment.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) || indexInfo.SegmentIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment) || indexInfo.ExtIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.content.data.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.DataIndex >= len(bidRequest.DOOH.Content.Data) || indexInfo.ExtIndex >= len(bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Data[indexInfo.DataIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.content.network": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Network, nil
	},
	"bidrequest.dooh.content.network.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Network.ID, nil
	},
	"bidrequest.dooh.content.network.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Network.Name, nil
	},
	"bidrequest.dooh.content.network.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Network == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Network.Domain, nil
	},
	"bidrequest.dooh.content.network.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Network == nil || indexInfo.ExtIndex >= len(bidRequest.DOOH.Content.Network.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Network.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.content.channel": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Channel, nil
	},
	"bidrequest.dooh.content.channel.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Channel.ID, nil
	},
	"bidrequest.dooh.content.channel.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Channel.Name, nil
	},
	"bidrequest.dooh.content.channel.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Channel == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Channel.Domain, nil
	},
	"bidrequest.dooh.content.channel.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || bidRequest.DOOH.Content.Channel == nil || indexInfo.ExtIndex >= len(bidRequest.DOOH.Content.Channel.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Channel.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.content.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || bidRequest.DOOH.Content == nil || indexInfo.ExtIndex >= len(bidRequest.DOOH.Content.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Content.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.dooh.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.DOOH == nil || indexInfo.ExtIndex >= len(bidRequest.DOOH.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.DOOH.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.device": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Device, nil
	},
	"bidrequest.device.geo": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo, nil
	},
	"bidrequest.device.geo.lat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Lat, nil
	},
	"bidrequest.device.geo.lon": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Lon, nil
	},
	"bidrequest.device.geo.type": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Type, nil
	},
	"bidrequest.device.geo.accuracy": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Accuracy, nil
	},
	"bidrequest.device.geo.lastfix": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.LastFix, nil
	},
	"bidrequest.device.geo.ipservice": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.IPService, nil
	},
	"bidrequest.device.geo.country": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Country, nil
	},
	"bidrequest.device.geo.region": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Region, nil
	},
	"bidrequest.device.geo.regionfips104": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.RegionFIPS104, nil
	},
	"bidrequest.device.geo.metro": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Metro, nil
	},
	"bidrequest.device.geo.city": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.City, nil
	},
	"bidrequest.device.geo.zip": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.ZIP, nil
	},
	"bidrequest.device.geo.utcoffset": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.UTCOffset, nil
	},
	"bidrequest.device.geo.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.Geo == nil || indexInfo.ExtIndex >= len(bidRequest.Device.Geo.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Geo.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.device.dnt": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.DNT, nil
	},
	"bidrequest.device.lmt": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Lmt, nil
	},
	"bidrequest.device.ua": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.UA, nil
	},
	"bidrequest.device.sua": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA, nil
	},
	"bidrequest.device.sua.browsers": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Browsers, nil
	},
	"bidrequest.device.sua.browsers.brand": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || indexInfo.BrowsersIndex >= len(bidRequest.Device.SUA.Browsers) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Browsers[indexInfo.BrowsersIndex].Brand, nil
	},
	"bidrequest.device.sua.browsers.version": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || indexInfo.BrowsersIndex >= len(bidRequest.Device.SUA.Browsers) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Browsers[indexInfo.BrowsersIndex].Version, nil
	},
	"bidrequest.device.sua.browsers.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || indexInfo.BrowsersIndex >= len(bidRequest.Device.SUA.Browsers) || indexInfo.ExtIndex >= len(bidRequest.Device.SUA.Browsers[indexInfo.BrowsersIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Browsers[indexInfo.BrowsersIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.device.sua.platform": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Platform, nil
	},
	"bidrequest.device.sua.platform.brand": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || bidRequest.Device.SUA.Platform == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Platform.Brand, nil
	},
	"bidrequest.device.sua.platform.version": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || bidRequest.Device.SUA.Platform == nil || indexInfo.VersionIndex >= len(bidRequest.Device.SUA.Platform.Version) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Platform.Version[indexInfo.VersionIndex], nil
	},
	"bidrequest.device.sua.platform.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || bidRequest.Device.SUA.Platform == nil || indexInfo.ExtIndex >= len(bidRequest.Device.SUA.Platform.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Platform.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.device.sua.mobile": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || bidRequest.Device.SUA.Mobile == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Mobile, nil
	},
	"bidrequest.device.sua.architecture": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Architecture, nil
	},
	"bidrequest.device.sua.bitness": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Bitness, nil
	},
	"bidrequest.device.sua.model": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Model, nil
	},
	"bidrequest.device.sua.source": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Source, nil
	},
	"bidrequest.device.sua.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || bidRequest.Device.SUA == nil || indexInfo.ExtIndex >= len(bidRequest.Device.SUA.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.SUA.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.device.ip": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.IP, nil
	},
	"bidrequest.device.ipv6": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.IPv6, nil
	},
	"bidrequest.device.devicetype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.DeviceType, nil
	},
	"bidrequest.device.make": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Make, nil
	},
	"bidrequest.device.model": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Model, nil
	},
	"bidrequest.device.os": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.OS, nil
	},
	"bidrequest.device.osv": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.OSV, nil
	},
	"bidrequest.device.hwv": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.HWV, nil
	},
	"bidrequest.device.h": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.H, nil
	},
	"bidrequest.device.w": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.W, nil
	},
	"bidrequest.device.ppi": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.PPI, nil
	},
	"bidrequest.device.pxratio": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.PxRatio, nil
	},
	"bidrequest.device.js": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.JS, nil
	},
	"bidrequest.device.geofetch": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.GeoFetch, nil
	},
	"bidrequest.device.flashver": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.FlashVer, nil
	},
	"bidrequest.device.language": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Language, nil
	},
	"bidrequest.device.langb": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.LangB, nil
	},
	"bidrequest.device.carrier": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Carrier, nil
	},
	"bidrequest.device.mccmnc": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.MCCMNC, nil
	},
	"bidrequest.device.connectiontype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.ConnectionType, nil
	},
	"bidrequest.device.ifa": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.IFA, nil
	},
	"bidrequest.device.didsha1": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.DIDSHA1, nil
	},
	"bidrequest.device.didmd5": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.DIDMD5, nil
	},
	"bidrequest.device.dpidsha1": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.DPIDSHA1, nil
	},
	"bidrequest.device.dpidmd5": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.DPIDMD5, nil
	},
	"bidrequest.device.macsha1": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.MACSHA1, nil
	},
	"bidrequest.device.macmd5": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.MACMD5, nil
	},
	"bidrequest.device.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Device == nil || indexInfo.ExtIndex >= len(bidRequest.Device.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Device.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.user": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.User, nil
	},
	"bidrequest.user.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.ID, nil
	},
	"bidrequest.user.buyeruid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.BuyerUID, nil
	},
	"bidrequest.user.yob": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Yob, nil
	},
	"bidrequest.user.gender": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Gender, nil
	},
	"bidrequest.user.keywords": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Keywords, nil
	},
	"bidrequest.user.kwarray": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.KwarrayIndex >= len(bidRequest.User.KwArray) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.KwArray[indexInfo.KwarrayIndex], nil
	},
	"bidrequest.user.customdata": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.CustomData, nil
	},
	"bidrequest.user.geo": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo, nil
	},
	"bidrequest.user.geo.lat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil || bidRequest.User.Geo.Lat == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Lat, nil
	},
	"bidrequest.user.geo.lon": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil || bidRequest.User.Geo.Lon == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Lon, nil
	},
	"bidrequest.user.geo.type": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Type, nil
	},
	"bidrequest.user.geo.accuracy": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Accuracy, nil
	},
	"bidrequest.user.geo.lastfix": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.LastFix, nil
	},
	"bidrequest.user.geo.ipservice": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.IPService, nil
	},
	"bidrequest.user.geo.country": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Country, nil
	},
	"bidrequest.user.geo.region": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Region, nil
	},
	"bidrequest.user.geo.regionfips104": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.RegionFIPS104, nil
	},
	"bidrequest.user.geo.metro": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Metro, nil
	},
	"bidrequest.user.geo.city": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.City, nil
	},
	"bidrequest.user.geo.zip": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.ZIP, nil
	},
	"bidrequest.user.geo.utcoffset": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.UTCOffset, nil
	},
	"bidrequest.user.geo.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || bidRequest.User.Geo == nil || indexInfo.ExtIndex >= len(bidRequest.User.Geo.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Geo.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.user.data": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex], nil
	},
	"bidrequest.user.data.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].ID, nil
	},
	"bidrequest.user.data.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Name, nil
	},
	"bidrequest.user.data.segment": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) || indexInfo.SegmentIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex], nil
	},
	"bidrequest.user.data.segment.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) || indexInfo.SegmentIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].ID, nil
	},
	"bidrequest.user.data.segment.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) || indexInfo.SegmentIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Name, nil
	},
	"bidrequest.user.data.segment.value": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) || indexInfo.SegmentIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Segment) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Value, nil
	},
	"bidrequest.user.data.segment.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) || indexInfo.SegmentIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Segment) || indexInfo.ExtIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Segment[indexInfo.SegmentIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.user.data.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.DataIndex >= len(bidRequest.User.Data) || indexInfo.ExtIndex >= len(bidRequest.User.Data[indexInfo.DataIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Data[indexInfo.DataIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.user.consent": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Consent, nil
	},
	"bidrequest.user.eids": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs, nil
	},
	"bidrequest.user.eids.source": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.EidsIndex >= len(bidRequest.User.EIDs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs[indexInfo.EidsIndex].Source, nil
	},
	"bidrequest.user.eids.uids": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.EidsIndex >= len(bidRequest.User.EIDs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs, nil
	},
	"bidrequest.user.eids.uids.id": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.EidsIndex >= len(bidRequest.User.EIDs) || indexInfo.UidsIndex >= len(bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs[indexInfo.UidsIndex].ID, nil
	},
	"bidrequest.user.eids.uids.atype": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.EidsIndex >= len(bidRequest.User.EIDs) || indexInfo.UidsIndex >= len(bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs[indexInfo.UidsIndex].AType, nil
	},
	"bidrequest.user.eids.uids.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.EidsIndex >= len(bidRequest.User.EIDs) || indexInfo.UidsIndex >= len(bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs) || indexInfo.ExtIndex >= len(bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs[indexInfo.UidsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs[indexInfo.EidsIndex].UIDs[indexInfo.UidsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.user.eids.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.EidsIndex >= len(bidRequest.User.EIDs) || indexInfo.ExtIndex >= len(bidRequest.User.EIDs[indexInfo.EidsIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.EIDs[indexInfo.EidsIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.user.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.User == nil || indexInfo.ExtIndex >= len(bidRequest.User.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.User.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.test": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Test, nil
	},
	"bidrequest.at": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.AT, nil
	},
	"bidrequest.tmax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.TMax, nil
	},
	"bidrequest.wseat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.WseatIndex >= len(bidRequest.WSeat) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.WSeat[indexInfo.WseatIndex], nil
	},
	"bidrequest.bseat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.BSeat, nil
	},
	"bidrequest.allimps": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.AllImps, nil
	},
	"bidrequest.cur": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Cur, nil
	},
	"bidrequest.wlang": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.WLang, nil
	},
	"bidrequest.wlangb": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.WLangB, nil
	},
	"bidrequest.acat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.ACat, nil
	},
	"bidrequest.bcat": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.BCat, nil
	},
	"bidrequest.cattax": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.CatTax, nil
	},
	"bidrequest.badv": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.BAdv, nil
	},
	"bidrequest.bapp": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.BApp, nil
	},
	"bidrequest.source": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Source, nil
	},
	"bidrequest.source.fd": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.FD, nil
	},
	"bidrequest.source.tid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.TID, nil
	},
	"bidrequest.source.pchain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.PChain, nil
	},
	"bidrequest.source.schain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain, nil
	},
	"bidrequest.source.schain.complete": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Complete, nil
	},
	"bidrequest.source.schain.nodes": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes, nil
	},
	"bidrequest.source.schain.nodes.asi": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].ASI, nil
	},
	"bidrequest.source.schain.nodes.sid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].SID, nil
	},
	"bidrequest.source.schain.nodes.rid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].RID, nil
	},
	"bidrequest.source.schain.nodes.name": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].Name, nil
	},
	"bidrequest.source.schain.nodes.domain": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].Domain, nil
	},
	"bidrequest.source.schain.nodes.hp": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].HP, nil
	},
	"bidrequest.source.schain.nodes.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.NodesIndex >= len(bidRequest.Source.SChain.Nodes) || indexInfo.ExtIndex >= len(bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Nodes[indexInfo.NodesIndex].Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.source.schain.ver": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Ver, nil
	},
	"bidrequest.source.schain.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || bidRequest.Source.SChain == nil || indexInfo.ExtIndex >= len(bidRequest.Source.SChain.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.SChain.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.source.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Source == nil || indexInfo.ExtIndex >= len(bidRequest.Source.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Source.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.regs": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {

		return bidRequest.Regs, nil
	},
	"bidrequest.regs.coppa": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Regs == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Regs.COPPA, nil
	},
	"bidrequest.regs.gdpr": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Regs == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Regs.GDPR, nil
	},
	"bidrequest.regs.usprivacy": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Regs == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Regs.USPrivacy, nil
	},
	"bidrequest.regs.gpp": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Regs == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Regs.GPP, nil
	},
	"bidrequest.regs.gppsid": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Regs == nil {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Regs.GPPSID, nil
	},
	"bidrequest.regs.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if bidRequest.Regs == nil || indexInfo.ExtIndex >= len(bidRequest.Regs.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Regs.Ext[indexInfo.ExtIndex], nil
	},
	"bidrequest.ext": func(bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo) (interface{}, error) {
		if indexInfo.ExtIndex >= len(bidRequest.Ext) {
			return nil, fmt.Errorf("invalid index present in indexInfo object '%+v'", indexInfo)
		}
		return bidRequest.Ext[indexInfo.ExtIndex], nil
	},
}

// LookUpBidRequest obtains value of type R using key from bidRequest. It uses indexIndo if required
// In case of failure returns error
func LookUpBidRequest[R any](key string, bidRequest openrtb2.BidRequest, indexInfo BidRequestIndexInfo, returnType R) (R, error) {
	fn := bidRequestLookup[key]
	if fn == nil {
		return returnType, fmt.Errorf("invalid OGNL key [%s]", key)
	}
	result, err := fn(bidRequest, indexInfo)
	if err != nil {
		// Note: returntype here is not actual value
		return returnType, err
	}
	if rTypeValue, ok := result.(R); ok {
		return rTypeValue, nil
	}
	return returnType, fmt.Errorf("data type mismatch.Expected [%T] but found [%T]", returnType, result)
}
