package haas

type HaasDevice interface {
	Entities() []Discovery
	StateTopic() string
}

type PwsDevice struct {
	Pws     HaasDevice
	Sensors []HaasDevice
}
