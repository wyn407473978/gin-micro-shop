package response

type UserPageReq struct {
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
	UserName string `json:"userName"`
}
