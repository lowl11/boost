package elk_service

import (
	"encoding/json"
	"github.com/lowl11/boost/data/entity"
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/internal/helpers/elk_parser"
	"net/http"
)

func (service *Service) SetAuth(username, password string) *Service {
	service.client.SetBasicAuth(username, password)
	return service
}

func (service *Service) GetIndices() ([]entity.ElasticIndex, error) {
	var indices []entity.ElasticIndex

	response, err := service.client.
		R().
		SetResult(&indices).
		GET("/_cat/indices?format=json")
	if err != nil {
		return nil, ErrorGetAllIndices(err)
	}

	if response.StatusCode() != http.StatusOK {
		return nil, errors.
			New("Status is not 200").
			AddContext("body", response.Body())
	}

	filtered := make([]entity.ElasticIndex, 0, len(indices))
	for _, index := range indices {
		if index.Index[0] == '.' {
			continue
		}

		filtered = append(filtered, index)
	}

	return filtered, nil
}

func (service *Service) CreateIndex(indexName string, object any, config ...entity.ElasticIndexConfig) error {
	// check for given type
	mappings, err := elk_parser.ParseObject(object)
	if err != nil {
		return err
	}

	// build request object
	request := createIndexRequest{
		Mappings: indexMappings{
			Properties: mappings,
		},
	}

	if len(config) > 0 {
		request.Settings.Index.NumberOfReplicas = config[0].NumberOfReplicas
		request.Settings.Index.NumberOfShards = config[0].NumberOfShards
	} else {
		request.Settings.Index.NumberOfReplicas = 2
		request.Settings.Index.NumberOfShards = 3
	}

	// send request
	response, err := service.client.
		R().
		SetBody(request).
		PUT("/" + indexName)
	if err != nil {
		return ErrorCreateIndex(err)
	}

	// check response code
	if response.StatusCode() != http.StatusCreated && response.StatusCode() != http.StatusOK {
		return errors.
			New("Status is not 201").
			AddContext("body", response.Body()).
			AddContext("status", response.StatusCode())
	}

	return nil
}

func (service *Service) DeleteIndex(indexName string) error {
	response, err := service.client.
		R().
		DELETE("/" + indexName)
	if err != nil {
		return ErrorDeleteIndex(err)
	}

	if response.StatusCode() == http.StatusNotFound {
		return nil
	}

	if response.StatusCode() != http.StatusOK {
		return errors.
			New("Status is not 200").
			AddContext("body", response.Body())
	}

	return nil
}

func (service *Service) BindAlias(pairs ...entity.ElasticAliasPair) error {
	request := bindAliasRequest{
		Actions: make([]entity.ElasticAliasAdd, 0, len(pairs)),
	}

	for _, pair := range pairs {
		request.Actions = append(request.Actions, entity.ElasticAliasAdd{
			Add: pair,
		})
	}

	response, err := service.client.
		R().
		SetBody(request).
		POST("/_aliases")
	if err != nil {
		return ErrorBindAlias(err)
	}

	if response.StatusCode() != http.StatusOK {
		return errors.
			New("Status is not 200").
			AddContext("body", response.Body())
	}

	return nil
}

func (service *Service) Insert(indexName string, object any) error {
	id, err := elk_parser.GetID(object)
	if err != nil {
		return err
	}

	response, err := service.client.
		R().
		SetBody(object).
		POST("/" + indexName + "/_doc/" + id)
	if err != nil {
		return ErrorInsertData(err)
	}

	if response.StatusCode() != http.StatusCreated {
		return errors.
			New("Status is not 201").
			AddContext("body", response.Body())
	}

	return nil
}

func (service *Service) Delete(indexName string, id string) error {
	response, err := service.client.
		R().
		DELETE("/" + indexName + "/_doc/" + id)
	if err != nil {
		return ErrorInsertData(err)
	}

	if response.StatusCode() != http.StatusOK {
		return errors.
			New("Status is not 200").
			AddContext("body", response.Body())
	}

	return nil
}

func (service *Service) GetAll(indexName string, export any) error {
	result := searchResult{}

	response, err := service.client.
		R().
		SetBody(map[string]any{
			"query": map[string]any{
				"match_all": map[string]string{},
			},
		}).
		SetResult(&result).
		POST("/" + indexName + "/_search")
	if err != nil {
		return ErrorGetAllDocuments(err)
	}

	if response.StatusCode() != http.StatusOK {
		return errors.
			New("Status is not 200").
			AddContext("body", response.Body())
	}

	sourceHits := make([]map[string]any, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		sourceHits = append(sourceHits, hit.Source)
	}

	sourceInBytes, err := json.Marshal(sourceHits)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(sourceInBytes, &export); err != nil {
		return err
	}

	return nil
}
