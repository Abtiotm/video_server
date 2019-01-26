package main

import (
	"github.com/Abtiotm/video_server/api/session"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)  //注册用户

	router.POST("/user/:username", Login)  //登陆

	router.GET("/user/:username", GetUserInfo)  //得到用户信息

	router.POST("/user/:username/videos", AddNewVideo)  //添加视频

	router.GET("/user/:username/videos", ListAllVideos)  //显示所有视频

	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)  删除视频

	router.POST("/videos/:vid-id/comments", PostComment)  //提交评论

	router.GET("/videos/:vid-id/comments", ShowComments)  //显示评论

	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}




