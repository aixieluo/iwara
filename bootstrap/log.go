package bootstrap

import (
	"github.com/kataras/iris/v12/middleware/accesslog"
	"os"
)

func makeAccessLog() * accesslog.AccessLog {
	// Initialize a new access log middleware.
	ac := accesslog.File("./log/access.log")
	// Remove this line to disable logging to console:
	ac.AddOutput(os.Stdout)

	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.SetFormatter(&accesslog.JSON{
		Indent:     "    ",
		HumanTime:  true,
	})

	return ac
}
