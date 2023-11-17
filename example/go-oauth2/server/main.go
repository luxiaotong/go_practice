package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	oe "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
)

var srv *server.Server

func main() {
	log.Println("test oauth server start")

	// 设置 client 信息
	client_store := store.NewClientStore()
	_ = client_store.Set("test", &models.Client{ID: "test", Secret: "dGVzdDEyMzQ1Njc4", Domain: "http://127.0.0.1:9094"})
	// 设置 manager, manager 参与校验 code/access token 请求
	manager := manage.NewDefaultManager()
	// 设置 token 配置信息
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	// manger 包含 client 信息
	manager.MapClientStorage(client_store)
	// server 也包含 manger, client 信息
	srv = server.NewServer(server.NewConfig(), manager)
	// 允许使用 get 方法请求授权
	srv.SetAllowGetAccessRequest(true)
	// authorization code 模式,  第一步获取code,然后再用code换取 access token, 而不是直接获取 access token
	srv.SetAllowedResponseType(oauth2.Code)
	// 设置为 authorization code 模式
	srv.SetAllowedGrantType(oauth2.AuthorizationCode)
	// 根据 client id 从 manager 中获取 client info, 在获取 access token 校验过程中会被用到
	srv.SetClientInfoHandler(clientInfoHandler)
	// 校验授权请求用户的handler, 会重定向到 登陆页面, 返回"", nil
	srv.SetUserAuthorizationHandler(userAuthorizationHandler)
	// 校验授权请求的用户的账号密码, 给 LoginHandler 使用, 简单起见, 只允许一个用户授权
	srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	srv.SetInternalErrorHandler(func(err error) (re *oe.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *oe.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// auth_server 授权入口
	http.HandleFunc("/oauth2/authorize", authorizeHandler)
	// auth_server 发现未登录状态, 跳转到的登录handler
	http.HandleFunc("/oauth2/login.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	// auth_server 发现未登录状态, 跳转到的登录handler
	http.HandleFunc("/oauth2/login", loginHandler)
	// 登录完成, 同意授权的页面
	http.HandleFunc("/oauth2/agree_auth.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "agree_auth.html")
	})
	// auth_server 处理由code 换取access token 的handler
	http.HandleFunc("/oauth2/token", tokenHandler)
	// access token 换取用户信息的handler
	http.HandleFunc("/oauth2/getuserinfo", getUserInfoHandler)

	log.Println("oauth server addr: http://127.0.0.1:9096")
	log.Fatal(http.ListenAndServe(":9096", nil))
}

func clientInfoHandler(r *http.Request) (clientID, clientSecret string, err error) {
	client_info, err := srv.Manager.GetClient(r.Context(), r.FormValue("client_id"))
	log.Printf("client info: %v, error: %v", client_info, err)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	return client_info.GetID(), client_info.GetSecret(), nil
}

// AuthorizeHandler 内部使用, 用于查看是否有登陆状态
func userAuthorizationHandler(w http.ResponseWriter, r *http.Request) (user_id string, err error) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uid, ok := store.Get("LoggedInUserId")
	log.Printf("uid in userAuthorizationHandler: %v", uid)
	// 如果没有查询到登陆状态, 则跳转到 登陆页面
	if !ok {
		if r.Form == nil {
			_ = r.ParseForm()
		}

		w.Header().Set("Location", "/oauth2/login.html")
		w.WriteHeader(http.StatusFound)
		return "", nil
	}
	// 若有登录状态, 返回 user id
	user_id = uid.(string)
	return user_id, nil
}

func passwordAuthorizationHandler(ctx context.Context, clientID, username, password string) (userID string, err error) {
	if username == "test" && password == "123456" {
		return "0001", nil
	}
	return "", errors.New("username or password error")
}

// 授权入口, client/test.html 和 agree-auth.html 按下 button 后
func authorizeHandler(w http.ResponseWriter, r *http.Request) {
	err := srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// 登录页面的handler
func loginHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user_id, err := srv.PasswordAuthorizationHandler(r.Context(), "", r.Form.Get("username"), r.Form.Get("password"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	store.Set("LoggedInUserId", user_id) // 保存登录状态
	if err := store.Save(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 跳转到 同意授权页面
	w.Header().Set("Location", "/oauth2/agree_auth.html")
	w.WriteHeader(http.StatusFound)
}

// code 换取 access token
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	log.Println("client_id: ", r.Form.Get("client_id"))
	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		log.Println("token handler error: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// access token 换取用户信息
func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 获取 access token
	access_token, ok := srv.BearerAuth(r)
	if !ok {
		log.Println("Failed to get access token from request")
		return
	}

	root_ctx := context.Background()
	ctx, cancle_func := context.WithTimeout(root_ctx, time.Second)
	defer cancle_func()

	// 从 access token 中获取 信息
	token_info, err := srv.Manager.LoadAccessToken(ctx, access_token)
	if err != nil {
		log.Println(err)
		return
	}

	// 获取 user id
	user_id := token_info.GetUserID()
	grant_scope := token_info.GetScope()

	// 根据 grant scope 决定获取哪些用户信息
	if grant_scope != "read_user_info" {
		log.Println("invalid grant scope")
		_, _ = w.Write([]byte("invalid grant scope"))
		return
	}
	type UserInfo struct {
		Username string `json:"username"`
		Gender   string `json:"gender"`
	}
	var user_info_map = make(map[string]UserInfo)
	user_info_map["0001"] = UserInfo{
		"Tom", "Male",
	}
	user_info := user_info_map[user_id]
	resp, _ := json.Marshal(user_info)
	_, _ = w.Write(resp)
}
