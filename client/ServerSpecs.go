package client

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~akilan1999/p2p-rendering-computation/server"
	"io/ioutil"
	"net/http"
)

func GetSpecs(IP string)(*server.SysInfo,error) {
	URL := "http://" + IP + ":" + serverPort + "/server_info"
	resp, err := http.Get(URL)
	if err != nil {
		return nil,err
	}

	// Convert response to byte value
	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	// Create variable for result response type
	var serverSpecsResult server.SysInfo

	// Adds byte value to docker.DockerVM struct
	json.Unmarshal(byteValue, &serverSpecsResult)
	if err != nil {
		return nil,err
	}

	return &serverSpecsResult, nil
}

// print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}