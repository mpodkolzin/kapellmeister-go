package main

import (
	"fmt"
	"kapellmeister-go/task"
	"net/http"
	"os/exec"
	"sync"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi there"))
	cmd := exec.Command("cmd", "python.exe", "-h")
	err := cmd.Run()
	//err := syscall.Exec("python.exe", []string{"-h"}, env)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	//app := Service{}
	serv, err := NewService()
	if err != nil {
		panic(err.Error())
	}

	wg := sync.WaitGroup{}

	task.Spawn(&wg, serv.Run)

	wg.Wait()

	//fmt.Scanln()
}
