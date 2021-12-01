package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// docker run --rm -v /Users/luxiaotong/code/go_practice/docker_wkhtml2pdf/fonts:/usr/share/fonts -v /Users/luxiaotong/code/go_practice/docker_wkhtml2pdf/data:/data madnight/docker-alpine-wkhtmltopdf /data/sample.html - > sample2.pdf
// docker run --rm -v $(pwd)/fonts:/usr/share/fonts -v $(pwd)/data:/data madnight/docker-alpine-wkhtmltopdf /data/sample.html - > sample2.pdf

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "madnight/docker-alpine-wkhtmltopdf:latest",
		Cmd:   []string{"/data/sample.html", "-", "/data/sample.pdf"},
	}, &container.HostConfig{
		Binds:      []string{"/Users/luxiaotong/code/go_practice/docker_wkhtml2pdf/data:/data", "/Users/luxiaotong/code/go_practice/docker_wkhtml2pdf/fonts:/usr/share/fonts"},
		AutoRemove: true,
	}, nil, nil, "wkhtml2pdf")
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	defer func() {
		_ = cli.ContainerStop(ctx, resp.ID, nil)
	}()
	fmt.Println("Done!")
}
