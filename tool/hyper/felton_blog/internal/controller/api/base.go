package api

import (
	"github.com/kataras/iris/v12"

	"github.com/wuyoushe/hyper-go/tool/hyper/felton_blog/internal/model"
	"github.com/wuyoushe/hyper-go/tool/hyper/felton_blog/internal/service"
)

type CApi struct {
	Ctx   iris.Context
	Svc   service.Service
	Pager *model.Pager
}

// Hello godoc
// @Description Hello
// @Tags Hello
// @Accept  json
// @Produce  json
// @Param page query int true "页码"
// @Param pagesize query int true "页码尺寸"
// @Success 200 {object} model.User
// @Failure 500 {object} model.JSON
// @Router /hello [get]
func (c *CApi) GetHello() {
	data := model.GiNana{
		Hello: "Hello GiNana!",
	}
	c.Ctx.JSON(model.PlusJson(data, nil))
}
