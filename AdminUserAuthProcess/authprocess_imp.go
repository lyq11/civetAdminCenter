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
	"time"
)

// AuthProcessImp servant implementation
type AuthProcessImp struct {
}

func (imp *AuthProcessImp) CheckToken(tarsCtx context.Context, token string, result *bool) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *AuthProcessImp) RefreshToken(tarsCtx context.Context, oldtoken string, newtoken *string) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}
func getAdminRight(roleID int32) (res []string, err error) {
	var AdminBasicRights []AdminPermissionMan.Permission
	var AdminRightsArray []string
	var results int32
	_, err = userPermission.QueryRolehasPermission(-1, -1, roleID, &AdminBasicRights, &results)
	if err != nil {
		return AdminRightsArray, errors.New("error")
	}

	if results == 1 {
		fmt.Println(AdminBasicRights, roleID)
		for _, basicRight := range AdminBasicRights {
			AdminRightsArray = append(AdminRightsArray, basicRight.Tag)
		}
		return AdminRightsArray, nil
	} else {
		return AdminRightsArray, errors.New("error")
	}

}
func (imp *AuthProcessImp) Login(tarsCtx context.Context, info *civetAdminCenter.AdminAuthDataInfo, result *int32, token *string) (ret int32, err error) {
	var rpcObjs []AdminUserManagment.AdminBasicInfo
	var getAdminResult int32
	var checkPwdResult bool

	_, err = userMgmt.AdminQueryByKeyWordStringList("username", info.Username, &rpcObjs, &getAdminResult)
	fmt.Println(rpcObjs)
	fmt.Println(getAdminResult)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("no exist")
	}
	if getAdminResult == 1 && len(rpcObjs) == 1 {
		basicInfo := AdminUserManagment.AdminBasicInfo{
			Id:       rpcObjs[0].Id,
			Username: info.Username,
			Password: info.Password,
		}
		_, pwcErr := userMgmt.AdminCheckPassword(&basicInfo, &checkPwdResult)
		fmt.Println(basicInfo)
		fmt.Println(checkPwdResult)
		if checkPwdResult == true {
			JwtConfig := CivetJWT.JwtConf{Issuer: "SundayTek", JwtSecret: []byte("adminsdf")}
			userRight, err := getAdminRight(rpcObjs[0].Id)
			fmt.Println(userRight)
			if err != nil {
				return 0, err
			}
			generateToken, err := JwtConfig.GenerateAdminToken(info.Logintype, rpcObjs[0].Id, rpcObjs[0].Role, userRight, 3)
			if err != nil {
				*result = -1
				return 0, err
			}
			stringUserID := strconv.Itoa(int(rpcObjs[0].Id))
			redis.CreateStr(tarsCtx, "Token:"+stringUserID+":"+info.Logintype, generateToken, 24*time.Hour)
			*token = generateToken
			*result = 1
			return 0, nil
		} else {
			*result = -1
			return 0, pwcErr
		}
	} else {
		*result = -1
		return 0, nil
	}

}

func (imp *AuthProcessImp) Logout(tarsCtx context.Context, token string, result *int32) (ret int32, err error) {
	JwtConfig := CivetJWT.JwtConf{Issuer: "SundayTek", JwtSecret: []byte("adminsdf")}
	parseToken, err := JwtConfig.ParseToken(token)
	if err != nil {
		return 0, err
	}
	stringUserID := strconv.Itoa(int(parseToken.UserID))
	compareRst := redis.CompareEqual(tarsCtx, "Token:"+stringUserID+":"+parseToken.LoginType, token)
	if compareRst == true {
		redis.DelStrByKey(tarsCtx, "Token:"+stringUserID+":"+parseToken.LoginType)
		deleteRst := redis.ExistKey(tarsCtx, "Token:"+stringUserID+":"+parseToken.LoginType)
		if deleteRst == false {
			*result = 1
			return 0, nil
		} else {
			*result = -1
			return 0, errors.New("logout fail 2")
		}
	} else {
		*result = -2
		return 0, errors.New("logout fail 1")
	}
}

func (imp *AuthProcessImp) LoginByWechat(tarsCtx context.Context, info *civetAdminCenter.AdminAuthDataInfo, result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *AuthProcessImp) LoginByPhone(tarsCtx context.Context, info *civetAdminCenter.AdminAuthDataInfo, result *int32) (ret int32, err error) {
	//TODO implement me
	panic("implement me")
}

var userMgmt *AdminUserManagment.UserManagment
var userPermission *AdminPermissionMan.PermissionM
var redis *CivetRedis.Worker

// Init servant init
func (imp *AuthProcessImp) Init() error {
	//initialize servant here:
	comm := tars.NewCommunicator()
	userMgmt = new(AdminUserManagment.UserManagment)
	userPermission = new(AdminPermissionMan.PermissionM)
	obj := "civetAdminCenter.AdminUserManagment.UserManagmentObj"
	comm.StringToProxy(obj, userMgmt)
	obj2 := "civetAdminCenter.AdminPermissionMan.PermissionMObj"
	comm.StringToProxy(obj2, userPermission)
	redisconf := BaseInstance.RedisBaseConfig{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "",
		UserName: "",
		Db:       0,
		Size:     1,
	}
	redis = CivetRedis.CreateWorker((*instance.RedisBaseConfig)(&redisconf))
	return nil
}

// Destroy servant destory
func (imp *AuthProcessImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *AuthProcessImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
func (imp *AuthProcessImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	return 0, nil
}
