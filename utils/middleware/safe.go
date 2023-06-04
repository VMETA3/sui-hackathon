package middleware

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/VMETA3/sui-hackathon/utils/response"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

// This part is used to verify the identity of users in the Metaverse
func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("VERIFY_START") == "true" {
			a := c.GetHeader("GameAccount")
			account, err := strconv.ParseInt(a, 10, 64)
			if err != nil {
				log.Debug("verifyToken account err", err)
				response.Response(c, &response.StandardResponse{
					Err: errors.New("account is invalid"),
				})
				c.Abort()
				return
			}
			token := c.GetHeader("Token")
			if !verifyToken(account, token, os.Getenv("VERIFY_KEY")) {
				response.Response(c, &response.StandardResponse{
					Err: errors.New("account is invalid"),
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func verifyToken(account int64, token string, key string) bool {
	strCount := bytes.Count([]byte(token), nil)
	if strCount < 60 {
		log.Debug("verifyToken token len < 60!")
		return false
	}
	accountStr := strconv.FormatInt(account, 10)
	data := []byte(token)
	randStr1 := string(data[0:20])
	timeStr1 := string(data[20:30])
	//
	timeUnixNow := time.Now().Unix()
	timeUnixToken, _ := strconv.Atoi(timeStr1)
	if (int)(timeUnixNow) > timeUnixToken+60 {
		log.Debug("verifyToken time error!")
		return false
	}
	//
	md5str := string(data[30:])
	//
	log.Debug("verifyToken randStr1 %s,timeStr1 %s,accountStr:%s,key:%s,md5Str:%s", randStr1, timeStr1, accountStr, key, md5str)
	allStr := randStr1 + timeStr1 + accountStr + key
	newData := []byte(allStr)
	newMd5CodeHex := md5.Sum(newData)
	newMd5Code := fmt.Sprintf("%x", newMd5CodeHex)
	log.Debug("verifyToken newMd5Code %s", newMd5Code)
	if newMd5Code == md5str {
		log.Debug("verifyToken token true")
		return true
	}
	log.Debug("verifyToken token false")
	return false
}
