package athletes

import "strings"

// Info ... (exported struct)
type Info struct {
	Country    string
	HairColour string
}

// Player ... (exported struct)
type Player struct {
	Name  string
	Sport string
	Age   string
	Info
}

// ToLowercase ... (exported method receiver)
func (p *Player) ToLowercase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairColour = strings.ToLower(p.HairColour)

	return p
}
