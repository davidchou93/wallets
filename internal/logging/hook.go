package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// WriterHook is a hook that writes log of specified LogLevels to specified Writer
type WriterHook struct {
	ErrorWriter io.Writer
	LogWriter   io.Writer
	LogLevels   []logrus.Level
}

// Fire will be called when some log function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	switch entry.Level {
	case logrus.PanicLevel:
		_, err = hook.ErrorWriter.Write([]byte(line))
		return err
	case logrus.FatalLevel:
		_, err = hook.ErrorWriter.Write([]byte(line))
		return err
	case logrus.ErrorLevel:
		_, err = hook.ErrorWriter.Write([]byte(line))
		return err
	case logrus.WarnLevel:
		_, err = hook.ErrorWriter.Write([]byte(line))
		return err
	case logrus.InfoLevel, logrus.DebugLevel, logrus.TraceLevel:
		_, err = hook.LogWriter.Write([]byte(line))
		return err
	}
	os.Stdout.Write([]byte(line))
	return nil
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// SetupHooks adds hooks to send logs to different destinations depending on level
// TODO: finish hooks
func SetupHooks(logger *logrus.Logger) (err error) {
	errFile, err := os.OpenFile("./error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Info("Failed to log to file")
		return
	}
	logFile, err := os.OpenFile("./info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Info("Failed to log to file")
		return
	}

	writerHook := new(WriterHook)
	writerHook.ErrorWriter = errFile
	writerHook.LogWriter = logFile
	logger.AddHook(writerHook)

	// // Add Slack hook
	// logger.SetFormatter(&logrus.JSONFormatter{})
	// cfg := lrhook.Config{
	// 	MinLevel: logrus.WarnLevel,
	// 	Message: chat.Message{
	// 		Channel: "#evolve",
	// 	},
	// }
	// slackHook := lrhook.New(cfg, "https://hooks.slack.com/services/TS7JY93KR/BS7VA7HPG/{accessKey}")
	// logger.AddHook(slackHook)

	return nil
}
