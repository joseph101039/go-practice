package path

import (
	"os"
	"path/filepath"
)

// BasePath 提取專案根路徑, 傳入路徑, 類似 laravel helper base_path()
func BasePath(paths ...string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths = append([]string{dir}, paths...) // prepend the directory into path
	return filepath.Join(paths...)
}

// StoragePath  傳入路徑, 類似 laravel helper storage_path()
func StoragePath(paths ...string) string {
	paths = append([]string{"storage"}, paths...) // prepend the directory into path
	return BasePath(paths...)
}

// AppPath  傳入路徑, 類似 laravel helper app_path()
func AppPath(paths ...string) string {
	paths = append([]string{"app"}, paths...) // prepend the directory into path
	return BasePath(paths...)
}
