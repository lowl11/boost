package s3_service

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (service *Service) get(ctx context.Context, prefix ...string) ([]*s3.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:     aws.String(service.bucket),
		StartAfter: aws.String(""),
	}

	if len(prefix) > 0 {
		//input.Prefix = aws.String(prefix[0])
		input.StartAfter = aws.String(prefix[0])
	}

	// List objects in a bucket
	var objects []*s3.Object
	if err := service.client.ListObjectsV2PagesWithContext(ctx, input,
		func(output *s3.ListObjectsV2Output, b bool) bool {
			objects = output.Contents
			return false
		},
	); err != nil {
		return nil, ErrorGetObjects(err, input.Prefix)
	}

	return objects, nil
}
