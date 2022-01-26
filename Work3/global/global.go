package global

type GetUserReq struct {
	UserID int64 `json:"userId"`
}

type GetUserResp struct {
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
}