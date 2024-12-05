package untils

import "encoding/json"

func GetString(obj interface{}) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(marshal)
}
