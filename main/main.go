package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/websocket"
	"flag"
	"github.com/vharitonsky/iniflags"
	"os"
	"io/ioutil"
	"regexp"
	gh "github.com/barakb/create-branch/github"
	"github.com/barakb/create-branch/session"
	"github.com/barakb/create-branch/handlers"
)

func init() {
	gs, err := session.NewManager("memory", "gosessionid", 10*1000*1000*1000)
	if err != nil {
		fmt.Printf("failed to create session manager error is: %#v\n", err)
		return
	}
	session.GlobalSessions = gs
}

var port = flag.Int("port", 4430, "Configure the server port")

var repos = flag.String("repos", defaultRepoFile(), "Configure the where the repositories file")

func main() {
	iniflags.Parse()
	repos, err := readRepos(*repos)
	if err != nil{
		fmt.Printf("Error:%s\n", err.Error())
		return
	}
	gh.ReposNames = repos
	go session.GlobalSessions.GC()
	http.Handle("/", handlers.MustAuth(handlers.MainHandler{File: "index.html"}))
	http.Handle("/events", handlers.MustAuth(websocket.Handler(handlers.EchoServer)))
	http.Handle("/logout", handlers.MustAuth(handlers.LogoutHandler{}))
	http.Handle("/githuboa_cb", handlers.GithubLoginHandler{})
	http.Handle("/api/create_branch/", handlers.MustAuth(handlers.CreateBranchHandler{}))
	http.Handle("/api/delete_branch/", handlers.MustAuth(handlers.DeleteBranchHandler{}))
	http.Handle("/api/get_branches/", handlers.MustAuth(handlers.GetBranchsHandler{}))
	http.Handle("/web/", handlers.MustAuth(http.StripPrefix("/web/", http.FileServer(http.Dir("web")))))
	fmt.Printf("Server started on https://localhost:%d\n", *port)
	fmt.Println(http.ListenAndServeTLS(fmt.Sprintf(":%d", *port), "keys/server.pem", "keys/server.key", nil))
}

func defaultRepoFile() string{
	dir, err := os.Getwd()
	if err != nil{
		return "repos.txt"
	}
	return fmt.Sprintf("%s/repos.txt", dir)
}

func readRepos(repos string) ([]string, error){
	bytes, err := ioutil.ReadFile(repos)
	if err != nil{
		return nil, err
	}
	str := string(bytes)
	return regexp.MustCompile("\\s+").Split(str, -1), nil
}