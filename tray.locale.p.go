package main

import "fmt"

type Messages map[string]string

var locale = "en"

func setLocale(name string) (ok bool) {
	switch name {
	case "en", "zh-Hans":
		locale = name
		return true
	}
	return false
}

func getMessage(key string) string {
	switch locale {
	case "en":
		return en[key]
	case "zh-Hans":
		return zh_Hans[key]
	}
	return "tray.locale.getMessage(key)???"
}

func getMessageWithParams(key string, params ...interface{}) string {
	switch locale {
	case "en":
		return fmt.Sprintf(en[key], params...)
	case "zh-Hans":
		return fmt.Sprintf(zh_Hans[key], params...)
	}
	return "tray.locale.getMessage(key,params)???"
}
