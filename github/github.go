package github

import (
	"fmt"
	"github.com/google/go-github/github"
	"strings"
	"sync"
"sync/atomic"
)

var ReposNames = []string{"CloudifySource/Cloudify-iTests-webuitf",
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


type Repo struct {
	owner string
	repo  string
	sha   string
}

func (repo Repo) String() string {
	return fmt.Sprintf("Repo{owner:%s, repo:%s, sha:%s}", repo.owner, repo.repo, repo.sha)
}

func createReposFromNames(names []string) []*Repo {
	ret := make([]*Repo, 0)
	for _, name := range names {
		if strings.TrimSpace(name) != "" && strings.Contains(name, "/"){
			components := strings.Split(name, "/")
			ret = append(ret, &Repo{owner: components[0], repo: components[1]})
		}
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

func createBranches(repos []*Repo, branch string, client *github.Client) (chan *Repo, chan struct{}, *int32) {
	var counter int32
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
				atomic.AddInt32(&counter, 1)
			}
			reposChan <- repo
		}(repo)
	}
	go func() {
		wg.Wait()
		close(reposChan)
		close(done)
	}()
	return reposChan, done, &counter
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

func DeleteBranch(branch string, client *github.Client) {
	repos := createReposFromNames(ReposNames)
	done := deleteBranches(repos, branch, client)

	select {
	case <-done:
	}
}

func CreateBranch(branch string, client *github.Client) (chan *Repo, chan struct{}, *int32) {
	repos := createReposFromNames(ReposNames)
	fillTipSha(repos, client).Wait()

	res, done, counter := createBranches(repos, branch, client)

	select {
	case <-done:
	}

	//done = deleteBranches(repos, branch, client);
	return res, done, counter
}

func ListRefs(client *github.Client, owner string, repo string) ([]github.Reference, error) {
	refs, _, err := client.Git.ListRefs(owner, repo, &github.ReferenceListOptions{Type:"heads/"});
	//refs, _, err := client.Git.ListRefs(owner, repo, &github.ReferenceListOptions{Type:"heads"});
	return refs, err;
}

type RefList struct {

}

type UIBranch struct {
	Name     string `json:"name"`
	Quantity int `json:"quantity"`
}

func ListAllRefs(client *github.Client) ([]*UIBranch, error) {
	repos := createReposFromNames(ReposNames)
	resChan := make(chan []github.Reference, len(repos))
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *Repo) {
			defer wg.Done()
			heads, err := ListRefs(client, repo.owner, repo.repo)
			if err == nil {
				resChan <- heads
			}else {
				fmt.Printf("List all ref error: %#v, on repo %#v\n", err, repo)
			}
		}(repo)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()

	branchesMap := collectBranches(resChan)
	return toSlice(branchesMap), nil;
}

func toSlice(m map[string]*UIBranch) []*UIBranch {
	res := make([]*UIBranch, len(m))
	var index = 0;
	for _, uiBranch := range m {
		res[index] = uiBranch
		index += 1
	}
	return res
}

func collectBranches(c chan []github.Reference) map[string]*UIBranch {
	res := make(map[string]*UIBranch)
	for {
		select {
		case refs, ok := <-c:
			if !ok {
				return res
			}
			for _, ref := range refs {
				name := strings.TrimPrefix(*ref.Ref, "refs/heads/")
				uiBranch, found := res[name]
				if !found {
					uiBranch = &UIBranch{Name:name, Quantity:0}
					res[name] = uiBranch
				}
				uiBranch.Quantity += 1
			}

		}
	}
	return res
}
