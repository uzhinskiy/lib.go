// Copyright © 2020 Uzhinskiy Boris <boris.ujinsky@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"net"
        "regexp"
	"strings"
	"bytes"
)


// Разделить строку по нескольким разделителям
func MultiSplit(word,delimiters string) []string {
    array := regexp.MustCompile("["+delimiters+"]+").Split(word, -1)
    return array
}

/* сортировка вставкой */
func InsertionSort(inp []rune, n int) []rune {
	j := 0
	for i := 1; i < n; i++ {
		for j = i; j > 0 && inp[j] < inp[j-1]; j-- {
			inp[j-1], inp[j] = inp[j], inp[j-1]
		}

	}
	return inp
}


func GetMaxValueInArray(v []int) int {
	m:=0	
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func GetMinValueInArray(v []int) int {
	m:=0	
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

// min <= i <= max
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

// конвертация строки в INT
func Atoi(s string) int {
	var (
		n uint64
		i int
		v byte
	)
	for ; i < len(s); i++ {
		d := s[i]
		if '0' <= d && d <= '9' {
			v = d - '0'
		} else if 'a' <= d && d <= 'z' {
			v = d - 'a' + 10
		} else if 'A' <= d && d <= 'Z' {
			v = d - 'A' + 10
		} else {
			n = 0
			break
		}
		n *= uint64(10)
		n += uint64(v)
	}
	return int(n)
}

// конвертация строки в INT
func Atoi32(s string) int32 {
	var (
		n uint64
		i int
		v byte
	)
	for ; i < len(s); i++ {
		d := s[i]
		if '0' <= d && d <= '9' {
			v = d - '0'
		} else if 'a' <= d && d <= 'z' {
			v = d - 'a' + 10
		} else if 'A' <= d && d <= 'Z' {
			v = d - 'A' + 10
		} else {
			n = 0
			break
		}
		n *= uint64(10)
		n += uint64(v)
	}
	return int32(n)
}

// конвертация INT во FLOAT
func Float(i int) float64 {
	return float64(i)
}

// Присваиваем INT в *INT
func IntPtr(i int) *int {
	return &i
}

func GetIP(ipport string, xrealip string, xffr string) string {
	if xrealip != "" {
		return xrealip
	} else if xffr != "" {
		return xffr
	} else {
		host, _, _ := net.SplitHostPort(ipport)
		return host
	}
}

// Конвертация strings.Reader в []byte
func ReaderToByte(stream *strings.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
