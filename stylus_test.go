// Copyright 2019 yhyzgn golus
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-11-04 7:38 下午
// version: 1.0.0
// desc   : 

package golus

import (
	"fmt"
	"testing"
)

type Test struct {
	Info string
}

func TestStyle(t *testing.T) {
	test := &Test{Info: "Test"}

	a := New().Apply(test)
	fmt.Println(a)

	a = New().FontStyle(styleBold).Apply(test, test)
	fmt.Println(a)

	b := New().FontColor(fontCyan).Apply(test)
	fmt.Println(b)

	c := New().BackColor(backCyan).Apply(test)
	fmt.Println(c)

	d := New().FontStyle(styleBold, styleItalic, styleUnderLine, styleReverse).Apply(test)
	fmt.Println(d)

	e := New().FontColor(fontCyan).BackColor(backYellow).FontStyle(styleBold, styleItalic, styleUnderLine).Apply(test)
	fmt.Println(e)
}
