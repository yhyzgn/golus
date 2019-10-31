// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-10-31 上午9:43
// version: 1.0.0
// desc   : 利用 ANSI 控制控制台输出信息的颜色和风格

package golus

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ESC ASCII 码
	// 也可用 ESC 转义字符 \u001b 代替
	ESC = rune(27)
)

// 字体颜色，背景颜色，字体风格
type Stylus int

// 字体颜色
const (
	// 黑色
	FontBlack Stylus = 30
	// 红色
	FontRed Stylus = 31
	// 绿色
	FontGreen Stylus = 32
	// 黄色
	FontYellow Stylus = 33
	// 蓝色
	FontBlue Stylus = 34
	// 紫色
	FontPurple Stylus = 35
	// 蓝绿
	FontBlueGreen Stylus = 36
	// 白色
	FontWhite Stylus = 37
)

// 背景颜色
const (
	// 黑色
	BackBlack Stylus = 40
	// 红色
	BackRed Stylus = 41
	// 绿色
	BackGreen Stylus = 42
	// 黄色
	BackYellow Stylus = 43
	// 蓝色
	BackBlue Stylus = 44
	// 紫色
	BackPurple Stylus = 45
	// 蓝绿
	BackBlueGreen Stylus = 46
	// 白色
	BackWhite Stylus = 47
)

// 字体风格
const (
	// 粗体
	StyleBold Stylus = 1
	// 斜体
	StyleItalic Stylus = 3
	// 下划线
	StyleUnderLine Stylus = 4
	// 反转
	StyleReverse Stylus = 7
)

// 给任何将要输出的对象加上颜色风格
//
// 结果： ESC [${stylus...}m${value.ToString()}ESC [0m
func Style(value interface{}, stylus ...Stylus) string {
	if stylus == nil || len(stylus) == 0 {
		return convert(value)
	}
	lth := len(stylus)
	var sb strings.Builder
	sb.WriteRune(ESC)
	sb.WriteString("[")
	for index, item := range stylus {
		sb.WriteString(strconv.Itoa(int(item)))
		if index < lth-1 {
			sb.WriteString(";")
		} else {
			sb.WriteString("m")
		}
	}
	sb.WriteString(convert(value))
	sb.WriteRune(ESC)
	sb.WriteString("[0m")
	return sb.String()
}

// 将任何类型 ToString
func convert(value interface{}) string {
	return fmt.Sprint(value)
}
