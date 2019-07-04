package apiv1

import (
	"fmt"

	"github.com/DOSNetwork/explorer-Api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetLogValidationResult(c *gin.Context) {
	tIdx := c.Params.ByName("traffic_id")
	fmt.Println("GetLogValidationResult traffic_id ", tIdx)
	validation := models.LogValidationResult{}
	if err := models.DB.Where("traffic_id = ?", tIdx).First(&validation).Error; gorm.IsRecordNotFoundError(err) {
		c.JSON(404, validation)
	} else {
		c.JSON(200, validation)
	}
}

///test http://localhost:8080/api/v1/validationResult/bb14823effa49c05f3b3f970aec6ffcab4da4cb1c6596044bfc6ba95b83a79b
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/validationResult/:traffic_id", GetLogValidationResult)
	}
}
