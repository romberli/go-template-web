package resp

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/go-template-web/config"
	"github.com/romberli/go-template-web/pkg/message"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
	"github.com/spf13/viper"
)

const (
	responseCodeKey    = "code"
	responseNOKMessage = `{"code": %d, "message": "%s"}`
	// responseContentType is json so that http clients render the body as json automatically
	responseContentType = "application/json; charset=utf-8"
)

// ResponseNOK responses with given code and values,
// if code is between 400000 and 500000, it will log error and resp 500 to client
// otherwise, it will log info and resp 200 to client
func ResponseNOK(c *gin.Context, code int, values ...any) {
	err := message.NewMessage(code, values...)
	log.Errorf(constant.LogWithStackString, err)

	resp := NewErrResponse(code, err.Error())
	jsonBytes, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		log.Errorf("response: marshal error response failed. error:\n%+v", marshalErr)
		c.String(viper.GetInt(config.ServerRouterHTTPErrorCodeKey), marshalErr.Error())
		return
	}

	c.Data(viper.GetInt(config.ServerRouterHTTPErrorCodeKey), responseContentType, jsonBytes)
}

func ResponseOK(c *gin.Context, respMessage string, code int, values ...any) {
	msg := message.NewMessage(code, values...).Error()
	log.Info(msg)

	httpErrorCode := viper.GetInt(config.ServerRouterHTTPErrorCodeKey)
	if httpErrorCode != config.DefaultServerRouterHTTPErrorCode {
		_, t, _, err := jsonparser.Get([]byte(respMessage), responseCodeKey)
		if err != nil && err != jsonparser.KeyPathNotFoundError {
			errMessage := message.NewMessage(message.ErrGetResponseCode, errors.Trace(err)).Error()
			log.Errorf(constant.LogWithStackString, errMessage)
			c.Data(http.StatusOK, responseContentType, []byte(respMessage))
			return
		}
		if t == jsonparser.NotExist {
			// server router http error code is set to 200, we need to add extra code field to the response message
			// to tell the client if the request is successfully processed
			messageBytes, err := jsonparser.Set([]byte(respMessage), []byte(strconv.Itoa(constant.ZeroInt)), responseCodeKey)
			if err != nil {
				errMessage := message.NewMessage(message.ErrSetResponseCode, errors.Trace(err)).Error()
				log.Errorf(constant.LogWithStackString, errMessage)
				c.Data(http.StatusOK, responseContentType, []byte(respMessage))
			}

			respMessage = string(messageBytes)
		}
	}

	// http error code is set to 200, there is no need to add extra code field,
	// the client could examine the http code to evaluate if the request is successfully processed
	c.Data(http.StatusOK, responseContentType, []byte(respMessage))
}

func ResponseOKWithDebug(c *gin.Context, respMessage string, code int, values ...any) {
	msg := message.NewMessage(code, values...).Error()
	log.Debug(msg)

	c.Data(http.StatusOK, responseContentType, []byte(respMessage))
}
