package eth_sw_KSZ9567S

import "encoding/binary"

func GetRegisterForPort(register int, port int) uint16 {

	return uint16(register | port<<12)
}

func (ethSw *EthSwKSZ9567S) WriteMMDRegister(port int, mmdValue uint16, registerValue uint16, dataValue uint16) {

	setupRegister := GetRegisterForPort(PHY_MMD_SETUP_REGISTER, port)
	dataRegister := GetRegisterForPort(PHY_MMD_DATA_REGISTER, port)

	DeviceSelect := make([]byte, 2)
	binary.BigEndian.PutUint16(DeviceSelect, mmdValue)
	_ = ethSw.i2cDevice.WriteToTwoByteRegister(setupRegister, DeviceSelect)

	RegisterSelect := make([]byte, 2)
	binary.BigEndian.PutUint16(RegisterSelect, registerValue)
	_ = ethSw.i2cDevice.WriteToTwoByteRegister(dataRegister, RegisterSelect)

	DataSelect := make([]byte, 2)
	binary.BigEndian.PutUint16(DataSelect, mmdValue|0x400)
	_ = ethSw.i2cDevice.WriteToTwoByteRegister(setupRegister, DataSelect)

	ValueSelect := make([]byte, 2)
	binary.BigEndian.PutUint16(ValueSelect, dataValue)
	_ = ethSw.i2cDevice.WriteToTwoByteRegister(dataRegister, ValueSelect)
}

func (ethSw *EthSwKSZ9567S) WriteSGMIIRegister(registerValue uint32, dataValue uint16) {

	RegisterSelect := make([]byte, 4)
	binary.BigEndian.PutUint32(RegisterSelect, registerValue)
	_ = ethSw.i2cDevice.WriteToTwoByteRegister(SGMII_CONTROL_REGISTER, RegisterSelect)

	ValueSelect := make([]byte, 2)
	binary.BigEndian.PutUint16(ValueSelect, dataValue)
	_ = ethSw.i2cDevice.WriteToTwoByteRegister(SGMII_DATA_REGISTER, ValueSelect)

}
