package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ibyond/go-start/config"
	"github.com/ibyond/go-start/global"
	"github.com/ibyond/go-start/utils"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	logging "github.com/op/go-logging"
	"io"
	"os"
	"strings"
	"time"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "go-start"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

func init() {
	c := global.GstConfig.Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := logging.MustGetLogger(module)
	var backends []logging.Backend
	registerStdout(c, &backends)
	if fileWriter := registerFile(c, &backends); fileWriter != nil {
		gin.DefaultWriter = io.MultiWriter(fileWriter, os.Stdout)
	}
	logging.SetBackend(backends...)
	global.GstLog = logger
}

func registerStdout(c config.Log, backends *[]logging.Backend) {
	if c.Stdout != "" {
		level, err := logging.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		*backends = append(*backends, createBackend(os.Stdout, c, level))
	}
}

func registerFile(c config.Log, backends *[]logging.Backend) io.Writer {
	if c.File != "" {
		if ok, _ := utils.PathExists(logDir); !ok {
			// directory not exist
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		fileWriter, err := rotatelogs.New(
			logDir+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
			// generate soft link, point to latest log file
			rotatelogs.WithLinkName(logSoftLink),
			// maximum time to save log files
			rotatelogs.WithMaxAge(7*24*time.Hour),
			// time period of log file switching
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
		}
		level, err := logging.LogLevel(c.File)
		if err != nil {
			fmt.Println(err)
		}
		*backends = append(*backends, createBackend(fileWriter, c, level))

		return fileWriter
	}
	return nil
}

func createBackend(w io.Writer, c config.Log, level logging.Level) logging.Backend {
	backend := logging.NewLogBackend(w, c.Prefix, 0)
	stdoutWriter := false
	if w == os.Stdout {
		stdoutWriter = true
	}
	format := getLogFormatter(c, stdoutWriter)
	backendLeveled := logging.AddModuleLevel(logging.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled
}

func getLogFormatter(c config.Log, stdoutWriter bool) logging.Formatter {
	pattern := defaultFormatter
	if !stdoutWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		// Remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return logging.MustStringFormatter(pattern)
}
