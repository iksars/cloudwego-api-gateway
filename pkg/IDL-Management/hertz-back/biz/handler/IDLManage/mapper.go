package IDLManageService

import (
	"time"

	IDLManage "github.com/iksars/cloudwego-api-gateway/pkg/IDL-Management/hertz-back/biz/model/IDLManage"
)

func RecorderToEntity(r *IDLRecorder) *IDLManage.IDLEntity {
	return &IDLManage.IDLEntity{
		Date:        timeToString(r.CreatedAt),
		Name:        r.ServiceName,
		Description: r.Description,
	}
}

func timeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
