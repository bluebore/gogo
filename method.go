package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB;
	a = PhoneConnecter{"PhoneCNT"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println(usb.Name());
}
