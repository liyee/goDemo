package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("------字符串s中是否包含substr，返回bool值-------")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true

	fmt.Println("-------字符串链接，把slice a通过sep链接起来------")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ",")) //foo,bar,baz

	fmt.Println("-------把s字符串按照sep分割，返回slice------")
	fmt.Printf("%q\n", strings.Split("aa,b,c", ","))                       //["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) //["" "man " "plan " "canal panama"]]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         //[" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            //[""]

	fmt.Println("-------在字符串s中查找sep所在的位置，返回位置值，找不到返回-1------")
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr")) //-1

	fmt.Println("-------重复s字符串count次，最后返回重复的字符串------")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana

	fmt.Println("--------在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换-----")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo

	fmt.Println("--------在s字符串的头部和尾部去除cutset指定的字符串-----")
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) //["Achtung"]

	//去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //Output:Fields are: ["foo" "bar" "baz"]

}

//字符串拼接：性能比较
//strings.Join ≈ strings.Builder > bytes.Buffer > "+" > fmt.Sprintf
func str1() {
	a := []string{"a", "b", "c"}
	//方式1：+
	ret := a[0] + a[1] + a[2]
	fmt.Println(ret)

	//方式2：fmt.Sprintf
	fmt.Sprintf("%s%s%s", a[0], a[1], a[2])

	//方式3：strings.Builder
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	ret = sb.String()
	fmt.Println(ret)

	//方式4：bytes.Buffer
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))
	ret = buf.String()
	fmt.Println(ret)

	//方式5：strings.Join
	ret = strings.Join(a, "")
	fmt.Println(ret)
}
