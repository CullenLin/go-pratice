package main

import (
	"net/http"
	"runtime"
	"fmt"
	"github.com/gorilla/context"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var route = mux.NewRouter()
const ApiPath = "/api/v1"

// init request by 'curl -XPOST http://localhost:8080/api/v1/job'
func main() {

	router := route.PathPrefix(ApiPath).Subrouter()

	router.Path("/job").Methods("POST").HandlerFunc(AssignBakeJob)

	http.ListenAndServe("localhost:8080", route)
}

// intercept request, just like filter in java
func ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		context.Clear(req)
		// handler the unexpected panic
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				logrus.Errorf("The following error occurs: %s", r.(error).Error())
			case string:
				logrus.Errorf("The following error occurs: %s", r.(string))
			default:
				logrus.Errorf("The following error occurs: %#v", r)
			}
			trace := make([]byte, 4028)
			bNum := runtime.Stack(trace, false)
			fmt.Printf("Stack of error trace: %s\n", trace[0:bNum])
		}
	}()

	logrus.Debugf("Start to handler request: %s", req.URL)

	route.ServeHTTP(resp, req)

	// flush the response.
	if flusher, ok := resp.(http.Flusher); ok {
		flusher.Flush()
	}
	logrus.Debugf("Complete to handler request: %s", req.URL)
}

func AssignBakeJob(resp http.ResponseWriter, req *http.Request) () {
	logrus.Infof("receive job")
}
