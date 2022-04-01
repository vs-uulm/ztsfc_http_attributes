package attributes

type System struct {
    ThreatLevel int `json:"threatLevel"`
}

func NewEmptySystem() *System {
    newSystem := new(System)
    newSystem.ThreatLevel = 0
    return newSystem
}

func NewSystem(_threatLevel int) *System {
    newSystem := new(System)
    // TODO: check if _threatLevel makes sense
    newSystem.ThreatLevel = _threatLevel
    return newSystem
}
