package catalog

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esutil"
)

type Repository interface {
	Close()
	PutProduct(ctx context.Context, p Product) error
	GetProductByID(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type elasticRepository struct {
	client *elasticsearch.Client
}

type productDocument struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewElasticRepository(url string) (Repository, error) {
	client, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{url},
		},
	)
	if err != nil {
		return nil, err
	}
	return &elasticRepository{client: client}, nil
}

func (r *elasticRepository) Close() {
	r.client.
}

func (r *elasticRepository) PutProduct(ctx context.Context, p Product) error {
	doc := productDocument{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
	_, err := r.client.Index(
		"catalog",
		esutil.NewJSONReader(&doc),
		r.client.Index.WithDocumentID(p.ID),
		r.client.Index.WithContext(ctx),
	)
	return err
}

func (r *elasticRepository) GetProductByID(ctx context.Context, id string) (*Product, error) {
	res, err := r.client.Get(
		"catalog",
		id,
		r.client.Get.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, nil // Not found or error
	}

	var doc struct {
		Source productDocument `json:"_source"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, err
	}

	return &Product{
		ID:          id,
		Name:        doc.Source.Name,
		Description: doc.Source.Description,
		Price:       doc.Source.Price,
	}, nil
}

func (r *elasticRepository) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	res,err:=r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("catalog"),
		r.client.Search.WithFrom(int(skip)),
		r.client.Search.WithSize(int(take)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	if res.IsError() {
		return nil, nil // Not found or error
	}
	var rj struct {
		Hits struct {
			Hits []struct {
				ID     string          `json:"_id"`
				Source productDocument `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&rj); err != nil {
		return nil, err
	}
	products := make([]Product, len(rj.Hits.Hits))
	for i, hit := range rj.Hits.Hits {
		products[i] = Product{
			ID:          hit.ID,
			Name:        hit.Source.Name,
			Description: hit.Source.Description,
			Price:       hit.Source.Price,
		}
	}
	return products, nil

}

func (r *elasticRepository) ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	res, err := r.client.Mget(
		esutil.NewJSONReader(map[string]interface{}{
			"ids": ids,
		}),
		r.client.Mget.WithIndex("catalog"),
		r.client.Mget.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, nil // Not found or error
	}

	var rj struct {
		Docs []struct {
			ID     string          `json:"_id"`
			Found  bool            `json:"found"`
			Source productDocument `json:"_source"`
		} `json:"docs"`
	}
	if err := json.NewDecoder(res.Body).Decode(&rj); err != nil {
		return nil, err
	}

	products := make([]Product, 0, len(rj.Docs))
	for _, doc := range rj.Docs {
		if doc.Found {
			products = append(products, Product{
				ID:          doc.ID,
				Name:        doc.Source.Name,
				Description: doc.Source.Description,
				Price:       doc.Source.Price,
			})
		}
	}

	return products, nil
}

func (r *elasticRepository) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	res, err := r.client.Search(
		esutil.NewJSONReader(map[string]interface{}{
			"query": map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":  query,
					"fields": []string{"name", "description"},
				},
			},
		}),
		r.client.Search.WithIndex("catalog"),
		r.client.Search.WithFrom(int(skip)),
		r.client.Search.WithSize(int(take)),
		r.client.Search.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, nil // Not found or error
	}

	var rj struct {
		Hits struct {
			Hits []struct {
				ID     string          `json:"_id"`
				Source productDocument `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&rj); err != nil {
		return nil, err
	}

	products := make([]Product, len(rj.Hits.Hits))
	for i, hit := range rj.Hits.Hits {
		products[i] = Product{
			ID:          hit.ID,
			Name:        hit.Source.Name,
			Description: hit.Source.Description,
			Price:       hit.Source.Price,
		}
	}

	return products, nil
}
