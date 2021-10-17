package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var mExtsImageFile map[string]string
var mExtsVideoFile map[string]string

func init() {
	mExtsImageFile = map[string]string{".png": "", ".PNG": "", ".jpg": "", ".JPG": "", ".jpeg": "", ".JPEG": "",
		".heic": "", ".HEIC": "", ".gif": "", ".GIF": "",
	}
	mExtsVideoFile = map[string]string{".MOV": "", ".mov": "", ".MP4": "", ".mp4": ""}
}

func IsMediaFile(filename string) bool {
	ext := path.Ext(filename)
	return IsImage(ext) || IsVideo(ext)
}

func IsImage(filename string) bool {
	ext := path.Ext(filename)
	_, exist := mExtsImageFile[ext]
	return exist
}

func IsVideo(filename string) bool {
	ext := path.Ext(filename)
	_, exist := mExtsVideoFile[ext]
	return exist
}

func IsVideoExt(ext string) bool {
	if strings.EqualFold(ext, ".mp4") || strings.EqualFold(ext, ".mov") || strings.EqualFold(ext, ".mpeg") ||
		strings.EqualFold(ext, ".mpg") || strings.EqualFold(ext, ".wmv") ||
		strings.EqualFold(ext, ".rm") || strings.EqualFold(ext, ".rmvb") ||
		strings.EqualFold(ext, ".swf") || strings.EqualFold(ext, ".flv") ||
		strings.EqualFold(ext, ".3GP") || strings.EqualFold(ext, ".mkv") ||
		strings.EqualFold(ext, ".m4v") || strings.EqualFold(ext, ".ogg") ||
		strings.EqualFold(ext, ".avi") || strings.EqualFold(ext, ".dat") ||
		strings.EqualFold(ext, ".vob") || strings.EqualFold(ext, ".mpe") ||
		strings.EqualFold(ext, ".asf") || strings.EqualFold(ext, ".asx") ||
		strings.EqualFold(ext, ".f4v") {
		return true
	}
	return false
}

//  /a/b/c.txt -> /a/b/
func GetFilePath(fullfilename string) string {
	dir, _ := filepath.Split(fullfilename)
	return dir
}

//  /a/b/c.txt -> /a/b/, c.txt
func GetFilePathAndName(fullfilename string) (string, string) {
	dir, file := filepath.Split(fullfilename)
	return dir, file
}

//  /a/b/c.txt -> c.txt
func GetFilename(fullfilename string) string {
	return path.Base(fullfilename)
}

//  /a/b/c.txt -> c
func GetFilenameOnly(fullfilename string) string {
	return strings.TrimSuffix(fullfilename, path.Ext(fullfilename))
}

//  /a/b/c.txt -> .txt
func GetFileSuffix(fullfilename string) string {
	return path.Ext(fullfilename)
}

func FileNameToJpgFileName(onlyFilename string) string {
	name := GetFilenameOnly(onlyFilename)
	return name + ".JPG"
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func IsFileExist(f string) bool {
	return !IsFileNotExist(f)
}
func IsFileNotExist(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}

func AddPathSepIfNeed(path string) (newPath string) {
	newPath = path
	if len(path) > 0 {
		if path[len(path)-1:] != "/" {
			newPath += "/"
		}
	} else {
		newPath += "/"
	}
	return
}

func CreateDirRecursive(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) { // not exist
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func WriteToFile(filename string, content []byte, truncateIfExist bool) error {
	flag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	if truncateIfExist {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}

	if IsFileNotExist(GetFilePath(filename)) {
		err := CreateDirRecursive(GetFilePath(filename))
		if err != nil {
			return fmt.Errorf("failed CreateDirRecursive, err:%v", err)
		}
	}

	fileObj, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()

	n, err := fileObj.Write(content)
	if err != nil {
		return err
	}
	if n != len(content) {
		return errors.New("written length error")
	}
	return nil
}

func WriteToFileWithFlag(filename string, content []byte, flag int) error {
	if IsFileNotExist(GetFilePath(filename)) {
		err := CreateDirRecursive(GetFilePath(filename))
		if err != nil {
			return fmt.Errorf("failed CreateDirRecursive, err:%v", err)
		}
	}

	fileObj, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()

	n, err := fileObj.Write(content)
	if err != nil {
		return err
	}
	if n != len(content) {
		return errors.New("written length error")
	}
	return nil
}

func ReadFromFile(filename string) ([]byte, error) {
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ReadFromFile, Open err=%v\n", err)
		return nil, err
	}
	defer fileObj.Close()

	content, err := ioutil.ReadAll(fileObj)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func MoveFile(file, targetDirWithoutTargetName string) error {
	targetDir := targetDirWithoutTargetName
	err := CreateDirRecursive(targetDir)
	if err != nil {
		//fmt.Printf("CreateDirRecursive {%v} failed :%v \n", targetDir, err)
		return err
	}

	base := path.Base(file)
	t := path.Join(targetDir, base)
	i := 1

	for {
		if IsFileNotExist(t) {
			break
		}
		fileSuffix := path.Ext(base)                         //获取文件后缀
		filenameOnly := strings.TrimSuffix(base, fileSuffix) //获取文件名
		t = path.Join(targetDir, filenameOnly+fmt.Sprintf("(%v)", i)+fileSuffix)
		i++
	}

	err = os.Rename(file, t)
	if err != nil {
		//fmt.Printf("Rename {%v to %v} failed :%v \n", file, t, err)
		return err
	}

	return nil
}

func ListDir(folder string) ([]string, error) {
	var ret []string
	if IsFile(folder) {
		ret = append(ret, folder)
	} else {

		files, err := ioutil.ReadDir(folder)
		if err != nil {
			return ret, err
		}
		for _, fi := range files {
			t := path.Join(folder, fi.Name())
			if fi.IsDir() {
				r, _ := ListDir(t)
				ret = append(ret, r...)
			} else {
				ret = append(ret, t)
			}
		}
	}
	return ret, nil
}

func SearchEmptyFoldersAndFiles(d string) ([]string, error) {
	var result []string
	files, err := ioutil.ReadDir(d)
	if err != nil {
		return result, err
	}
	for _, fi := range files {
		subd := path.Join(d, fi.Name())
		if fi.IsDir() {
			sz, err := DirSize(subd)
			if err != nil {
				return result, err
			}
			if sz < 1 {
				result = append(result, subd)
			} else {
				SearchEmptyFoldersAndFiles(subd)
			}

		} else {
			if fi.Size() < 1 {
				result = append(result, subd)
			}
		}
	}
	return result, nil
}

func WriteToFileAsJson(filename string, v interface{}, indent string, truncateIfExist bool) error {
	buf, err := json.MarshalIndent(v, "", indent)
	if err != nil {
		return err
	}
	err = WriteToFile(filename, buf, truncateIfExist)
	if err != nil {
		return err
	}
	return nil
}

func ReadFileJsonToObject(filename string, obj interface{}) error {
	buf, err := ReadFromFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return err
	}
	return nil
}

// IsFile Checks if Path is Valid File
func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
