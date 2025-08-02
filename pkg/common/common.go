package common

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSONStrict(c *gin.Context, obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return err
	}
	return nil
}
