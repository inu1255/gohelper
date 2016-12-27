package gohelper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func Json2struct(s interface{}, classname string) error {
	var info map[string]interface{}
	switch v := s.(type) {
	case string:
		err := json.Unmarshal(Str2bytes(v), &info)
		if err != nil {
			ioutil.WriteFile(classname+".html", Str2bytes(v), 0655)
			return err
		}
	case map[string]interface{}:
		info = v
	default:
		return errors.New("unknow")
	}
	var f func(interface{}) string
	f = func(value interface{}) string {
		switch v := value.(type) {
		case bool:
			return "bool\n"
		case int:
			return "int\n"
		case int8:
			return "int8\n"
		case int16:
			return "int16\n"
		case int32:
			return "int32\n"
		case int64:
			return "int64\n"
		case float32:
			return "float32\n"
		case float64:
			return "float64\n"
		case string:
			return "string\n"
		case []interface{}:
			if len(v) > 0 {
				return "[]" + f(v[0])
			}
			return "[]interface{}\n"
		case map[string]interface{}:
			s := " struct{\n"
			for k, value := range v {
				s += UpperFirstLetter(k) + " " + f(value)
			}
			s += "}\n"
			return s
		default:
			return "interface{}\n"
		}
	}
	ioutil.WriteFile(classname+".go", Str2bytes("package main\ntype "+UpperFirstLetter(classname)+" "+f(info)), 0655)
	return nil
}
