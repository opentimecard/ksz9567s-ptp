package eth_sw_KSZ9567S

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"time"
)

func (ethSw *EthSwKSZ9567S) EnableTransparentClockMode() error {

	var GlobalPTPMessageConfigOneRegister uint16
	GlobalPTPMessageConfigOneRegister |= 1 << 6 // Enable IEEE 1588 PTP Mode
	GlobalPTPMessageConfigOneRegister |= 1 << 5 // Enable Detection of IEEE 802.3 Ethernet PTP Messages
	GlobalPTPMessageConfigOneRegister |= 1 << 4 // Enable Detection of IPv4/UDP PTP Messages
	GlobalPTPMessageConfigOneRegister |= 1 << 3 // Enable Detection of IPv6/UDP PTP Messages
	GlobalPTPMessageConfigOneRegister |= 1 << 0 // Selection of One-step or Two-step Operation (one-step)

	var GlobalPTPMessageConfigTwoRegister uint16
	GlobalPTPMessageConfigTwoRegister |= 1 << 12 // Enable Unicast PTP
	GlobalPTPMessageConfigTwoRegister |= 1 << 11 // Enable Alternate Master
	GlobalPTPMessageConfigTwoRegister |= 1 << 10 // PTP Messages Priority TX Queue
	GlobalPTPMessageConfigTwoRegister |= 1 << 2  // Enable IPv4/UDP Checksum Calculation for Egress Packets

	if err := ethSw.i2cDevice.WriteToTwoByteRegister(GLOBAL_PTP_MESSAGE_CONFIG_1_REGISTER, []byte{
		byte(GlobalPTPMessageConfigOneRegister >> 8),
		byte(GlobalPTPMessageConfigOneRegister)}); err != nil {
		fmt.Printf("unable to apply global PTP message config one register")
	}
	if err := ethSw.i2cDevice.WriteToTwoByteRegister(GLOBAL_PTP_MESSAGE_CONFIG_2_REGISTER, []byte{
		byte(GlobalPTPMessageConfigTwoRegister >> 8),
		byte(GlobalPTPMessageConfigTwoRegister)}); err != nil {
		fmt.Printf("unable to apply global PTP message config two register")
	}

	return nil
}

func (ethSw *EthSwKSZ9567S) DisableTransparentClockMode() error {

	// Could probably just send zeros
	var GlobalPTPMessageConfigOneRegister uint16
	GlobalPTPMessageConfigOneRegister |= 0 << 6 // Enable IEEE 1588 PTP Mode
	GlobalPTPMessageConfigOneRegister |= 0 << 5 // Enable Detection of IEEE 802.3 Ethernet PTP Messages
	GlobalPTPMessageConfigOneRegister |= 0 << 4 // Enable Detection of IPv4/UDP PTP Messages
	GlobalPTPMessageConfigOneRegister |= 0 << 3 // Enable Detection of IPv6/UDP PTP Messages
	GlobalPTPMessageConfigOneRegister |= 0 << 0 // Selection of One-step or Two-step Operation (one-step)

	var GlobalPTPMessageConfigTwoRegister uint16
	GlobalPTPMessageConfigTwoRegister |= 0 << 12 // Enable Unicast PTP
	GlobalPTPMessageConfigTwoRegister |= 0 << 11 // Enable Alternate Master
	GlobalPTPMessageConfigTwoRegister |= 0 << 10 // PTP Messages Priority TX Queue
	GlobalPTPMessageConfigTwoRegister |= 0 << 2  // Enable IPv4/UDP Checksum Calculation for Egress Packets

	if err := ethSw.i2cDevice.WriteToTwoByteRegister(GLOBAL_PTP_MESSAGE_CONFIG_1_REGISTER, []byte{
		byte(GlobalPTPMessageConfigOneRegister >> 8),
		byte(GlobalPTPMessageConfigOneRegister)}); err != nil {
		fmt.Printf("unable to apply global PTP message config one register")
	}
	if err := ethSw.i2cDevice.WriteToTwoByteRegister(GLOBAL_PTP_MESSAGE_CONFIG_2_REGISTER, []byte{
		byte(GlobalPTPMessageConfigTwoRegister >> 8),
		byte(GlobalPTPMessageConfigTwoRegister)}); err != nil {
		fmt.Printf("unable to apply global PTP message config two register")
	}

	return nil
}

func (ethSw *EthSwKSZ9567S) SetPTPClock() error {

	tsNow := unix.NsecToTimespec(time.Now().UnixNano())

	if err := ethSw.i2cDevice.WriteToTwoByteRegister(GLOBAL_PTP_RTC_CLOCK_TIME_REGISTER, []byte{
		byte(tsNow.Nsec >> 24),
		byte(tsNow.Nsec << 16),
		byte(tsNow.Nsec >> 8),
		byte(tsNow.Nsec),
		byte(tsNow.Sec >> 24),
		byte(tsNow.Sec >> 16),
		byte(tsNow.Sec >> 8),
		byte(tsNow.Sec),
	}); err != nil {
		fmt.Printf("unable to apply time to time registers")
	}

	var controlRegister uint16
	controlRegister |= 1 << 3
	controlRegister |= 1 << 1

	if err := ethSw.i2cDevice.WriteToTwoByteRegister(GLOBAL_PTP_CLOCK_CONTROL_REGISTER, []byte{
		byte(controlRegister >> 8),
		byte(controlRegister),
	}); err != nil {
		fmt.Printf("unable to load time and enable ptp clock")

	}

	return nil
}

func (ethSw *EthSwKSZ9567S) Run1PPSInRunloop(done chan struct{}) error {

	ticker := time.NewTicker(1 * time.Second)
Runloop:
	for {
		select {
		case <-ticker.C:
			ethSw.dumpTimehead(TIMESTAMP_1ST_SAMPLE_TIME_NSEC_REGISTER)
			ethSw.ResetInputEventTrigger()

		case <-done:
			break Runloop
		}
	}

	return nil
}

func (ethSw *EthSwKSZ9567S) dumpTimehead(firstRegister int) {

	nsecBuf, _ := ethSw.i2cDevice.ReadFromTwoByteRegister(uint16(firstRegister), 4)
	secBuf, _ := ethSw.i2cDevice.ReadFromTwoByteRegister(uint16(firstRegister+4), 4)
	phaseBuf, _ := ethSw.i2cDevice.ReadFromTwoByteRegister(uint16(firstRegister+8), 4)
	if len(nsecBuf) < 4 || len(secBuf) < 4 || len(phaseBuf) < 4 {
		fmt.Printf("Buffer short\n")
	} else {
		sec := binary.BigEndian.Uint32(secBuf)
		nsec := binary.BigEndian.Uint32(nsecBuf)
		fmt.Printf("Time captured: %s:%d\n", time.Unix(int64(sec), int64(nsec)).String(), binary.BigEndian.Uint32(phaseBuf))
	}
}

func (ethSw *EthSwKSZ9567S) ResetInputEventTrigger() {

	var tsStatusAndControlRegister uint32

	tsStatusAndControlRegister |= 1 << 7 // Enable Rising Edge Detection (TS_RISING_EDGE_ENB) - 1 = Enable rising edge detection
	if err := ethSw.i2cDevice.WriteToTwoByteRegister(TIMESTAMP_STATUS_AND_CONTROL_REGISTER, []byte{
		byte(tsStatusAndControlRegister >> 24),
		byte(tsStatusAndControlRegister >> 16),
		byte(tsStatusAndControlRegister >> 8),
		byte(tsStatusAndControlRegister),
	}); err != nil {
		fmt.Printf("unable to write timestamp status and control register")
	}

	var tsControlAndStatusRegister uint32

	tsControlAndStatusRegister |= 1 << 6 // GPIO Output Enable (GPIO_OEN) - 1 = Enables the GPIO pin as a timestamp input
	tsControlAndStatusRegister |= 1 << 0 // Event Timestamp Input Unit Software Reset (TS_SW_RESET)
	// 1 = Resets the timestamp unit to the inactive state and default settings

	if err := ethSw.i2cDevice.WriteToTwoByteRegister(TIMESTAMP_CONTROL_AND_STATUS_REGISTER, []byte{
		byte(tsControlAndStatusRegister >> 24),
		byte(tsControlAndStatusRegister >> 16),
		byte(tsControlAndStatusRegister >> 8),
		byte(tsControlAndStatusRegister),
	}); err != nil {
		fmt.Printf("unable to write timestamp control and status register")
	}

	tsControlAndStatusRegister = 0
	tsControlAndStatusRegister |= 1 << 6 // GPIO Output Enable (GPIO_OEN) - 1 = Enables the GPIO pin as a timestamp input
	tsControlAndStatusRegister |= 1 << 1 // Event Timestamp Input Unit Enable (TS_ENB) -  1 = Enables the selected event
	// timestamp input unit. Writing “1” to this bit will clear the TS_EVENT_DET_CNT
	// of the associated unit.

	if err := ethSw.i2cDevice.WriteToTwoByteRegister(TIMESTAMP_CONTROL_AND_STATUS_REGISTER, []byte{
		byte(tsControlAndStatusRegister >> 24),
		byte(tsControlAndStatusRegister >> 16),
		byte(tsControlAndStatusRegister >> 8),
		byte(tsControlAndStatusRegister),
	}); err != nil {
		fmt.Printf("unable to write timestamp control and status register")
	}
}
