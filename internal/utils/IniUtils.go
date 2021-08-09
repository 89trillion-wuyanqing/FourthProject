package utils

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"strconv"
	"strings"
)

var (
	//配置信息
	iniFile *ini.File
)

func init() {
	file, e := ini.Load("././config/app.ini")
	if e != nil {
		log.Fatal("Fail to load conf/app.ini" + e.Error())
	}
	iniFile = file
}

func GetSection(sectionName string) *ini.Section {
	section, e := iniFile.GetSection(sectionName)
	if e != nil {
		log.Fatal("未找到对应的配置信息:" + sectionName + e.Error())

	}
	return section
}

func GetSectionMap(sectionName string) map[string]string {
	section, e := iniFile.GetSection(sectionName)
	if e != nil {
		log.Fatal("未找到对应的配置信息:" + sectionName + e.Error())

	}
	section_map := make(map[string]string, 0)
	for _, e := range section.Keys() {
		section_map[e.Name()] = e.Value()
	}
	return section_map
}

/**
获取字符串值
*/
func GetVal(sectionName string, key string) string {
	var temp_val string
	section := GetSection(sectionName)
	if nil != section {
		temp_val = section.Key(key).Value()
	}
	return temp_val
}

/**
获取字符串数组,通过,分割
*/
func GetArr(sectionName string, key string) []string {
	temp_val := GetVal(sectionName, key)
	value := strings.Split(temp_val, ",")
	return value
}

/**
获取布尔值
*/
func GetBool(sectionName string, key string) bool {
	temp_val := GetVal(sectionName, key)
	value, e := strconv.ParseBool(temp_val)
	if nil != e {
		log.Fatal(e)
	}
	return value
}

/**
获取int
*/
func GetInt(sectionName string, key string) int {
	temp_val := GetVal(sectionName, key)
	value, e := strconv.Atoi(temp_val)
	if nil != e {
		log.Fatal(e)
	}
	return value
}

/**
获取int64
*/
func GetInt64(sectionName string, key string) int64 {
	temp_val := GetVal(sectionName, key)
	value, e := strconv.ParseInt(temp_val, 0, 64)
	if nil != e {
		log.Fatal(e)
	}
	return value
}

/**
获取float
*/
func GetFloat(sectionName string, key string) float64 {
	temp_val := GetVal(sectionName, key)
	value, e := strconv.ParseFloat(fmt.Sprintf("%.2f", temp_val), 64)
	if nil != e {
		log.Fatal(e)
	}
	return value
}
