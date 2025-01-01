package generic_serial_device

import (
	"encoding/binary"
	"fmt"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
	"strconv"
	"strings"
	"sync"
	"time"
)

//"golang.org/x/exp/io/i2c"

const (
	DevicePathPrefix = "/dec/i2c-"
)

type I2CDevice struct {
	path    string
	address byte
	device  i2c.Dev
	debug   bool
	lock    sync.Mutex
}

func NewI2CDevice(path string, address byte) *I2CDevice {

	i2cDevice := &I2CDevice{
		path:    path,
		address: address,
	}

	pathUpdated := i2cDevice.fixDevicePath()
	if pathUpdated {
		fmt.Printf("i2c path updated")
	}

	if _, err := host.Init(); err != nil {
		fmt.Printf("i2c init failed")
	}

	if i2cDevice.debug {
		OutputI2CDevices()
	}
	bus, err := i2creg.Open(i2cDevice.path)
	device := i2c.Dev{bus, uint16(i2cDevice.address)}

	if err != nil {
		fmt.Printf("failed to open I2C device: %s", err.Error())
		return nil
	}

	i2cDevice.device = device
	return i2cDevice
}

func OutputI2CDevices() {
	fmt.Print("I²C buses available:\n")
	for _, ref := range i2creg.All() {

		fmt.Printf("- %s\n", ref.Name)

		if ref.Number != -1 {
			fmt.Printf(" %d\n", ref.Number)
		}

		if len(ref.Aliases) != 0 {
			fmt.Printf(" %s\n", strings.Join(ref.Aliases, " "))
		}
	}
}

func (i2cDevice *I2CDevice) WriteThenRead(txBuf []byte, rxLen int) (rxBuf []byte, err error) {

	rxBuf = make([]byte, rxLen)

	time.Sleep(5 * time.Millisecond) // Random guess
	err = i2cDevice.Write(txBuf)
	if err != nil {
		return
	}

	time.Sleep(5 * time.Millisecond) // Random guess
	err = i2cDevice.device.Tx(nil, rxBuf)

	return
}

func (i2cDevice *I2CDevice) Write(txBuf []byte) error {
	time.Sleep(5 * time.Millisecond) // Random guess
	_, err := i2cDevice.device.Write(txBuf)
	return err
}

func (i2cDevice *I2CDevice) fixDevicePath() (changed bool) {

	if deviceNumber, err := strconv.Atoi(i2cDevice.path); err == nil {

		i2cDevice.path = fmt.Sprintf("%s%d", i2cDevice.path, deviceNumber)
		//i2cDevice.logger.Info("")
		changed = true
	}

	return
}

func (i2cDevice *I2CDevice) WriteToTwoByteRegister(register uint16, txBuf []byte) error {

	sendBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(sendBuf, register)
	sendBuf = append(sendBuf, txBuf...)
	//fmt.Printf("WR: % 2X\n", sendBuf)
	return i2cDevice.Write(sendBuf)
}

func (i2cDevice *I2CDevice) ReadFromTwoByteRegister(register uint16, rxLen int) (rxBuf []byte, err error) {
	sendBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(sendBuf, register)
	rxBuf, err = i2cDevice.WriteThenRead(sendBuf, rxLen)
	return
}

func (i2cDevice *I2CDevice) ReadRegister(register []byte, rxLen int) (rxBuf []byte, err error) {
	rxBuf = make([]byte, rxLen)
	err = i2cDevice.device.Tx(register, rxBuf)
	fmt.Printf("rxBuf (%2X): % 2X\n", register, rxBuf)
	return
}
