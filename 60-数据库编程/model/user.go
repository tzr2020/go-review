package model

import (
	"log"
	"test/60-数据库编程/utils"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// 新增用户
func (u *User) Add() error {
	// 准备SQL语句
	sql := `insert into users(
		username, password, email
		) values(?, ?, ?)`

	// 预处理SQL语句
	stmt, err := utils.DB.Prepare(sql)
	if err != nil {
		log.Println("SQL语句预处理失败")
		return err
	}

	// 执行SQL语句
	_, err = stmt.Exec(u.Username, u.Password, u.Email)
	if err != nil {
		log.Println("SQL语句执行失败")
		return err
	}

	return nil
}

// 获取用户，根据用户ID
func GetUserByID(id int) (*User, error) {
	// 准备SQL语句
	sql := `select id, username, password, email
		from users
		where id = ?`

	// 预处理SQL语句
	stmt, err := utils.DB.Prepare(sql)
	if err != nil {
		log.Println("SQL语句预处理失败")
		return nil, err
	}

	// 执行SQL语句，获取结果并扫描到结构体
	u := new(User)
	row := stmt.QueryRow(id)
	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		log.Println("扫描结果失败")
		return nil, err
	}

	return u, nil
}

// 获取所有用户
func GetAllUser() ([]*User, error) {
	sql := `select id, username, password, email
		from users`

	// 执行SQL语句，获取结果集
	rows, err := utils.DB.Query(sql)
	if err != nil {
		log.Println("SQL语句执行失败")
		return nil, err
	}

	// 遍历结果集，扫描到结构体
	var us []*User
	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email)
		if err != nil {
			log.Println("扫描结果失败")
			return nil, err
		}
		us = append(us, u)
	}

	return us, nil
}
