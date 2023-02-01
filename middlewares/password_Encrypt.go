package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"qingxunyin/bytedance-tiktok/models"
	"unicode"
)

// EncryptPassword 将密码加密，需要传入密码返回的是加密后的密码
func EncryptPassword(password string) (string, error) {
	// 加密密码，使用 bcrypt 包当中的 GenerateFromPassword 方法，bcrypt.DefaultCost 代表使用默认加密成本
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果有错误则返回异常，加密后的空字符串返回为空字符串，因为加密失败
		return "", err
	} else {
		// 返回加密后的密码和空异常
		return string(encryptPassword), nil
	}
}

func EncryptMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Query("password")
		if password == "" {
			password = c.PostForm("password")
		}
		//对密码的规范进行检测，
		if len(password) > 20 {
			c.JSON(http.StatusOK, models.ResponseStatus{
				StatusCode: 1,
				StatusMsg:  "密码长度过长",
			})
			c.Abort()
			return
		}
		if !IsChinese(password) {
			c.JSON(http.StatusOK, models.ResponseStatus{
				StatusCode: 1,
				StatusMsg:  "密码包含中文",
			})
			c.Abort()
			return
		}
		password, err := EncryptPassword(password)
		if err != nil {
			c.JSON(http.StatusOK, models.ResponseStatus{
				StatusCode: 1,
				StatusMsg:  "密码加密错误",
			})
			c.Abort()
			return
		}

		c.Set("password", password)
		c.Next()
	}
}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
