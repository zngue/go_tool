package common

import "encoding/json"

func StuckToMap(i interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	b,_:=json.Marshal(i)
	json.Unmarshal(b,&m)
	return m
}
