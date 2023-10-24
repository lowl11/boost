package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lowl11/boost/internal/services/s3_service"
	"github.com/lowl11/boost/pkg/io/paths"
)

func Init(region, bucket string) error {
	return service().
		Region(region).
		Bucket(bucket).
		Connect()
}

func MustInit(region, bucket string) {
	if err := Init(region, bucket); err != nil {
		panic(err)
	}
}

func CreateFile(ctx context.Context, path string, body []byte) error {
	return service().CreateFile(ctx, path, body)
}

func CreateFolder(ctx context.Context, path string) error {
	return service().CreateFolder(ctx, path)
}

func Delete(ctx context.Context, path string) error {
	return service().Delete(ctx, path)
}

func GetAll(ctx context.Context) ([]Object, error) {
	objects, err := service().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return entityToModel(objects), nil
}

func GetPath(ctx context.Context, path string) ([]Object, error) {
	objects, err := service().GetPath(ctx, path)
	if err != nil {
		return nil, err
	}

	return entityToModel(objects), nil
}

func service() *s3_service.Service {
	return s3_service.Get()
}

func entityToModel(objects []*s3.Object) []Object {
	list := make([]Object, 0, len(objects))
	for _, obj := range objects {
		_, name := paths.GetFolderName(*obj.Key)

		list = append(list, Object{
			Name: name,
			Path: *obj.Key,
			Size: *obj.Size,
		})
	}

	return list
}
