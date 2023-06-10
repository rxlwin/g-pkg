package validate

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type f struct{}

func (f f) Str(v string, min int, max int) bool {
	v1 := strings.Trim(v, " ")
	runeCountInString := utf8.RuneCountInString(v1)
	if runeCountInString < min {
		return false
	}
	if max > 0 && runeCountInString > max {
		return false
	}
	return true
}

func (f f) Int(v int, min int, max int) bool {
	if v < min {
		return false
	}
	if max > 0 && v > max {
		return false
	}
	return true
}

func (f f) Int64(v int64, min int, max int) bool {
	if v < int64(min) {
		return false
	}
	if max > 0 && v > int64(max) {
		return false
	}
	return true
}

func (f f) Phone(phone string) bool {
	reg := `^1\d{10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

func (f f) Email(email string) bool {
	//reg:=`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	reg := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(email)
}

func (f f) Date(date string) bool {
	reg := `^[0-9]{4}-[0-9]{2}-[0-9]{2}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(date)
}

func (f f) Enum(val interface{}, list []interface{}) bool {
	for _, l := range list {
		if val == l {
			return true
		}
	}
	return false
}

func (f f) IsMap(val map[string]string) bool {
	if val != nil {
		return true
	}
	return true
}

func (f f) List(val interface{}) bool {
	_, ok := val.([]interface{})
	if ok {
		return true
	}
	return false
}

func (f f) Id(val int64) bool {
	valInt := int(val)
	return f.Int(valInt, 1, 0)
}
