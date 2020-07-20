package request

import "gin-vue-admin/model"

type NielsenSoftwareSearch struct{
    model.NielsenSoftware
    PageInfo
}