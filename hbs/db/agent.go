package db

import (
	"fmt"
	"github.com/open-falcon/hbs/g"
	"github.com/open-falcon/common/model"
	"log"
)

func UpdateAgent(agentInfo *model.AgentUpdateInfo) {
	sql := ""
	if g.Config().Hosts == "" {
		sql = fmt.Sprintf(
			"insert into host(hostname, ip, agent_version, plugin_version) values ('%s', '%s', '%s', '%s') on duplicate key update ip='%s', agent_version='%s', plugin_version='%s'",
			agentInfo.ReportRequest.Hostname,
			agentInfo.ReportRequest.IP,
			agentInfo.ReportRequest.AgentVersion,
			agentInfo.ReportRequest.PluginVersion,
			agentInfo.ReportRequest.IP,
			agentInfo.ReportRequest.AgentVersion,
			agentInfo.ReportRequest.PluginVersion,
		)
	} else {
		// sync, just update
		sql = fmt.Sprintf(
			"update host set hostname='%s', agent_version='%s', plugin_version='%s' where ip='%s'",
			agentInfo.ReportRequest.AgentVersion,
			agentInfo.ReportRequest.PluginVersion,
			agentInfo.ReportRequest.Hostname,
			agentInfo.ReportRequest.IP,
		)
	}

	_, err := DB.Exec(sql)
	if err != nil {
		log.Println("exec", sql, "fail", err)
	}
		log.Print(DB)

}

