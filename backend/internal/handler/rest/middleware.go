package rest

import (
	"errors"
	"github.com/foxleren/auth-service/backend/pkg/authService"
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	authCtxUserID       = "userID"
	authCtxUserRole     = "userRole"
	ctxId               = "id"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, authentication, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Got empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Got invalid auth header")
		return
	}

	userData, err := authService.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(authCtxUserID, userData.Id)
	c.Set(authCtxUserRole, userData.Role)
}

func getContextId(c *gin.Context) (int, error) {
	id, ok := c.GetQuery(ctxId)
	if !ok {
		logrus.Printf("Level: middleware; func getContextId(): id not found")
		return 0, errors.New("id not found" + id)
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Printf("Level: middleware; func getContextId(): id is of invalid type")
		return 0, errors.New("id is of invalid type")
	}

	logrus.Printf("Level: middleware; func getContextId(): id=%d", id)
	return idInt, nil
}
