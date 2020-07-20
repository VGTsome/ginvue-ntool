import service from '@/utils/request'

// @Tags NielsenSoftware
// @Summary 创建NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "创建NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ns/createNielsenSoftware [post]
export const createNielsenSoftware = (data) => {
     return service({
         url: "/ns/createNielsenSoftware",
         method: 'post',
         data
     })
 }


// @Tags NielsenSoftware
// @Summary 删除NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "删除NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ns/deleteNielsenSoftware [delete]
 export const deleteNielsenSoftware = (data) => {
     return service({
         url: "/ns/deleteNielsenSoftware",
         method: 'delete',
         data
     })
 }

// @Tags NielsenSoftware
// @Summary 删除NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ns/deleteNielsenSoftware [delete]
 export const deleteNielsenSoftwareByIds = (data) => {
     return service({
         url: "/ns/deleteNielsenSoftwareByIds",
         method: 'delete',
         data
     })
 }

// @Tags NielsenSoftware
// @Summary 更新NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "更新NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ns/updateNielsenSoftware [put]
 export const updateNielsenSoftware = (data) => {
     return service({
         url: "/ns/updateNielsenSoftware",
         method: 'put',
         data
     })
 }


// @Tags NielsenSoftware
// @Summary 用id查询NielsenSoftware
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NielsenSoftware true "用id查询NielsenSoftware"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ns/findNielsenSoftware [get]
 export const findNielsenSoftware = (params) => {
     return service({
         url: "/ns/findNielsenSoftware",
         method: 'get',
         params
     })
 }


// @Tags NielsenSoftware
// @Summary 分页获取NielsenSoftware列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取NielsenSoftware列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ns/getNielsenSoftwareList [get]
 export const getNielsenSoftwareList = (params) => {
     return service({
         url: "/ns/getNielsenSoftwareList",
         method: 'get',
         params
     })
 }