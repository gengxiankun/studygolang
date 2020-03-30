// 抽象语法表示法
// 抽象语法表示法/1(ASN.1)最初出现在1984年，它是一个为电信行业设计的复杂标准，
// Go的标准包asn1实现了它的一个子集，
// 它可以将复杂的数据结构序列化成自描述的数据。
// 在当前的网络系统中，它主要用于对认证系统中普遍使用的X.509证书的编码。
// Go对ASN.1的支持主要是X.509证书的读写上。
package main

import (
	"fmt"
	"os"
	"encoding/asn1"
)

func main() {
	// 将数据值编组成序列化的字节数组
	// 编组时，我们只需要传递某个类型的变量的值即可
	// 整型的值很容易被编、解组。类似的boolean与real等基本类型处理手法也类似。
	mdata, err := asn1.Marshal(13)
	checkError(err)
	fmt.Println("Marshal: ", mdata)

	// 解组它，则需要一个与被序列化过的数据匹配的确定类型的变量
	// 除了有确定类型的变量外，我们同时需要保证那个变量的内存已经被分配，以使被解组后的数据能有实际被写入的地址。
	var n int
	_, err = asn1.Unmarshal(mdata, &n)
	checkError(err)
	fmt.Println("Unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel error: ", err.Error())
		os.Exit(1)
	}
}