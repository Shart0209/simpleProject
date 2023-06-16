package externalserver

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strconv"
)

type updTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *updTransport) Handler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request by ID {_} not found"})
	}

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		t.log.Error().Err(err).Msg("bad index")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: fmt.Sprintf("bad request by ID {%s} not integer", id)})
		return
	}

	var res map[string]interface{}

	body, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(body, &res); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	if err := t.svc.Update(res, idx); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
