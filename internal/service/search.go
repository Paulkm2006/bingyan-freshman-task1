package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"

	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/dto"
)

var es *elasticsearch.Client

func InitES() {
	cfg := elasticsearch.Config{
		Addresses: []string{config.Config.ES.Host},
		Username:  config.Config.ES.Username,
		Password:  config.Config.ES.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !config.Config.ES.VerifyTls,
			},
		},
	}
	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
}

func IndexPost(post *dto.Post, content *string) error {

	body := map[string]interface{}{
		"pid":     post.PID,
		"title":   post.Title,
		"content": content,
		"uid":     post.UID,
		"created": post.Created,
	}

	jsonString, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "posts",
		DocumentID: strconv.FormatInt(int64(post.PID), 10),
		Body:       strings.NewReader(string(jsonString)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.String())
	}

	return nil
}

func SearchPost(keyword string) ([]dto.Post, error) {

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  "*" + keyword + "*",
				"fields": []string{"title", "content"},
				"type":   "phrase_prefix",
			},
		},
	}

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("posts"),
		es.Search.WithBody(strings.NewReader(string(jsonQuery))),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	posts := make([]dto.Post, 0, len(hits))

	for _, hit := range hits {
		source := hit.(map[string]interface{})
		pid, _ := strconv.ParseInt(hit.(map[string]interface{})["_id"].(string), 10, 64)

		post := dto.Post{
			PID:   int(pid),
			Title: source["_source"].(map[string]interface{})["title"].(string),
			UID:   int(source["_source"].(map[string]interface{})["uid"].(float64)),
			Created: func() time.Time {
				createdStr := source["_source"].(map[string]interface{})["created"].(string)
				createdTime, _ := time.Parse(time.RFC3339, createdStr)
				return createdTime
			}(),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
