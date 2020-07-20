package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitNielsenSoftwareRouter(Router *gin.RouterGroup) {
	NielsenSoftwareRouter := Router.Group("ns").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NielsenSoftwareRouter.POST("createNielsenSoftware", v1.CreateNielsenSoftware)   // 新建NielsenSoftware
		NielsenSoftwareRouter.DELETE("deleteNielsenSoftware", v1.DeleteNielsenSoftware) // 删除NielsenSoftware
		NielsenSoftwareRouter.DELETE("deleteNielsenSoftwareByIds", v1.DeleteNielsenSoftwareByIds) // 批量删除NielsenSoftware
		NielsenSoftwareRouter.PUT("updateNielsenSoftware", v1.UpdateNielsenSoftware)    // 更新NielsenSoftware
		NielsenSoftwareRouter.GET("findNielsenSoftware", v1.FindNielsenSoftware)        // 根据ID获取NielsenSoftware
		NielsenSoftwareRouter.GET("getNielsenSoftwareList", v1.GetNielsenSoftwareList)  // 获取NielsenSoftware列表
	}
}
