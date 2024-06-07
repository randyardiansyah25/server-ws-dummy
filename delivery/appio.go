package delivery

import (
	"server-ws-dummy/entities"

	"github.com/kpango/glg"
)

func PrintoutObserver() {
	for po := range entities.PrintOutChan {
		if po.Type == entities.PRINTOUT_TYPE_ERR {
			_ = glg.Error(po.Message...)
		} else if po.Type == entities.PRINTOUT_TYPE_LOG {
			_ = glg.Log(po.Message...)
		}
	}
}
