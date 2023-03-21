package path

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Mkdir 遞迴創造多層目錄
func Mkdir(path string) error {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) { // directory not found error
		if mkErr := os.MkdirAll(path, os.ModePerm); mkErr != nil {
			return mkErr
		}
	} else {
		return err // another error
	}

	return nil
}

// FileExists 檢查檔案或目錄使否已經存在
func FileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// FilePutContent 寫檔案 append 開啟則
func FilePutContent(filename string, content string, append bool) error {

	var flag int = 0
	if append {
		flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	} else {
		flag = os.O_CREATE | os.O_WRONLY
	}

	dir := filepath.Dir(filename)
	if err := Mkdir(dir); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filename, flag, os.ModePerm|0755)
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(content)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func FileGetContent(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	return string(content), err
}
