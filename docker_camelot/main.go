package main

import (
	"fmt"
	"time"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	cli, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	listener := make(chan *docker.APIEvents, 3)
	filters := map[string][]string{
		"type":  {"container"},
		"event": {"exec_create", "exec_start", "exec_die"},
	}
	// // now := time.Now().Unix()
	eo := docker.EventsOptions{
		// 	// Since:   "1374067970",
		// 	// Until:   "1442421700",
		Filters: filters,
	}
	if err = cli.AddEventListenerWithOptions(eo, listener); err != nil {
		panic(err)
	}

	exec, err := cli.CreateExec(docker.CreateExecOptions{
		Container: "asset_parser",
		// Container:    "zen_carson",
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: false,
		Tty:          true,
		Cmd:          []string{"python3", "/app/parse_pdf.py", "/medias/SetDatassetsApply.pdf", "/medias/SetDatassetsApply.pdf.json"},
		// Cmd: []string{"echo", "hello world"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("exec id: ", exec.ID)
	err = cli.StartExec(exec.ID, docker.StartExecOptions{
		Detach: false,
		Tty:    true,
	})
	if err != nil {
		panic(err)
	}

	timeout := time.After(500 * time.Second)
loop:
	for i := range []int{0, 1, 2} {
		select {
		case msg, ok := <-listener:
			if !ok {
				break loop
			}
			fmt.Println("msg action: ", msg.Action)
			if msg.Action == "exec_die" {
				break loop
			}
		case <-timeout:
			fmt.Printf("echo: timed out waiting on events after %d events\n", i)
			break loop
		}
	}

}
