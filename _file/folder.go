package _file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

type Folder struct {
}

func NewFolder() *Folder {
	return &Folder{}
}

//获取当前文件夹中的所有文件
func (*Folder) ReadFilesToStr(dir string) (fileList []string, err error) {
	if dir == "" {
		return nil, nil
	}
	fileList = make([]string, 0, 30)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name()[:1] != "." {
			fileList = append(fileList, dir+info.Name())
		}
		return nil
	})
	return
}

//获取当前文件夹中的所有文件
func (*Folder) ReadFiles(dir string) (fileList []os.FileInfo, err error) {
	if dir == "" {
		return nil, nil
	}
	fileList = make([]os.FileInfo, 0, 30)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name()[:1] != "." {
			fileList = append(fileList, info)
		}
		return nil
	})
	return
}

//读取目录文件并按文件名中的数字排序
func (*Folder) ReadFilesByOrderNum(dir string) (fileList []string, err error) {
	fileList = make([]string, 0, 10)
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	sort.Slice(list, func(i, j int) bool {
		s, _ := strconv.Atoi(regexp.MustCompile(`[0-9]+`).FindString(list[i].Name()))
		e, _ := strconv.Atoi(regexp.MustCompile(`[0-9]+`).FindString(list[j].Name()))
		return s < e
	})
	for _, v := range list {
		if !v.IsDir() {
			fileList = append(fileList, dir+"/"+v.Name())
		}
	}
	return
}

//读取目录中的目录
func (*Folder) ReadDirs(dir string) ([]string, error) {
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fileList := make([]string, 0, 10)
	for _, v := range list {
		if v.IsDir() {
			fileList = append(fileList, dir+"/"+v.Name())
		}
	}
	return fileList, err
}

func (*Folder) IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}
	if nil != err {
		return false
	}
	return fio.IsDir()
}
