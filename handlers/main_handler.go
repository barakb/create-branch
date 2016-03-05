package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	"html/template"
	"net/http"
)

type MainHandler struct {
	File string
}

func (h MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MainHandler")
	sess := session.GlobalSessions.SessionStart(w, r)
	//client := sess.Get("*github.client")
	user := sess.Get("user")
	t, _ := template.ParseFiles(h.File)
	m := make(map[string]interface{})
	m["user"] = user
	m["request"] = r;
	t.Execute(w, m)
}
