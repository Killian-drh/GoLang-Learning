package main

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"regexp"
)

func isPrivate(ip net.IP) bool {
	rfc1918 := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}
	for _, cidr := range rfc1918 {
		_, net, _ := net.ParseCIDR(cidr)
		if net.Contains(ip) {
			return true
		}
	}
	return false
}

func ProcessData(list_ip []string) {

	for i, ip := range list_ip {
		ip = ip + "/32"
		ipTmp, _, err := net.ParseCIDR(ip)
		if err != nil {
			continue
		} else if isPrivate(ipTmp) {
			fmt.Printf("%s is a private ip \n", list_ip[i])
			result := ExecCommandPrivateIp(list_ip[i])
			fmt.Printf("%s", result)
			if regex(result) {
				fmt.Printf("%s is up \n", list_ip[i])
			}
		}

	}

}

func ExecCommandPrivateIp(ip string) string {

	arg0 := "nmap -Pn -p 1-65535 -T4 --max-retries 3 --min-rtt-timeout 30ms --initial-rtt-timeout 100ms -sS --version-intensity 7 -O --osscan-limit -oX scan.xml " + ip

	cmd := exec.Command("bash", "-c", arg0)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())

	}
	fmt.Println("Result: " + out.String())

	return out.String()
}

func regex(s string) bool {
	matched, err := regexp.MatchString(`up`, s)
	if err == nil {
		fmt.Printf("ERROR while matching string %s \n", err)
	}
	return matched
}
