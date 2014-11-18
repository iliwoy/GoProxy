package main

import (
    "gossl"
    "fmt"
    "bufio"
    "os"
    "strings"
)

func useage() {
    fmt.Println(`General commands:
start               Start services
stop                Stop  services
load  [path]        Load configurations from the path
`)
}

var proxy gossl.Proxy
var httpMirror gossl.HttpMirror
var httpsMirror gossl.HttpsMirror

func start() {
    proxy.Start(8181)
    httpMirror.Start(80)
    httpsMirror.Start(443)

    fmt.Println("success!")
}

func initServer() {
    proxy = gossl.Proxy{}
    httpMirror = gossl.HttpMirror{}
    httpsMirror = gossl.HttpsMirror{}
}

func loadConfigs(filePath string) {
    gossl.LoadConfigs(filePath)
    proxy.KillConnections()
    fmt.Println("success!")
}

func startRenego() {
    //renego := gossl.Renego{}
    //renego.Start(443)
}

func main() {
    fmt.Println(`Gossl v1.0`);
    reader := bufio.NewReader(os.Stdin)
    running := true
    for running {
        fmt.Print(">>")
        data, _, _ := reader.ReadLine()
        commands := strings.Split(string(data), " ")
        switch commands[0] {
        case "help":
            useage()
        case "start":
            start()
        case "load":
            if len(commands) == 2 {
                loadConfigs(commands[1])
            } else {
                fmt.Println("unknown command")
            }
        case "renego":
            startRenego()
        case "quit":
            running = false
            fmt.Println("bye!")
        default:
            fmt.Println("unknown command")
        }
    }
}
