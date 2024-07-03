package middleware

import (
	"context"
	"net/http"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/server/response"
	"strings"
)

type CtxString string

func Auth(client *grpcclient.Client, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenBearer := r.Header.Get("Authorization")
		if tokenBearer == "" {
			response.Forbidden(w, "missing token")
			return
		}

		tokenFields := strings.Fields(tokenBearer)
		if len(tokenFields) != 2 || tokenFields[0] != "Bearer" {
			response.Forbidden(w, "unknown type of jwt token")
		}

		ctx := r.Context()
		resp, err := client.Auth().CheckAuth(ctx, &auth.CheckAuthRequest{TokenStr: tokenFields[1]})
		if err != nil {
			response.Unauthorized(w, err.Error())
			return
		}

		ctx = context.WithValue(ctx, CtxString("user_id"), resp.GetUserid())

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
