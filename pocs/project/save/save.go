package save

import (
	"demo/global"
	"demo/model"
	"demo/utils"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SavePoc(c *gin.Context) {
	// 获取绝对路径
	rootPath, err := filepath.Abs("pocs")
	if err != nil {
		log.Fatalf("无法获取绝对路径: %v", err)
	}

	// 找到所有 yml 和 yaml 文件
	files, err := utils.FindYAMLFiles(rootPath)
	if err != nil {
		log.Fatalf("无法找到 YAML 文件: %v", err)
	}

	// 处理每个文件
	for _, file := range files {
		// 读取文件内容
		content, err := utils.ReadFileContent(file)
		if err != nil {
			log.Printf("无法读取文件 %s: %v", file, err)
			continue
		}

		// 检查文件内容长度
		if len(content) > 65535 {
			log.Printf("文件 %s 的内容超出长度限制，跳过保存", file)
			continue
		}

		// 获取第一层子文件夹的名称
		fromDir := utils.GetFirstLevelDir(rootPath, file)
		if fromDir == "afrog" {
			fromDir = "https://github.com/zan8in/afrog-pocs/tree/main/CVE"
		}
		if fromDir == "fscan" {
			fromDir = "https://github.com/shadow1ng/fscan/tree/main/WebScan/pocs"
		}
		// 创建文件记录
		record := model.FileRecord{Name: filepath.Base(file), From: fromDir, Content: content}

		// 检查数据库中是否已有相同 Name 的记录
		var existingRecord model.FileRecord
		if err := global.GetGlobalDB().Where("name = ?", record.Name).First(&existingRecord).Error; err == nil {
			log.Printf("文件 %s 已存在，跳过保存", file)
			continue
		}

		// 将记录存入数据库
		if err := global.GetGlobalDB().Create(&record).Error; err != nil {
			log.Printf("无法保存文件 %s 的记录: %v", file, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "POC 获取加载完成"})
}
