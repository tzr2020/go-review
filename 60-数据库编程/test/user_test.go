package test

import (
	"log"
	"testing"

	"github.com/tzr2020/go-review/60-数据库编程/model"
)

func TestUser(t *testing.T) {
	// t.Run("新增用户", testAddUser)
	// t.Run("获取用户，根据用户ID", testGetUserByID)
	t.Run("获取所有用户", testGetAllUser)
}

func testAddUser(t *testing.T) {
	u := &model.User{
		Username: "zhangsan",
		Password: "123456",
		Email:    "zhangsan@qq.com",
	}
	err := u.Add()
	if err != nil {
		log.Printf("新增用户失败: %v\n", err)
	}

	u2 := &model.User{
		Username: "lisi",
		Password: "123456",
		Email:    "lisi@qq.com",
	}
	err = u2.Add()
	if err != nil {
		log.Printf("新增用户失败: %v\n", err)
	}

	u3 := &model.User{
		Username: "wangwu",
		Password: "123456",
		Email:    "wangwu@qq.com",
	}
	err = u3.Add()
	if err != nil {
		log.Printf("新增用户失败: %v\n", err)
	}
}

func testGetUserByID(t *testing.T) {
	u, err := model.GetUserByID(1)
	if err != nil {
		log.Printf("获取用户失败: %v\n", err)
	}
	log.Printf("user, value: %#v\n", u)
}

func testGetAllUser(t *testing.T) {
	us, err := model.GetAllUser()
	if err != nil {
		log.Printf("获取所有用户失败: %v\n", err)
	}
	for _, v := range us {
		log.Printf("user, value: %#v\n", v)
	}
}
