package structs

type Pagination struct {
	PageNum int `json:"pageNum"`
	Limit   int `json:"limit"`
}
