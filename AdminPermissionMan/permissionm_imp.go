package main

import (
	"Civets/CivetTarsDataBase"
	"civetAdminCenter"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"os"
	"strconv"
)

// PermissionMImp servant implementation
type PermissionMImp struct {
}

func (imp *PermissionMImp) QueryRole(tarsCtx context.Context, offset int32, limit int32, role *[]civetAdminCenter.AdminRole, count *int32, result *int32) (ret int32, err error) {

	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, role) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *PermissionMImp) QueryRolePermission(tarsCtx context.Context, offset int32, limit int32, binds *[]civetAdminCenter.AdminRoleHasPermission, count *int32, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, binds) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *PermissionMImp) QueryPermission(tarsCtx context.Context, offset int32, limit int32, Rights *[]civetAdminCenter.Permission, count *int32, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), count, Rights) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *PermissionMImp) CheckRoleHasPermissionByPath(tarsCtx context.Context, roleID int32, Path string, res *int32) (ret int32, err error) {
	var ressult []civetAdminCenter.Permission
	var result int32
	_, err = imp.QueryRolehasPermission(tarsCtx, -1, -1, roleID, &ressult, &result)
	if err != nil {
		return 0, err
	}
	if result == 1 {
		for _, right := range ressult {
			if right.Path == Path {
				fmt.Println(right.Tag)
				fmt.Println(Path)
				*res = 1
				return 0, nil
			}
		}
		*res = -1
		return 0, nil
	}
	return 0, nil
}

func (imp *PermissionMImp) UpdateRole(tarsCtx context.Context, newRole *civetAdminCenter.AdminRole, keys string, value string, result *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(newRole, keys, value) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

//func (imp *PermissionMImp) QueryRole(tarsCtx context.Context, offset int32, limit int32, role *[]civetAdminCenter.AdminRole, result *int32) (ret int32, err error) {
//	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), role) {
//		*result = 1
//		return 0, nil
//	} else {
//		*result = -1
//		return 0, errors.New("queryFail")
//	}
//}

func (imp *PermissionMImp) QueryRoleByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, role *[]civetAdminCenter.AdminRole, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), role, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

//func (imp *PermissionMImp) QueryRolePermission(tarsCtx context.Context, offset int32, limit int32, binds *[]civetAdminCenter.AdminRoleHasPermission, result *int32) (ret int32, err error) {
//	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), binds) {
//		*result = 1
//		return 0, nil
//	} else {
//		*result = -1
//		return 0, errors.New("queryFail")
//	}
//}

func (imp *PermissionMImp) QueryRolePermissionByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, binds *[]civetAdminCenter.AdminRoleHasPermission, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), binds, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *PermissionMImp) UpdatePermission(tarsCtx context.Context, newCRight *civetAdminCenter.Permission, keys string, value string, result *int32) (ret int32, err error) {
	if TarDB.EditRowByConditionbyModel(newCRight, keys, value) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("edit Fail")
	}
}

//func (imp *PermissionMImp) QueryPermission(tarsCtx context.Context, offset int32, limit int32, Rights *[]civetAdminCenter.Permission, result *int32) (ret int32, err error) {
//	if TarDB.QueryRowsAllWithOSLM(int(offset), int(limit), Rights) {
//		*result = 1
//		return 0, nil
//	} else {
//		*result = -1
//		return 0, errors.New("queryFail")
//	}
//}

func (imp *PermissionMImp) QueryPermissionByCondition(tarsCtx context.Context, offset int32, limit int32, keys string, value string, Rights *[]civetAdminCenter.Permission, result *int32) (ret int32, err error) {
	if TarDB.QueryRowsAllWithConditionOSLM(int(offset), int(limit), Rights, keys, value) {
		*result = 1
		return 0, nil
	} else {
		*result = -1
		return 0, errors.New("queryFail")
	}
}

func (imp *PermissionMImp) CreateRole(tarsCtx context.Context, newRole *civetAdminCenter.AdminRole, c *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("admin_roles") {
		if TarDB.CreateRow(newRole) {
			*c = 1
			return 0, nil
		} else {
			*c = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(civetAdminCenter.AdminRole{}) {
			if TarDB.CreateRow(newRole) {
				*c = 1
				return 0, nil
			} else {
				*c = -1
				return 0, err
			}
		} else {
			return 0, err
		}
	}
}

func (imp *PermissionMImp) DeleteRole(tarsCtx context.Context, roleID int32, result *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&civetAdminCenter.AdminRole{}, "id", roleID) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *PermissionMImp) BindRoleAndRights(tarsCtx context.Context, newRoleHasRight *civetAdminCenter.AdminRoleHasPermission, result *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("admin_role_has_permissions") {
		var roles []civetAdminCenter.AdminRole
		var permissions []civetAdminCenter.Permission
		var checkroles int32
		var checkpermission int32
		_, err := imp.QueryRoleByCondition(tarsCtx, -1, -1, "id", strconv.Itoa(int(newRoleHasRight.Role_id)), &roles, &checkroles)
		if err != nil {
			return 0, err
		}
		if checkroles != 1 || len(roles) == 0 {
			*result = -1
			return 0, errors.New("can not find role")
		}
		_, err = imp.QueryPermissionByCondition(tarsCtx, -1, -1, "id", strconv.Itoa(int(newRoleHasRight.Permission_id)), &permissions, &checkpermission)
		if err != nil {
			return 0, err
		}
		if checkpermission != 1 || len(permissions) == 0 {
			*result = -1
			return 0, errors.New("can not find permission")
		}
		if TarDB.CreateRow(newRoleHasRight) {
			*result = 1
			return 0, nil
		} else {
			*result = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(civetAdminCenter.AdminRoleHasPermission{}) {
			if TarDB.CreateRow(newRoleHasRight) {
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

func (imp *PermissionMImp) UnBindRoleAndRights(tarsCtx context.Context, RoleID int32, result *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&civetAdminCenter.AdminRoleHasPermission{}, "id", RoleID) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *PermissionMImp) CreatePermission(tarsCtx context.Context, newCRight *civetAdminCenter.Permission, c *int32) (ret int32, err error) {
	if TarDB.CheckTableExist("permissions") {
		if TarDB.CreateRow(newCRight) {
			*c = 1
			return 0, nil
		} else {
			*c = -1
			return 0, err
		}
	} else {
		if TarDB.CreateTable(civetAdminCenter.Permission{}) {
			if TarDB.CreateRow(newCRight) {
				*c = 1
				return 0, nil
			} else {
				*c = -1
				return 0, err
			}
		} else {
			return 0, err
		}
	}
}

func (imp *PermissionMImp) DeletePermission(tarsCtx context.Context, RightsID int32, result *int32) (ret int32, err error) {
	if TarDB.DelRowByCondition(&civetAdminCenter.Permission{}, "id", RightsID) {
		*result = 1
		return 0, nil
	} else {
		return 0, errors.New("del Fail")
	}
}

func (imp *PermissionMImp) CheckRoleHasPermissionByID(tarsCtx context.Context, roleID int32, rightID int32, res *int32) (ret int32, err error) {
	var ressult []civetAdminCenter.Permission
	var result int32
	_, err = imp.QueryRolehasPermission(tarsCtx, -1, -1, roleID, &ressult, &result)
	if err != nil {
		return 0, err
	}
	if result == 1 {
		for _, right := range ressult {
			if right.Id == rightID {
				fmt.Println(right.Tag)
				fmt.Println(rightID)
				*res = 1
				return 0, nil
			}
		}
		*res = -1
		return 0, nil
	}
	return 0, nil
}

func (imp *PermissionMImp) CheckRoleHasPermissionByTag(tarsCtx context.Context, roleID int32, rightsTag string, res *int32) (ret int32, err error) {
	var ressult []civetAdminCenter.Permission
	var result int32
	_, err = imp.QueryRolehasPermission(tarsCtx, -1, -1, roleID, &ressult, &result)
	if err != nil {
		return 0, err
	}
	if result == 1 {
		for _, right := range ressult {
			if right.Tag == rightsTag {
				fmt.Println(right.Tag)
				fmt.Println(rightsTag)
				*res = 1
				return 0, nil
			}
		}
		*res = -1
		return 0, nil
	}
	return 0, nil
}
func (imp *PermissionMImp) QueryRolehasPermission(tarsCtx context.Context, offset int32, limit int32, roleID int32, Rights *[]civetAdminCenter.Permission, result *int32) (ret int32, err error) {
	var ressult []civetAdminCenter.Permission
	var rows *sql.Rows
	if limit == -1 {
		rows, err = TarDB.Db.Raw("SELECT permissions.id,permissions.name,permissions.summarize,permissions.path,permissions.tag FROM admin_roles INNER JOIN admin_role_has_permissions on admin_role_has_permissions.role_id = admin_roles.id INNER JOIN permissions on permissions.id = admin_role_has_permissions.permission_id where admin_roles.id = ?", roleID).Rows()
	} else {
		rows, err = TarDB.Db.Raw("SELECT permissions.id,permissions.name,permissions.summarize,permissions.path,permissions.tag FROM admin_roles INNER JOIN admin_role_has_permissions on admin_role_has_permissions.role_id = admin_roles.id INNER JOIN permissions on permissions.id = admin_role_has_permissions.permission_id where admin_roles.id = ?  LIMIT ? OFFSET ?", roleID, limit, offset).Rows()
	}

	defer rows.Close()

	fmt.Println(err)
	for rows.Next() {
		// ScanRows 将一行扫描至 user
		tmp := civetAdminCenter.Permission{}
		TarDB.Db.ScanRows(rows, &tmp)
		ressult = append(ressult, tmp)
		// 业务逻辑...
	}
	fmt.Println(rows)
	*Rights = ressult
	*result = 1
	fmt.Println(ressult)
	return 0, nil
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
func (imp *PermissionMImp) Init() error {
	var TarsConfig = tars.NewRConf("civetAdminCenter", "AdminPermissionMan", tars.GetServerConfig().BasePath)
	config, _ := TarsConfig.GetConfig("AdminPermissionMan.conf")
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
func (imp *PermissionMImp) Destroy() {
	//destroy servant here:
	//...
}
