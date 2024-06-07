package models

type Flags struct {
	Ip   string
	Port string
}

var flagsObj *Flags

func NewFlags(ip, port string) *Flags {
	if flagsObj == nil {
		flagsObj = &Flags{
			Ip:   ip,
			Port: port,
		}
	}

	return flagsObj
}
