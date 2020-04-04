package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luxiaotong/go_practice/grpc_example/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			log.Fatal("Param a error: ", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A "})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			log.Fatal("Param b error: ", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B "})
			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprintf("%d", response.Result),
			})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})

	g.GET("/multiply/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			log.Fatal("Param a error: ", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A "})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			log.Fatal("Param b error: ", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B "})
			return
		}

		req := &proto.Request{A: int64(a), B: int64(b)}

		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprintf("%d", response.Result),
			})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
