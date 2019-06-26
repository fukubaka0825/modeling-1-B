package room

// Neme 部屋名
type Name int

const (
	Susanoo = iota
	Amaterasu
	Tsukuyomi
	Unknown
)

func NameFrom(nameString string) Name {
	switch nameString {
	case "スサノオ":
		return Susanoo
	case "アマテラス":
		return Amaterasu
	case "ツクヨミ":
		return Tsukuyomi
	default:
		return Unknown
	}
}

func (name Name) String() string {
	switch name {
	case Susanoo:
		return "スサノオ"
	case Amaterasu:
		return "アマテラス"
	case Tsukuyomi:
		return "ツクヨミ"
	default:
		return ""
	}
}
