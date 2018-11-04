package main

import (
	"strings"
	"fmt"
	"unicode"
)

func testEqualFold()  {
	str := "hello"
	bool := strings.EqualFold(str,"HeLLO")
	fmt.Println(bool)//true
}

func testHasPrefix()  {
	str := "hello"
	bool := strings.HasPrefix(str,"he")
	fmt.Println(bool)//true
	bool = strings.HasPrefix(str,"el")
	fmt.Println(bool)//false
}

func testSuffix()  {
	str := "hello"
	bool := strings.HasSuffix(str,"lo")
	fmt.Println(bool)//true
	bool = strings.HasSuffix(str,"el")
	fmt.Println(bool)//false
}

func testContains()  {
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", "")) //true
	fmt.Println(strings.Contains("", "")) //true
}

func testContainRune()  {
	fmt.Println(strings.ContainsRune("我是王八蛋",'蛋'))
	fmt.Println(strings.ContainsRune("我是王八蛋",'蛋'))
}

func testContainAny()  {
	fmt.Println(strings.ContainsAny("team", "i")) //false
	fmt.Println(strings.ContainsAny("failure", "u&i")) //true
	fmt.Println(strings.ContainsAny("foo", "")) // false
	fmt.Println(strings.ContainsAny("", "")) //false
}

func testCount(){
	fmt.Println(strings.Count("cheese", "e")) //2
	fmt.Println(strings.Count("five", "")) // before & after each rune 5
}

func testIndex()  {
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr")) //-1
}

func testIndexByte()  {
	//中文占3个字节
	fmt.Println(strings.IndexByte("我是a王八蛋",'a')) //6
	fmt.Println(strings.IndexByte("chicken",'c')) //3
}

func testIndexRune(){
	fmt.Println(strings.IndexRune("我是a王八蛋",'王')) //7
	fmt.Println(strings.IndexRune("chicken", 'c')) //3
}

func testFunc()  {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello，世界", f)) //8
	fmt.Println(strings.IndexFunc("Hello，world", f))//5 错的 -> -1
	fmt.Println(strings.IndexFunc("Hello,world", f))//-1
}

func main()  {
	testEqualFold()
}

