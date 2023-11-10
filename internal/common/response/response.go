package response

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Response(r *http.Request, w http.ResponseWriter, resp any, err error) {
	body := &Body{
		Code:    0,
		Message: "ok",
		Data:    resp,
	}
	if err != nil {
		body.Code = 1
		body.Message = err.Error()
		logx.WithContext(r.Context()).Errorf("error: %s", err.Error())
	}
	httpx.OkJsonCtx(r.Context(), w, body)
}
