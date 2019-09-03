package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//Resp for requeest and response
type Resp struct {
	X   []int     `json:"x"`
	Y   []int     `json:"y"`
	Pol []float64 `json:"pol"`
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("REQUEST -------- HOST - %s METHOD - %s Params - %s \n", r.Host, r.Method, r.URL.Query())

		myX := r.URL.Query().Get("x")
		myY := r.URL.Query().Get("y")

		//Split strings to slice of strings
		Xstring := strings.Split(myX, ",")
		Ystring := strings.Split(myY, ",")

		reqlen := len(Xstring)

		//Convert slice of strings to slice of ints
		Xint := make([]int, reqlen)
		Yint := make([]int, reqlen)
		for i := 0; i < reqlen; i++ {
			Xint[i], _ = strconv.Atoi(Xstring[i])
			Yint[i], _ = strconv.Atoi(Ystring[i])
		}

		//Set some headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		var res Resp

		//If its not single point data
		if reqlen > 1 {
			res.Pol = make([]float64, reqlen)
			PolynomialCoefficents(&res.Pol, Xint, Yint)
		}

		resJSON, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		fmt.Println("Result polynomial coefficents: ", res.Pol)

		w.Write(resJSON)
	})

	fmt.Println("starting server at :9000")
	http.ListenAndServe(":9000", nil)
}

/*
Template to deeper work with http reqests
func runGetFullReq() {

	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"Notebook/golang"},
		},
	}

	req.URL, _ = url.Parse("http://192.168.0.98:9000/?msg=1234556789")
	req.URL.Query().Add("param", "param1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close() // важный пункт!

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("testGetFullReq resp %#v\n\n\n", string(respBody))
}*/

//VERY HARD FUNC
func main() {
	startServer()
}
