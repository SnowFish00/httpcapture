package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// findYAMLFiles 递归查找 root 目录下的所有 yml 和 yaml 文件
func FindYAMLFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断文件扩展名
		if !info.IsDir() && (filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// readFileContent 读取指定路径的文件内容
func ReadFileContent(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// getFirstLevelDir 获取文件所在的第一层子目录的名称
func GetFirstLevelDir(root, path string) string {
	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		log.Fatalf("无法获取相对路径: %v", err)
	}
	parts := strings.Split(relativePath, string(os.PathSeparator))
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}
