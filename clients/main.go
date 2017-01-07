package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/choueric/clog"
	"github.com/choueric/homeKit/homeKit"
)

const WEBSERVER = "http://ericnode.info:8088"
var gInfoArray []homeKit.IfaceInfo

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

	_, err = client.Post(WEBSERVER+"/save/", "application/json", strings.NewReader(string(data)))
	check(err)
	return true
}

func isBlobChanged(b *homeKit.IfaceInfoBlob) bool {
	a := b.InfoArray
	if len(gInfoArray) != len(a) {
		gInfoArray = a
		clog.Printf("len changed\n")
		return true 
	}

	for i, v := range gInfoArray {
		if (v.Name != a[i].Name) || (v.IP.Equal(a[i].IP) == false) {
			gInfoArray = a
			clog.Printf("%v != %v\n", v, a[i])
			return true 
		}
	}
	return false
}

func main() {
	clog.SetFlags(clog.Lshortfile | clog.LstdFlags)

	/*
		if data, err := blob.ToJson(); err == nil {
			fmt.Println(string(data))
		}
	*/
	for {
		time.Sleep(1 * time.Second)
		ifaces := getIfaces()
		blob, err := homeKit.NewIfaceInfoBlob(ifaces)
		check(err)

		if isBlobChanged(blob) == false {
			continue;
		}

		clog.Printf("send blob\n")
		sendBlobToServer(blob)
	}
}

func checkFatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
