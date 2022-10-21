package server

type Slave struct {
	RunId string `json:"runId"`
	Host  string `json:"host"`
	Port  string `json:"port"`
}
