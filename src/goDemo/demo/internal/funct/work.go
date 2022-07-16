package funct

import (
	"errors"
	"fmt"
	"strconv"
)

func Monthwork(month string) error {
	mnumb, err := strconv.Atoi(month)
	if err != nil {
		return errors.New("转换异常, " + month)
	} else {
		fmt.Println(mnumb)
		return nil
	}
}

func Work() {

}
