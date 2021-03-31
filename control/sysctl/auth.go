package sysctl

import (
	"blog/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// AuthGet doc
// @Tags sysauth
// @Summary 通过id获取单条认证信息
// @Param id path int true "pk id" default(1)
// @Router /sys/auth/get/{id} [get]
func AuthGet(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	}
	mod, has := model.AuthGet(id)
	if !has {
		return ctx.JSON(utils.ErrOpt("未查询到认证信息"))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// AuthAll doc
// @Tags sysauth
// @Summary 获取所有认证信息
// @Router /sys/auth/all [get]
func AuthAll(ctx echo.Context) error {
	mods, err := model.AuthAll()
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到认证信息", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

// AuthPage doc
// @Tags sysauth
// @Summary 获取认证分页信息
// @Param cid path int true "分类id" default(1)
// @Param pi query int true "分页数" default(1)
// @Param ps query int true "每页条数[5,20]" default(5)
// @Router /sys/auth/page/{cid} [get]
func AuthPage(ctx echo.Context) error {
	// cid, err := strconv.Atoi(ctx.Param("cid"))
	// if err != nil {
	//  return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	// }
	ipt := &model.Page{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	if ipt.Ps > 20 || ipt.Ps < 5 {
		return ctx.JSON(utils.ErrIpt("分页大小输入错误", ipt.Ps))
	}
	count := model.AuthCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", " count < 1"))
	}
	mods, err := model.AuthPage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("查询数据错误", err.Error()))
	}
	if len(mods) < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", "len(mods) < 1"))
	}
	return ctx.JSON(utils.Page("succ", mods, int(count)))
}

// AuthAdd doc
// @Tags sysauth
// @Summary 添加认证信息
// @Param token query string true "hmt" default(token)
// @Router /sys/auth/add [post]
func AuthAdd(ctx echo.Context) error {
	ipt := &model.Auth{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	// ipt.Utime = time.Now()
	err = model.AuthAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// AuthEdit doc
// @Tags sysauth
// @Summary 修改认证信息
// @Param token query string true "hmt" default(token)
// @Router /sys/auth/edit [post]
func AuthEdit(ctx echo.Context) error {
	ipt := &model.Auth{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	// ipt.Utime = time.Now()
	err = model.AuthEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// AuthDrop doc
// @Tags sysauth
// @Summary 通过id删除单条认证信息
// @Param id path int true "pk id" default(1)
// @Param token query string true "hmt" default(token)
// @Router /sys/auth/drop/{id} [get]
func AuthDrop(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	}
	err = model.AuthDrop(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}
