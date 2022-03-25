package attributes

import (
    logger "github.com/vs-uulm/ztsfc_http_logger"
)

type Device struct {
    DeviceID string `json:"deviceID"`
    CurrentIP string `json:"currentIP"`
    Revoked bool `json:"revoked"`
}

func NewEmptyDevice() (*Device, error) {
    newDevice := new(Device)
    newDevice.DeviceID = ""
    newDevice.CurrentIP = ""
    newDevice.Revoked = true
    return newDevice, nil
}

func NewDevice(_deviceID string, _currentIP string, _revoked bool) (*Device, error) {
    newDevice := new(Device)
    newDevice.DeviceID = _deviceID
    newDevice.CurrentIP = _currentIP
    newDevice.Revoked = _revoked
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
