package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CopyRequest struct {
	SourceBucket string   `json:"sourceBucket"`
	Objects      []string `json:"objectKeys"`
	DestBucket   string   `json:"destBucket"`
	DestLocation string   `json:"destLocation"`
}

func (copyRequest *CopyRequest) Validate() error {
	return validation.ValidateStruct(copyRequest,
		validation.Field(&copyRequest.SourceBucket, validation.Required),
		validation.Field(&copyRequest.DestBucket, validation.Required),
		validation.Field(&copyRequest.Objects, validation.Required),
	)
}
