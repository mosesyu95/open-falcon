package funcs

import (
	"io/ioutil"
	"strings"
	"os/exec"
	"github.com/toolkits/net"
	"log"
)

func GetKernel()string{
	kernel := "none"
	Kernel, err := exec.Command("/bin/bash", "-c", "uname -r").Output()
	if err != nil {
		log.Println(err)
	}
		kernel = strings.TrimSpace(string(Kernel))
	return kernel
}
func GetSys()string{
	Sys:="none"
	contents, err := ioutil.ReadFile("/etc/redhat-release")
	if err == nil {
		Sys = strings.TrimSpace(string(contents))
	}
	return Sys
}
//获取本地IP
func GetIp()string{
	ip,err := net.IntranetIP()
	if err != nil {
                log.Println(err)
        }
        return ip[0]
}
func GetIlo()string{
	ip:= "none"
	ilo_ip, err := exec.Command("/bin/bash", "-c", "sudo  ipmitool lan  print  | grep 'IP Address'|grep -oP '[1-9][0-9]{0,2}(\\.[0-9]{1,3}){3}'").Output()
	if err != nil {
		log.Println(err)
	}
		ip = strings.TrimSpace(string(ilo_ip))
	return ip
}