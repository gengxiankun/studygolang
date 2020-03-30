// ASN.1还包含一些未在上边列表中出现的“有用的类型(useful types)”, 
// 比如UTC时间类型，GO支持此UTC时间类型。
// 就是说你可以用一种特有的类型来传递时间值。
// ASN.1不支持指针,Go中却有指向时间值的指针。
// 比如函数GetLocalTime返回*time.Time。
// asn1包编组这个time结构，也使用这个包解组到一个time.Time对象指针中。
package main

import (
	"time"
	"encoding/asn1"
	"fmt"
)

func main() {
	t := time.Now()
	mdata, _ := asn1.Marshal(t)
	fmt.Println("Marshal: ", mdata)

	// Now与new函数都返回的是*time.Time类型的指针，GO将内部对这些特殊类型进行处理。
	var newtime = new(time.Time)
	asn1.Unmarshal(mdata, &newtime)
	fmt.Println("Unmarshal: ", newtime)
}