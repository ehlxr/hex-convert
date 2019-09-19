// Copyright © 2018 ehlxr <ehlxr.me@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package converter

import (
	"github.com/ehlxr/hex-convert/metadata"
	"strconv"
)

func ToDecimal(scale int, data string) (int, error) {
	// var newNum float64
	// newNum = 0.0
	// nNum := len(strings.Split(data, "")) - 1
	// for _, value := range strings.Split(data, "") {
	// 	tmp := float64(findkey(value))
	// 	if tmp != -1 {
	// 		newNum = newNum + tmp*math.Pow(float64(scale), float64(nNum))
	// 		nNum = nNum - 1
	// 	} else {
	// 		break
	// 	}
	// }
	// return int(newNum), nil

	intSize := 32 << uint(^uint(0)>>63)
	num, err := strconv.ParseInt(data, scale, intSize)

	return int(num), err
}

func ToBinary(scale int, data string) (string, error) {
	result, err := ToDecimal(scale, data)
	if err != nil {
		return "", err
	}

	return fromDecimal(2, result)
}

func ToHex(scale int, data string) (string, error) {
	result, err := ToDecimal(scale, data)
	if err != nil {
		return "", err
	}

	return fromDecimal(16, result)
}

func ToOctal(scale int, data string) (string, error) {
	result, err := ToDecimal(scale, data)
	if err != nil {
		return "", err
	}

	return fromDecimal(8, result)
}

func findkey(in string) int {
	result := -1
	for k, v := range metadata.TEN_TO_ANY {
		if in == v {
			result = k
		}
	}
	return result
}

func fromDecimal(scale, data int) (string, error) {
	newNumStr := ""
	var remainder int
	var remainderString string
	for data != 0 {
		remainder = data % scale
		if 76 > remainder && remainder > 9 {
			remainderString = metadata.TEN_TO_ANY[remainder]
		} else {
			remainderString = strconv.Itoa(remainder)
		}
		newNumStr = remainderString + newNumStr
		data = data / scale
	}

	return newNumStr, nil
}
