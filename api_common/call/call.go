package call

import (
	"log"
	"log/syslog"

	"github.com/sheikhrachel/future/api_common/papertrail"
)

type Call struct {
	pt *syslog.Writer
}

func New() Call {
	return Call{
		pt: papertrail.NewLogger(),
	}
}

type Service interface {
	InfoF(msg string)
	TraceF(msg string)
}

func (cc *Call) InfoF(msg string) {
	cc.pt.Info(msg)
	log.Println(msg)
}

func (cc *Call) TraceF(msg string) {
	log.Println(msg)
}
