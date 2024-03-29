package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lowl11/boost/pkg/io/paths"
)

func Init(region, bucket string) error {
	return getService().
		Region(region).
		Bucket(bucket).
		Connect()
}

func MustInit(region, bucket string) {
	if err := Init(region, bucket); err != nil {
		panic(err)
	}
}

func CreateFile(ctx context.Context, path string, body []byte, acl ...string) error {
	return getService().CreateFile(ctx, path, body, acl...)
}

func CreateFolder(ctx context.Context, path string, acl ...string) error {
	return getService().CreateFolder(ctx, path, acl...)
}

func Delete(ctx context.Context, path string) error {
	return getService().Delete(ctx, path)
}

func GetAll(ctx context.Context) ([]Object, error) {
	objects, err := getService().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return entityToModel(objects), nil
}

func GetPath(ctx context.Context, path string) ([]Object, error) {
	objects, err := getService().GetPath(ctx, path)
	if err != nil {
		return nil, err
	}

	return entityToModel(objects), nil
}

func GetFile(ctx context.Context, path string) ([]byte, error) {
	return getService().GetFile(ctx, path)
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
