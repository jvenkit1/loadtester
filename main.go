package main

import (
	"github.com/sirupsen/logrus"
	"loadtester/client"
)

func main(){
	logrus.Info("Hello World")
	url := "http://localhost:8080/get"
	client.GetRequest(url)
}