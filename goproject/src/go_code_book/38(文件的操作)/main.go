package main

import (
	"os"
	"fmt"
	"io"
)

func main() {

	path := "./demo.txt" //当前路径下的demo.txt
	//writeFile(path)
	readFile(path)
}

//写文件
func writeFile(path string) {
	//1、打开文件 没有的话 新建
	file, error := os.Create(path)
	if error != nil {
		fmt.Println(error)
		return
	}
	//2、使用完毕后要进行关闭文件
	defer file.Close()

	//3、写内容
	var buf string
	for i := 0; i < 10; i++ {
		//将数据格式化然后赋值给buf
		buf = fmt.Sprintf("i=%d\n", i)
		n, error := file.WriteString(buf)
		if error != nil {
			fmt.Println(error)
			return
		}
		//写入的长度
		fmt.Println(n)
	}

}

func readFile(path string) {
	file, error := os.Open(path)
	// io.EOF 代表文章的末尾
	//如果文件出错，并且没有到结尾
	if error != nil && error != io.EOF {
		fmt.Println(error)
		return
	}
	//定义一个长度
	buf := make([]byte, 1024*2)
	// n: 读取的长度
	n, error := file.Read(buf);
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(string((buf[:n])))

	fmt.Println(n)
	defer file.Close()
}
