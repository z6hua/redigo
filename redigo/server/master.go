package server

type Master struct {
	RunId string `json:"runId"`
	Host  string `json:"host"`
	Port  string `json:"port"`
}

func (m *Master) RemoteAddr() string {
	return m.Host + ":" + m.Port
}
