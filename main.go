package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Method struct
type Method struct {
	Service  string `json:"service"`
	Method   string `json:"method"`
	IsPublic bool   `json:"ispublic"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	urlPart := strings.Split(r.URL.Path, "/")

	serviceName := urlPart[1]

	path := "http://" + serviceName + ":8080/"

	serveReverseProxy(path, w, r)
}

func wrapper(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		jsonFile, _ := os.Open("method.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var method []Method
		json.Unmarshal(byteValue, &method)
		urlPart := strings.Split(r.URL.Path, "/")

		if len(urlPart) < 3 { // Kiem tra url theo cau truc /service/method/xyz
			http.Error(w, "URL is Invalid!", http.StatusBadRequest)
			return
		}

		var IsExist, IsPublic bool
		for _, m := range method { // Kiem tra method ton tai
			if urlPart[1] == m.Service && urlPart[2] == m.Method {
				IsExist = true
				IsPublic = m.IsPublic
				break
			}
		}
		if !IsExist {
			http.Error(w, "Method does not exist!", http.StatusNotFound)
			return
		}

		if IsPublic { // Kiem tra method public
			h.ServeHTTP(w, r)
			return
		}

		reqToken := getJwtToken(r.Header)
		if len(reqToken) < 1 {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		claims, err := verifyToken(reqToken)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		if claims.Role == 1 { // Kiem tra role user
			h.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Unauthorized!", http.StatusUnauthorized)
	}
}

func main() {

	http.HandleFunc("/", wrapper(handler))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
