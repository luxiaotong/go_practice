package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	// docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v /Users/luxiaotong/code/go_practice/docker_postgres/data:/var/lib/postgresql/data -d postgres
	port := "5432/tcp"
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "postgres:latest",
		Env:   []string{"POSTGRES_PASSWORD=postgres"},
		ExposedPorts: nat.PortSet{
			nat.Port(port): {},
		},
		Healthcheck: &container.HealthConfig{
			Test:     []string{"CMD-SHELL", "pg_isready -U postgres"},
			Interval: 5 * time.Second,
			Retries:  0,
		},
	}, &container.HostConfig{
		PublishAllPorts: true,
		Binds:           []string{"/Users/luxiaotong/code/go_practice/docker_postgres/data:/var/lib/postgresql/data"},
		AutoRemove:      true,
	}, nil, nil, "postgres_test")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	defer func() {
		_ = cli.ContainerStop(ctx, resp.ID, nil)
	}()

	inspect, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	hostPort := inspect.NetworkSettings.Ports[nat.Port(port)][0].HostPort
	fmt.Println("health: ", inspect.State.Health)

	connStr := fmt.Sprintf("user=postgres password=postgres host=127.0.0.1 port=%s sslmode=disable", hostPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			fmt.Println("ping failed", err)
			time.Sleep(5 * time.Second)
		}
	}
	if _, err = db.Exec("CREATE USER test"); err != nil {
		fmt.Println("exec failed", err)
		return
	}
	if _, err = db.Exec("ALTER USER test WITH PASSWORD 'datassets'"); err != nil {
		fmt.Println("password failed", err)
		return
	}
	if _, err = db.Exec("CREATE DATABASE test"); err != nil {
		fmt.Println("create db failed", err)
		return
	}
	if _, err = db.Exec("GRANT ALL PRIVILEGES ON DATABASE test TO test"); err != nil {
		fmt.Println("create db failed", err)
		return
	}

	db2, err := sql.Open("postgres", fmt.Sprintf("user=test password=datassets host=127.0.0.1 port=%s sslmode=disable", hostPort))
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()
	if _, err = db2.Exec("create table city(id int,name text)"); err != nil {
		fmt.Println("create table failed", err)
		return
	}
	fmt.Println("container")
}
