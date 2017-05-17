package cron

import (
	"fmt"
        "github.com/open-falcon/alarm/g"
	"github.com/open-falcon/common/model"
	"github.com/open-falcon/common/utils"

)

func BuildCommonSMSContent(event *model.Event) string {
	system_name, manager_name := g.GetSysName(event.Endpoint)
	return fmt.Sprintf(
		"[P%d][%s][%s][%s][%s][%s][%s %s %s %s%s%s][O%d %s]",
		event.Priority(),
		event.Status,
		event.Endpoint,
		system_name,
		manager_name,
		event.Note(),
		event.Func(),
		event.Metric(),
		utils.SortedTags(event.PushedTags),
		utils.ReadableFloat(event.LeftValue),
		event.Operator(),
		utils.ReadableFloat(event.RightValue()),
		event.CurrentStep,
		event.FormattedTime(),
	)
}

func BuildCommonMailContent(event *model.Event) string {
	system_name, manager_name := g.GetSysName(event.Endpoint)
	counter:=g.GetC(event.Metric(),utils.SortedTags(event.PushedTags))
	id :=g.GetMap(event.Endpoint,counter,)
	link :="http://10.249.1.44:8081/chart/big?id="+id
	return fmt.Sprintf(
		"状态%s  级别P%d\r\n设备:%s\r\n系统:%s %s\n类别:%s\r\n内容:%s%s: %s%s%s\n备注:%s\r\n最多报警%d次, 第%d次报警\r\n报警时间:%s\r\n%s\r\n",
		event.Status,
		event.Priority(),
		event.Endpoint,
		system_name,
		manager_name,
		event.Metric(),
		utils.SortedTags(event.PushedTags),
		event.Func(),
		utils.ReadableFloat(event.LeftValue),
		event.Operator(),
		utils.ReadableFloat(event.RightValue()),
		event.Note(),
		event.MaxStep(),
		event.CurrentStep,
		event.FormattedTime(),
		link,
	)
}

func GenerateSmsContent(event *model.Event) string {
	return BuildCommonSMSContent(event)
}

func GenerateMailContent(event *model.Event) string {
	return BuildCommonMailContent(event)
}
