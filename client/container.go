package client

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~akilan1999/p2p-rendering-computation/server/docker"
	"io/ioutil"
	"net/http"
)

const (
	serverPort = "8089"
)


var client = http.Client{}


// Start container using REST api Implementation
// From the selected server IP address
func StartContainer(Ip string) (*docker.DockerVM ,error) {
	URL := "http://" + Ip + ":8088/startcontainer"
	resp, err := http.Get(URL)

	// Convert response to byte value
	byteValue, _ := ioutil.ReadAll(resp.Body)

	// Create variable for result response type
	var dockerResult docker.DockerVM

	// Adds byte value to docker.DockerVM struct
	json.Unmarshal(byteValue, &dockerResult)
	if err != nil {
		return nil,err
	}

	return &dockerResult, nil
}

// Prints results Generated container
func PrintStartContainer(d *docker.DockerVM){
	fmt.Println("ID : " + fmt.Sprint(d.ID))
	fmt.Println("SSH port: " + fmt.Sprint(d.SSHPort))
	fmt.Println("SSH username: " + fmt.Sprint(d.SSHUsername))
	fmt.Println("SSH password: " + fmt.Sprint(d.SSHPassword))
	fmt.Println("VNC port: " + fmt.Sprint(d.VNCPort))
	fmt.Println("VNC password: " + fmt.Sprint(d.VNCPassword))
	fmt.Println("Ports Open")
	for i := range d.Ports {
		fmt.Println(d.Ports[i])
	}
}


// TODO implementation using RPC calls
//func StartContainer(Ip string) (*docker.DockerVM,error){
//	client, err := rpc.DialHTTP("tcp",Ip + ":" + serverPort)
//
//	if err != nil {
//		return nil,err
//	}
//
//	in := bufio.NewReader(os.Stdin)
//
//	for {
//		line, _, err := in.ReadLine()
//		if err != nil {
//			return nil,err
//		}
//
//		var reply Docker
//
//		err = client.Call("Listener.StartContainer", line, &reply)
//		if err != nil {
//			return nil,err
//		}
//		log.Printf("Reply: %v, Data: %v", reply, reply.docker)
//		return reply.docker, nil
//	}
//
//}