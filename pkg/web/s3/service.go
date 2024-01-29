package s3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lowl11/boost/data/errors"
	"io"
)

const (
	defaultMaxRetries = 3
)

type service struct {
	maxRetries int
	region     string
	bucket     string

	connection *session.Session
	client     *s3.S3
}

var instance *service

func getService() *service {
	if instance != nil {
		return instance
	}

	instance = &service{
		maxRetries: defaultMaxRetries,
	}
	return instance
}

func (service *service) MaxRetries(retries int) *service {
	service.maxRetries = retries
	return service
}

func (service *service) Region(region string) *service {
	service.region = region
	return service
}

func (service *service) Bucket(bucket string) *service {
	service.bucket = bucket
	return service
}

func (service *service) Connect() error {
	var err error
	service.connection, err = session.NewSession(&aws.Config{
		Region:     aws.String(service.region),
		MaxRetries: &service.maxRetries,
	})
	if err != nil {
		return errors.
			New("Connect to S3 error").
			SetType("ConnectS3Error").
			SetError(err)
	}

	// Create an S3 service client
	service.client = s3.New(service.connection)

	return nil
}

func (service *service) CreateFolder(ctx context.Context, path string, acl ...string) error {
	var aclValue *string
	if len(acl) > 0 {
		aclValue = aws.String(acl[0])
	}

	if _, err := service.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(path + "/"),
		ACL:    aclValue,
	}); err != nil {
		return errors.
			New("Create folder error").
			SetType("S3_CreateFolderError").
			SetError(err).
			AddContext("path", path)
	}

	return nil
}

func (service *service) CreateFile(ctx context.Context, path string, body []byte, acl ...string) error {
	var aclValue *string
	if len(acl) > 0 {
		aclValue = aws.String(acl[0])
	}

	if _, err := service.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(path),
		Body:   bytes.NewReader(body),
		ACL:    aclValue,
	}); err != nil {
		return errors.
			New("Create file error").
			SetType("S3_CreateFileError").
			SetError(err).
			AddContext("path", path)
	}

	return nil
}

func (service *service) Delete(ctx context.Context, name string) error {
	if _, err := service.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(name),
	}); err != nil {
		return errors.
			New("Delete object error").
			SetType("S3_DeleteObjectError").
			SetError(err).
			AddContext("path", name)
	}

	return nil
}

func (service *service) Rename(ctx context.Context, oldName, newName string) error {
	// create new based on original with new name
	if _, err := service.client.CopyObjectWithContext(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(service.bucket),
		CopySource: aws.String(oldName),
		Key:        aws.String(newName),
	}); err != nil {
		return errors.
			New("Copy object error").
			SetType("S3_CopyObjectError").
			SetError(err).
			AddContext("old", oldName).
			AddContext("new", newName)
	}

	// delete original file
	return service.Delete(ctx, oldName)
}

func (service *service) GetAll(ctx context.Context) ([]*s3.Object, error) {
	return service.get(ctx)
}

func (service *service) GetPath(ctx context.Context, path string) ([]*s3.Object, error) {
	return service.get(ctx, path)
}

func (service *service) GetFile(ctx context.Context, path string) ([]byte, error) {
	result, err := service.client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(service.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, err
	}

	return io.ReadAll(result.Body)
}

func (service *service) get(ctx context.Context, prefix ...string) ([]*s3.Object, error) {
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
		return nil, errors.
			New("Get objects error").
			SetType("S3_GetObjectsError").
			SetError(err).
			AddContext("path", input.StartAfter)
	}

	return objects, nil
}
