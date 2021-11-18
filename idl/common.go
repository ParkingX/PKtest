package pvt

// InputCommon input paramter common field
//
// example:
// type HelloInput struct {
// 	InputCommon
// 	Body struct {
// 		deviceID string
// 	}
// }
type InputCommon struct {
	RequestID string `json:"request_id"`
	Cmd       string `json:"cmd"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
}

// OutputCommon output paramter common field
//
// example
// type HelloOutput struct {
// 	OutputCommon
// 	Body struct {
// 		Status uint32
// 	}
// }
type OutputCommon struct {
	RequestID string `json:"request_id"`
	Code      int64  `json:"code"`
	CodeMsg   string `json:"code_msg"`
	TraceID   string `json:"trace_id"`
}
