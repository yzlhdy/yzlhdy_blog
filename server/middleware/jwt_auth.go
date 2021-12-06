package middleware

import (
	"log"
	"net/http"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse(401, "Unauthorized", helper.EmptyObj{}, "token不存在")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse(401, "Token is not valid", helper.EmptyObj{}, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
