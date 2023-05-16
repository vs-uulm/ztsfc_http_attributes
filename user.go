package attributes

import (
	"golang.org/x/time/rate"
)

// TODO: Implement JSON comptible time; Implement Marhsaler Interface for that.
// See: https://stackoverflow.com/questions/23695479/how-to-format-timestamp-in-outgoing-json
type User struct {
	UserID                 string     `json:"userID" yaml:"user_id"`
	FailedPWAuthentication int        `json:"failedPWAuthentication" yaml:"failed_pw_authentication"`
	EnterprisePresence     bool       `json:"enterprisePresence" yaml:"enterprise_presence"`
	UsualTimeBegin         int        `json:"usualTimeBegin" yaml:"usual_time_begin"`
	UsualTimeEnd           int        `json:"usualTimeEnd" yaml:"usual_time_end"`
	UsualAcessRate         rate.Limit `json:"usualAccessRate" yaml:"usual_access_rate"` // describes 'UsualAccessRate' requests per second
	UsualServices          []string   `json:"usualServices" yaml:"usual_services"`
	AllowedServices        []string   `json:"allowedServices" yaml:"allowed_services"`
	UsualDevices           []string   `json:"usualDevices" yaml:"usual_devices"`
	TrustHistory           int        `json:"trustHistory" yaml:"trust_history"`
}

func NewEmptyUser() *User {
	newUser := new(User)
	newUser.UserID = ""
	newUser.FailedPWAuthentication = 0
	newUser.EnterprisePresence = false
	newUser.UsualTimeBegin = 0
	newUser.UsualTimeEnd = 0
	newUser.UsualAcessRate = 0
	newUser.UsualServices = make([]string, 0)
	newUser.AllowedServices = make([]string, 0)
	newUser.UsualDevices = make([]string, 0)
	newUser.TrustHistory = 0
	return newUser
}

func NewUser(_userID string, _failedPWAuthentication int, _enterprisePresence bool, _usualTimeBegin, _usualTimeEnd int, _usualAccessRate rate.Limit,
	_usualServices, _allowedServices, _usualDevices []string, _trustHistory int) *User {
	newUser := new(User)
	newUser.UserID = _userID
	newUser.FailedPWAuthentication = _failedPWAuthentication
	newUser.EnterprisePresence = _enterprisePresence
	newUser.UsualTimeBegin = _usualTimeBegin
	newUser.UsualTimeEnd = _usualTimeEnd
	newUser.UsualAcessRate = _usualAccessRate
	newUser.UsualServices = _usualServices
	newUser.AllowedServices = _allowedServices
	newUser.UsualDevices = _usualDevices
	newUser.TrustHistory = _trustHistory
	return newUser
}
