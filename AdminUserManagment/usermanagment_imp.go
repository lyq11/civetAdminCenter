package main

import (
	"Civets/CivetTarsDataBase"
	"civetAdminCenter"
	"civetAdminManagement/AdminDataProcessServer/civetAdminManagement"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"os"
	"strconv"
	"time"
)

// UserManagmentImp servant implementation
type UserManagmentImp struct {
}

//func (imp *UserManagmentImp) AdminQueryAll(tarsCtx context.Context, numbers int32, Pages int32, AdminMemberList *[]civetAdminCenter.AdminBasicInfo, result *int32) (ret int32, err error) {
//	if TarDB.QueryRowsAllWithOSLM(int(Pages), int(numbers), AdminMemberList) {
//		*result = 1
//		return 0, nil
//	} else {
//		*result = -1
//		return 0, errors.New("queryFail")
//	}
//}
func (imp *UserManagmentImp) AdminQueryAll(tarsCtx context.Context, numbers int32, Pages int32, AdminMemberList *[]civetAdminCenter.AdminBasicInfo, count *int32, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(Pages), int(numbers), count, AdminMemberList) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *UserManagmentImp) AdminAdd(tarsCtx context.Context, info *civetAdminCenter.AdminBasicInfo, result *int32) (ret int32, err error) {
	if CheckUsername(info.Username) != true {
		*result = -1
		return 0, errors.New("error 071")
	}
	if CheckPassword(info.Password) != true {
		*result = -2
		return 0, errors.New("error 072")
	}
	if CheckEmailFormat(info.Email) != true {
		*result = -3
		return 0, errors.New("error 073")
	}
	if CheckMobile(info.Number) != true {
		*result = -4
		return 0, errors.New("error 074")
	}
	info.Enable = "0"
	info.CreateTime = time.Now().Unix()
	info.UpdateTime = 0
	info.Password = HashAndSalt([]byte(info.Password))
	if TarDB.CheckTableExist("admin_basic_infos") {
		if TarDB.CreateRow(info) {
			*result = 1
			return 0, nil
		} else {
			*result = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(civetAdminCenter.AdminBasicInfo{}) {
			if TarDB.CreateRow(info) {
				*result = 1
				return 0, nil
			} else {
				*result = -1
				return 0, err
			}
		} else {
			return 0, err
		}
	}
}

func (imp *UserManagmentImp) AdminEditPassWord(tarsCtx context.Context, info *civetAdminCenter.AdminBasicInfo, result *int32) (ret int32, err error) {
	info.Password = HashAndSalt([]byte(info.Password))
	if TarDB.EditRowByConditionbyModel(info, "id", strconv.Itoa(int(info.Id))) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *UserManagmentImp) AdminDelete(tarsCtx context.Context, id int32, result *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&civetAdminCenter.AdminBasicInfo{}, "id", id) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *UserManagmentImp) AdminCheckPassword(tarsCtx context.Context, info *civetAdminCenter.AdminBasicInfo, result *bool) (ret int32, err error) {
	if CheckUsername(info.Username) != true {
		*result = false
		return 0, errors.New("error 061")
	}
	if CheckPassword(info.Password) != true {
		*result = false
		return 0, errors.New("error 062")
	}
	var temp = civetAdminManagement.AdminBasicInfo{}
	res := TarDB.QueryRowWithID(&temp, info.Id)
	if temp.Password == "" && res == true {
		fmt.Println("error")
		*result = false
		return 0, errors.New("error 063")
	} else {
		if ValidatePasswords(temp.Password, []byte(info.Password)) == true {
			*result = true
			return 0, nil
		} else {
			*result = false
			return 0, errors.New("wrong password")
		}
	}
}

func (imp *UserManagmentImp) AdminUpdateByString(tarsCtx context.Context, info *civetAdminCenter.AdminBasicInfo, searchKey string, searchValue string, result *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(info, searchKey, searchValue) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

func (imp *UserManagmentImp) AdminQueryByKeyWordStringList(tarsCtx context.Context, keyword string, stringWord string, AdminMemberList *[]civetAdminCenter.AdminBasicInfo, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(-1, -1, AdminMemberList, keyword, stringWord) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

var TarDB *CivetTarsDataBase.CivetData

type Config struct {
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBHost     string `json:"DBHost"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
	RDUserName string `json:"RDUserName"`
	RDPassword string `json:"RDPassword"`
	RDHost     string `json:"RDHost"`
	RDPort     string `json:"RDPort"`
	RDSize     int    `json:"RDSize"`
}

// Init servant init
func (imp *UserManagmentImp) Init() error {
	var TarsConfig = tars.NewRConf("civetAdminCenter", "AdminUserManagment", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("AdminUserManagment.conf")
	var mc Config
	json.Unmarshal([]byte(config), &mc)
	if mc.DBName == "" || mc.DBHost == "" || mc.DBUserName == "" {
		fmt.Println("The Mysql Config error")
		os.Exit(1)
	} else {
		fmt.Println("the config is", config, mc.DBHost, mc.DBPassword, mc.DBUserName, mc.DBName)
	}
	TarDB = CivetTarsDataBase.CreateCivetData(mc.DBUserName, mc.DBPassword, mc.DBHost, mc.DBPort)
	TarDB.ConnectOrCreateDataBase(mc.DBName)
	return nil
}

// Destroy servant destory
func (imp *UserManagmentImp) Destroy() {
	//destroy servant here:
	//...
}
