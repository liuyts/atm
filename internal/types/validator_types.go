package types

import (
	"errors"
	"regexp"
	"strconv"
)

func (r *UserRegisterRequest) Validate() error {
	match, _ := regexp.MatchString(`^[1-9]\d{5}(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[1-2]\d|3[0-1])\d{3}(\d|X)$`, r.IdCard)
	if !match {
		return errors.New("身份证号码格式不正确")
	}
	match, _ = regexp.MatchString(`^1[3456789]\d{9}$`, r.Phone)
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
