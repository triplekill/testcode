package lshw

import (
	"encoding/json"
	"log"
	"os/exec"
)

//sudo lshw -C network -json

type Entry struct {
	Id            string `json:"id"`
	Class         string `json:"class"`
	Claimed       bool   `json:"claimed"`
	Handle        string `json:"handle"`
	Description   string `json:"description"`
	Product       string `json:"product"`
	Vendor        string `json:"vendor"`
	Physid        string `json:"physid"`
	Businfo       string `json:"businfo"`
	Logicalname   string `json:"logicalname"`
	Version       string `json:"version"`
	Serial        string `json:"serial"`
	Units         string `json:"units"`
	Size          int    `json:"size"`
	Capacity      int    `json:"capacity"`
	Width         int    `json:"width"`
	Clock         int    `json:"clock"`
	Configuration struct {
		Autonegotiation string `json:"autonegotiation"`
		Broadcast       string `json:"broadcast"`
		Driver          string `json:"driver"`
		Driverversion   string `json:"driverversion"`
		Duplex          string `json:"duplex"`
		Firmware        string `json:"firmware"`
		Ip              string `json:"ip"`
		Latency         string `json:"latency"`
		Link            string `json:"link"`
		Multicast       string `json:"multicast"`
		Port            string `json:"port"`
		Speed           string `json:"speed"`
	} `json:"configuration"`
	Capabilities struct {
		Pm              string `json:"pm"`
		Msi             string `json:"msi"`
		Pciexpress      string `json:"pciexpress"`
		Msix            string `json:"msix"`
		BusMaster       string `json:"bus_master"`
		CapList         string `json:"cap_list"`
		Ethernet        bool   `json:"ethernet"`
		Physical        string `json:"physical"`
		Tp              string `json:"tp"`
		Mii             string `json:"mii"`
		Bt              string `json:"10bt"`
		BtFd            string `json:"10bt-fd"`
		Bt1             string `json:"100bt"`
		BtFd1           string `json:"100bt-fd"`
		BtFd2           string `json:"1000bt-fd"`
		Autonegotiation string `json:"autonegotiation"`
	} `json:"capabilities"`
}

func GetLSHW() {
	var entries []Entry

	cmd := exec.Command("sudo", "lshw", "-C", "network", "-json")
	output, err := cmd.Output()
	if err != nil {
		log.Println("run Exec Err:", err)
	}
	log.Println("output:", string(output))
	// json unmarshal output
	err = json.Unmarshal(output, &entries)
	if err != nil {
		log.Println("JSON UNMARSHAL Err:", err)
	}
	log.Println(entries)
}
