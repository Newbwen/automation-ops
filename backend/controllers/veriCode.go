package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateCaptcha(c *gin.Context) {
	// 生成验证码ID
	captchaID := captcha.New()

	// 生成验证码图片
	var content bytes.Buffer
	if err := captcha.WriteImage(&content, captchaID, 240, 80); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成验证码失败"})
		return
	}

	// 转换为base64
	base64Data := base64.StdEncoding.EncodeToString(content.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"captcha_id": captchaID,
		"image_data": "data:image/png;base64," + base64Data,
	})
}
