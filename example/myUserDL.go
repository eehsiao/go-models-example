// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"database/sql"
	"fmt"

	model "github.com/eehsiao/go-models"
	mysql "github.com/eehsiao/go-models-mysql"
)

const (
	userTable = "user"
)

// MyUserDao : extend from mysql.Dao
type MyUserDao struct {
	*mysql.Dao
}

// UserTb : sql table struct that to store into mysql
type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}

// GetFirstUser : this is a data logical function, you can write more logical in there
// sample data logical function to get the first user
func (m *MyUserDao) GetFirstUser() (user *User, err error) {

	m.Select("Host", "User", "Select_priv").From("user").Limit(1)
	fmt.Println("GetFirstUser", m.BuildSelectSQL().BuildedSQL())
	var (
		val interface{}
		row *sql.Row
	)

	if row, err = m.GetRow(); err == nil {
		if val, err = m.ScanRowType(row, (*UserTb)(nil)); err == nil {
			u, _ := val.(*UserTb)

			user = &User{
				Host:       model.Iif(u.Host.Valid, u.Host.String, "").(string),
				User:       model.Iif(u.User.Valid, u.User.String, "").(string),
				SelectPriv: model.Iif(u.SelectPriv.Valid, u.SelectPriv.String, "").(string),
			}
		}
	}
	row, val = nil, nil

	return
}

// GetUsers : this is a data logical function, you can write more logical in there
// sample data logical function to get the all users
func (m *MyUserDao) GetUsers() (users []*User, err error) {

	m.Select(model.Struce4QuerySlice(m.DaoStructType)...).From(m.GetTbName()).Limit(3)
	fmt.Println("GetUsers", m.BuildSelectSQL().BuildedSQL())
	var (
		vals []interface{}
		rows *sql.Rows
	)

	if rows, err = m.Get(); err == nil {
		if vals, err = m.Scan(rows); err == nil {
			for _, v := range vals {
				u, _ := v.(*UserTb)

				user := &User{
					Host:       model.Iif(u.Host.Valid, u.Host.String, "").(string),
					User:       model.Iif(u.User.Valid, u.User.String, "").(string),
					SelectPriv: model.Iif(u.SelectPriv.Valid, u.SelectPriv.String, "").(string),
				}
				users = append(users, user)
			}
		}
		rows.Close()
	}
	rows, vals = nil, nil

	return
}
