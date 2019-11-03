package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type User struct {
	Name string `json:"name"`
}

// 参考连接https://www.cnblogs.com/craneboos/p/9849714.html
// 文件操作
func main() {
	//createFile()
	//ReadFile()
	//stat()
	//rename()
	//removeFile()
	//move()
	//openFile()
	//isExistFile()
	//readFile()
	//writeFile()
	//writeAppendFile()
	//copyFile()
	//checkPermission()
	//fileAndDirMoodify()
	//linkAndSymLink()
	//seek()
	//scanner()
	//mkdir()
	//zipTest()
	//httpDownLoad()
	crptoTest()
}

//哈希和摘要
func crptoTest() {

	TestString := "Hi, pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	// 得到文件内容
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 计算Hash
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))

	//上面的例子复制整个文件内容到内存中，传递给hash函数。
	//另一个方式是创建一个hash writer, 使用Write、WriteString、Copy将数据传给它。
	//下面的例子使用 md5 hash,但你可以使用其它的Writer。
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//创建一个新的hasher,满足writer接口
	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	// 计算hash并打印结果。
	// 传递 nil 作为参数，因为我们不通参数传递数据，而是通过writer接口。
	sum := hasher.Sum(nil)
	fmt.Printf("Md5 checksum: %x\n", sum)
}

// 通过HTTP下载文件
func httpDownLoad() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	url := "http://www.devdungeon.com/archive"
	response, err := http.Get(url)
	defer response.Body.Close()

	// 将HTTP response Body中的内容写入到文件
	// Body满足reader接口，因此我们可以使用ioutil.Copy
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}

//打包(zip) 文件
func zipTest() {
	// 创建一个打包文件
	outFile, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// 创建zip writer
	zipWriter := zip.NewWriter(outFile)

	// 往打包文件中写文件。
	// 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file"},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	// 下面将要打包的内容写入到打包文件中，依次写入。
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 解压缩
func unZip() {
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// 遍历打包文件中的每一文件/文件夹
	for _, file := range zipReader.Reader.File {
		// 打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// 指定抽取的文件名。
		// 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
		// 我们这个例子使用打包文件中相同的文件名。
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// 抽取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			// 创建文件夹并设置同样的权限
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			//抽取正常的文件
			log.Println("Extracting file:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// 通过io.Copy简洁地复制文件内容
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// 还支持 zlib, bz2, flate, lzw
func gzipTest() {
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
	// 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Compressed data written to file.")
}

func unGzipTest() {
	// 打开一个gzip文件。
	// 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，
	// 它的内容不是一个文件，而是一个内存流
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	// 解压缩到一个writer,它是一个file writer
	outfileWriter, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()

	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}

// 创建目录.临时目录
func mkdir() {
	uploadDir := filepath.Join("static/upload/", time.Now().Format("2006/01/02/"))
	os.MkdirAll(uploadDir, os.ModePerm)
	join := filepath.Join(uploadDir, "java")
	os.MkdirAll(join, os.ModePerm)
	os.Create(join + "/a.txt")
	// 创建单层目录
	os.Mkdir("java", os.ModePerm)
	os.Create("java/te.txt")
	// 移除java目录的文件，不会删除java目录，如果需要删除java，需要单独删除
	os.Remove("java/te.txt")
	//// 移除指定目录包括当前指定目录、当前目录下所有目录以及文件
	os.RemoveAll(uploadDir)
	//// 获取当前目录
	log.Println(os.Getwd())
	// 创建临时目录
	ioutil.TempDir("java", "a")
	// 创建临时文件
	ioutil.TempFile("java", "*")

	// 检查文件夹是否是目录
	dir, _ := os.Getwd()
	file, e := os.Stat(dir)
	if e != nil {
		log.Fatalln(e)
	}
	log.Println(file.IsDir())
}

// 使用scanner去取
func scanner() {
	//Scanner是bufio包下的类型,在处理文件中以分隔符分隔的文本时很有用。
	//通常我们使用换行符作为分隔符将文件内容分成多行。在CSV文件中，逗号一般作为分隔符。
	//os.File文件可以被包装成bufio.Scanner，它就像一个缓存reader。
	//我们会调用Scan()方法去读取下一个分隔符，使用Text()或者Bytes()获取读取的数据。
	//分隔符可以不是一个简单的字节或者字符，有一个特殊的方法可以实现分隔符的功能，以及将指针移动多少，返回什么数据。
	//如果没有定制的SplitFunc提供，缺省的ScanLines会使用newline字符作为分隔符，其它的分隔函数还包括ScanRunes和ScanWords,皆在bufio包中。
	file, _ := os.Open("test.txt")
	newScanner := bufio.NewScanner(file)
	// 设置以什么方式分割
	newScanner.Split(bufio.ScanWords)
	for {
		if scan := newScanner.Scan(); scan {
			log.Println(newScanner.Text()) // newScanner.Bytes() 读取字节
		}
	}
}

// 指定读取文件的位置
func seek() {
	file, e := os.Open("test.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 26
	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	var whence int = 0
	position, e := file.Seek(offset, whence)
	log.Println(position)

	// 读取光标之后的文件，每次读取多少个字节，这里表示7个字节
	bytes := make([]byte, 9)
	read, e := file.Read(bytes)
	if e != nil {
		log.Fatal(e)
	}
	bs := bytes[:read]
	log.Println("读取的文字", string(bs))

	// 从当前位置回退两个字节
	newPosition, err := file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPosition)

	// 使用下面的技巧得到当前的位置
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("Current position:", currentPosition)

	// 转到文件开始处
	newPosition1, err := file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition1)
}

// 创建软连接和硬链接,如果存在软连接或者硬链接会报错
func linkAndSymLink() {
	// 创建一个硬链接
	// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
	// 删除和重命名不会影响另一个。
	link := os.Link("test.txt", "test_link.txt")
	if link != nil {
		log.Fatal(link)
	}

	// 创建一个软连接
	symlink := os.Symlink("test.txt", "t.txt")
	if symlink != nil {
		log.Fatal(symlink)
	}

	// Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
	// Symlink在Windows中不工作。
	fileInfo, err := os.Lstat("t.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	//改变软链接的拥有者不会影响原始文件。
	err = os.Lchown("t.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}

//改变文件或者文件的权限、拥有者、时间戳
func fileAndDirMoodify() {
	// 改变文件的权限 该文件只有只读权限,改成可读可写
	chmod := os.Chmod("test_permisson.txt", 0620)
	if chmod != nil {
		log.Fatal("修改文件权限失败")
	}

	// 修改拥有者
	e := os.Chown("test_permisson.txt", os.Geteuid(), os.Getegid())
	if e != nil {
		log.Fatal(e)
	}

	// 修改文件时间
	x := os.Chtimes("test_permisson.txt", time.Now(), time.Now())
	if x != nil {
		log.Fatal(e)
	}
}

// 检查读写权限
func checkPermission() {
	// w 权限
	os.OpenFile("test_permisson.txt", os.O_WRONLY|os.O_CREATE, 0222)
	// r+w 权限
	os.OpenFile("test_permisson1.txt", os.O_WRONLY|os.O_CREATE, 0622)
	// rw+r+r 权限
	os.OpenFile("test_permisson2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// rwx+r+r 权限
	os.OpenFile("test_permisson3.txt", os.O_WRONLY|os.O_CREATE, 0766)
	// 这个例子测试写权限，如果没有写权限则返回error。
	// 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
	file, e := os.OpenFile("test_permisson.txt", os.O_RDONLY, 0666)
	if e != nil {
		if os.IsPermission(e) {
			log.Fatalln(e)
		}
	}
	defer file.Close()
}

func copyFile() {
	srcFile, e := os.Open("test_read.txt")
	if e != nil {
		log.Fatalln(e)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	destFile, e := os.OpenFile("test_read2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		log.Fatalln(e)
	}
	defer destFile.Close()
	writer := bufio.NewWriter(destFile)

	written, e := io.Copy(writer, reader)
	// 一定要刷新缓存到磁盘里面
	writer.Flush()
	if e != nil {
		log.Fatalln(e)
	}
	log.Fatalln(written)
}

// 打开文件
func openFile() {
	// 默认是以只读的方式打开文件
	file, e := os.Open("test.txt")
	if e != nil {
		log.Fatalf("%v", e)
	}
	log.Println(file)
	defer file.Close()
	// rwx = 4 + 2 + 1 = 7
	// 用户 分组 其他
	f, e := os.OpenFile("test_1.txt", os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("%v\n", f)

}

// 检查文件是否存在
func isExistFile() {
	info, e := os.Stat("test_2.txt")
	if e != nil {
		if os.IsNotExist(e) {
			log.Fatalf("没有找到文件=%v", e)
		}
	}
	log.Println(info)
}

// 创建新文件
func createFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("创建文件失败")
	}
	defer file.Close()
	fmt.Println("创建完成")
	log.Println(file)
}

// 写入文件
func writeFile() {
	file, e := os.OpenFile("test_read.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		log.Fatal("%v", e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "hello world\t\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

// 文件必须存在，追加字符文件末尾
func writeAppendFile() {
	file, e := os.OpenFile("test_read.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatal("%v", e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "system go \t\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

// 一次性读取到内存中，适合小文件读取
func readFile() {
	file, e := ioutil.ReadFile("test_read.txt")
	if e != nil {
		log.Fatalf("%v", e)
	}
	log.Printf("%v", string(file))
}

//获取文件信息
func stat() {
	info, e := os.Stat("test.txt")
	if e != nil {
		log.Fatalf("获取文件信息失败 %v", e)
	}
	log.Printf("name = %v,size = %v,isDir = %v,mode=%v,modTime=%v,sys=%T,sys=%+v", info.Name(), info.Size(),
		info.IsDir(), info.Mode(), info.ModTime(), info.Sys(), info.Sys())
}

// 文件重名
func rename() {
	e := os.Rename("test.txt", "test.bak.txt")
	if e != nil {
		log.Fatalf("文件重名失败 %v", e)
	}
}

// 移动文件
func move() {
	e := os.Rename("test.bak.txt", "/Users/lgfcxx/Documents/test.txt")
	if e != nil {
		log.Fatalf("文件重名失败 %v", e)
	}
}

// 删除文件
func removeFile() {
	e := os.Remove("test.txt")
	if e != nil {
		log.Fatalf("删除文件失败 %v", e)
	}
}

// 截取文件
func truncateFile() {
	// 裁剪一个文件到100个字节。
	// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过100个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的100个字节的文件。
	// 传入0则会清空文件。
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Println("截断失败")
	}
}

// 读取文件
func ReadFile() {
	getgid := os.Getgid()
	fmt.Println(getgid)
	dir, _ := os.Getwd()
	fmt.Printf("%v", dir)
	// 获取当前目录
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	var path string = pwd + "/src/go_dev/io/file/test.log"
	file, _ := os.Open(path)
	defer file.Close()
	fmt.Printf("%T\n", file)
	// 创建一个带缓冲的Reader ,默认缓冲是4096
	reader := bufio.NewReader(file)
	for {
		// 通过空格读取，需要注意文件内容最后必须多一个空格，不然最后一个字符不能读取
		str, err := reader.ReadString(' ')
		// str, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(str))
	}
}
