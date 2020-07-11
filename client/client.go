package client

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io/ioutil"
)

// sends a get request
func GetRequest(url string) {
	// Create Pipeline Client
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		logrus.WithError(err).Error("Can't send http request")
	}

	body := resp.Body()
	logrus.WithFields(
		logrus.Fields{
			"Count": string(body),
		}).Info("Printing Metadata")

}

// send a post request
func PostRequest(url string, bodyPath string){
	// read the post body
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	jsonFile, err := ioutil.ReadFile(bodyPath)
	if err!=nil {
		logrus.WithError(err).Error("Unable to open file")
	}

	req.SetBody(jsonFile)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err = client.Do(req, resp)
	if err != nil {
		logrus.WithError(err).Error("Error sending http request")
	}

	body := resp.Body()
	logrus.WithFields(logrus.Fields{
			"Count": string(body),
		}).Info("Printing Metadata")
}