package v1

import (
	"fmt"
	"net/http"
	"storing_events/internal/db"
)

func (s *HTTPServer) finish(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.String())
	if !db.MongoParam.FinishEvent(r.FormValue("type")) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}