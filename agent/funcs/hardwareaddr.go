package funcs

import (
        "os/exec"
	"github.com/toolkits/nux"
        "log"
	"strings"
	"runtime"
	"os"
	"io/ioutil"
	"strconv"
)
//获取序列号
func GetSN()string{
        var sn string
	contents, err := ioutil.ReadFile("/sys/class/dmi/id/product_serial")
	if err == nil {
		sn = strings.TrimSpace(string(contents))
        }else{
        	SN, err := exec.Command("/bin/bash", "-c", "sudo dmidecode -t 1 |grep 'Serial Number:' |awk '{print $3}'").Output()
        	if err != nil {
        	        log.Println(err)
        	}
		sn = strings.TrimSpace(string(SN))
        }
        return sn
}
//获取厂商
func GetCom()string {
	var com string
	contents, err := ioutil.ReadFile("/sys/class/dmi/id/bios_vendor")
	if err == nil {
		arr := strings.Split(strings.TrimSpace(string(contents))," ")
		com  =arr[0]
	} else {
		Com, err := exec.Command("/bin/bash", "-c", "sudo dmidecode -t 1 |grep 'Manufacturer:' |awk '{print $2}'").Output()
		if err != nil {
			log.Println(err)
		}
		com = strings.TrimSpace(string(Com))
	}
	return com
}
//获取型号
func GetPro()string{
	var pro string
	contents, err := ioutil.ReadFile("/sys/class/dmi/id/product_name")
	if err == nil {
		pro = strings.TrimSpace(string(contents))
	} else {
        	Pro,err := exec.Command("/bin/bash","-c","dmidecode -t 1 |grep 'Product Name:' |awk -F : '{print $2}'").Output()
        	if err != nil {
                	log.Println(err)
        	}
		pro = strings.TrimSpace(string(Pro))
	}
	return pro
}
//获取内存
func GetMem()string{
        m, err := nux.MemInfo()
        if err != nil {
                log.Println(err)
        }
        return strconv.FormatUint(m.MemTotal/1024/1024/1000,10)
}
//获取主机名
func GetHoNa()string{
        ho,err := os.Hostname()
        if err != nil {
                log.Println(err)
        }
        return ho
}
//获取CPU核心数
func GetCpuNum()string{
        return strconv.Itoa(runtime.NumCPU())
}
//获取CPU型号
func GetCpuModel()string{
	model,err := nux.CpuModel()
	if err != nil {
                log.Println(err)
        }
        return model
}
//获取MAC地址字典类型
func GetMac()string{

        Pro,err := exec.Command("/bin/bash","-c","grep -oP '(\\w{2}:){5}\\w{2}' /var/log/dmesg").Output()
        if err != nil {
                log.Println(err)
        }
        return strings.Replace(strings.TrimSpace(string(Pro)),"\n",";",-1)
}

//获取disk大小
func GetDisk()string {
        var disk []byte
        disk, _ = exec.Command("/bin/bash", "-c", "fdisk -l | grep 'GB'|awk '{print $2}' | awk -F '：' '{print  $2}'").Output()
        if strings.TrimSpace(string(disk)) == ""{
                disk, _ = exec.Command("/bin/bash", "-c", "fdisk -l | grep 'GB'|awk '{print $3}'").Output()
        }
        Len := strings.Split(strings.TrimSpace(string(disk)),"\n")
        sum := 0.0
        for i:=0;i<len(Len);i++{
                value,_ := strconv.ParseFloat(Len[i],64)
                sum += value
        }
        return strconv.FormatInt(int64(sum),10)
}

