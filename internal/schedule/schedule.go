package schedule

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	mp     = map[string]*TLogger{}
	mpLock sync.RWMutex
	logger *TLogger
)

// 为 Xor 初始
func init() {
	logger = initCurrentLogger()
	if logger == nil {
		panic(nil)
	}
}

func X() *TLogger {
	mapStr := currentLogger()

	mpLock.RLock()
	v, ok := mp[mapStr]
	mpLock.RUnlock()
	if ok {
		return v
	}

	logs := initCurrentLogger()
	if logs == nil {
		return logger
	}

	mpLock.RLock()
	mp = make(map[string]*TLogger, 0)
	mp[mapStr] = logs
	mpLock.RUnlock()

	// 更换文件后，清理
	runtime.GC()
	return logs
}

func currentLogger() string {
	szCurrent := time.Now()
	szData := szCurrent.Format("20060102")
	szFile := szCurrent.Format("2006010215")

	mapPath := fmt.Sprintf("./log/%s/%s.log", szData, szFile)
	szHash := md5.New()
	szHash.Write([]byte(mapPath))
	return hex.EncodeToString(szHash.Sum(nil))
}
