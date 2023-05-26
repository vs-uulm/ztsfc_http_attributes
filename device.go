package attributes

import (
	logger "github.com/vs-uulm/ztsfc_http_logger"
)

type Device struct {
	DeviceID           string   `json:"deviceID" yaml:"device_id"`
	FailedAuthAttempts int      `json:"FailedAuthAttempts" yaml:"failed_auth_attempts"`
	CurrentIP          string   `json:"currentIP" yaml:"current_ip"`
	Fingerprint        string   `json:"fingerprint" yaml:"fingerprint"`
	Revoked            bool     `json:"revoked" yaml:"revoked"`
	EnterprisePresence bool     `json:"enterprisePresence" yaml:"enterprise_presence"`
	UsualServices      []string `json:"usualServices" yaml:"usual_services"`
	UsualUser          []string `json:"usualUser" yaml:"usual_user"`
	TrustHistory       int      `json:"trustHistory" yaml:"trust_history"`
}

func NewEmptyDevice() *Device {
	newDevice := new(Device)
	newDevice.DeviceID = ""
	newDevice.FailedAuthAttempts = 0
	newDevice.CurrentIP = ""
	newDevice.Fingerprint = ""
	newDevice.Revoked = false
	newDevice.EnterprisePresence = false
	newDevice.UsualServices = make([]string, 0)
	newDevice.UsualUser = make([]string, 0)
	newDevice.TrustHistory = 0
	return newDevice
}

func NewDevice(_deviceID string, _failedAuthAttempts int, _currentIP string, _fingerprint string, _revoked bool, _enterprisePresence bool, _usualServices, _usualUser []string,
	_trustHistory int) (*Device, error) {
	newDevice := new(Device)
	newDevice.DeviceID = _deviceID
	newDevice.FailedAuthAttempts = _failedAuthAttempts
	newDevice.CurrentIP = _currentIP
	newDevice.Fingerprint = _fingerprint
	newDevice.Revoked = _revoked
	newDevice.EnterprisePresence = _enterprisePresence
	newDevice.UsualServices = _usualServices
	newDevice.UsualUser = _usualUser
	newDevice.TrustHistory = _trustHistory
	return newDevice, nil
}

func PrintDevices(sysLogger *logger.Logger, devices map[string]*Device) {
	for _, deviceObj := range devices {
		sysLogger.Infof("%v\n", deviceObj)
	}
}

func FindDeviceByIPInIDMap(sysLogger *logger.Logger, ip string, devices map[string]*Device) *Device {
	for _, device := range devices {
		if device.CurrentIP == ip {
			sysLogger.Debugf("attributes: FindDeviceByIPInMap(): device '%s' has currently IP addr '%s'", device.DeviceID, device.CurrentIP)
			return device
		}
	}

	return nil
}
