package health

import (
	"github.com/romberli/go-template-web/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initHealthDebugMessage()
	initHealthInfoMessage()
	initHealthErrorMessage()
}

const (
	// info
	InfoHealthPing = 200101
)

func initHealthDebugMessage() {

}

func initHealthInfoMessage() {
	message.Messages[InfoHealthPing] = config.NewErrMessage(message.DefaultMessageHeader, InfoHealthPing, "health: ping completed")
}

func initHealthErrorMessage() {

}
