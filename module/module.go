package module

import (
	"fmt"
	"net/http"
)

func Init() error {
	return nil
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"code":0, "data":"hello!", "msg":""}`)
}
