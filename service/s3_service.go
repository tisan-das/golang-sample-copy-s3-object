package service

import (
	"copy-s3-object/client"
	"copy-s3-object/dto"
)

type S3Service interface {
	CopyAllObjects(dto.CopyRequest) error
}

type s3Service struct {
	client client.S3Client
}

func NewS3Service(client client.S3Client) S3Service {
	return &s3Service{client: client}
}

func (svc *s3Service) CopyAllObjects(payload dto.CopyRequest) error {
	for _, object := range payload.Objects {
		err := svc.client.CopyObject(payload.SourceBucket, object,
			payload.DestBucket, payload.DestLocation, "")
		if err != nil {
			return err
		}
	}
	return nil
}
