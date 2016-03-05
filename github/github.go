package github

import (
	"fmt"
	"github.com/google/go-github/github"
	"strings"
	"sync"
)

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

func fillTipSha(repos []*Repo, client *github.Client) *sync.WaitGroup {
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()
			//fmt.Printf("working on repo %#v\n", repo)
			ref, _, err := client.Git.GetRef(repo.owner, repo.repo, "refs/heads/master")
			if err != nil {
				fmt.Printf("error: %#v\n", err)
			}
			//fmt.Printf("Ref: %s\nURL: %s\n", *ref.Ref, *ref.URL)
			//fmt.Printf("Type: %s\nURL: %s\nSHA: %s\n", *ref.Object.Type, *ref.Object.URL, *ref.Object.SHA)
			repo.sha = *ref.Object.SHA
		}(repo)
	}
	return &wg
}

func createBranches(repos []*Repo, branch string, client *github.Client) (chan *Repo, chan struct{}) {
	var wg sync.WaitGroup
	reposChan := make(chan *Repo, len(repos))
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
				fmt.Printf("error createing branch %s in repo: %#v, error is %s:\n", branch, repo, err.Error())
			}else {
				fmt.Printf("branch %s created in repo %#v\n", branch, repo)
			}
			reposChan <- repo
		}(repo)
	}
	go func() {
		wg.Wait()
		close(reposChan)
		close(done)
	}()
	return reposChan, done
}

func deleteBranches(repos []*Repo, branch string, client *github.Client) chan struct{} {
	var wg sync.WaitGroup
	done := make(chan struct{})
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()

			_, err := client.Git.DeleteRef(repo.owner, repo.repo, "refs/heads/" + branch)
			if err != nil {
				fmt.Printf("error deleting branch %s in repo: %#v, error is %s:\n", branch, repo, err.Error())
			}else {
				fmt.Printf("branch %s deleted from repo %#v\n", branch, repo)
			}
		}(repo)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}

func CreateBranch(branch string, client *github.Client) (chan *Repo, chan struct{}) {
	repos := createReposFromNames(reposNames)
	fillTipSha(repos, client).Wait()

	fmt.Println("After fill tip sha !\n")
	res, done := createBranches(repos, branch, client)

	select {
	case <-done:
	}

	done = deleteBranches(repos, branch, client);
	return res, done
}

func ListRefs(client *github.Client, owner string, repo string) ([]github.Reference,  error){
	refs, _, err := client.Git.ListRefs(owner, repo, &github.ReferenceListOptions{Type:"heads/"});
	//refs, _, err := client.Git.ListRefs(owner, repo, &github.ReferenceListOptions{Type:"heads"});
	return refs, err;
}
