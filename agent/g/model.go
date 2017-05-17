package g
import (
	"fmt"
)


type AgentReportRequest struct {
	Hostname      string
	IP            string
	SN  	      string
	Company       string
	Model         string
	CPU 	      string
	CPU_num	      string
	Mem	      string
	System        string
	Kernel	      string
	ILO 	      string
	Disk	      string
	Mac  	      string
	Timestamp     string

}

func (this *AgentReportRequest) String() string {

		return fmt.Sprintf(
			"<hostname:%s, ip:%s, sn:%s, company:%s, model:%s, cpu_type:%s, cpu_num:%s, memory:%s, os_version:%s, kernel:%s, ilo_ip:%s, disk:%g, mac:%s, timestamp:%s>",
			this.Hostname,
			this.IP,
			this.SN,
			this.Company,
			this.Model,
			this.CPU,
			this.CPU_num,
			this.Mem,
			this.System,
			this.Kernel,
			this.ILO,
			this.Disk,
			this.Mac,
			this.Timestamp,

		)}

