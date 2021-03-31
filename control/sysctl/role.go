package sysctl

// import (
// 	"blog/model"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// 	"github.com/zxysilent/utils"
// )

// // RoleGet doc
// // @Tags sysrole
// // @Summary 通过id获取单条角色信息
// // @Param id path int true "pk id" default(1)
// // @Router /sys/role/get/{id} [get]
// func RoleGet(ctx echo.Context) error {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
// 	}
// 	mod, has := model.RoleGet(id)
// 	if !has {
// 		return ctx.JSON(utils.ErrOpt("未查询到角色信息"))
// 	}
// 	return ctx.JSON(utils.Succ("succ", mod))
// }

// // RoleAll doc
// // @Tags sysrole
// // @Summary 获取所有角色信息
// // @Router /sys/role/all [get]
// func RoleAll(ctx echo.Context) error {
// 	mods, err := model.RoleAll()
// 	if err != nil {
// 		return ctx.JSON(utils.ErrOpt("未查询到角色信息", err.Error()))
// 	}
// 	return ctx.JSON(utils.Succ("succ", mods))
// }

// // RolePage doc
// // @Tags sysrole
// // @Summary 获取角色分页信息
// // @Param cid path int true "分类id" default(1)
// // @Param pi query int true "分页数" default(1)
// // @Param ps query int true "每页条数[5,20]" default(5)
// // @Router /sys/role/page/{cid} [get]
// func RolePage(ctx echo.Context) error {
// 	// cid, err := strconv.Atoi(ctx.Param("cid"))
// 	// if err != nil {
// 	//  return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
// 	// }
// 	ipt := &model.Page{}
// 	err := ctx.Bind(ipt)
// 	if err != nil {
// 		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
// 	}
// 	if ipt.Ps > 20 || ipt.Ps < 5 {
// 		return ctx.JSON(utils.ErrIpt("分页大小输入错误", ipt.Ps))
// 	}
// 	count := model.RoleCount()
// 	if count < 1 {
// 		return ctx.JSON(utils.ErrOpt("未查询到数据", " count < 1"))
// 	}
// 	mods, err := model.RolePage(ipt.Pi, ipt.Ps)
// 	if err != nil {
// 		return ctx.JSON(utils.ErrOpt("查询数据错误", err.Error()))
// 	}
// 	if len(mods) < 1 {
// 		return ctx.JSON(utils.ErrOpt("未查询到数据", "len(mods) < 1"))
// 	}
// 	return ctx.JSON(utils.Page("succ", mods, int(count)))
// }

// // RoleAdd doc
// // @Tags sysrole
// // @Summary 添加角色信息
// // @Param token query string true "hmt" default(token)
// // @Router /sys/role/add [post]
// func RoleAdd(ctx echo.Context) error {
// 	ipt := &model.Role{}
// 	err := ctx.Bind(ipt)
// 	if err != nil {
// 		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
// 	}
// 	// ipt.Utime = time.Now()
// 	err = model.RoleAdd(ipt)
// 	if err != nil {
// 		return ctx.JSON(utils.Fail("添加失败", err.Error()))
// 	}
// 	return ctx.JSON(utils.Succ("succ"))
// }

// // RoleEdit doc
// // @Tags sysrole
// // @Summary 修改角色信息
// // @Param token query string true "hmt" default(token)
// // @Router /sys/role/edit [post]
// func RoleEdit(ctx echo.Context) error {
// 	ipt := &model.Role{}
// 	err := ctx.Bind(ipt)
// 	if err != nil {
// 		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
// 	}
// 	// ipt.Utime = time.Now()
// 	err = model.RoleEdit(ipt)
// 	if err != nil {
// 		return ctx.JSON(utils.Fail("修改失败", err.Error()))
// 	}
// 	return ctx.JSON(utils.Succ("succ"))
// }

// // RoleDrop doc
// // @Tags sysrole
// // @Summary 通过id删除单条角色信息
// // @Param id path int true "pk id" default(1)
// // @Param token query string true "hmt" default(token)
// // @Router /sys/role/drop/{id} [get]
// func RoleDrop(ctx echo.Context) error {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
// 	}
// 	err = model.RoleDrop(id)
// 	if err != nil {
// 		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
// 	}
// 	return ctx.JSON(utils.Succ("succ"))
// }
