package urlhdl

import (
	"errors"
	"espad/internal/core/ports"
	"espad/pkg/appErrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	urlService ports.UrlService
}

func NewHTTPHandler(urlService ports.UrlService) *HTTPHandler {
	return &HTTPHandler{
		urlService: urlService,
	}
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	c.BindJSON(&body)

	u, err := hdl.urlService.Create(body.OriginalUrl, body.Username)
	if err != nil {
		if errors.Is(err, appErrors.InvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
		}
		if errors.Is(err, appErrors.Internal) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, MakeResponseUrl(u))
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	path := c.Param("path")

	u, err := hdl.urlService.Get(path)
	if err != nil {
		if errors.Is(err, appErrors.NotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
		}
		if errors.Is(err, appErrors.Internal) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, MakeResponseUrl(u))
}

func (hdl *HTTPHandler) Redirect(c *gin.Context) {
	path := c.Param("path")

	u, err := hdl.urlService.Get(path)
	if err != nil {
		if errors.Is(err, appErrors.NotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
		}
		if errors.Is(err, appErrors.Internal) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
		}
		return
	}

	c.Redirect(http.StatusFound, u.Original)
}
