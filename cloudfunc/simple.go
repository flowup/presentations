package main

import (
	"github.com/flowup/cloudfunc/api"
)

func main() {
	input := make(map[string]interface{})

	cloudFunction := api.NewCloudFunc()

	// parse request and its body
	req, _ := cloudFunction.GetRequest()
	req.BindBody(&input)

	// send it back
	cloudFunction.SendResponse(&input)
}
