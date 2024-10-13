package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 = "124"
	var num1, err1 = strconv.Atoi(str1)

	if err1 == nil {
		fmt.Println(num1) // 124
	}

	var num2 = 124
	var str2 = strconv.Itoa(num2)

	fmt.Println(str2) // "124"

	var str3 = "124"
	var num3, err3 = strconv.ParseInt(str3, 10, 64)

	if err3 == nil {
		fmt.Println(num3) // 124
	}

	var str4 = "1010"
	var num4, err4 = strconv.ParseInt(str4, 2, 8)

	if err4 == nil {
		fmt.Println(num4) // 10
	}

	var num5 = int64(24)
	var str5 = strconv.FormatInt(num5, 8)

	fmt.Println(str5) // 30

	var str6 = "24.12"
	var num6, err6 = strconv.ParseFloat(str6, 32)

	if err6 == nil {
		fmt.Println(num6) // 24.1200008392334
	}

	var num7 = float64(24.12)
	var str7 = strconv.FormatFloat(num7, 'f', 6, 64)

	fmt.Println(str7) // 24.120000

	var str8 = "true"
	var bul1, err8 = strconv.ParseBool(str8)

	if err8 == nil {
		fmt.Println(bul1) // true
	}

	var bul2 = true
	var str9 = strconv.FormatBool(bul2)

	fmt.Println(str9) // true

	// konversi nilai 24 bertipe int ke float64
	var a float64 = float64(24)
	fmt.Println(a) // 24

	// konversi nilai 24.00 bertipe float32 ke int32
	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	var text1 = "halo"
	var b1 = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b1[0], b1[1], b1[2], b1[3])
	// 104 97 108 111

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)
	// halo

	var c int64 = int64('h')
	fmt.Println(c) // 104

	var d string = string(104)
	fmt.Println(d) // h

	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
