package main

import (
	"fmt"
	"github.com/barakb/github-branch/handlers"
	"github.com/barakb/github-branch/session"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
	"sync"
)

func init() {
	gs, err := session.NewManager("memory", "gosessionid", 10*1000*1000*1000)
	if err != nil {
		fmt.Printf("failed to create session manager error is: %#v\n", err)
		return
	}
	session.GlobalSessions = gs
}

var reposNames = []string{"CloudifySource/Cloudify-iTests-webuitf",
	"CloudifySource/Cloudify-iTests",
	"GigaSpaces-QA/devops",
	"Gigaspaces/xap-ui-desktop",
	"Gigaspaces/RESTData",
	"Gigaspaces/xap-session-sharing-manager",
	"CloudifySource/iTests-Framework",
	"GigaSpaces/mongo-datasource-itests",
	"GigaSpaces/mongo-datasource",
	"GigaSpaces/xap-mule",
	"Gigaspaces/xap-scala",
	"GigaSpaces/petclinic-jpa",
	"GigaSpaces/xap-apm-introscope",
	"GigaSpaces/xap-blobstore-mapdb",
	"GigaSpaces/xap-blobstore-rocksdb",
	"GigaSpaces/xap-cassandra",
	"GigaSpaces/xap-cpp",
	"GigaSpaces/xap-dotnet",
	"GigaSpaces/xap-example-data",
	"GigaSpaces/xap-example-helloworld",
	"GigaSpaces/xap-example-tutorials",
	"GigaSpaces/xap-example-web",
	"GigaSpaces/xap-jetty",
	"GigaSpaces/xap-jms",
	"GigaSpaces/xap-maven-plugin",
	"GigaSpaces/xap-rest",
	"GigaSpaces/xap-session-sharing-manager-itests",
	"GigaSpaces/xap-spatial",
	"GigaSpaces/xap-spring-data",
	"GigaSpaces/xap-ui-web",
	"GigaSpaces/xap"}

//var reposNames = []string{"barakb/foo"}

type Repo struct {
	owner string
	repo  string
	sha   string
}

func (repo Repo) String() string {
	return fmt.Sprintf("Repo{owner:%s, repo:%s, sha:%s}", repo.owner, repo.repo, repo.sha)
}

func createReposFromNames(names []string) []*Repo {
	ret := make([]*Repo, len(names))
	for index, name := range names {
		components := strings.Split(name, "/")
		ret[index] = &Repo{owner: components[0], repo: components[1]}
	}
	return ret
}

func fillTipSha(repos []*Repo, client *github.Client) chan struct{} {
	var wg sync.WaitGroup
	done := make(chan struct{})
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()
			//fmt.Printf("working on repo %#v\n", repo)
			ref, _, err := client.Git.GetRef(repo.owner, repo.repo, "refs/heads/master")
			if err != nil {
				fmt.Errorf("error: %#v\n", err)
			}
			//fmt.Printf("Ref: %s\nURL: %s\n", *ref.Ref, *ref.URL)
			//fmt.Printf("Type: %s\nURL: %s\nSHA: %s\n", *ref.Object.Type, *ref.Object.URL, *ref.Object.SHA)
			repo.sha = *ref.Object.SHA
		}(repo)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}

func createBranches(repos []*Repo, branch string, client *github.Client) chan struct{} {
	var wg sync.WaitGroup
	done := make(chan struct{})
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()

			_, _, err := client.Git.CreateRef(repo.owner, repo.repo, &github.Reference{
				Ref: github.String("refs/heads/" + branch),
				Object: &github.GitObject{
					SHA: &repo.sha,
				},
			})
			if err != nil {
				fmt.Errorf("error: %#v\n", err)
			}
		}(repo)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}

func deleteBranches(repos []*Repo, branch string, client *github.Client) chan struct{} {
	var wg sync.WaitGroup
	done := make(chan struct{})
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()

			_, err := client.Git.DeleteRef(repo.owner, repo.repo, "refs/heads/"+branch)
			if err != nil {
				fmt.Errorf("error: %#v\n", err)
			}
		}(repo)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}

//func getTheTipOfTheMaster(repos []string) chan

var oauthConf *oauth2.Config
var oauthStateString string

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
	http.Handle("/logout", handlers.MustAuth(handlers.LogoutHandler{}))
	http.Handle("/githuboa_cb", handlers.GithubLoginHandler{})
	http.Handle("/api/create_branch/", handlers.CreateBranchHandler{})
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	fmt.Print("Started running on http://localhost:8080\n")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
