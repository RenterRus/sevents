package v1

import (
	"fmt"
	"net/http"
)

func (s *HTTPServer) finish(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())

	answer := make(chan bool, 1)

	go func() {
		answer <- s.Mongo.FinishEvent(r.FormValue("type"))
	}()

	if !<-answer {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}
