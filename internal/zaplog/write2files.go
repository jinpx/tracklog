package zaplog

import (
	"go.uber.org/zap"
)

// TLogger 日志信息
type TLogger struct {
	initStart    bool
	newFiles     bool
	szData       string
	szFile       string
	hashDirPath  string
	hashFilePath string
	logger       *zap.Logger
	sugarLogger  *zap.SugaredLogger
}

// Printf 为了兼容
func (m *TLogger) Printf(template string, args ...interface{}) {
	m.sugarLogger.Debugf(template, args...)
}

// Print 为了兼容
func (m *TLogger) Print(args ...interface{}) {
	m.Println(args...)
}

// Println 为了兼容
func (m *TLogger) Println(args ...interface{}) {
	m.sugarLogger.Info(args...)
}

// Debug ()
func (m *TLogger) Debug(args ...interface{}) {
	m.sugarLogger.Debug(args...)
}

// Debugf ()
func (m *TLogger) Debugf(template string, args ...interface{}) {
	m.sugarLogger.Debugf("[+] "+template, args...)
}

// Info ()
func (m *TLogger) Info(args ...interface{}) {
	m.sugarLogger.Info(args...)
}

// Infof ()
func (m *TLogger) Infof(template string, args ...interface{}) {
	m.sugarLogger.Infof("[√] "+template, args...)
}

// Warn ()
func (m *TLogger) Warn(args ...interface{}) {
	m.sugarLogger.Warn(args...)
}

// Warnf () 警告
func (m *TLogger) Warnf(template string, args ...interface{}) {
	m.sugarLogger.Warnf("[!] "+template, args...)
}

// Error ()
func (m *TLogger) Error(args ...interface{}) {
	m.sugarLogger.Error(args...)
}

// Errorf ()
func (m *TLogger) Errorf(template string, args ...interface{}) {
	m.sugarLogger.Errorf("[x] "+template, args...)
}

// DPanic ()
func (m *TLogger) DPanic(args ...interface{}) {
	m.sugarLogger.DPanic(args...)
}

// DPanicf ()
func (m *TLogger) DPanicf(template string, args ...interface{}) {
	m.sugarLogger.DPanicf(`[D] `+template, args...)
}

// Panic ()
func (m *TLogger) Panic(args ...interface{}) {
	m.sugarLogger.Panic(args...)
}

// Panicf ()
func (m *TLogger) Panicf(template string, args ...interface{}) {
	m.sugarLogger.Panicf(`[P] `+template, args...)
}

// Fatal ()
func (m *TLogger) Fatal(args ...interface{}) {
	m.sugarLogger.Fatal(args...)
}

// Fatalf ()
func (m *TLogger) Fatalf(template string, args ...interface{}) {
	m.sugarLogger.Fatalf(`[F] `+template, args...)
}

// Flush ()
func (m *TLogger) Flush() {
	err := m.sugarLogger.Sync()
	if err != nil {
		return
	}
}

// Sync ()
func (m *TLogger) Sync() {
	err := m.sugarLogger.Sync()
	if err != nil {
		return
	}
}
