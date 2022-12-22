package main

import (
	"encoding/json"
	"log"
	"net/http"
	"scheduler-extender/pkg/filter"
	"scheduler-extender/pkg/prioritize"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/filter", filterHandler)
	r.HandleFunc("/prioritize", prioritizeHandler)

	log.Fatal(http.ListenAndServe(":8888", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("scheduler extender is running"))
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	args := extenderv1.ExtenderArgs{}
	response := extenderv1.ExtenderFilterResult{}

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		response.Error = err.Error()
	} else {
		response = filter.Filter(args)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error(err)
	}
}

func prioritizeHandler(w http.ResponseWriter, r *http.Request) {
	args := extenderv1.ExtenderArgs{}
	response := make(extenderv1.HostPriorityList, 0)

	if err := json.NewDecoder(r.Body).Decode(&args); err == nil {
		response = prioritize.Prioritize(args)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error(err)
	}
}
