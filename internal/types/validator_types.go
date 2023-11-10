package types

import (
	"errors"
	"regexp"
	"strconv"
)

func (r *UserRegisterRequest) Validate() error {
	regex := `^1[3456789]\d{9}$`
	match, _ := regexp.MatchString(regex, r.Phone)
	if !match {
		return errors.New("手机号格式不正确")
	}
	if len(r.Password) != 6 {
		return errors.New("密码长度必须为6位")
	}
	_, err := strconv.Atoi(r.Password)
	if err != nil {
		return errors.New("密码必须为数字")
	}
	return nil
}
