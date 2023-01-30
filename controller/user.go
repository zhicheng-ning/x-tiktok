package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"x-tiktok/dao"
	"x-tiktok/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}


func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username,password)

	usi := service.UserServiceImpl{}
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}else {
		newUser := dao.UserBasicInfo{
			Name: username,
			Password: service.EnCoder(password),
		}
		if usi.InsertUser(&newUser) != true {
			fmt.Println("Insert Fail")
		}
		// 得到用户id
		user := usi.GetUserBasicInfoByName(username)
		token := service.GenerateToken(username)
		log.Println("注册返回的id: ", user.Id)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := service.EnCoder(password)
	log.Println("encoderPassword:", encoderPassword)
	// 登录逻辑：使用jwt，根据用户信息生成token
	usi := service.UserServiceImpl{}

	user := usi.GetUserBasicInfoByName(username)

	if encoderPassword == user.Password {
		token := service.GenerateToken(username)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: user.Id,
			Token: token,
		})
	}else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User or Password Error"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	// token做权限校验
	usi := service.UserServiceImpl{}
	id, _ := strconv.ParseInt(userId, 10, 64)
	if user, err := usi.GetUserLoginInfoById(id); err!= nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User: User(user),
		})
	}
}