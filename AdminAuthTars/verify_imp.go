package main

import (
	instance "CivetRedis/CivetRedis/BaseInstance"
	"Civets/CivetJWT"
	"Civets/CivetRedis"
	"Civets/CivetRedis/BaseInstance"
	"civetAdminCenter"
	AdminPermissionMan "civetAdminCenter/AdminPermissionMan/civetAdminCenter"
	AdminUserManagment "civetAdminCenter/AdminUserManagment/civetAdminCenter"
	"context"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"strconv"
)

// VerifyImp servant implementation
type VerifyImp struct {
}

func getAdminRight(roleID int32) (res []string, paths []string, err error) {
	var AdminBasicRights []AdminPermissionMan.Permission
	var AdminPermissionPathArray []string
	var AdminRightsArray []string
	var results int32
	_, err = userPermission.QueryRolehasPermission(-1, -1, roleID, &AdminBasicRights, &results)
	if err != nil {
		return AdminRightsArray, AdminPermissionPathArray, errors.New("error")
	}

	if results == 1 {
		fmt.Println(AdminBasicRights, roleID)
		for _, basicRight := range AdminBasicRights {
			AdminRightsArray = append(AdminRightsArray, basicRight.Tag)
			AdminPermissionPathArray = append(AdminPermissionPathArray, basicRight.Path)
		}
		return AdminRightsArray, AdminPermissionPathArray, nil
	} else {
		return AdminRightsArray, AdminPermissionPathArray, errors.New("error")
	}

}
func (imp *VerifyImp) Verify(tarsCtx context.Context, req *civetAdminCenter.VeifyReq, rsp *civetAdminCenter.VeifyRsp) (ret int32, err error) {
	JwtConfig := CivetJWT.JwtConf{Issuer: "SundayTek", JwtSecret: []byte("adminsdf")}
	PathFromReq := req.VerifyHeaders["paths"]
	token, err := JwtConfig.ParseToken(req.Token)
	if err != nil {
		newRsp := civetAdminCenter.VeifyRsp{
			Ret:     -2,
			Uid:     "",
			Context: "",
		}
		*rsp = newRsp
		return 0, err
	}
	_, paths, err := getAdminRight(token.RoleID)
	if err != nil {
		return 0, err
	}
	fmt.Println("the path is ", PathFromReq, "the list is", paths)
	for _, path := range paths {
		if path == PathFromReq {
			newRsp := civetAdminCenter.VeifyRsp{
				Ret:     0,
				Uid:     strconv.Itoa(int(token.UserID)),
				Context: "success",
			}
			*rsp = newRsp
			return 0, nil
		}
	}
	newRsp := civetAdminCenter.VeifyRsp{
		Ret:     -1,
		Uid:     strconv.Itoa(int(token.UserID)),
		Context: "do not has permission",
	}
	*rsp = newRsp
	return 0, nil
}

var userMgmt *AdminUserManagment.UserManagment
var userPermission *AdminPermissionMan.PermissionM
var redis *CivetRedis.Worker

// Init servant init
func (imp *VerifyImp) Init() error {
	comm := tars.NewCommunicator()
	//comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 172.25.0.3 -t 60000 -p 17890")
	userMgmt = new(AdminUserManagment.UserManagment)
	userPermission = new(AdminPermissionMan.PermissionM)
	obj := "civetAdminCenter.AdminUserManagment.UserManagmentObj"
	comm.StringToProxy(obj, userMgmt)
	obj2 := "civetAdminCenter.AdminPermissionMan.PermissionMObj"
	comm.StringToProxy(obj2, userPermission)
	redisConf := BaseInstance.RedisBaseConfig{
		Host:     "172.25.0.1",
		Port:     "6379",
		Password: "",
		UserName: "",
		Db:       0,
		Size:     1,
	}
	redis = CivetRedis.CreateWorker((*instance.RedisBaseConfig)(&redisConf))
	return nil
}

// Destroy servant destory
func (imp *VerifyImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *VerifyImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *VerifyImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
