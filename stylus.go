// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-10-31 上午9:43
// version: 1.0.0
// desc   : 利用 ANSI 控制控制台输出信息的颜色和风格
//
// https://blog.csdn.net/Mculover666/article/details/105433609
// https://blog.csdn.net/qq_18361349/article/details/107385022

package golus

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
)

const (
	// esc ASCII 码
	// 也可用 esc 转义字符 \u001b 代替
	esc = rune(27)
)

// fontColor 字体颜色
type fontColor uint8

// backColor 背景颜色
type backColor uint8

// fontStyle 字体风格
type fontStyle uint8

// 字体颜色
const (
	fontBlack   fontColor = 30 + iota // 黑色
	fontRed                           // 红色
	fontGreen                         // 绿色
	fontYellow                        // 黄色
	fontBlue                          // 蓝色
	fontMagenta                       // 紫红色
	fontCyan                          // 青蓝色
	fontWhite                         // 白色
)

// 背景颜色
const (
	backBlack   backColor = 40 + iota // 黑色
	backRed                           // 红色
	backGreen                         // 绿色
	backYellow                        // 黄色
	backBlue                          // 蓝色
	backMagenta                       // 紫红色
	backCyan                          // 青蓝色
	backWhite                         // 白色
)

// 字体风格
const (
	styleBold      fontStyle = 1 // 粗体
	styleItalic    fontStyle = 3 // 斜体
	styleUnderLine fontStyle = 4 // 下划线
	styleReverse   fontStyle = 7 // 反转
)

// Stylus 字体风格类
type Stylus struct {
	stylus []uint8
}

// New 创建个风格对象
func New() *Stylus {
	return &Stylus{
		stylus: make([]uint8, 0),
	}
}

// FontColor 字体颜色
func (s *Stylus) FontColor(fontColor fontColor) *Stylus {
	if fontColor >= fontBlack && fontColor <= fontWhite {
		s.stylus = append(s.stylus, uint8(fontColor))
	}
	return s
}

// BackColor 背景颜色
func (s *Stylus) BackColor(backColor backColor) *Stylus {
	if backColor >= backBlack && backColor <= backWhite {
		s.stylus = append(s.stylus, uint8(backColor))
	}
	return s
}

// FontStyle 字体风格
func (s *Stylus) FontStyle(styles ...fontStyle) *Stylus {
	if styles == nil || len(styles) == 0 {
		return s
	}
	for _, item := range styles {
		if item == styleBold || item == styleItalic || item == styleUnderLine || item == styleReverse {
			s.stylus = append(s.stylus, uint8(item))
		}
	}
	return s
}

// Apply 给任何将要输出的对象加上颜色风格
//
// 结果： esc[${fontColor;backColor;fontStyle}m${value.ToString()}esc[0m
func (s *Stylus) Apply(value ...interface{}) string {
	if len(s.stylus) == 0 {
		return sliceString(value...)
	}
	lth := len(s.stylus)
	var sb strings.Builder
	sb.WriteRune(esc)
	sb.WriteString("[")
	for index, item := range s.stylus {
		sb.WriteString(strconv.Itoa(int(item)))
		if index < lth-1 {
			sb.WriteString(";")
		} else {
			sb.WriteString("m")
		}
	}
	sb.WriteString(sliceString(value...))
	sb.WriteRune(esc)
	sb.WriteString("[0m")
	return sb.String()
}

func sliceString(values ...interface{}) string {
	if nil == values {
		return ""
	}

	var sb strings.Builder
	for _, val := range values {
		sb.WriteString(toString(val))
	}
	return sb.String()
}

func toString(value interface{}) string {
	value = indirectToStringerOrError(value)

	switch val := value.(type) {
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case int32:
		return strconv.Itoa(int(val))
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint:
		return strconv.FormatInt(int64(val), 10)
	case uint64:
		return strconv.FormatInt(int64(val), 10)
	case uint32:
		return strconv.FormatInt(int64(val), 10)
	case uint16:
		return strconv.FormatInt(int64(val), 10)
	case uint8:
		return strconv.FormatInt(int64(val), 10)
	case []byte:
		return string(val)
	case template.HTML:
		return string(val)
	case template.URL:
		return string(val)
	case template.JS:
		return string(val)
	case template.CSS:
		return string(val)
	case template.HTMLAttr:
		return string(val)
	case nil:
		return ""
	case fmt.Stringer:
		return val.String()
	case error:
		return val.Error()
	default:
		return fmt.Sprint(value)
	}
}

func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
