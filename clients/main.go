package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/choueric/clog"
	"github.com/choueric/homeKit/homeKit"
)

// fetch interfaces in system except: "lo"
func getIfaces() []net.Interface {
	ifaces, err := net.Interfaces()
	check(err)
	i := func(ifaces []net.Interface) int {
		for i, v := range ifaces {
			if v.Name == "lo" {
				return i
			}
		}
		return -1
	}(ifaces)

	ifaces = append(ifaces[:i], ifaces[i+1:]...)
	return ifaces
}

func sendBlobToServer(blob *homeKit.IfaceInfoBlob) bool {
	client := &http.Client{}

	data, err := blob.ToJson()
	check(err)

	_, err = client.Post("http://127.0.0.1:8088/save/", "application/json", strings.NewReader(string(data)))
	check(err)
	return true
}

func main() {
	clog.SetFlags(clog.Lshortfile | clog.LstdFlags)

	ifaces := getIfaces()

	blob, err := homeKit.NewIfaceInfoBlob(ifaces)
	check(err)

	/*
		if data, err := blob.ToJson(); err == nil {
			fmt.Println(string(data))
		}
	*/
	sendBlobToServer(blob)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
