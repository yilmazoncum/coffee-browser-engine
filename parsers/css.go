package parsers

import (
	"fmt"
	"strings"
)

func ParseCss(css string) (*Stylesheet, error) {
	var rule Rule
	var ss Stylesheet
	cursor := 0

	for i := 0; i < len(css); i++ {

		if css[i] == '{' {
			rule.Selector = strings.TrimSpace(css[cursor:i])
			cursor = i + 1
		} else if css[i] == '}' {
			rule.Properties = ParseProperties(css[cursor:i])
			ss.Rules = append(ss.Rules, rule)
			cursor = i + 1
		}
	}
	return &ss, nil
}

func ParseProperties(css string) map[string]string {
	var tmp string
	cursor := 0
	properties := make(map[string]string)

	for i := 0; i < len(css); i++ {
		if css[i] == ';' {
			tmp = strings.TrimSpace(css[cursor:i])
			tmpArr := strings.SplitN(tmp, ":", 2)
			properties[tmpArr[0]] = tmpArr[1]
			cursor = i + 1
		}
	}

	return properties
}

func PrintStyle(style *Stylesheet) {
	for i := 0; i < len(style.Rules); i++ {
		fmt.Println(style.Rules[i].Selector)
		for key, value := range style.Rules[i].Properties {
			fmt.Println("   " + key + " " + value)
		}
	}
}
