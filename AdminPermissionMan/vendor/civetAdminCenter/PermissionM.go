// Package civetAdminCenter comment
// This file was generated by tars2go 1.1.4
// Generated from PermissionM.tars
package civetAdminCenter

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// AdminRole struct implement
type AdminRole struct {
	Id       int32  `json:"id"`
	RoleName string `json:"roleName"`
	Name     string `json:"name"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

func (st *AdminRole) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AdminRole) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.Id, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.RoleName, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Name, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Created, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Updated, 4, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AdminRole) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AdminRole, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AdminRole) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.Id, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.RoleName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Name, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Created, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Updated, 4)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AdminRole) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// AdminRoleHasPermission struct implement
type AdminRoleHasPermission struct {
	Id            int32  `json:"id"`
	Role_id       int32  `json:"role_id"`
	Permission_id int32  `json:"Permission_id"`
	Create_time   int32  `json:"create_time"`
	Update_time   int32  `json:"update_time"`
	Uni_id        string `json:"uni_id"`
}

func (st *AdminRoleHasPermission) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AdminRoleHasPermission) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.Id, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Role_id, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Permission_id, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Create_time, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Update_time, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Uni_id, 5, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *AdminRoleHasPermission) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AdminRoleHasPermission, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *AdminRoleHasPermission) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.Id, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Role_id, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Permission_id, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Create_time, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Update_time, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Uni_id, 5)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AdminRoleHasPermission) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// Permission struct implement
type Permission struct {
	Id        int32  `json:"id"`
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	Summarize string `json:"summarize"`
	Path      string `json:"path"`
}

func (st *Permission) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *Permission) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int32(&st.Id, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Tag, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Name, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Summarize, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Path, 4, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *Permission) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Permission, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *Permission) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.Id, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Tag, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Name, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Summarize, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Path, 4)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *Permission) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
