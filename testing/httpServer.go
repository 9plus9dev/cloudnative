package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {

	http.HandleFunc("/", testServerHandler)

	http.ListenAndServe(":5080", nil)
}

func testServerHandler(wr http.ResponseWriter, req *http.Request) {
	out, _ := exec.Command("uname", "-r").Output() // get system version using linux command

	remoteIp := req.RemoteAddr                                   //get remote host IP from request
	log.Printf("get a request from remote host: %v\n", remoteIp) //send to log

	code := http.StatusOK //code=200
	wr.WriteHeader(code)  //explictly set status code to 200
	log.Printf("the return http status code is %d", code)

	for k, v := range req.Header {
		io.WriteString(wr, k+": "+strings.Join(v[:], ",")+"\n")
	} // write request head infomation to response

	io.WriteString(wr, "\nsystem-version: "+string(out)) // write system version to response
}
