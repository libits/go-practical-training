package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	//默认大小为100MB
	re, _ := regexp.Compile("[0-9]+")
	//单位
	unit := string(re.ReplaceAll([]byte(size), []byte("")))

	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	//大小写转化
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB

	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}
	if num == 0 {
		log.Println("ParseSize仅支持B、KB、MB、GB、TB、PB")
		num = 100
		byteNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}
func GetValueSize(val interface{}) int64 {
	//TODO
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	//size := unsafe.Sizeof(val)
	fmt.Println(size)
	return 0
}