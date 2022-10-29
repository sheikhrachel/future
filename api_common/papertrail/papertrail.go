package papertrail

import (
	"log"
	"log/syslog"

	"github.com/sheikhrachel/future/api_common/util/errutil"
)

var (
	network       = "udp"
	papertrailUrl = "logs4.papertrailapp.com:22249"
	tag           = "appointment-service"
)

func NewLogger() (pt *syslog.Writer) {
	pt, err := syslog.Dial(network, papertrailUrl, syslog.LOG_EMERG|syslog.LOG_KERN, tag)
	if errutil.HandleError(err) {
		log.Panic(err.Error())
	}
	return pt
}
