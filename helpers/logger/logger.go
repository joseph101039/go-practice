package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"rdm/google_organization/helpers/path"
)

// 記錄最近程序的 file path
var logFilePath string = ""

// SetFilePath 設置 log 檔案路徑, append 表示是否覆寫檔案內容
func SetFilePath(logPath string, append bool) {
	if logPath != "" {
		// Create the folder if not found
		dir := filepath.Dir(logPath)
		if err := path.Mkdir(dir); err != nil {
			panic(err)
		}

		//If the file doesn't exist, create it, or append to the file
		var flag int = 0
		if append {
			flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		} else {
			flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
		}
		f, err := os.OpenFile(logPath, flag, 0755)
		if err != nil {
			log.Fatal(err)
		}
		//log format
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)

		logFilePath = logPath
	}
}

//GetFilePath 為取得最近程序的 file path
func GetFilePath() string {
	return logFilePath
}
