package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/peer"
	"io/ioutil"
	"log"
	"os"
)

func main()  {

	arg := os.Args

	if len(arg) < 3 {
		log.Fatal("Not enough arguments")
	}

	if arg[1] == "" {
		log.Fatal("Provide path to chaincode package")
	}

	data, err := ioutil.ReadFile(arg[1])

	if err != nil {
		log.Fatal(err)
	}

	ccSpec := &peer.ChaincodeDeploymentSpec{}

	err = proto.Unmarshal(data, ccSpec)

	if err != nil {
		log.Fatal(err)
	}

	filePath := arg[2]
	if filePath == "" {
		log.Fatal("Provide path where to save connection.json")
	}

	err = os.MkdirAll(filePath, 0644)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/connection.json", filePath), ccSpec.CodePackage, 0644)

	if err != nil {
		log.Fatal(err)
	}

}



