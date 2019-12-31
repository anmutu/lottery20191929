package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"math/rand"
	"strings"
	"time"
)

var userList []string

func main() {
	app := NewApp()
	userList = make([]string, 0)
	app.Run(iris.Addr("8123"))
}

func NewApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

type lotteryController struct {
	Ctx iris.Context
}

//得到参数抽奖人数的信息
//curl http://localhost:8080/
func (*lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("参加抽奖人数为%d", count)
}

//导入参数抽奖的人
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",") //这里的是strings，有个s。
	originCount := len(userList)
	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	newCount := len(userList)
	return fmt.Sprintf("新加进入%d个用户，共有%d个用户参与抽奖", newCount-originCount, originCount)
}

func (c *lotteryController) GetLucky() string {
	count := len(userList)
	if count > 1 {
		//rand内部运算随机数
		seed := time.Now().UnixNano()
		//rand计算得到随机数
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		//抽取用户
		user := userList[index]
		//用切片，将抽取到的用户移除
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("恭喜%s中将，剩余抽奖用户%s人。", user, len(userList))
	} else if count == 1 {
		user := userList[0]
		userList = userList[0:0]
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
	} else {
		return fmt.Sprintf("无用户参与抽奖，请先导入用户 \n")
	}
}
