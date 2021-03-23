package test

import (
	"gitee.com/openeuler/go-gitee/gitee"
	"github.com/antihax/optional"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"testing"
)

var client *gitee.APIClient

func init() {
	token := "your gitee token"

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	conf := gitee.NewConfiguration()
	conf.HTTPClient = oauth2.NewClient(context.Background(), ts)

	client = gitee.NewAPIClient(conf)
}

func TestPutV5ReposOwnerRepoPullsNumberLabels(t *testing.T) {
	op := &gitee.PutV5ReposOwnerRepoPullsNumberLabelsOpts{
		Body: optional.NewInterface([]string{"feature", "go-gitee-test"}),
	}
	labels, _, err := client.PullRequestsApi.PutV5ReposOwnerRepoPullsNumberLabels(context.Background(), "cve-manage-test", "phpp", 1, op)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(labels)
	}
}

func TestPutV5ReposOwnerRepoIssuesNumberLabels(t *testing.T) {
	op := &gitee.PutV5ReposOwnerRepoIssuesNumberLabelsOpts{
		Body: optional.NewInterface([]string{"feature", "go-gitee-test"}),
	}
	labels, _, err := client.LabelsApi.PutV5ReposOwnerRepoIssuesNumberLabels(context.Background(), "cve-manage-test", "config", "I3A2AO", op)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(labels)
	}

}
