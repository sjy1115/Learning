package proto

type StudentListReq struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type StudentListItem struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Sex   int    `json:"sex"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}

type StudentListResp struct {
	Count   int               `json:"count"`
	Results []StudentListItem `json:"results"`
}

type StudentCreateReq struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Sex   int    `json:"sex"`
	Phone string `json:"phone"`
}

type StudentCreateResp struct {
	Id int `json:"id"`
}

type StudentDetailReq struct {
	Id int `json:"id"`
}

type StudentDetailResp struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Sex   int    `json:"sex"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}

type StudentUpdateReq struct {
}

type StudentUpdateResp struct {
	Id int `json:"id"`
}

type StudentDeleteReq struct {
	Id int `json:"id"`
}
