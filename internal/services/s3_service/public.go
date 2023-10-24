package s3_service

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (service *Service) MaxRetries(retries int) *Service {
	service.maxRetries = retries
	return service
}

func (service *Service) Region(region string) *Service {
	service.region = region
	return service
}

func (service *Service) Bucket(bucket string) *Service {
	service.bucket = bucket
	return service
}

func (service *Service) Connect() error {
	var err error
	service.connection, err = session.NewSession(&aws.Config{
		Region:     aws.String(service.region),
		MaxRetries: &service.maxRetries,
	})
	if err != nil {
		return ErrorConnect(err)
	}

	// Create an S3 service client
	service.client = s3.New(service.connection)

	return nil
}

func (service *Service) CreateFolder(ctx context.Context, path string) error {
	if _, err := service.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(path + "/"),
	}); err != nil {
		return ErrorCreateFolder(err, path)
	}

	return nil
}

func (service *Service) CreateFile(ctx context.Context, path string, body []byte) error {
	if _, err := service.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(path),
		Body:   bytes.NewReader(body),
	}); err != nil {
		return ErrorCreateFile(err, path)
	}

	return nil
}

func (service *Service) Delete(ctx context.Context, name string) error {
	if _, err := service.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(name),
	}); err != nil {
		return ErrorDeleteObject(err, name)
	}

	return nil
}

func (service *Service) Rename(ctx context.Context, oldName, newName string) error {
	// create new based on original with new name
	if _, err := service.client.CopyObjectWithContext(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(service.bucket),
		CopySource: aws.String(oldName),
		Key:        aws.String(newName),
	}); err != nil {
		return ErrorCopyObject(err, oldName, newName)
	}

	// delete original file
	return service.Delete(ctx, oldName)
}

func (service *Service) GetAll(ctx context.Context) ([]*s3.Object, error) {
	return service.get(ctx)
}

func (service *Service) GetPath(ctx context.Context, path string) ([]*s3.Object, error) {
	return service.get(ctx, path)
}
