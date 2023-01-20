package externalserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type docManagementTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *docManagementTransport) Addhandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, id)
	//isError := ctx.Query("is_error")
	//if isError != "" {
	//	t.log.Error().Msg("fake error")
	//
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
	//		Error: "fake error",
	//	})
	//
	//	return
	//}
	//
	//type Parameters struct {
	//	ID int `uri:"id"`
	//}
	//var param Parameters
	//
	//switch ctx.Request.RequestURI {
	//case "/documents/list":
	//	data, err := t.svc.GetAll()
	//	t.errorData(ctx, err)
	//	ctx.JSON(http.StatusOK, data)
	//case "/documents/id/":
	//
	//	if err := ctx.ShouldBindUri(&param); err != nil {
	//		t.errorData(ctx, err)
	//	}
	//	fmt.Println(param.ID)
	//	data, err := t.svc.GetID(param.ID)
	//	t.errorData(ctx, err)
	//
	//	ctx.JSON(http.StatusOK, data)
	//
	//case "/documents/add":
	//
	//case "/documents/delete":
	//
	//case "/documents/update":
	//
	//}

}

func (t *docManagementTransport) errorData(ctx *gin.Context, err error) {
	if err != nil {
		t.log.Error().Err(err).Msg("error response list doc")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Error: err.Error(),
		})

		return
	}
}
