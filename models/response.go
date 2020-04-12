package models

// JSONResponse API统一返回信息格式
type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// AttackMapDataResponse 给球的数据格式
type AttackMapDataResponse struct {
	SourceType string `json:"type"`
	EventType  string `json:"event_type"`
	TimeStamp  string `json:"@timestamp"`

	SrcIP     string  `json:"src_ip"`
	SrcLat    float64 `json:"src_lat"`
	SrcLng    float64 `json:"src_lng"`
	SrcRegion string  `json:"src_reg"`

	DstIP     string  `json:"dest_ip"`
	DstLat    float64 `json:"dest_lat"`
	DstLng    float64 `json:"dest_lng"`
	DstRegion string  `json:"dest_reg"`
}

// AttackMapCtrResponse 给球的统计格式
type AttackMapCtrResponse struct {
	SourceType string `json:"type"`
	CtrAllTime int    `json:"all_time"`
	Ctr7d      int    `json:"7d"`
	Ctr1d      int    `json:"1d"`
	Ctr1h      int    `json:"1h"`
	Ctr1m      int    `json:"1m"`
}
