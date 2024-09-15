package halpers

import (
	"encoding/json"
	"strconv"
)

func DataParser1[T1 any, T2 any](src T1, dst T2) error {
	bytData, err := json.Marshal(src)
	if err != nil {
		return err
	}
	json.Unmarshal(bytData, dst)
	return nil
}

func GetLimit(s string)int32{
	limit, err := strconv.Atoi(s)
	if err != nil {
		limit = 10
	}
	return int32(limit)
}

func GetPage(s string)int32{
	page, err := strconv.Atoi(s)
	if err != nil {
		page = 1
	}
	return int32(page)
}