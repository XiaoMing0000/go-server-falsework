package utils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func GetUUID() string {
	var uuid [16]byte
	if _, err := io.ReadFull(rand.Reader, uuid[:]); err != nil {
		// 系统随机源故障，属于不可恢复的严重错误
		panic(err) // 或 return err
	}
	return hex.EncodeToString(uuid[:])
}
