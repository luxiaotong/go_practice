package search

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

const elasticUrl = "http://139.9.119.21:59200/"

type Comment struct {
	PKID        int    `json:"pk_id"`
	UserURL     string `json:"user_url"`
	CommentTime string `json:"comment_time"`
	RatingNum   string `json:"rating_num"`
	Content     string `json:"content"`
	UserName    string `json:"user_name"`
	VoteCount   int    `json:"vote_count"`
}

func InitElastic() {
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(elasticUrl).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(elasticUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

func Search(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	if err != nil {
		panic(err)
	}
	fmt.Printf("client: %v\n", client)

	query := elastic.NewMatchQuery("content", "歌词")
	searchResult, err := client.Search().
		Index("mojito_jay_zhou").
		Type("comment").
		Query(query).
		Size(size).
		From((page - 1) * size).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var comment Comment
	for _, item := range searchResult.Each(reflect.TypeOf(comment)) {
		if c, ok := item.(Comment); ok {
			fmt.Printf("%d, Comment by %s(%s): %s\n", c.PKID, c.UserName, c.RatingNum, c.Content)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d comments\n", searchResult.TotalHits())
}

func Index(comment Comment) error {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	if err != nil {
		return err
	}
	idx, err := client.Index().
		Index("mojito_jay_zhou").
		Type("comment").
		Id(strconv.Itoa(comment.PKID)).
		BodyJson(comment).
		Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Indexed comment %s to index %s, type %s\n", idx.Id, idx.Index, idx.Type)
	return nil
}

func Get(id string) bool {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	if err != nil {
		return false
	}
	get, err := client.Get().
		Index("mojito_jay_zhou").
		Type("comment").
		Id(id).
		Do(ctx)
	if err != nil {
		return false
	}
	if get.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, *get.Version, get.Index, get.Type)
	}
	return get.Found
}

func Update(id string, voteCount int) error {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	res, err := client.Update().
		Index("mojito_jay_zhou").
		Type("comment").
		Id(id).
		Doc(map[string]interface{}{"vote_count": voteCount}).
		Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("update vote count %s\n", res.Result)
	return nil
}

func Delete(id string) error {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(elasticUrl))
	res, err := client.Delete().Index("mojito_jay_zhou").
		Type("comment").
		Id(id).
		Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("delete result %s\n", res.Result)
	return nil
}
