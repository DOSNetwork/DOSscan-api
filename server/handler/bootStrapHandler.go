package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BootStrapHandler struct {
	BootStrapIPs []string
}

func NesBootStrapHandler(bootStrapIPs []string) *BootStrapHandler {
	return &BootStrapHandler{
		BootStrapIPs: bootStrapIPs,
	}
}

func (s *BootStrapHandler) BootStrap(c *gin.Context) {

	r := ""
	for _, ip := range s.BootStrapIPs {
		r = r + ip + ","
	}

	c.String(http.StatusOK, r)
}
