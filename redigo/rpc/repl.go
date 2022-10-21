package rpc

import (
	"redigo/redigo/server"
	"redigo/redigo/utils"
)

type Repl interface {
	Psync(masterRunId string, runId string, offset int) map[string]any
}

type ReplRPC struct {
	RemoteAddr string
	URL        string
}

func NewReplRPC() *ReplRPC {
	return &ReplRPC{
		RemoteAddr: server.GetMaster().RemoteAddr(),
		URL:        utils.GenURL(server.GetMaster().Host, server.GetMaster().Port, ""),
	}
}

func (r ReplRPC) Psync(masterRunId string, runId string, offset int) map[string]any {
	path := "/psync_r"
	params := map[string]any{
		"runId":       runId,
		"masterRunId": masterRunId,
		"offset":      offset,
	}
	return utils.RequestGetData(r.URL+path, params)
}
