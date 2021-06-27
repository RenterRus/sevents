package v1

import (
	"fmt"
	"net/http"
	"storing_events/internal/db"
	"strconv"
	"time"
)

func (s *HTTPServer) start(w http.ResponseWriter, r *http.Request){
//Если в базе уже есть незавершенное событие переданного типа, то новое событие создавать не нужно, и в ответ не должно приходить сообщение об ошибке
	//Не знаю, какой ответ в данном случае тогда отсылать, т.к. Already use, фактически, тоже, сообщение об ошибке, так что всегда ответ OK
	fmt.Println(r.URL.String())
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, db.MongoParam.WriteEvent(db.DocStruct{
		Id: strconv.Itoa(int(time.Now().Unix())),
		Start: time.Now().Format("01-02-2006 15:04:05.000000"),
		State: 0,
		Finish: "",
		TypeEvent: r.FormValue("type"),
	}))
}