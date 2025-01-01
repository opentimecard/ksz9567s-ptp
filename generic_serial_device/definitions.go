package generic_serial_device

type ProtocolId int

const (
	PROTO_UNKNOWN ProtocolId = iota
	PROTO_I2C
	PROTO_SPI
)

const ETH_SW_PREFIX = "eth"

type SerialDeviceConfig struct {
	Protocol   ProtocolId
	Path       string
	Address    byte
	CardConfig []string
}
