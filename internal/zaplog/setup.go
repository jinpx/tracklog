package zaplog

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type TCurrentLogger struct {
	FilePath string // 文件路径
	TimePath string // 时间路径
	LogName  string // 日志名称
	logFile  *os.File
	logger   *log.Logger
}

// NewCurrentLogger 构造函数，返回一个 *TCurrentLogger 实例
func NewCurrentLogger(filePath string, timePath string, logName string) (*TCurrentLogger, error) {
	// 构建日志文件路径
	logDirectory := filepath.Join(filePath, timePath)
	err := os.MkdirAll(logDirectory, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	// 构建完整的日志文件路径
	logFilePath := filepath.Join(logDirectory, fmt.Sprintf("%s.log", logName))

	// 打开或创建日志文件
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	// 创建 logger 实例
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// 返回日志实例
	return &TCurrentLogger{
		FilePath: filePath,
		TimePath: timePath,
		LogName:  logName,
		logFile:  logFile,
		logger:   logger,
	}, nil
}

// Close 关闭日志文件
func (logger *TCurrentLogger) Close() {
	if err := logger.logFile.Close(); err != nil {
		fmt.Printf("failed to close log file: %v\n", err)
	}
}

// 创建日志文件
func (m *TLogger) createFile(dirPath string) (*os.File, error) {
	filePath := fmt.Sprintf("%s/%s.log", dirPath, m.szFile)
	hashFilePath := m.generateHash(filePath)

	// 如果文件路径有变化，则重新创建文件
	if m.hashFilePath != hashFilePath {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
		if err != nil {
			return nil, fmt.Errorf("failed to create/open log file: %w", err)
		}
		m.hashFilePath = hashFilePath
		return file, nil
	}

	return nil, nil
}

// 生成路径的 md5 哈希值
func (m *TLogger) generateHash(path string) string {
	hash := md5.New()
	hash.Write([]byte(path))
	return hex.EncodeToString(hash.Sum(nil))
}

// 日志处理
func (m *TLogger) setupLogger() error {

	var (
		file *os.File
		err  error
	)

	szCurrent := time.Now()
	m.szData = szCurrent.Format("20060102")
	m.szFile = szCurrent.Format("2006010215")

	// 路径生成
	dirPath := fmt.Sprintf("./logs/%s", m.szData)
	szHash := md5.New()
	szHash.Write([]byte(dirPath))
	hashDirPath := hex.EncodeToString(szHash.Sum(nil))
	if m.hashDirPath != hashDirPath {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
		m.hashDirPath = hashDirPath
	}

	// 文件生成
	filePath := fmt.Sprintf("%s/%s.log", dirPath, m.szFile)
	szHash = md5.New()
	szHash.Write([]byte(filePath))
	hashFilePath := hex.EncodeToString(szHash.Sum(nil))
	if m.hashFilePath != hashFilePath {
		file, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
		if err != nil {
			return err
		}
		m.hashFilePath = hashFilePath
	}

	// 初始化
	if !m.initStart || m.newFiles {

		ws := io.MultiWriter(file, os.Stdout)
		writeSyncer := zapcore.AddSync(ws)

		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder := zapcore.NewConsoleEncoder(encoderConfig)

		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
		m.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

		m.sugarLogger = m.logger.Sugar()
	}

	return nil
}
