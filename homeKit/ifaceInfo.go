package homeKit

import (
	"encoding/json"
	"net"
	"os"

	"github.com/choueric/clog"
)

type IfaceInfo struct {
	Name string `json:"name"`
	IP   net.IP `json:"ip_addr"`
}

type IfaceInfoBlob struct {
	InfoArray []IfaceInfo `json:"info_array"`
}

func (b *IfaceInfoBlob) ToJson() ([]byte, error) {
	data, err := json.MarshalIndent(b, "  ", "  ")
	if err != nil {
		clog.Fatal(err)
	}

	return data, nil
}

func (b *IfaceInfoBlob) FromJson(data []byte) error {
	if err := json.Unmarshal(data, b); err != nil {
		return err
	}

	return nil
}

func (b *IfaceInfoBlob) Save(filename string) error {
	data, err := b.ToJson()

	file, err := os.Create(filename)
	if err != nil {
		clog.Fatal(err)
	}
	defer file.Close()

	file.Write(data)

	return nil
}

func NewIfaceInfoBlob(ifaces []net.Interface) (*IfaceInfoBlob, error) {
	var blob IfaceInfoBlob

	for _, v := range ifaces {
		addrs, err := v.Addrs()
		if err != nil {
			clog.Println(err)
			continue
		}
		for _, a := range addrs {
			var ip net.IP
			switch v := a.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			clog.Printf("%s: %s\n", v.Name, a.String())
			info := IfaceInfo{v.Name, ip}
			blob.InfoArray = append(blob.InfoArray, info)
		}
	}

	return &blob, nil
}
