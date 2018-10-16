package main

import (
	"strings"
	"bufio"
	"fmt"
	"bytes"
)

func TestReader() {

	//ReadString delim 也读出来
	//ReadLine \n是不都出来的

	rd := strings.NewReader("hello world\nyes\ryesr")
	brd := bufio.NewReader(rd)

	if data, _, err := brd.ReadLine(); err == nil {
		fmt.Println(brd.Buffered()) //8
		fmt.Println(string(data))   //hello world
	}
	if data, _, err := brd.ReadLine(); err == nil {
		fmt.Println(string(data)) //yesr
	}

	brd.Reset(strings.NewReader("ba ga ya lu"))

	if data, _, err := brd.ReadLine(); err == nil {
		fmt.Println(string(data)) //ba ga ya lu
	}
	fmt.Println(brd.Buffered()) //0 why?

	//下面测试WriteTo
	brd.Reset(strings.NewReader("hi \n xl \n yes \n"))
	buf := bytes.NewBuffer(make([]byte, 4, 8))
	brd.WriteTo(buf)

	data, err := buf.ReadString('\n') //\n 也保函的哦
	if err != nil {
		panic(nil)
	}
	fmt.Println("1: " + data) // hi
	data, _ = buf.ReadString('\n')
	fmt.Println("2: " + data) // xl

	data, err = brd.ReadString('\n')
	if err != nil {
		fmt.Println(data) //因为都写进了writer 所以返回EOF错误
	}

}

func TestWriter() {
	rd := bufio.NewReader(strings.NewReader("我hi \n xl \n yes \n"))
	w := bufio.NewWriter(bytes.NewBuffer(make([]byte, 10, 20)))
	fmt.Println(w.Available()) //4096

	ru, size, _ := rd.ReadRune()
	fmt.Println("size : ", size) // size : 3
	w.WriteRune(ru)
	fmt.Println(w.Available()) //4093
	fmt.Println(w.Buffered())  //3

	// after flush
	w.Flush()
	fmt.Println(w.Available()) //4096
	fmt.Println(w.Buffered())  //0


}

func main() {
	TestWriter()

}
