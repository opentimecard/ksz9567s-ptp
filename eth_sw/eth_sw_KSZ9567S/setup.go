package eth_sw_KSZ9567S

import (
	"fmt"
	"github.com/opentimecard/ksz9567s-ptp/generic_serial_device"
	"strings"
)

type Internals struct {
	useTc bool
}

type EthSwKSZ9567S struct {
	i2cDevice    *generic_serial_device.I2CDevice
	deviceConfig generic_serial_device.SerialDeviceConfig
	internals    Internals
}

func NewEthSwKSZ9567S(serialDeviceConfig *generic_serial_device.SerialDeviceConfig) *EthSwKSZ9567S {

	device := generic_serial_device.NewI2CDevice(serialDeviceConfig.Path, serialDeviceConfig.Address)
	if device == nil {
		return nil
	}

	ethSwKSZ9567S := &EthSwKSZ9567S{
		i2cDevice:    device,
		deviceConfig: *serialDeviceConfig,
		internals: Internals{
			useTc: true,
		},
	}
	return ethSwKSZ9567S
}

func (ethSw *EthSwKSZ9567S) Start(done chan struct{}) error {

	ethSw.ParseConfig()

	err := ethSw.VerifyConnectedToDevice()
	if err != nil {
		return err
	}

	ethSw.RunAllErrata()
	_ = ethSw.SetSerDesMode()

	_ = ethSw.SetPTPClock()
	if ethSw.internals.useTc {
		_ = ethSw.EnableTransparentClockMode()
	} else {
		_ = ethSw.DisableTransparentClockMode()
	}

	//go ethSw.Run1PPSInRunloop(done)

	return nil
}

func (ethSw *EthSwKSZ9567S) VerifyConnectedToDevice() error {

	response, err := ethSw.i2cDevice.WriteThenRead(VerificationReqOne[:], 1)
	if err != nil || int(response[0]) != VerificationRespOne {
		return err
	}

	response, err = ethSw.i2cDevice.WriteThenRead(VerificationReqTwo[:], 1)
	if err != nil || int(response[0]) != VerificationRespTwo {
		return err
	}

	response, err = ethSw.i2cDevice.WriteThenRead(VerificationReqThree[:], 1)
	if err != nil || int(response[0]) != VerificationRespThree {
		return err
	}

	return nil
}

func (ethSw *EthSwKSZ9567S) ParseConfig() {

	for _, configOption := range ethSw.deviceConfig.CardConfig {
		if strings.HasPrefix(configOption, generic_serial_device.ETH_SW_PREFIX) {

			configOptionParts := strings.Split(configOption, ":")
			if len(configOptionParts) < 2 {
				continue
			}

			switch configOptionParts[1] {

			case "sfp_accuracy":

				if len(configOptionParts) <= 3 {
					if strings.ToLower(configOptionParts[2]) == "precise" {
						ethSw.internals.useTc = false
						fmt.Printf("Setting SFP accuracy to precise\n")
					}
				}
			}
		}
	}
}
