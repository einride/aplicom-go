package dprotocol

type GPSFlag uint8

//go:generate gobin -m -run golang.org/x/tools/cmd/stringer -type GPSFlag -trimprefix GPSFlag

const (
	GPSFlagSpeedOverflow      GPSFlag = 0b0000_0001
	GPSFlagMaxSpeedOverflow   GPSFlag = 0b0000_0010
	GPSFlagAssistedGPS        GPSFlag = 0b0000_0100
	GPSFlagJammingDetected    GPSFlag = 0b0000_1000
	GPSFlagDualModeGPSFix     GPSFlag = 0b0001_0000
	GPSFlagDualModeGLONASSFix GPSFlag = 0b0010_0000
	GPSFlagFix                GPSFlag = 0b0100_0000
	GPSFlagCurrentFix         GPSFlag = 0b1000_0000
)