package main

import (
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"

	"github.com/conductor-sdk/conductor-go/sdk/worker"
)


const SERVER_URL = "_server_url_"
const KEY_ID = "_auth_key_"
const SECRET = "_auth_secret_"
const TASK_NAME = "_taskname_"

func workerFn(task *model.Task) (any, error) {
	return map[string]interface{}{
		"greetings": "Hello " + fmt.Sprintf("%v", task.InputData["name"]),
	}, nil
}

var (
	apiClient = client.NewAPIClient(
		settings.NewAuthenticationSettings(KEY_ID, SECRET),
		settings.NewHttpSettings(SERVER_URL))
	taskRunner = worker.NewTaskRunnerWithApiClient(apiClient)
)

func main() {
	taskRunner.StartWorker(TASK_NAME, workerFn, 1, time.Millisecond*200)
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
