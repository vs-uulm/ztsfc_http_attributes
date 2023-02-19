package attributes

type System struct {
    ThreatLevel int64 `json:"threatLevel" yaml:"threat_level"`
}

func NewEmptySystem() *System {
    newSystem := new(System)
    newSystem.ThreatLevel = 0
    return newSystem
}

func NewSystem(_threatLevel int64) *System {
    newSystem := new(System)
    // TODO: check if _threatLevel makes sense
    newSystem.ThreatLevel = _threatLevel
    return newSystem
}
