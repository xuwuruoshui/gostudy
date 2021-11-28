package controller

import (
	"encoding/json"
	"gostudy/25mysql/04practice/logreg/entity"
	"gostudy/25mysql/04practice/logreg/service"
	"gostudy/25mysql/04practice/logreg/vo"
	"io/ioutil"
	"log"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-10 10:55:59
* @content: 控制器
 */

func Index(w http.ResponseWriter, r *http.Request) {

	// 1.从header中获取用户token
	token := r.Header.Get("token")
	// 2.存在cookie查询用户信息,不存在跳转登录
	var result []byte
	log.Printf("%v\n", token)
	log.Println(len(token))
	if token != "null" {
		var uservo vo.UserVo
		err := json.Unmarshal([]byte(token), &uservo)
		if err != nil {
			log.Panic("json转换异常", err)
		}

		// 3.查询id进行展示用户信息
		user := service.FetchOne(uservo.Id)

		result, _ = json.Marshal(vo.Result{Code: 200, Msg: "", Data: user})
	} else {
		result, _ = json.Marshal(vo.Result{Code: 401, Msg: "请先登录"})
	}
	w.Write(result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic("前端传入失败", err)
	}
	var u entity.User
	err = json.Unmarshal(content, &u)
	log.Println(u)
	if err != nil {
		log.Panic("json转换异常:", err)
	}
	var result []byte
	user, ok := service.FetchOneByUsernamePasswd(u)
	log.Println(user)
	if ok {
		userValue, _ := json.Marshal(user)
		log.Println(user)
		log.Printf("%T", string(userValue))
		result, _ = json.Marshal(vo.Result{Code: 200, Msg: "登录成功", Data: string(userValue)})
	} else {
		result, _ = json.Marshal(vo.Result{Code: 401, Msg: "用户名或密码错误", Data: ""})
	}
	w.Write(result)
}

func Reg(w http.ResponseWriter, r *http.Request) {
	// 1.存入cookie,存在跳转到首页
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic("前端传入失败", err)
	}
	log.Println(content)
	log.Println(string(content))

	var result []byte
	var u entity.User
	err = json.Unmarshal(content, &u)
	if err != nil {
		log.Panic("json转换异常:", err)
	}

	id, ok := service.AddUser(u)
	if ok {
		user := entity.User{Id: int(id)}
		value, _ := json.Marshal(user)
		result, _ = json.Marshal(vo.Result{Code: 200, Msg: "注册成功", Data: string(value)})
	} else {
		result, _ = json.Marshal(vo.Result{Code: 400, Msg: "注册失败,用户名冲突"})
	}
	w.Write(result)
}
