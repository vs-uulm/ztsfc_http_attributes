package attributes

import (
    "time"
)

// TODO: Implement JSON comptible time; Implement Marhsaler Interface for that.
// See: https://stackoverflow.com/questions/23695479/how-to-format-timestamp-in-outgoing-json
type User struct {
    UserID string `json:"userID"`
    UsualTimeBegin int `json`
    UsualTimeEnd int
    UsualServices []string
}

func NewEmptyUser() *User {
    newUser := new(User)
    newUser.UserID = ""
    newUser.UsualTimeBegin = time.Now()
    newUser.UsualTimeEnd = time.Now()
    newUser.UsualServices = make([]string, 0)
    return newUser
}

func NewUser(_userID string, _usualTimeBegin, _usualTimeEnd time.Time, _usualServices []string) *User {
    newUser := new(User)
    newUser.UserID = _userID
    newUser.UsualTimeBegin = _usualTimeBegin
    newUser.UsualTimeEnd = _usualTimeEnd
    newUser.UsualServices = _usualServices
    return newUser
}
