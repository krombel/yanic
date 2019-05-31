package data

type DomainDirector struct {
	Enabled            bool   `json:"enabled"`
	SwitchAfter        string `json:"switch_after"`
	SwitchAfterOffline string `json:"switch_after_offline"`
	Target             string `json:"target,omitempty"`
}
