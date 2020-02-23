package FkGitOps

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/url"
)

const (
	STATUSPENDING	string = "pending"
	STATUSSUCCESS   string = "success"
	STATUSFAILURE   string = "failure"
	STATUSERROR		string = "error"
)

type FkGitConfig struct {
	AccessToken	string `yaml:"token" json:"token"`
	ApiURL 		string `yaml:"url" json:"url"`
}

//func NewGitClient (fkGitConfig *FkGitConfig) (gitClient *github.Client, err error) {
//
//	//fkGitConfig.Ctxt = context.Background()
//	ts := oauth2.StaticTokenSource(
//		&oauth2.Token{AccessToken: fkGitConfig.AccessToken},
//	)
//	tc := oauth2.NewClient(fkGitConfig.Ctxt, ts)
//	gitClient,err = github.NewEnterpriseClient(fkGitConfig.ApiURL,fkGitConfig.ApiURL,tc)
//	return gitClient,err
//}

func NewGitClient2 (fkGitConfig FkGitConfig,ctx context.Context) (gitClient *github.Client, err error) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: fkGitConfig.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	//gitClient,_= github.NewEnterpriseClient(fkGitConfig.ApiURL,fkGitConfig.ApiURL,tc)
	gitClient = github.NewClient(tc)
	u, _ := url.Parse(fkGitConfig.ApiURL)
	gitClient.BaseURL = u
	gitClient.UploadURL = u
	return gitClient,err
}

