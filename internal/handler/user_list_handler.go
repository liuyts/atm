package handler

import (
	"net/http"

	"ATM/internal/logic"
	"ATM/internal/svc"
	"ATM/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserListRequest
		if err := httpx.Parse(r, &req); err != nil {
			Response(r, w, nil, err)

			return
		}

		l := logic.NewUserListLogic(r.Context(), svcCtx)
		resp, err := l.UserList(&req)
		Response(r, w, resp, err)

	}
}
