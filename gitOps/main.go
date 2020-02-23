package main

import (
	"context"
	"fmt"
	"github.com/ambsflip/goDemos/gitOps/FkGitOps"
	"github.com/google/go-github/github"
)

//Ref : https://godoc.org/github.com/google/go-github/github#pkg-examples
func main() {

	ctx := context.Background()

	//ambsflip
	var fkGitConfig FkGitOps.FkGitConfig
	owner := "ambsflip" // is this mandatory?
	repo := "goDemos"
	ref := "318d290522bad83cc2cfb198a29f40230015dd7c"
	fkGitConfig.AccessToken = "df11350f67fba49899e05bf807cdd4b5f8b9218b"
	fkGitConfig.ApiURL = "https://api.github.com/"

	gitClient,_ := FkGitOps.NewGitClient2(fkGitConfig,ctx)

	statusState     := "pending" //"pending" "success" "error"
	targetUrl   := "https://xyz.ngrok.com/status"
	pendingDesc := "Build/testing in progress, please wait."
	appName     := "TEST-CI"

	status1 := &github.RepoStatus{State: &statusState, TargetURL: &targetUrl, Description: &pendingDesc, Context: &appName}
	status, resp , err := gitClient.Repositories.CreateStatus(ctx,owner,repo,ref,status1)
	if err != nil{
		fmt.Println("err:	",err)
	}
	fmt.Println("status:	",status)
	fmt.Println("resp:	",resp)


	comment := "Test Result \n"+"{\n  \"failed\": 0,\n  \"broken\": 0,\n  \"skipped\": 0,\n  \"passed\": 1,\n  \"unknown\": 0,\n  \"total\": 1\n}"
	commentObj := github.IssueComment{Body: &comment}
	issueComment,resp,err := gitClient.Issues.CreateComment(ctx,owner,repo,1,&commentObj)
	fmt.Println("issueComment: ",issueComment)
	fmt.Println("resp: ",resp)
	fmt.Println("err: ",err)
	//prid := 1
	//fileName1 := "gitOps/commentBody.json"
	//fileContentBytes1, err := ioutil.ReadFile(fileName1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var commentObj github.PullRequestComment
	//err = json.Unmarshal(fileContentBytes1, &commentObj)
	//if err != nil {
	//	fmt.Printf("Unmarshal: %v", err)
	//}
	//
	//prComment, prResponse, err := gitClient.PullRequests.CreateComment(ctx,owner,repo,prid,&commentObj)
    //fmt.Println("prComment: ", prComment)
	//fmt.Println("prResponse: ", prResponse)
	//fmt.Println("err: ", err)
	//	client := github.NewClient(nil)

	//orgs, _, err := client.Organizations.List(context.Background(), "ambsflip", nil)
	//if err == nil {
	//	fmt.Println(orgs)
	//}

	//opt := &github.RepositoryListAllOptions{Since:0}
	//repos, _, err := client.Repositories...ListAll(context.Background(),opt)
	//if err == nil {
	//	for i := range repos {
	//		fmt.Println(repos[i].Name)
	//	}
	//}

}
