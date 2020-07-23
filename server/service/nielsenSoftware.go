package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateNielsenSoftware
// @description   create a NielsenSoftware
// @param     ns               model.NielsenSoftware
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateNielsenSoftware(ns model.NielsenSoftware) (err error) {
	err = global.GVA_DB.Create(&ns).Error
	return err
}

// @title    DeleteNielsenSoftware
// @description   delete a NielsenSoftware
// @auth                     （2020/04/05  20:22）
// @param     ns               model.NielsenSoftware
// @return                    error

func DeleteNielsenSoftware(ns model.NielsenSoftware) (err error) {
	err = global.GVA_DB.Delete(ns).Error
	return err
}

// @title    DeleteNielsenSoftwareByIds
// @description   delete NielsenSoftwares
// @auth                     （2020/04/05  20:22）
// @param     ns               model.NielsenSoftware
// @return                    error

func DeleteNielsenSoftwareByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.NielsenSoftware{}, "id in (?)", ids.Ids).Error
	return err
}

// @title    UpdateNielsenSoftware
// @description   update a NielsenSoftware
// @param     ns          *model.NielsenSoftware
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateNielsenSoftware(ns *model.NielsenSoftware) (err error) {
	err = global.GVA_DB.Save(ns).Error
	return err
}

// @title    GetNielsenSoftware
// @description   get the info of a NielsenSoftware
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    NielsenSoftware        NielsenSoftware

func GetNielsenSoftware(id uint) (err error, ns model.NielsenSoftware) {
	err = global.GVA_DB.Where("id = ?", id).First(&ns).Error
	return
}

// @title    GetNielsenSoftwareInfoList
// @description   get NielsenSoftware list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetNielsenSoftwareInfoList(nsw model.NielsenSoftware, info request.NielsenSoftwareSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.NielsenSoftware{})
	var nss []model.NielsenSoftware
	// 如果有条件搜索 下方会自动创建搜索语句

	if nsw.SoftName != "" {
		db = db.Where("soft_name LIKE ?", "%"+nsw.SoftName+"%")
	}

	// if api.Method != "" {
	// 	db = db.Where("method = ?", api.Method)
	// }

	// if api.ApiGroup != "" {
	// 	db = db.Where("api_group = ?", api.ApiGroup)
	// }

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&nss).Error
	return err, nss, total
}
