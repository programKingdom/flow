package flow

import (
	"net/http"
	"strconv"

	"gitee.com/antlinker/flow/schema"
	"github.com/teambition/gear"
)

// API 提供API管理
type API struct {
	engine *Engine
}

// Init 初始化
func (a *API) Init(engine *Engine) *API {
	a.engine = engine
	return a
}

// 获取分页的页索引
func (a *API) pageIndex(ctx *gear.Context) uint {
	if v := ctx.Query("current"); v != "" {
		i, _ := strconv.Atoi(v)
		return uint(i)
	}
	return 1
}

// 获取分页的页大小
func (a *API) pageSize(ctx *gear.Context) uint {
	if v := ctx.Query("pageSize"); v != "" {
		i, _ := strconv.Atoi(v)
		if i > 40 {
			i = 40
		}
		return uint(i)
	}
	return 10
}

// QueryPage 查询分页数据
func (a *API) QueryPage(ctx *gear.Context) error {
	pageIndex, pageSize := a.pageIndex(ctx), a.pageSize(ctx)
	params := schema.FlowQueryParam{
		Code: ctx.Query("code"),
		Name: ctx.Query("name"),
	}

	total, items, err := a.engine.flowBll.QueryFlowPage(params, pageIndex, pageSize)
	if err != nil {
		return gear.ErrInternalServerError.From(err)
	}

	response := map[string]interface{}{
		"list": items,
		"pagination": map[string]interface{}{
			"total":    total,
			"current":  pageIndex,
			"pageSize": pageSize,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}
