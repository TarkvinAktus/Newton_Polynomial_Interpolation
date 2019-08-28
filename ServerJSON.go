package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
)

//Resp ОТВЕТНЫЕ ТОЧКИ
type Resp struct {
	X []int `json:"x"`
	Y []int `json:"y"`
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("REQUEST -------- HOST - %s METHOD - %s Params - %s \n", r.Host, r.Method, r.URL.Query())

		myX := r.URL.Query().Get("x")
		myY := r.URL.Query().Get("y")
		fmt.Println("x = ", myX, " y = ", myY)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		var res Resp

		res.X = make([]int, 10)
		res.Y = make([]int, 10)

		for i := 0; i < 10; i++ {
			res.X[i] = int(math.Pow(2, float64(i)))
			res.Y[i] = i
		}

		resJSON, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		w.Write(resJSON)
	})
	/*
		http.HandleFunc("/raw_body", func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close() // важный пункт!
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			fmt.Fprintf(w, "postHandler: raw body %s\n", string(body))
		})
	*/
	fmt.Println("starting server at :9000")
	http.ListenAndServe(":9000", nil)
}
/*
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

func main() {
	startServer()
}
