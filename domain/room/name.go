package room

// Name 部屋名
type Name int

// 部屋名定義
const (
	Susanoo Name = iota
	Amaterasu
	Tsukuyomi
	Unknown
)

// NameFrom 文字列からNameに変換する
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

// String Nameを文字列化する
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
