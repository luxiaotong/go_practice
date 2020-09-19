package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
)

func GInitElastic() error {
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("Error creating the client: %s", err)
		return err
	}

	res, err := es.Info()
	if err != nil {
		fmt.Printf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()
	fmt.Println(res)

	return nil
}

func GSearch() error {
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return errors.Errorf("error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  "歌词",
				"fields": []string{"user_name", "content"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return errors.Errorf("error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("mojito_jay_zhou"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return errors.Errorf("error encoding query: %s", err)
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return errors.Errorf("error parsing the response body: %s", err)
	}

	if res.IsError() {
		return errors.Errorf("index error: [%s] %s: %s", res.Status(),
			r["error"].(map[string]interface{})["type"],
			r["error"].(map[string]interface{})["reason"],
		)
	}

	// Print the response status, number of results, and request duration.
	fmt.Printf("[%s] %d hits; took: %dms\n", res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		fmt.Printf(" * ID=%s, %s\n", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	return nil
}

func GIndex(comment Comment) error {
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return errors.Errorf("error creating the client: %s", err)
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(comment); err != nil {
		return errors.Errorf("error encoding query: %s", err)
	}
	// fmt.Printf("buf: %s\n", buf.Bytes())

	req := esapi.IndexRequest{
		Index:        "mojito_jay_zhou",
		DocumentType: "comment",
		DocumentID:   strconv.Itoa(comment.PKID),
		Body:         bytes.NewReader(buf.Bytes()),
		Refresh:      "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("[%s] error indexing document ID=%d", res.Status(), comment.PKID)
	}

	// Deserialize the response into a map.
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return errors.Errorf("Error parsing the response body: %s", err)
	}

	// Print the response status and indexed document version.
	log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
	return nil
}

func GGet(id string) error {
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return errors.Errorf("error creating the client: %s", err)
	}

	res, err := es.Get(
		"mojito_jay_zhou",
		id,
		es.Get.WithPretty(),
	)
	if err != nil {
		return errors.Errorf("error getting the response: %s", err)
	}
	defer res.Body.Close()

	// Deserialize the response into a map.
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return errors.Errorf("error parsing the response body: %s", err)
	}

	// Print the response status and indexed document version.
	fmt.Printf("got document %s in version %d from index %s, type %s\n", r["_id"], int(r["_version"].(float64)), r["_index"], r["_type"])
	return nil
}

func GUpdate(comment Comment) error {
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return errors.Errorf("error creating the client: %s", err)
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(comment); err != nil {
		return errors.Errorf("error encoding query: %s", err)
	}

	res, err := es.Index(
		"mojito_jay_zhou",
		bytes.NewReader(buf.Bytes()),
		es.Index.WithDocumentID(strconv.Itoa(comment.PKID)),
		es.Index.WithPretty(),
	)
	if err != nil {
		return errors.Errorf("error getting the response: %s", err)
	}
	defer res.Body.Close()

	// Deserialize the response into a map.
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return errors.Errorf("error parsing the response body: %s", err)
	}
	fmt.Printf("update vote count %s, version: %d\n", r["result"], int(r["_version"].(float64)))
	return nil
}

func GDelete(id string) error {
	cfg := elasticsearch.Config{
		Addresses: []string{elasticUrl},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return errors.Errorf("error creating the client: %s", err)
	}

	res, err := es.DeleteByQuery(
		[]string{"mojito_jay_zhou"},
		strings.NewReader(fmt.Sprintf(`{
			"query": {
				"term": {
					"pk_id": %s
				}
			}
		}`, id)),
	)
	if err != nil {
		return errors.Errorf("Error getting the response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res, err)
	return nil
}
