package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	wm "git.intra.weibo.com/openapi_rd/weibo-motan-go"
	"github.com/weibocom/motan-go"
	m "github.com/weibocom/motan-go"
)

func StopMotanAgent(agent *motan.Agent) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		s := agent.GetAgentServer()
		fmt.Printf("%+v\n", s)
		s.Destroy()
		io.WriteString(writer, "motan agent stoped")
	}
}

func main() {
	extFactory := wm.GetWeiboExtentionFactory()
	agent := m.NewAgent(extFactory)
	agent.RegisterManageHandler("/stop_motan_agent", StopMotanAgent(agent))
	agent.StartMotanAgent()
	time.Sleep(time.Hour * 999999)
}
