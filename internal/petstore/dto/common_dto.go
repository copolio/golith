package dto

import (
	"github.com/copolio/golith/golithgin"
)

type DefaultResponse struct {
	Code    golithgin.HttpStatus `json:"code"`
	Type    string               `json:"type"`
	Message string               `json:"message"`
}
