
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"reflect"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", HealthzHandler)
	err := http.ListenAndServe(":1080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
	code := reflect.ValueOf(w).Elem().FieldByName("status")
	logrus.Info("ip: ", r.RemoteAddr, ", http code: ", code)
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	code := reflect.ValueOf(w).Elem().FieldByName("status")
	fmt.Fprintf(w, "响应码 %v\n", code)

}
