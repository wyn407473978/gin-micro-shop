package response

type UserPageReq struct {
	PageSize int    `json:"pageSize"`
	PageNum  int    `json:"pageNum"`
	UserName string `json:"userName"`
}

type UserCreateReq struct {
	UserName string `json:"userName"`
	Age      int32  `json:"age"`
	Sex      int32  `json:"sex"`
}
