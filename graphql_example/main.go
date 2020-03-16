package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/luxiaotong/go_practice/graphql_example/blockchain"
    "os"
    "github.com/graphql-go/graphql"
)

var chain *blockchain.BlockChain

var blockType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Block",
        Fields: graphql.Fields{
            "PrevHash": &graphql.Field{
                Type: graphql.String,
            },
            "Hash": &graphql.Field{
                Type: graphql.String,
            },
            "Data": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        /* Get (read) blockchain
           http://localhost:8080/product?query={list{PrevHash,Hash,Data}}
        */
        "list": &graphql.Field{
            Type:        graphql.NewList(blockType),
            Description: "Get block list",
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                iter := chain.Iterator()
                var result []interface{}
                for {
                    block := iter.Next()
                    prevHash := fmt.Sprintf("%x", block.PrevHash)
                    hash := fmt.Sprintf("%x", block.Hash)
                    data := fmt.Sprintf("%s", block.Data)
                    row := map[string]string {
                        "PrevHash": prevHash,
                        "Hash": hash,
                        "Data": data,
                    }
                    result = append(result, row)
                    fmt.Printf("PrevHash: %x\n", block.PrevHash)
                    fmt.Printf("Hash: %x\n", block.Hash)
                    fmt.Printf("Data: %s\n", block.Data)
                    if len(block.PrevHash) <= 0 {
                        break
                    }
                }
                return result, nil
            },
        },
    },
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Mutation",
    Fields: graphql.Fields{
        /* Create new block
        http://localhost:8080/graphql?query=mutation+_{create(data:"first block"){data}}
        */
        "create": &graphql.Field{
            Type:        blockType,
            Description: "Create new block",
            Args: graphql.FieldConfigArgument{
                "data": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                data, _ := params.Args["data"].(string)
                chain.AddBlock(data)

                iter := chain.Iterator()
                block := iter.Next()

                prevHash := fmt.Sprintf("%x", block.PrevHash)
                hash := fmt.Sprintf("%x", block.Hash)
                data = fmt.Sprintf("%s", block.Data)

                result := map[string]string {
                    "PrevHash": prevHash,
                    "Hash": hash,
                    "Data": data,
                }
                fmt.Printf("PrevHash: %x\n", block.PrevHash)
                fmt.Printf("Hash: %x\n", block.Hash)
                fmt.Printf("Data: %s\n", block.Data)

                return result, nil
            },
        },
    },
})
var schema, _ = graphql.NewSchema(
    graphql.SchemaConfig{
        Query:    queryType,
        Mutation: mutationType,
    },
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
    result := graphql.Do(graphql.Params{
        Schema:        schema,
        RequestString: query,
    })
    if len(result.Errors) > 0 {
        fmt.Printf("errors: %v", result.Errors)
    }
    return result
}

func main() {
    defer os.Exit(0)
    fmt.Println("hello,world")

    chain = blockchain.InitBlockChain()
    defer chain.Database.Close()

    http.HandleFunc("/rest_print", func(w http.ResponseWriter, r *http.Request) {
        iter := chain.Iterator()
        var result []interface{}
        for {
            block := iter.Next()
            prevHash := fmt.Sprintf("%x", block.PrevHash)
            hash := fmt.Sprintf("%x", block.Hash)
            data := fmt.Sprintf("%s", block.Data)
            row := map[string]string {
                "PrevHash": prevHash,
                "Hash": hash,
                "Data": data,
            }
            result = append(result, row)
            fmt.Printf("PrevHash: %x\n", block.PrevHash)
            fmt.Printf("Hash: %x\n", block.Hash)
            fmt.Printf("Data: %s\n", block.Data)
            if len(block.PrevHash) <= 0 {
                break
            }
        }
        json.NewEncoder(w).Encode(result)
    })
    http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
        result := executeQuery(r.URL.Query().Get("query"), schema)
        json.NewEncoder(w).Encode(result)
    })

    fmt.Println("Now server is running on port 8080")
    fmt.Println("Test with Get      : curl -g 'http://localhost:8080/rest_print'")
    fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={list{PrevHash,Hash,Data}}'")
    fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query=mutation+_{create(data:\"first block\"){Data}}'")

    http.ListenAndServe(":8080", nil)
}
