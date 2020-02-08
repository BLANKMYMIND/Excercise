package main

// import log "github.com/sirupsen/logrus" // logrus 全适配原始 log 的 api，可直接替换
import (
	"github.com/sirupsen/logrus"
	"os"
)

// 更日常的使用：单个应用记录日志到多处 new
var log = logrus.New()

func main() {
	// 绑定 log 输出到默认输出
	log.Out = os.Stdout

	// 设置 文本格式
	log.SetFormatter(&logrus.TextFormatter{})

	// 设置 日志触发位置记录 (func & file)
	log.SetReportCaller(true)

	// 设置 最低记录的日志等级 (高于等于该等级则记录)
	log.SetLevel(logrus.InfoLevel)

	// 取得 log 对象的 writer， 输出到任意位置
	log.Writer()

	// level 日志等级
	// 指示
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	// 警告
	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// 致命
	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// hook 钩子（触发器）官方用途：在指定等级日志时发送到日志到其他地方 实际上怎么写都行
	// 需要重写 hook 类型并创建，然后 AddHook
	// 【hook】
	//type Hook interface {
	//	Levels() []Level
	//	Fire(*Entry) error
	//}
	// 【hook 官方示例 syslog】
	// 结构，可自定义
	//type SyslogHook struct {
	//	Writer        *syslog.Writer
	//	SyslogNetwork string
	//	SyslogRaddr   string
	//}
	// 创建 hook (非必要接口)
	//func NewSyslogHook(network, raddr string, priority syslog.Priority, tag string) (*SyslogHook, error) {
	//	w, err := syslog.Dial(network, raddr, priority, tag)
	//	return &SyslogHook{w, network, raddr}, err
	//}
	//
	// 主要 hook 业务 (必要接口，返回值是 err)
	//func (hook *SyslogHook) Fire(entry *logrus.Entry) error {
	//	line, err := entry.String()
	//	if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
	//	return err
	//}
	//	switch entry.Level {
	//case logrus.PanicLevel:
	//	return hook.Writer.Crit(line)
	//case logrus.FatalLevel:
	//	return hook.Writer.Crit(line)
	//case logrus.ErrorLevel:
	//	return hook.Writer.Err(line)
	//case logrus.WarnLevel:
	//	return hook.Writer.Warning(line)
	//case logrus.InfoLevel:
	//	return hook.Writer.Info(line)
	//case logrus.DebugLevel, logrus.TraceLevel:
	//	return hook.Writer.Debug(line)
	//default:
	//	return nil
	//}
	//}
	//
	// 关联的日志等级 (一般不改变)
	//func (hook *SyslogHook) Levels() []logrus.Level {
	//	return logrus.AllLevels
	//}

}
