package models

type Response struct {
	State int    `json:"state,omitempty"`
	Msg   string `json:"msg,omitempty"`
}
