package elastic

import (
	"context"
	"github.com/lowl11/boost/data/entity"
	"github.com/lowl11/boost/log"
)

func Init(host, username, password string) error {
	err := getService(host).
		SetAuth(username, password).
		Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func MustInit(host, username, password string) {
	if err := Init(host, username, password); err != nil {
		log.Fatal("Initialize Elasticsearch error:", err)
	}
}

func DeleteIndex(ctx context.Context, indexName string) error {
	return getService().DeleteIndex(ctx, indexName)
}

func CreateIndex(ctx context.Context, indexName string, object any, config ...entity.ElasticIndexConfig) error {
	return getService().CreateIndex(ctx, indexName, object, config...)
}

func GetIndices(ctx context.Context) ([]entity.ElasticIndex, error) {
	return getService().GetIndices(ctx)
}

func BindAlias(ctx context.Context, pairs ...entity.ElasticAliasPair) error {
	if len(pairs) == 0 {
		return nil
	}

	return getService().BindAlias(ctx, pairs)
}

func Insert(ctx context.Context, indexName string, object any) error {
	return getService().Insert(ctx, indexName, object)
}

func Delete(ctx context.Context, indexName string, id string) error {
	return getService().Delete(ctx, indexName, id)
}

func GetAll(ctx context.Context, indexName string, export any) error {
	return getService().GetAll(ctx, indexName, export)
}

func Search(ctx context.Context, indexName string, query map[string]any, export any) error {
	return getService().Search(ctx, indexName, query, export)
}

func Exist(ctx context.Context, indexName string) (bool, error) {
	return getService().Exist(ctx, indexName)
}

func MatchAll() map[string]any {
	return map[string]any{
		"query": map[string]any{
			"match_all": map[string]any{},
		},
	}
}

func MultiMatch(query string, fields []string) map[string]any {
	return map[string]any{
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":  query,
				"fields": fields,
			},
		},
	}
}

func QueryBool() *BoolQuery {
	return queryBool()
}
