package logger

import (
	"github.com/kpango/glg"
	"github.com/pluveto/site-deploy/pkg/setting"
)

func Setup() {
	infolog := glg.FileWriter(setting.AppSetting.LogPath+"/info.log", 0666)
	errlog := glg.FileWriter(setting.AppSetting.LogPath+"/error.log", 0666)
	defer infolog.Close()
	defer errlog.Close()
	glg.Get().
		SetMode(glg.BOTH).
		AddLevelWriter(glg.INFO, infolog). // add info log file destination
		AddLevelWriter(glg.ERR, errlog)    // add error log file destination

}
