package helper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetAuthorzied(ctx *gin.Context) (*string, error) {

	userData, _ := ctx.Get("claims")
	userBytes, _ := json.Marshal((userData))
	var usr *string
	err := json.Unmarshal(userBytes, &usr)

	return usr, err
}
