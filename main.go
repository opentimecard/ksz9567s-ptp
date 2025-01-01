package main

import (
	"github.com/opentimecard/ksz9567s-ptp/generic_serial_device"
	"os"
)

func main() {

	generic_serial_device.NewSPIDevice()
	/*
		...
	*/
	os.Exit(0)

}
