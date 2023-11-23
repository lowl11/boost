package elastic

import (
	"context"
	"github.com/lowl11/boost/data/entity"
	"github.com/lowl11/boost/internal/services/elk_service"
	"github.com/lowl11/boost/log"
)

func Init(host, username, password string) error {
	err := elk_service.
		Get(host).
		SetAuth(username, password).Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func MustInit(host, username, password string) {
	if err := Init(host, username, password); err != nil {
		log.Fatal(err, "Initialize Elasticsearch error")
	}
}

func DeleteIndex(ctx context.Context, indexName string) error {
	return elk_service.Get().DeleteIndex(ctx, indexName)
}

func CreateIndex(ctx context.Context, indexName string, object any, config ...entity.ElasticIndexConfig) error {
	return elk_service.Get().CreateIndex(ctx, indexName, object, config...)
}

func GetIndices(ctx context.Context) ([]entity.ElasticIndex, error) {
	return elk_service.Get().GetIndices(ctx)
}

func BindAlias(ctx context.Context, pairs ...entity.ElasticAliasPair) error {
	if len(pairs) == 0 {
		return nil
	}

	return elk_service.Get().BindAlias(ctx, pairs)
}

func Insert(ctx context.Context, indexName string, object any) error {
	return elk_service.Get().Insert(ctx, indexName, object)
}

func Delete(ctx context.Context, indexName string, id string) error {
	return elk_service.Get().Delete(ctx, indexName, id)
}

func GetAll(ctx context.Context, indexName string, export any) error {
	return elk_service.Get().GetAll(ctx, indexName, export)
}

func Search(ctx context.Context, indexName string, query map[string]any, export any) error {
	return elk_service.Get().Search(ctx, indexName, query, export)
}

func Exist(ctx context.Context, indexName string) (bool, error) {
	return elk_service.Get().Exist(ctx, indexName)
}
