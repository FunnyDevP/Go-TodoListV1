package model

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

type ResponseFailed struct {
	Error Error `json:"error"`
}

type Error struct {
	Status int `json:"status"`
	Detail string `json:"detail"`
}