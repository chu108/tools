package _file

import (
	"bufio"
	"bytes"
	"github.com/chu108/tools/_str"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}

//获取文件大小
func (*File) FileSize(file string) int64 {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

//判断所给路径文件/文件夹是否存在
func (*File) IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//如果文件夹不存在，则递归创建文件夹
func (obj *File) CreateFileByNot(filePath string) error {
	if !obj.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

//一次性加载文件，按行读取
func (*File) ReadAllOnLine(fielPath string, callBak func(row string) bool) {
	file, err := os.OpenFile(fielPath, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		b, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		rowStr := _str.NewBytes().ToStr(b)
		if rowStr != "" && !callBak(rowStr) {
			break
		}
	}
}

//逐行读取文件
func (*File) ReadOnLine(fielPath string, callBak func(row string) bool) {
	file, err := os.OpenFile(fielPath, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowStr := _str.NewBytes().ToStr(scanner.Bytes())
		if rowStr != "" && !callBak(rowStr) {
			break
		}
	}
}

//一次性读取文件所有内容
func (*File) ReadAll(fielPath string) ([]byte, error) {
	file, err := os.OpenFile(fielPath, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

//覆盖写入文件
func (*File) CreateFile(filePath string, body []byte) (writerLen int64, err error) {
	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	return io.Copy(file, bytes.NewReader(body))
}

//追加写入
func (*File) AppendWriteFile(filePath string, body []byte) (writerLen int, err error) {
	var file *os.File
	file, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(body)
}

// WriteFileSafer writes the data to a temp file and atomically move if everything else succeeds.
func (File) WriteFileSafer(writePath string, data []byte, perm os.FileMode) error {
	// credits: https://github.com/vitessio/vitess/blob/master/go/ioutil2/ioutil.go

	dir, name := filepath.Split(writePath)
	f, err := ioutil.TempFile(dir, name)
	if nil != err {
		return err
	}

	if _, err = f.Write(data); nil == err {
		err = f.Sync()
	}

	if closeErr := f.Close(); nil == err {
		err = closeErr
	}

	if permErr := os.Chmod(f.Name(), perm); nil == err {
		err = permErr
	}

	if nil == err {
		var renamed bool
		for i := 0; i < 3; i++ {
			err = os.Rename(f.Name(), writePath) // Windows 上重命名是非原子的
			if nil == err {
				renamed = true
				break
			}

			if errMsg := strings.ToLower(err.Error()); strings.Contains(errMsg, "access is denied") || strings.Contains(errMsg, "used by another process") { // 文件可能是被锁定
				time.Sleep(100 * time.Millisecond)
				continue
			}
			break
		}

		if !renamed {
			// 直接写入
			err = ioutil.WriteFile(writePath, data, perm)
		}
	}

	if nil != err {
		os.Remove(f.Name())
	}
	return err
}

// IsBinary determines whether the specified content is a binary file content.
func (*File) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}
