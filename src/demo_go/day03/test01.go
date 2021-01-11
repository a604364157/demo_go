// file的操作
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//os.FileInfo()接口定义了文件的相关信息

//File操作(带权限)
/*//File代表一个打开的文件对象。

func Create(name string) (file *File, err error)
//Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。

func Open(name string) (file *File, err error)
//Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。

func NewFile(fd uintptr, name string) *File
//NewFile使用给出的Unix文件描述符和名称创建一个文件。

func Pipe() (r *File, w *File, err error)
//Pipe返回一对关联的文件对象。从r的读取将返回写入w的数据。本函数会返回两个文件对象和可能的错误。

func (f *File) Name() string
//Name方法返回（提供给Open/Create等方法的）文件名称。

func (f *File) Stat() (fi FileInfo, err error)
//Stat返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError。

func (f *File) Fd() uintptr
//Fd返回与文件f对应的整数类型的Unix文件描述符。

func (f *File) Chdir() error
//Chdir将当前工作目录修改为f，f必须是一个目录。如果出错，错误底层类型是*PathError。

func (f *File) Chmod(mode FileMode) error
//Chmod修改文件的模式。如果出错，错误底层类型是*PathError。

func (f *File) Chown(uid, gid int) error
//Chown修改文件的用户ID和组ID。如果出错，错误底层类型是*PathError。

func (f *File) Close() error
//Close关闭文件f，使文件不能用于读写。它返回可能出现的错误。

func (f *File) Readdir(n int) (fi []FileInfo, err error)
//Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。如果n>0，Readdir函数会返回一个最多n个成员的切片。这时，如果Readdir返回一个空切片，它会返回一个非nil的错误说明原因。如果到达了目录f的结尾，返回值err会是io.EOF。如果n<=0，Readdir函数返回目录中剩余所有文件对象的FileInfo构成的切片。此时，如果Readdir调用成功（读取所有内容直到结尾），它会返回该切片和nil的错误值。如果在到达结尾前遇到错误，会返回之前成功读取的FileInfo构成的切片和该错误。

func (f *File) Readdirnames(n int) (names []string, err error)
//Readdir读取目录f的内容，返回一个有n个成员的[]string，切片成员为目录中文件对象的名字，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。如果n>0，Readdir函数会返回一个最多n个成员的切片。这时，如果Readdir返回一个空切片，它会返回一个非nil的错误说明原因。如果到达了目录f的结尾，返回值err会是io.EOF。如果n<=0，Readdir函数返回目录中剩余所有文件对象的名字构成的切片。此时，如果Readdir调用成功（读取所有内容直到结尾），它会返回该切片和nil的错误值。如果在到达结尾前遇到错误，会返回之前成功读取的名字构成的切片和该错误。

func (f *File) Truncate(size int64) error
//Truncate改变文件的大小，它不会改变I/O的当前位置。 如果截断文件，多出的部分就会被丢弃。如果出错，错误底层类型是*PathError。*/
func func01() {
	//获取运行时路径
	fmt.Println(os.Args[0])
	//获取当前路径(开发工具的当前路径事配置的GOPATH,实际运行时看环境)
	fmt.Println(os.Getwd())
	fileInfo, err := os.Stat("src\\study\\day03\\resources\\1.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	//文件名
	fmt.Println(fileInfo.Name())
	//文件大小
	fmt.Println(fileInfo.Size())
	//是否是目录
	fmt.Println(fileInfo.IsDir())
	//修改时间
	fmt.Println(fileInfo.ModTime())
	//权限
	fmt.Println(fileInfo.Mode())
}

/*
   文件操作：
   1.路径：
       相对路径：relative
           ab.txt
           相对于当前工程
       绝对路径：absolute
           /Users/ruby/Documents/pro/a/aa.txt

       .当前目录
       ..上一层
   2.创建文件夹，如果文件夹存在，创建失败
       os.MkDir()，创建一层
       os.MkDirAll()，可以创建多层

   3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
       os.Create()，创建文件

   4.打开文件：让当前的程序，和指定的文件之间建立一个连接
       os.Open(filename)
       os.OpenFile(filename,mode,perm)

   5.关闭文件：程序和文件之间的链接断开。
       file.Close()

   5.删除文件或目录：慎用，慎用，再慎用
       os.Remove()，删除文件和空目录
       os.RemoveAll()，删除所有
*/
func func02() {
	fmt.Println("-----------------func02------------------------")
	//绝对路径
	fileName1 := "E:\\workspaces\\demo\\demo_go\\src\\study\\day03\\resources\\1.txt"
	//相对路径(看项目配置,是文件运行还是包运行,相对的起点不一样)
	fileName2 := "src\\study\\day03\\resources\\1.txt"
	//判断是否为绝对路径
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	//返回文件的绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	//创建目录(路径,权限)
	err := os.Mkdir("E:\\workspaces\\demo\\demo_go\\src\\study\\day03\\resources\\test", os.ModePerm)
	if err != nil {
		fmt.Println("发生error:", err)
	}

	//可以创建多层目录
	err2 := os.MkdirAll("E:\\workspaces\\demo\\demo_go\\src\\study\\day03\\resources\\test", os.ModePerm)
	if err2 != nil {
		fmt.Println("发生error:", err2)
	}

	//创建文件,默认0666
	file1, err3 := os.Create("src\\study\\day03\\resources\\test\\1.txt")
	if err3 != nil {
		fmt.Println("发生error:", err3)
	}
	fmt.Println(file1)

	//打开文件(默认只读)
	file2, err4 := os.Open(fileName2)
	if err4 != nil {
		fmt.Println("发生error:", err4)
	}
	fmt.Println(file2)

	//打开文件(文件名,打开方式[读写],文件的权限,不存在则创建)
	file3, err := os.OpenFile(fileName2, os.O_RDWR, os.ModePerm)
	fmt.Println(file3)

	//关闭文件
	if file3 != nil {
		file3.Close()
	}

	//删除
	//os.Remove("")
	//os.RemoveAll("")
}

//io操作 Reader和Writer接口
func func03() {
	fmt.Println("----------------func03-------------------")
	//读取数据
	fileName := "src\\study\\day03\\resources\\1.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	data := make([]byte, 100, 100)
	i := -1
	for {
		i, err = file.Read(data)
		if i == 0 || err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		fmt.Println(string(data[:i]))
	}
}

//文件拷贝
func func04(srcFile, destFile string) (int, error) {
	file, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	fileCopy, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	defer fileCopy.Close()
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("复制完毕")
			break
		} else if err != nil {
			fmt.Println("发生异常")
			return total, err
		}
		total += n
		fileCopy.Write(bs[:n])
	}
	return total, nil
}

//io包下的copy函数
func func05(srcFile, destFile string) (int64, error) {
	file, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	fileCopy, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	defer fileCopy.Close()
	return io.Copy(fileCopy, file)
	//Copy（dst,src） 为复制src 全部到 dst 中。
	//CopyN(dst,src,n) 为复制src 中 n 个字节到 dst
	//CopyBuffer（dst,src,buf）为指定一个buf缓存区，以这个大小完全复制。
}

//ioutil包进行小文件复制(因为一次性读取,文件过大会内存溢出)
func func06(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = ioutil.WriteFile(destFile, input, 0777)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return len(input), nil
}

//基于Seeker接口实现断点续传
func func07(srcFile, destFile string) {
	tempFile := destFile + "temp"
	file, _ := os.Open(srcFile)
	fileCopy, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	fileTemp, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	defer fileCopy.Close()
	//1.读取临时文件中的数据(seek)
	fileTemp.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := fileTemp.Read(bs)
	if err != nil {
		return
	}
	fmt.Println(n1)
	countStr := string(bs[:n1])
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	//2.设置独写的偏移量
	file.Seek(count, 0)
	fileCopy.Seek(count, 0)
	data := make([]byte, 1024, 1024)
	n2 := -1
	n3 := -1
	total := int(count)
	for {
		//3.读取数据
		n2, err = file.Read(data)
		if err == io.EOF {
			fmt.Println("复制完毕")
			fileTemp.Close()
			os.Remove(tempFile)
			break
		}
		//将数据写入目标文件
		n3, _ = fileCopy.Write(data[:n2])
		total += n3
		//将复制的总量,存到临时文件
		fileTemp.Seek(0, io.SeekStart)
		fileTemp.WriteString(strconv.Itoa(total))
	}
}

//bufio包:带缓存的io操作
func func08() {
	fileName := "src\\study\\day03\\resources\\1.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//bufio高效读取
	reader := bufio.NewReader(file)
	bytes := make([]byte, 1024)
	read, err := reader.Read(bytes)
	fmt.Println(string(bytes[:read]))

	//读取一行
	line, prefix, err := reader.ReadLine()
	fmt.Println(line, prefix, err, string(line))

	//读取字符
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(line)
	}

	//读取byte
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		fmt.Println(string(line))
	}
}

//bufio包:带缓存的io操作
func func09() {
	fileName := "src\\study\\day03\\resources\\2.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//bufio高效写入
	writer := bufio.NewWriter(file)
	for i := 1; i < 100; i++ {
		writer.WriteString(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d", i, i, i, i, i, i, i, i, i, i))
	}
	writer.Flush()
}

func func10() {
	// Discard 是一个 io.Writer 接口，调用它的 Write 方法将不做任何事情
	// 并且始终成功返回。
	//var Discard io.Writer = devNull(0)

	// ReadAll 读取 r 中的所有数据，返回读取的数据和遇到的错误。
	// 如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
	// 所有数据，所以不会把 EOF 当做错误处理。
	//func ReadAll(r io.Reader) ([]byte, error)

	// ReadFile 读取文件中的所有数据，返回读取的数据和遇到的错误。
	// 如果读取成功，则 err 返回 nil，而不是 EOF
	//func ReadFile(filename string) ([]byte, error)

	// WriteFile 向文件中写入数据，写入前会清空文件。
	// 如果文件不存在，则会以指定的权限创建该文件。
	// 返回遇到的错误。
	//func WriteFile(filename string, data []byte, perm os.FileMode) error

	// ReadDir 读取指定目录中的所有目录和文件（不包括子目录）。
	// 返回读取到的文件信息列表和遇到的错误，列表是经过排序的。
	//func ReadDir(dirname string) ([]os.FileInfo, error)

	// NopCloser 将 r 包装为一个 ReadCloser 类型，但 Close 方法不做任何事情。
	//func NopCloser(r io.Reader) io.ReadCloser

	// TempFile 在 dir 目录中创建一个以 prefix 为前缀的临时文件，并将其以读
	// 写模式打开。返回创建的文件对象和遇到的错误。
	// 如果 dir 为空，则在默认的临时目录中创建文件（参见 os.TempDir），多次
	// 调用会创建不同的临时文件，调用者可以通过 f.Name() 获取文件的完整路径。
	// 调用本函数所创建的临时文件，应该由调用者自己删除。
	//func TempFile(dir, prefix string) (f *os.File, err error)

	// TempDir 功能同 TempFile，只不过创建的是目录，返回目录的完整路径。
	//func TempDir(dir, prefix string) (name string, err error)
}

func main() {
	func01()
	func02()
	func03()
	func04("src\\study\\day03\\resources\\1.txt", "src\\study\\day03\\resources\\2.txt")
	func05("src\\study\\day03\\resources\\1.txt", "src\\study\\day03\\resources\\3.txt")
	func06("src\\study\\day03\\resources\\1.txt", "src\\study\\day03\\resources\\4.txt")
}
