package router

import (
	"encoding/json"
	"server-ws-dummy/entities"
	"server-ws-dummy/entities/request"

	"github.com/kpango/glg"
	"github.com/randyardiansyah25/wsbase-handler"
)

func WSMessageListener(msg string) {
	wsmessage := wsbase.Message{}
	er := json.Unmarshal([]byte(msg), &wsmessage)
	if er != nil {
		_ = glg.Error("Error while unmarshal wsbase message :", er.Error())
		return
	}
	ch, ok := request.Pool[wsmessage.SenderId]
	if ok {
		respmsg, er := json.Marshal(wsmessage.Body)
		if er != nil {
			_ = glg.Error("Error while marshal ws message body for response : ", er.Error())
			responseObject := entities.BaseResponseString{
				BaseResponse: entities.BaseResponse{
					ResponseCode:    "1111",
					ResponseMessage: "Internal service error",
				},
				ResponseData: "",
			}
			responseError, _ := json.Marshal(responseObject)
			ch <- string(responseError)
		} else {
			ch <- string(respmsg)
		}
	} else {
		_ = glg.Infof("%s %s %s", "Request", wsmessage.SenderId, "has been deleted due time limit was exceeded...")
	}
}
