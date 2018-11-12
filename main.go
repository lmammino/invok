package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"

	"github.com/aws/aws-lambda-go/lambda/messages"
)

var host = flag.String("host", "localhost:1234", "The host where the lambda is running")
var ping = flag.Bool("ping", false, "Set the flag if you want to issue just a ping request")

func main() {
	flag.Parse()

	c, err := rpc.Dial("tcp", *host)
	if err != nil {
		log.Fatalln(err)
	}

	if *ping == true {
		var result messages.PingResponse
		request := &messages.PingRequest{}
		err = c.Call("Function.Ping", request, &result)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		os.Exit(0)
	}

	data, _ := ioutil.ReadAll(bufio.NewReader(os.Stdin))

	var result messages.InvokeResponse
	request := &messages.InvokeRequest{
		Payload: data,
	}
	err = c.Call("Function.Invoke", request, &result)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	if result.Error != nil {
		fmt.Fprintf(os.Stderr, "%s\n", result.Error.Message)
		os.Exit(1)
	}

	fmt.Printf("result: %v\n", string(result.Payload))
	os.Exit(0)
}
