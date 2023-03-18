package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"simpleProject/pkg/model"
)

type addTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *addTransport) handler(ctx *gin.Context) {

	var bindForm model.BindForm

	//var res interface{}
	//b, _ := io.ReadAll(ctx.Request.Body)
	//_ = json.Unmarshal(b, &res)

	if err := ctx.ShouldBind(&bindForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	if err := t.svc.Create(&bindForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	// TODO не работает redirect - разобраться почему!
	//location := url.URL{Path: "/documents/"}
	//fmt.Println(location)
	//ctx.Redirect(http.StatusMovedPermanently, location.RequestURI())

	ctx.Status(http.StatusOK)
	return
}
