# Introduction

This is a general load testing tool for any REST endpoint. 

This activity is performed by repeatedly sending HTTP Requests to the specified endpoint. User
is required to also provide values for the total duration of the test and the total number of requests
to be sent.


Currently this tool supports load testing GET and POST APIs.

# Building

Run the following command to generate an executable:

```
make dep
make build
```

This will download all dependencies and generate the executable.

# Usage:

Following are the flags that are to be set for a general use:
1. --X = Type of the API (GET/POST) [Default GET]
2. --body = Path to the JSON body for the request [Default NULL]
3. --request = Number of requests to be sent [Default 1]
4. --time = Duration of testing [Default 1]
5. --url = Endpoint URL [Default NULL]

Sample usages are as follows:

### Typical GET API Usage

```
./ltest dispatch --X GET --request 100 --time 10 --url http://localhost:8080/get
```

This will send a total of 100 GET API requests over 10 seconds, to the specified URL.

### Typical POST API Usage

```
./test dispatch --X POST --request 100 --time 10 --url http://localhost:8080/post --body = utils/bZ/

```
This request sends a total of 100 POST API requests over 10 seconds.