package gin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/pkg/apierrors"
)

func errorHandler(c *gin.Context) {
	c.Next()

	lastGinError := c.Errors.Last()
	if lastGinError == nil {
		return
	}

	apierr := toApiError(lastGinError.Err)

	log.Println(apierr)

	c.JSON(apierr.Status(), apierr)
}

func toApiError(err error) apierrors.ApiError {
	apierr, ok := err.(apierrors.ApiError)
	if !ok {
		return apierrors.NewInternalServerApiError("Internal Server Error", err)
	}

	return apierr
}
