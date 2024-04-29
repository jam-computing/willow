package player

import (
	"encoding/json"
	"fmt"
)

type Animation struct {
	Name      string        `json:"name"`
	Artist    string        `json:"artist"`
	Tick_rate uint16        `json:"tick_rate"`
	Frames    []interface{} `json:"frames"`
}

func MakeAnimation(data string) Animation {
	var a Animation
	err := json.Unmarshal([]byte(data), &a)

	if err != nil {
		fmt.Println(err)
		return Animation{}
	}

	return a
}
