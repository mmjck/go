package models

import "fmt"

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

func (f *Flags) GetApplicationUrl() (*string, *ErrorDetail) {
	f, err := GetFlags()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s:%s", f.Ip, f.Port)
	return &url, nil
}

func GetFlags() (*Flags, *ErrorDetail) {
	if flagsObj == nil {
		return nil, &ErrorDetail{
			ErrorType:    ErrorTypeFatal,
			ErrorMessage: "Flags not set",
		}
	}

	return flagsObj, nil
}
