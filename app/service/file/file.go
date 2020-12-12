package file

import (
	"io/ioutil"
	"path/filepath"
)

type (
	File struct {
	}
	Info struct { // 文件信息
		Name string
		Size float64
	}
)

const ( // 文件大小单位
	_  = iota
	KB = 1 << (10 * iota)
	// MB
)

func New() *File {
	return &File{}
}

func (file *File) Walk(dirname string) ([]*Info, error) {
	op, err := filepath.Abs(dirname) // 获取目录的绝对路径
	if nil != err {
		return nil, err
	}
	files, err := ioutil.ReadDir(op) // 获取目录下所有文件的信息，包括文件和文件夹
	if nil != err {
		return nil, err
	}
	var fileInfos []*Info // 返回值，存储读取的文件信息
	for _, f := range files {
		if f.IsDir() { // 如果是目录，那么就递归调用
			fs, err := file.Walk(op + `/` + f.Name()) // 路径分隔符，linux 和 windows 不同
			if nil != err {
				return nil, err
			}
			fileInfos = append(fileInfos, fs...) // 将 slice 添加到 slice
		} else {
			fi := &Info{op + `/` + f.Name(), float64(f.Size()) / KB}
			fileInfos = append(fileInfos, fi) // slice 中添加成员
		}
	}
	return fileInfos, nil
}
