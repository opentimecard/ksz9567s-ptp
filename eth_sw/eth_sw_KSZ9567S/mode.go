package eth_sw_KSZ9567S

func (ethSw *EthSwKSZ9567S) SetSerDesMode() error {

	ethSw.WriteSGMIIRegister(0x001F8001, 0x0019)
	ethSw.WriteSGMIIRegister(0x001F0004, 0x01A0)
	ethSw.WriteSGMIIRegister(0x001F0000, 0x1340)
	return nil
}
