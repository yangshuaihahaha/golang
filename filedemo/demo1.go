package main

func main() {
	//file, error := os.Open("/Users/yangshuai/go/src/demo3/filedemo/text.text")
	//if error != nil {
	//	fmt.Println("error", error)
	//}
	//defer file.Close()
	//读取文件(带缓冲的方式)
	//reader := bufio.NewReader(file);
	//for {
	//	str, error := reader.ReadString('\n')
	//	if error != nil {
	//		fmt.Println(error)
	//	}
	//	if error == io.EOF {
	//		break
	//	}
	//	fmt.Print(str)
	//}

	//不带缓冲的方式(不用打开关闭文件 file.open() file.close() )
	//content, error := ioutil.ReadFile("/Users/yangshuai/go/src/demo3/filedemo/text.text")
	//if error != nil {
	//	fmt.Println(error)
	//}
	//fmt.Println(string(content))// []byte

	//创建新的文件
	//filePath := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"
	//file, error := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	//if error != nil {
	//	fmt.Println("创建文件失败", error)
	//	return
	//}
	////及时关闭file, 防止句柄泄漏
	//defer file.Close()
	////准备进行写入
	//writer := bufio.NewWriter(file)
	//writer.WriteString("hello\n")
	//writer.WriteString("hello\n")
	//writer.WriteString("hello\n")
	//writer.WriteString("hello\n")
	//writer.WriteString("hello\n")
	////因为是带缓存的, 所以要调用flush
	//writer.Flush()

	//创建一个文件覆盖里面的内容
	//filePath := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"
	////O_TRUNC会覆盖原本的内容
	//file, error := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	//if error != nil {
	//	fmt.Println(error)
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(file)
	//writer.WriteString("hahahha")
	//writer.Flush()

	//创建一个文件追加里面的内容
	//filePath := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"
	////O_TRUNC会覆盖原本的内容
	//file, error := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	//if error != nil {
	//	fmt.Println(error)
	//}
	//defer file.Close()
	//writer := bufio.NewWriter(file)
	//writer.WriteString("nihao\n")
	//writer.Flush()

	//打开一个文件读出里面的内容并且追加文字
	//filePath := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"
	//file, error := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	//if error != nil {
	//	fmt.Println(error)
	//}
	//defer file.Close()
	////读取文件
	//reader := bufio.NewReader(file)
	//for {
	//	str, error := reader.ReadString('\n')
	//	if error == io.EOF {
	//		break
	//	}
	//	fmt.Print(str)
	//}
	////追加文件
	//writer := bufio.NewWriter(file)
	//writer.WriteString("oqwieurqoeiru\n")
	//writer.Flush()

	//判断文件是否存在
	//filePath := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"
	//_, error := os.Stat(filePath)
	//if error == nil {
	//	fmt.Println("文件存在")
	//}
	//if os.IsNotExist(error) {
	//	fmt.Println("文件不存在")
	//}

	//拷贝文件
	//srcFileName := "/Users/yangshuai/go/src/demo3/filedemo/abc.text"
	//srcFile, erro := os.Open(srcFileName)
	//if erro != nil {
	//	fmt.Println(erro)
	//}
	//defer srcFile.Close()
	////通过srcFile获取到reader
	//reader := bufio.NewReader(srcFile)
	//dstFileName := "/Users/yangshuai/go/src/demo3/filedemo/def.text"
	//dstFile, erro := os.Create(dstFileName)
	//if erro != nil {
	//	fmt.Println(dstFile)
	//}
	//defer dstFile.Close()
	////writer := bufio.NewWriter(dstFile)
	//
	//_, error := io.Copy(dstFile, reader)
	//if error != nil {
	//	fmt.Println(error)
	//}
}
