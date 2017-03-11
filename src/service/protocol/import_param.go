package protocol

type ImportParam struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	File       string `json:"file"`
}
