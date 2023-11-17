package handler

import (
	"net/http"

	"ATM/internal/logic"
	"ATM/internal/svc"
	"ATM/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TakeMoneyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TakeMoneyRequest
		if err := httpx.Parse(r, &req); err != nil {
			Response(r, w, nil, err)
			return
		}

		l := logic.NewTakeMoneyLogic(r.Context(), svcCtx)
		resp, err := l.TakeMoney(&req)
		Response(r, w, resp, err)
	}
}
