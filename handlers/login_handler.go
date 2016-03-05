package handlers

import (
	"fmt"
	"github.com/barakb/github-branch/session"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
)

type GithubLoginHandler struct {
}

func (h GithubLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	fmt.Println("GithubLoginHandler")
	state := r.FormValue("state")

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	sess.Set("code", code)
	// On success, exchange this for an access token
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		panic(err)
	}
	sess.Set("user", user)
	sess.Set("*github.client", client)
	sess.Set("token", token)
	redirect := sess.Get("redirect")
	if redirect != nil && redirect != "" {
		fmt.Printf("Redirect to %v\n", redirect)
		sess.Delete("redirect")
		http.Redirect(w, r, redirect.(string), http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	fmt.Printf("user is %v\n", user)

	//user, _, err := client.Users.Get("")
	//if err != nil {
	//	fmt.Printf("client.Users.Get() failed with '%s'\n", err)
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}
	//buf := []string{"GitHub login id: ", *user.Login, "| GitHub email id: ", *user.Email}
	//userInfoTemplate.Execute(w, buf)

}
