package room

import (
	"errors"
)

// Name 和名: 部屋名
type Name struct {
	value string
}

// NewName : 部屋名はこの関数を通してインスタンス化します
func NewName(name string) (*Name, error) {
	switch name {
	case "スサノオ":
		return &Name{"スサノオ"}, nil
	case "アマテラス":
		return &Name{"アマテラス"}, nil
	case "ツクヨミ":
		return &Name{"ツクヨミ"}, nil
	default:
		return nil, errors.New("invalid name: " + name)
	}
}
