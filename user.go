package attributes

// TODO: Implement JSON comptible time; Implement Marhsaler Interface for that.
// See: https://stackoverflow.com/questions/23695479/how-to-format-timestamp-in-outgoing-json
type User struct {
    UserID string `json:"userID" yaml:"user_id"`
    FailedPWAuthentication int `json:"failedPWAuthentication" yaml:"failed_pw_authentication"`
    UsualTimeBegin int `json:"usualTimeBegin" yaml:"usual_time_begin"`
    UsualTimeEnd int `json:"usualTimeEnd" yaml:"usual_time_end"`
    UsualServices []string `json:"usualServices" yaml:"usual_services"`
    AllowedServices []string `json:"allowedServices" yaml:"allowed_services"`
}

func NewEmptyUser() *User {
    newUser := new(User)
    newUser.UserID = ""
    newUser.FailedPWAuthentication = 0
    newUser.UsualTimeBegin = 0
    newUser.UsualTimeEnd = 0
    newUser.UsualServices = make([]string, 0)
    newUser.AllowedServices = make([]string, 0)
    return newUser
}

func NewUser(_userID string, _failedPWAuthentication, _usualTimeBegin, _usualTimeEnd int, _usualServices, _allowedServices []string) *User {
    newUser := new(User)
    newUser.UserID = _userID
    newUser.FailedPWAuthentication = _failedPWAuthentication
    newUser.UsualTimeBegin = _usualTimeBegin
    newUser.UsualTimeEnd = _usualTimeEnd
    newUser.UsualServices = _usualServices
    newUser.AllowedServices = _allowedServices
    return newUser
}
