package client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"time"
)

// Main Get handler. Performs the operation depending on the rate per second obtained
func Get(url string, numRequests int, time int){
	// calculate rate and send that per second
	rate := numRequests / time

	ch := make(chan string)

	for i:=0;i<rate;i++{
		go getRequest(url, ch)
	}

	// printing metadata
	for i:=0;i<rate;i++{
		logrus.Infof("%s", <-ch)
	}

}

// Main POST handler. Makes POST request depending on the rate per second provided.
func Post(url string, numRequests int, time int, bodyPath string) {
	jsonFile, err := ioutil.ReadFile(bodyPath)
	if err!=nil {
		logrus.WithError(err).Fatal("Unable to open file")
	}

	rate := numRequests / time

	ch := make(chan string)

	for i:=0;i<rate;i++ {
		go postRequest(url, jsonFile, ch)
	}
	// printing metadata
	for i:=0;i<rate;i++{
		logrus.Infof("%s", <-ch)
	}
}

// sends a get request. to be used a a goroutine
func getRequest(url string, ch chan<-string) {

	start := time.Now()

	// Create Pipeline Client
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	//client := &fasthttp.PipelineClient{
	//	MaxConns: rate,
	//	Addr: "localhost:8080",
	//}
	client := &fasthttp.Client{}

	err := client.Do(req, resp)
	if err != nil {
		logrus.WithError(err).Fatal("Can't send http request")
	}

	body := resp.Body()
	timeElapsed := time.Since(start).Seconds()
	logrus.WithFields(logrus.Fields{
		"Body": string(body),
	}).Info("Metadata")
	ch<-fmt.Sprintf("%.2f elapsed with response length: %d %s", timeElapsed, len(body), url)
}

// send a post request. to be used as a goroutine
func postRequest(url string, jsonFile []byte, ch chan<-string){
	start := time.Now()
	// read the post body
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.SetBody(jsonFile)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)


	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		logrus.WithError(err).Fatal("Error sending http request")
	}

	body := resp.Body()
	timeElapsed := time.Since(start).Seconds()
	logrus.WithFields(logrus.Fields{
			"Count": string(body),
		}).Info("Printing Metadata")
	ch<-fmt.Sprintf("%.2f elapsed with response length: %d %s", timeElapsed, len(body), url)
}