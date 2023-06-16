package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/afiifatuts/simple_bank/token"
	"github.com/gin-gonic/gin"
)

// extract authorization
const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// its not middleware but just only higher order function that will return authentication middleware function
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	//this anonymous function is in fact the authentication middleware function that we want to implement
	return func(ctx *gin.Context) {
		//extract outhorization header
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}
		//in case the authorization type is Barier token
		//then the access token should be the second element of the fields slice
		accessToken := fields[1]
		//parse and verify this access token to get the payload
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		//store the payload the context before passing it to the
		//next handler
		ctx.Set(authorizationPayloadKey, payload)
		//forward the request to the next handler
		ctx.Next()

	}
}
