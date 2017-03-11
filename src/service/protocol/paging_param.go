package protocol

type PagingParam struct {
	PageIndex uint32 `json:"pageIndex"`
	PageSize  uint32 `json:"pageSize"`
}
