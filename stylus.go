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
	FontBlack     Stylus = 30 // 黑色
	FontRed                   // 红色
	FontGreen                 // 绿色
	FontYellow                // 黄色
	FontBlue                  // 蓝色
	FontPurple                // 紫色
	FontBlueGreen             // 蓝绿
	FontWhite                 // 白色
)

// 背景颜色
const (
	BackBlack     Stylus = 40 // 黑色
	BackRed                   // 红色
	BackGreen                 // 绿色
	BackYellow                // 黄色
	BackBlue                  // 蓝色
	BackPurple                // 紫色
	BackBlueGreen             // 蓝绿
	BackWhite                 // 白色
)

// 字体风格
const (
	StyleBold      Stylus = 1 // 粗体
	StyleItalic    Stylus = 3 // 斜体
	StyleUnderLine Stylus = 4 // 下划线
	StyleReverse   Stylus = 7 // 反转
)

// 给任何将要输出的对象加上颜色风格
//
// 结果： ESC[${stylus...}m${value.ToString()}ESC[0m
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
