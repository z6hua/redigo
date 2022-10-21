package enums

type ServerType int
type ReplSYNCType string

const (
	Master ServerType = iota
	Slave
)

const (
	FullReSync ReplSYNCType = "FullReSync"
	PartReSync ReplSYNCType = "PartReSync"
)
