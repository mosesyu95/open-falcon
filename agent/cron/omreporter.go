package cron

import (
	"github.com/open-falcon/agent/g"
	"github.com/open-falcon/agent/funcs"
	"log"
	"time"
	"strconv"
	"github.com/open-falcon/common/model"
)

func ReportOmAgentStatus() {
	if g.Config().Omserver.Enabled && g.Config().Omserver.Addr != "" {
		go reportOmAgentStatus(time.Duration(g.Config().Omserver.Interval) * time.Hour)
	}
}

func reportOmAgentStatus(interval time.Duration) {
	for {
		req := g.AgentReportRequest{
			Hostname:	funcs.GetHoNa(),
			IP:		funcs.GetIp(),
			SN:		funcs.GetSN(),
			Company:	funcs.GetCom(),
			Model:		funcs.GetPro(),
			CPU:		funcs.GetCpuModel(),
			CPU_num:	funcs.GetCpuNum(),
			Mem:		funcs.GetMem(),
			System:		funcs.GetSys(),
			Kernel:		funcs.GetKernel(),
			ILO:		funcs.GetIlo(),
			Mac:		funcs.GetMac(),
			Disk:		funcs.GetDisk(),
			Timestamp:	strconv.FormatInt(int64(time.Now().Unix()),10),
		}
		if g.Config().Debug {
			log.Println("server information :",req)
		}
		var resp model.SimpleRpcResponse
		err := g.OmClient.Call("Agent.ReportStatus", req, &resp)
		if err != nil || resp.Code != 0 {
			log.Println("call Agent.ReportStatus fail:", err, "Request:", req, "Response:", resp)
		}

		time.Sleep(interval)
	}
}
