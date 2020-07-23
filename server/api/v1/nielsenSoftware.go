package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

// @Tags NielsenSoftware
// @Summary 创建NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "创建NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ns/createNielsenSoftware [post]
func CreateNielsenSoftware(c *gin.Context) {
	var ns model.NielsenSoftware
	_ = c.ShouldBindJSON(&ns)
	err := service.CreateNielsenSoftware(ns)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags NielsenSoftware
// @Summary 删除NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "删除NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ns/deleteNielsenSoftware [delete]
func DeleteNielsenSoftware(c *gin.Context) {
	var ns model.NielsenSoftware
	_ = c.ShouldBindJSON(&ns)
	err := service.DeleteNielsenSoftware(ns)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NielsenSoftware
// @Summary 批量删除NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ns/deleteNielsenSoftwareByIds [delete]
func DeleteNielsenSoftwareByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteNielsenSoftwareByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags NielsenSoftware
// @Summary 更新NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "更新NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ns/updateNielsenSoftware [put]
func UpdateNielsenSoftware(c *gin.Context) {
	var ns model.NielsenSoftware
	_ = c.ShouldBindJSON(&ns)
	err := service.UpdateNielsenSoftware(&ns)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags NielsenSoftware
// @Summary 用id查询NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "用id查询NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ns/findNielsenSoftware [get]
func FindNielsenSoftware(c *gin.Context) {
	var ns model.NielsenSoftware
	_ = c.ShouldBindQuery(&ns)
	err, rens := service.GetNielsenSoftware(ns.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rens": rens}, c)
	}
}

// @Tags NielsenSoftware
// @Summary 分页获取NielsenSoftware列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NielsenSoftwareSearch true "分页获取NielsenSoftware列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ns/getNielsenSoftwareList [get]
func GetNielsenSoftwareList(c *gin.Context) {
	var pageInfo request.NielsenSoftwareSearch

	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetNielsenSoftwareInfoList(pageInfo.NielsenSoftware, pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
