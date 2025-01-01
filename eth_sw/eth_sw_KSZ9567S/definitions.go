package eth_sw_KSZ9567S

var VerificationReqOne = [...]byte{0x5f, 0x00, 0x00}
var VerificationRespOne = 0x00

var VerificationReqTwo = [...]byte{0x5f, 0x00, 0x01}
var VerificationRespTwo = 0x95

var VerificationReqThree = [...]byte{0x5f, 0x00, 0x02}
var VerificationRespThree = 0x67

const (
	PHY_BASIC_CONTROL_REGISTER = 0x0100
	PHY_MMD_SETUP_REGISTER     = 0x011A
	PHY_MMD_DATA_REGISTER      = 0x011C

	SGMII_CONTROL_REGISTER = 0x7200
	SGMII_DATA_REGISTER    = 0x7206

	GLOBAL_PTP_CLOCK_CONTROL_REGISTER   = 0x0500
	GLOBAL_PTP_RTC_CLOCK_PHASE_REGISTER = 0x0502
	GLOBAL_PTP_RTC_CLOCK_TIME_REGISTER  = 0x0504 // Made up name, just to have a common register for the
	// next 4 which are filled concurrently
	GLOBAL_PTP_RTC_CLOCK_NSEC_HIGH_REGISTER  = 0x0504
	GLOBAL_PTP_RTC_CLOCK_NSEC_LOW_REGISTER   = 0x0506
	GLOBAL_PTP_RTC_CLOCK_SEC_HIGH_REGISTER   = 0x0508
	GLOBAL_PTP_RTC_CLOCK_SEC_LOW_REGISTER    = 0x050A
	GLOBAL_PTP_MESSAGE_CONFIG_1_REGISTER     = 0x0514
	GLOBAL_PTP_MESSAGE_CONFIG_2_REGISTER     = 0x0516
	GLOBAL_PTP_DOMAIN_AND_VERSION_REGISTER   = 0x0518
	GLOBAL_PTP_UNIT_INDEX_REGISTER           = 0x0520
	GPIO_STATUS_MONITOR_REGISTER_0           = 0x0524
	GPIO_STATUS_MONITOR_REGISTER_1           = 0x0528
	TIMESTAMP_CONTROL_AND_STATUS_REGISTER    = 0x052C
	TIMESTAMP_STATUS_AND_CONTROL_REGISTER    = 0x0550
	TIMESTAMP_1ST_SAMPLE_TIME_NSEC_REGISTER  = 0x0554
	TIMESTAMP_1ST_SAMPLE_TIME_SEC_REGISTER   = 0x0558
	TIMESTAMP_1ST_SAMPLE_TIME_PHASE_REGISTER = 0x055C
	TIMESTAMP_2ND_SAMPLE_TIME_NSEC_REGISTER  = 0x0560
	TIMESTAMP_2ND_SAMPLE_TIME_SEC_REGISTER   = 0x0564
	TIMESTAMP_2ND_SAMPLE_TIME_PHASE_REGISTER = 0x0568
	TIMESTAMP_3RD_SAMPLE_TIME_NSEC_REGISTER  = 0x056C
	TIMESTAMP_3RD_SAMPLE_TIME_SEC_REGISTER   = 0x0570
	TIMESTAMP_3RD_SAMPLE_TIME_PHASE_REGISTER = 0x0574
	TIMESTAMP_4TH_SAMPLE_TIME_NSEC_REGISTER  = 0x0578
	TIMESTAMP_4TH_SAMPLE_TIME_SEC_REGISTER   = 0x057C
	TIMESTAMP_4TH_SAMPLE_TIME_PHASE_REGISTER = 0x0580
	TIMESTAMP_5TH_SAMPLE_TIME_NSEC_REGISTER  = 0x0584
	TIMESTAMP_5TH_SAMPLE_TIME_SEC_REGISTER   = 0x0588
	TIMESTAMP_5TH_SAMPLE_TIME_PHASE_REGISTER = 0x058C
	TIMESTAMP_6TH_SAMPLE_TIME_NSEC_REGISTER  = 0x0590
	TIMESTAMP_6TH_SAMPLE_TIME_SEC_REGISTER   = 0x0594
	TIMESTAMP_6TH_SAMPLE_TIME_PHASE_REGISTER = 0x0598
	TIMESTAMP_7TH_SAMPLE_TIME_NSEC_REGISTER  = 0x059C
	TIMESTAMP_7TH_SAMPLE_TIME_SEC_REGISTER   = 0x05A0
	TIMESTAMP_7TH_SAMPLE_TIME_PHASE_REGISTER = 0x05A4
	TIMESTAMP_8TH_SAMPLE_TIME_NSEC_REGISTER  = 0x05A8
	TIMESTAMP_8TH_SAMPLE_TIME_SEC_REGISTER   = 0x05AC
	TIMESTAMP_8TH_SAMPLE_TIME_PHASE_REGISTER = 0x05B0
)