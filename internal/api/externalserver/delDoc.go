package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type delTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *delTransport) handler(ctx *gin.Context) {
	//delete by id document /documents/delete/:id
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request errors"})
	}



		idx, err := strconv.Atoi(id)
		if err != nil {
			t.log.Error().Err(err).Msg("bad index")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request errors"})
			return
		}

		err = t.svc.Delete(idx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, Response{Data: "entry deleted"})
		return
	}
}
