package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sagar23sj/benchmark-grpc-rest/sample"
)

func hello(w http.ResponseWriter, r *http.Request) {

	var req sample.BenchLarge

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respMsg := fmt.Sprintf("\nName: %s, Age: %d, Height: %f,Weight: %f,Alive: %v,Desc: %v,Nickname: %s,Num: %d,Flt: %f,Dbl: %f,Tru: %v,Data: %v",
		req.Name,
		req.Age,
		req.Height,
		req.Weight,
		req.Alive,
		req.Desc,
		req.Nickname,
		req.Num,
		req.Flt,
		req.Dbl,
		req.Tru,
		req.Data,
	)

	json.NewEncoder(w).Encode(sample.Response{
		Message: respMsg,
	})
}

func RestStart() {
	fmt.Println("Starting rest server ....")
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
