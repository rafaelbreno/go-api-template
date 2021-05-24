package utils

import (
	"fmt"
	"strconv"
)

func StringToUint(str string) (id uint, err error) {
	if str == "" {
		err = fmt.Errorf("Field empty")
	}

	id64, err := strconv.ParseUint(str, 10, 64)

	id = uint(id64)

	return
}
