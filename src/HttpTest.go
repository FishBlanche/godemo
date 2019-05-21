package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

type Resp struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
}

type  Auth struct {
	Username string `json:"username"`
	Pwd      string   `json:"password"`
}
var db *sql.DB
func init() {
            db, _ = sql.Open("mysql", "admin:stariot@tcp(127.0.0.1:3306)/student?charset=utf8")
	        db.SetMaxOpenConns(2000)
            db.SetMaxIdleConns(1000)
            db.Ping()
}

func main () {
 // 设置路由
//	http.HandleFunc("/login1", login1)
	http.HandleFunc("/login2", login2)
 // 路由注册完，开始运行
 err := http.ListenAndServe(":6055", nil)
 if err != nil {
 log.Fatal(err)
	 }
}
func login2(writer http.ResponseWriter,  request *http.Request)  {
	request.ParseForm()
	username, uError :=  request.Form["username"]
	pwd, pError :=  request.Form["password"]
	var result  Resp
	if !uError || !pError {
		result.Code = "401"
		result.Msg = "失败"
	} else   {
		stmt, err := db.Prepare(`INSERT gouser (user_name,user_pass) values (?,?)`)
		checkErr(err)
		res, err := stmt.Exec(username[0],pwd[0])
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)
		result.Code = "200"
		result.Msg = "成功"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}
//接收x-www-form-urlencoded类型的post请求或者普通get请求
/*
func login2(writer http.ResponseWriter,  request *http.Request)  {
	request.ParseForm()
	username, uError :=  request.Form["username"]
	pwd, pError :=  request.Form["password"]

	var result  Resp
	if !uError || !pError {
		result.Code = "401"
		result.Msg = "登录失败"
	} else if username[0] == "admin" && pwd[0] == "123456" {
		stmt, err := db.Prepare(`INSERT gouser (user_name,user_pass) values (?,?)`)
		checkErr(err)
		res, err := stmt.Exec("tony", 20)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)
		result.Code = "200"
		result.Msg = "登录成功"
	} else {
		result.Code = "203"
		result.Msg = "账户名或密码错误"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}*/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func sayOne (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is version 1")
}

//post接口接收json数据
func login1(writer http.ResponseWriter,  request *http.Request)  {
	var auth Auth
	if err := json.NewDecoder(request.Body).Decode(&auth); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	var result  Resp
	if auth.Username == "admin" && auth.Pwd == "123456" {
		result.Code = "200"
		result.Msg = "登录成功"
	} else {
		result.Code = "401"
		result.Msg = "账户名或密码错误"
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
	}
}