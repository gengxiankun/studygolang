// 由ASCII字符构成的字符串也很容易。
// 但当处理 "hello \u00bc"这种含有 '¼'这个非ASCII字符的字符串，则会出现错误：“ASN.1 结构错误:PrintableString包含非法字符”。
// 以下的代码仅在处理由可打印字符（printable characters）构成的字符串时,工作良好。
package main

import (
	"fmt"
	"encoding/asn1"
)

func main() {
	s := "hello"
	mdata, _ := asn1.Marshal(s)
	fmt.Println("Marshal: ", mdata)

	var newstr string
	asn1.Unmarshal(mdata, &newstr)
	fmt.Println("Unmarshal: ", newstr)
}