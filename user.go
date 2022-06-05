package attributes

// TODO: Implement JSON comptible time; Implement Marhsaler Interface for that.
// See: https://stackoverflow.com/questions/23695479/how-to-format-timestamp-in-outgoing-json
type User struct {
    UserID string `json:"userID"`
    FailedPWAuthentication int `json:"failedPWAuthentication"`
    UsualTimeBegin int `json:"usualTimeBegin"`
    UsualTimeEnd int `json:"usualTimeEnd"`
    UsualServices []string `json:"usualServices"`
}

func NewEmptyUser() *User {
    newUser := new(User)
    newUser.UserID = ""
    newUser.FailedPWAuthentication = 0
    newUser.UsualTimeBegin = 0
    newUser.UsualTimeEnd = 0
    newUser.UsualServices = make([]string, 0)
    return newUser
}

func NewUser(_userID string, _failedPWAuthentication, _usualTimeBegin, _usualTimeEnd int, _usualServices []string) *User {
    newUser := new(User)
    newUser.UserID = _userID
    newUser.FailedPWAuthentication = _failedPWAuthentication
    newUser.UsualTimeBegin = _usualTimeBegin
    newUser.UsualTimeEnd = _usualTimeEnd
    newUser.UsualServices = _usualServices
    return newUser
}
