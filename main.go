package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

var testcases = map[string]interface{}{
	"cel-app": run.InitializedTestCaseFn(runSync),
}

func main() {
	run.InvokeMap(testcases)
}

func runSync(runenv *runtime.RunEnv, initCtx *run.InitContext) error {
	// cmd := exec.Command(
	// 	"ls", "-la", "appconfig",
	// )
	cmd := exec.Command(
		"./celestia-appd",
		"start",
		"--moniker", "core0",
		"--home", "/appconfig/core0",
	)
	// "start",
	// "--moniker", "core0",
	// "--home", "celestia-app/core0",
	// )

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}
	time.Sleep(30 * time.Second)
	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	fmt.Printf("%s\n", string(data))

	return nil
}
