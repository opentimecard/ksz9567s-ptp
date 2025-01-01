package eth_sw_KSZ9567S

import (
	"encoding/binary"
	"fmt"
)

// Pointless code required by MicroChip fumbling

func (ethSw *EthSwKSZ9567S) RunAllErrata() {

	_ = ethSw.Set100MbpsNoAutoNegotiation()
	_ = ethSw.ErrataOne()
	_ = ethSw.ErrataTwo()
	_ = ethSw.ErrataFour()
	_ = ethSw.ErrataSeven()
	_ = ethSw.ErrataNine()
	_ = ethSw.Set1000MbpsAutoNegotiation()
}

func (ethSw *EthSwKSZ9567S) Set100MbpsNoAutoNegotiation() error {

	OneHundredMegsNoAutoNegotiation := make([]byte, 2)
	binary.BigEndian.PutUint16(OneHundredMegsNoAutoNegotiation, 0x2100)

	for port := 1; port <= 5; port++ {

		register := GetRegisterForPort(PHY_BASIC_CONTROL_REGISTER, port)
		if err := ethSw.i2cDevice.WriteToTwoByteRegister(register, OneHundredMegsNoAutoNegotiation); err != nil {
			fmt.Printf("unable to apply set 100 megs amd no auto negotiation register")
		}
	}
	return nil
}

func (ethSw *EthSwKSZ9567S) Set1000MbpsAutoNegotiation() error {

	AGigAndAutoNegotiation := make([]byte, 2)
	binary.BigEndian.PutUint16(AGigAndAutoNegotiation, 0x1340)

	for port := 1; port <= 5; port++ {

		register := GetRegisterForPort(PHY_BASIC_CONTROL_REGISTER, port)
		//fmt.Printf("Bytes to re-enable full speed: %X - % 2X\n", register, AGigAndAutoNegotiation)

		if err := ethSw.i2cDevice.WriteToTwoByteRegister(register, AGigAndAutoNegotiation); err != nil {
			fmt.Printf("unable to apply set 1000 megs and auto negotiation register")
		}
	}
	return nil
}

// ErrataOne
// Register settings are needed to improve PHY receive performance
func (ethSw *EthSwKSZ9567S) ErrataOne() error {

	mmd := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x1C, 0x1C}
	register := []byte{0x6F, 0x8F, 0x9D, 0x75, 0xD3, 0x06, 0x08}
	data := []uint16{0xDD0B, 0x6032, 0x248C, 0x0060, 0x7777, 0x3008, 0x2000}

	for port := 1; port <= 5; port++ {

		for x := 0; x < len(mmd); x++ {

			//fmt.Printf("Running errate one: %d - %d\n", port, x)
			mmdValue := uint16(mmd[x])
			registerValue := uint16(register[x])
			dataValue := data[x]

			ethSw.WriteMMDRegister(port, mmdValue, registerValue, dataValue)
		}
	}
	return nil
}

// ErrataTwo
// Transmit waveform amplitude can be improved (1000BASE-T, 100BASE-TX, 10BASE- Te)
func (ethSw *EthSwKSZ9567S) ErrataTwo() error {

	mmdValue := uint16(0x1C)
	registerValue := uint16(0x4)
	dataValue := uint16(0x00D0)

	for port := 1; port <= 5; port++ {

		//fmt.Printf("Running errate two: %d\n", port)
		ethSw.WriteMMDRegister(port, mmdValue, registerValue, dataValue)
	}
	return nil
}

// ErrataFour
// Energy Efficient Ethernet (EEE) feature select must be manually disabled
func (ethSw *EthSwKSZ9567S) ErrataFour() error {

	mmdValue := uint16(0x7)
	registerValue := uint16(0x3C)
	dataValue := uint16(0x0000)

	for port := 1; port <= 5; port++ {

		//fmt.Printf("Running errate four: %d\n", port)
		ethSw.WriteMMDRegister(port, mmdValue, registerValue, dataValue)
	}
	return nil
}

// ErrataSeven
// SGMII auto-negotiation does not set bit 0 in the auto-negotiation code word
func (ethSw *EthSwKSZ9567S) ErrataSeven() error {

	ethSw.WriteSGMIIRegister(0x001F0004, 0x01A0)
	return nil
}

// ErrataEight
// SGMII port link details from the connected SGMII PHY are not passed properly to the port 7 GMAC
func (ethSw *EthSwKSZ9567S) ErrataEight() error {
	// Not implemented. Only 1000mbps possible on port 7 SGMII / SerDes
	return nil
}

// ErrataNine
// Register settings are required to meet data sheet supply current specifications
func (ethSw *EthSwKSZ9567S) ErrataNine() error {

	mmd := []byte{0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C, 0x1C}
	register := []byte{0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x20}
	data := []uint16{0x6EFF, 0xE6FF, 0x6EFF, 0xE6FF, 0x00FF, 0x43FF, 0xC3FF, 0x6FFF, 0x07FF, 0x0FFF, 0xE7FF, 0xEFFF, 0xEEEE}

	for port := 1; port <= 5; port++ {

		for x := 0; x < len(mmd); x++ {

			//fmt.Printf("Running errate seven: %d - %d\n", port, x)
			mmdValue := uint16(mmd[x])
			registerValue := uint16(register[x])
			dataValue := data[x]

			ethSw.WriteMMDRegister(port, mmdValue, registerValue, dataValue)
		}
	}
	return nil
}
