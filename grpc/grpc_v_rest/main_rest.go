package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func MainREST(addr string) {
	fmt.Println("Starting REST SErver......")

	// r := vestigo.NewRouter()
	// r.Post("/info", SetInfo)

	// log.Fatal(http.ListenAndServe(addr, r))

	fmt.Println("Starting rest server ....")
	http.HandleFunc("/info", SetInfo)

	http.ListenAndServe(addr, nil)
}

// SetInfo - Rest HTTP Handler
func SetInfo(w http.ResponseWriter, r *http.Request) {
	// var req apiInput

	// err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// fmt.Println("sssssssssssssssssssssssssssssssssssss")

	// respMsg := fmt.Sprintf("\nName: %s, Age: %d, Height: %f,Weight: %f,Alive: %v,Desc: %v,Nickname: %s,Num: %d,Flt: %f,Dbl: %f,Tru: %v,Data: %v",
	// 	req.Name,
	// 	req.Age,
	// 	req.Height,
	// 	req.Weight,
	// 	req.Alive,
	// 	req.Desc,
	// 	req.Nickname,
	// 	req.Num,
	// 	req.Flt,
	// 	req.Dbl,
	// 	req.Tru,
	// 	req.Data,
	// )

	// var response apiResponse
	// // validate input
	// if err := validate(req); err != nil {
	// 	response.Message = "error occured"
	// 	respBytes, _ := json.Marshal(response)

	// 	w.WriteHeader(400)
	// 	w.Write(respBytes)
	// 	return
	// }

	// // json.NewEncoder(w).Encode(apiResponse{
	// // 	Message: respMsg,
	// // })

	// respBytes, _ := json.Marshal(apiResponse{
	// 	Message: respMsg,
	// })
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(200)
	// w.Write(respBytes)
	// return

	var req apiInput

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

	json.NewEncoder(w).Encode(apiResponse{
		Message: respMsg,
	})
}

type apiResponse struct {
	Message string `json:"message"`
}

type apiInput struct {
	Name     string  `json:"name,omitempty"`
	Age      int64   `json:"age,omitempty"`
	Height   float32 `json:"height,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Alive    bool    `json:"alive,omitempty"`
	Desc     []byte  `json:"desc,omitempty"`
	Nickname string  `json:"nickname,omitempty"`
	Num      int64   `json:"num,omitempty"`
	Flt      float32 `json:"flt,omitempty"`
	Dbl      float64 `json:"dbl,omitempty"`
	Tru      bool    `json:"tru,omitempty"`
	Data     []byte  `json:"data,omitempty"`
}

// Validate - implementation of Validatable
func (ai apiInput) Validate() error {
	var err validationErrors
	if ai.Name == "" {
		err = append(err, errors.New("Name must be present"))
	}
	if ai.Age <= 0 {
		err = append(err, errors.New("Age must be real"))
	}
	if ai.Height <= 0 {
		err = append(err, errors.New("Height must be real"))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}
