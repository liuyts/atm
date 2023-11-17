package handler

import (
	"net/http"

	"ATM/internal/logic"
	"ATM/internal/svc"
	"ATM/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PutMoneyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PutMoneyRequest
		if err := httpx.Parse(r, &req); err != nil {
			Response(r, w, nil, err)
			return
		}

		l := logic.NewPutMoneyLogic(r.Context(), svcCtx)
		resp, err := l.PutMoney(&req)
		Response(r, w, resp, err)
	}
}
