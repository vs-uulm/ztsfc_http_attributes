package attributes

import (
    logger "github.com/vs-uulm/ztsfc_http_logger"
)

type Device struct {
    DeviceID string `json:"deviceID"`
    CurrentIP string `json:"currentIP"`
    Revoked bool `json:"revoked"`
}

func NewDevice(_deviceID string = "", _currentIP string = "", _revoked bool = false) (*Device, error) {
    newDevice := new(Device)
    newDevice.DeviceID = _deviceID
    newDevice.CurrentIP = _currentIP
    newDevice.Revoked = _revoked
    return newDevice, nil
}

func PrintDevices(sysLogger *logger.Logger) {
    for _, deviceObj := range DevicesByID {
        sysLogger.Infof("%v\n", deviceObj)
    }
}
