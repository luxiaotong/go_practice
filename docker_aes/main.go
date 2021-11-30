package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// ctx := context.Background()
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	panic(err)
	// }
	// done := make(chan struct{})
	// idResp, err := cli.ContainerExecCreate(ctx, "aes", types.ExecConfig{
	// 	Cmd:          []string{"java", "-cp", "commons-lang3-3.0.jar:.", "Main", "20211117001", "31c241309a9231f585bca20c9873b49a"},
	// 	AttachStdout: true,
	// 	Tty:          true,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("id resp: ", idResp.ID)
	// resp, err := cli.ContainerExecAttach(ctx, idResp.ID, types.ExecStartCheck{Tty: true})
	// if err != nil {
	// 	panic(err)
	// }
	// go func() {
	// 	defer resp..Close()
	// 	if _, err := stdcopy.StdCopy(os.Stdout, os.Stderr, resp.Reader); err != nil {
	// 		panic(err)
	// 	}
	// 	close(done)
	// }()
	// <-done
	// fmt.Printf("aes return %d bytes\n", n)
	// docker exec -it aes java -cp commons-lang3-3.0.jar:. Main 20211117001 31c241309a9231f585bca20c9873b49a
	// cmd := exec.Command("docker", "exec", "-it", "aes", "java", "-cp", "commons-lang3-3.0.jar:.", "Main", "20211117001", "31c241309a9231f585bca20c9873b49a")
	cmd := exec.Command("/bin/sh", "-c", "docker exec aes java -cp commons-lang3-3.0.jar:. Main 20211117001 31c241309a9231f585bca20c9873b49a")
	out, e := cmd.Output()
	if e != nil {
		panic(e)
	}
	fmt.Printf("output: %v\n", string(out))
}
