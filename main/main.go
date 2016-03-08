package main

import (
	"fmt"
	"github.com/barakb/github-branch/handlers"
	"github.com/barakb/github-branch/session"
	"net/http"
	"golang.org/x/net/websocket"
)

func init() {
	gs, err := session.NewManager("memory", "gosessionid", 10*1000*1000*1000)
	if err != nil {
		fmt.Printf("failed to create session manager error is: %#v\n", err)
		return
	}
	session.GlobalSessions = gs
}

func main() {
	/*
		//createBranches(client, []string{"xap", "b", "c", "d"})

		repos := createReposFromNames(reposNames)
		done := fillTipSha(repos, client)

		select {
		case <-done:
		}

		fmt.Println("Done reading!")
		for _, repo := range repos {
			fmt.Println(repo)
		}

		done = createBranches(repos, "barak-test-branch", client)
		select {
		case <-done:
		}

		fmt.Println("Done creating !")

		done = deleteBranches(repos, "barak-test-branch", client);

		select {
		case <-done:
		}

		fmt.Println("Done Deleting !")
	*/

	go session.GlobalSessions.GC()

	http.Handle("/", handlers.MustAuth(handlers.MainHandler{File: "index.html"}))
	http.Handle("/events", handlers.MustAuth(websocket.Handler(handlers.EchoServer)))
	http.Handle("/logout", handlers.MustAuth(handlers.LogoutHandler{}))
	http.Handle("/githuboa_cb", handlers.GithubLoginHandler{})
	http.Handle("/api/create_branch/", handlers.MustAuth(handlers.CreateBranchHandler{}))
	http.Handle("/api/delete_branch/", handlers.MustAuth(handlers.DeleteBranchHandler{}))
	http.Handle("/api/get_branches/", handlers.MustAuth(handlers.GetBranchsHandler{}))
	http.Handle("/web/", handlers.MustAuth(http.StripPrefix("/web/", http.FileServer(http.Dir("web")))))
	fmt.Print("Started running on https://localhost:4430\n")
	fmt.Println(http.ListenAndServeTLS(":4430", "keys/server.pem", "keys/server.key", nil))

}
