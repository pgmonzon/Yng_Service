package models

//Todo struct to todo
type Token struct {
	Raw        string         `json:"Raw"`
	Method     interface{}    `json:"Method"`
	Header     interface{}    `json:"Header"`
	Claims     interface{}    `json:"Claims"`
	Signature  string         `json:"Signature"`
  Valid      bool           `json:"Valid"`
}
