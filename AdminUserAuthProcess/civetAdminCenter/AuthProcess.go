// Package civetAdminCenter comment
// This file was generated by tars2go 1.1.4
// Generated from AuthProcess.tars
package civetAdminCenter

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// AdminAuthDataInfo struct implement
type AdminAuthDataInfo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Wechat     string `json:"wechat"`
	VerifyCode string `json:"verifyCode"`
	Logintype  string `json:"logintype"`
}

func (st *AdminAuthDataInfo) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *AdminAuthDataInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Username, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Password, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Phone, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Wechat, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.VerifyCode, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Logintype, 5, false)
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
func (st *AdminAuthDataInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require AdminAuthDataInfo, but not exist. tag %d", tag)
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
func (st *AdminAuthDataInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Username, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Password, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Phone, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Wechat, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.VerifyCode, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Logintype, 5)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *AdminAuthDataInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
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
