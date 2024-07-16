package search

import (
	"demo/global"
	"demo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPocByID 根据 ID 返回对应的 FileRecord
func GetPocByID(c *gin.Context) {
	idParam := c.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID 无效"})
		return
	}

	var record model.FileRecord
	if err := global.GetGlobalDB().First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "记录未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"poc": record})
}
