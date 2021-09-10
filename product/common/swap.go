package common

import "encoding/json"

func SwapTo(request, category interface{}) error {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, category)
}
