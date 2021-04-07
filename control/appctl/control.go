package appctl

import (
	"blog/model"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// Upload doc
// @Tags init
// @Summary 文件上传
// @Accept  mpfd
// @Param file formData file true "file"
// @Success 200 {object} model.Reply{data=string} "成功数据"
// @Router /api/upload/image [post]
func Upload(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`未发现文件,请重试`, err.Error()))
	}
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`文件打开失败,请重试`, err.Error()))
	}
	defer src.Close()
	basePath := "static/upload/" + time.Now().Format(utils.FmtyyyyMMdd) + "/"
	//确保文件夹存在
	os.MkdirAll(basePath, 0777)
	fileName := utils.RandStr(16) + filepath.Ext(file.Filename)
	filePathName := basePath + fileName
	dst, err := os.Create(filePathName)
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`目标文件创建失败,请重试`, err.Error()))
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return ctx.JSON(utils.ErrIpt(`文件写入失败,请重试`, err.Error()))
	}
	return ctx.JSON(utils.Succ(`文件上传成功`, "/"+filePathName))
}

// Sys 系统信息
func Sys(ctx echo.Context) error {
	state := struct {
		ARCH    string `json:"arch"`
		OS      string `json:"os"`
		Version string `json:"version"`
		NumCPU  int    `json:"num_cpu"`
	}{
		ARCH:    runtime.GOARCH,
		OS:      runtime.GOOS,
		Version: runtime.Version(),
		NumCPU:  runtime.NumCPU(),
	}
	return ctx.JSON(utils.Succ(`系统信息`, state))
}

// Collect 统计信息
func Collect(ctx echo.Context) error {
	if mod, has := model.Collect(); has {
		return ctx.JSON(utils.Succ(`统计信息`, mod))
	}
	return ctx.JSON(utils.Fail(`未查询到统计信息`))
}

// ExportMd
func ExportMd(ctx echo.Context) error {
	//TODO
	return nil
}
