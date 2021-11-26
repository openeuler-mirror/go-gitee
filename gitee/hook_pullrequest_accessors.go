package gitee


func (p *PullRequestEvent) GetAction() string{
	if p == nil || p.Action == nil {
		return ""
	}
	return *p.Action
}

func (p *PullRequestEvent) GetActionDesc() string{
	if p == nil || p.ActionDesc == nil {
		return ""
	}
	return *p.ActionDesc
}

func (p *PullRequestEvent) GetPullRequest() *PullRequestHook {
	if p == nil || p.PullRequest == nil {
		return nil
	}
	return p.PullRequest
}

func (p *PullRequestEvent) GetAuthor() *UserHook {
	if p == nil || p.Author == nil {
		return nil
	}
	return p.Author
}

func (p *PullRequestEvent) GetProject() *ProjectHook {
	if p == nil || p.Project == nil {
		return nil
	}
	return p.Project
}

func (p *PullRequestEvent) GetRepository() *ProjectHook {
	if p == nil || p.Repository == nil {
		return nil
	}
	return p.Repository
}

func (p *PullRequestEvent) GetTitle() string{
	if p == nil || p.Title == nil {
		return ""
	}
	return *p.Title
}

func (p *PullRequestEvent) GetBody() string{
	if p == nil || p.Body == nil {
		return ""
	}
	return *p.Body
}

func (p *PullRequestEvent) GetState() string{
	if p == nil || p.State == nil {
		return ""
	}
	return *p.State
}

func (p *PullRequestEvent) GetMergeStatus() string{
	if p == nil || p.MergeStatus == nil {
		return ""
	}
	return *p.MergeStatus
}

func (p *PullRequestEvent) GetMergeCommitSha() string{
	if p == nil || p.MergeCommitSha == nil {
		return ""
	}
	return *p.MergeCommitSha
}

func (p *PullRequestEvent) GetURL() string{
	if p == nil || p.URL == nil {
		return ""
	}
	return *p.URL
}

func (p *PullRequestEvent) GetSourceBranch() string{
	if p == nil || p.SourceBranch == nil {
		return ""
	}
	return *p.SourceBranch
}

func (p *PullRequestEvent) GetTargetBranch() string{
	if p == nil || p.TargetBranch == nil {
		return ""
	}
	return *p.TargetBranch
}

func (p *PullRequestEvent) GetSourceRepo() *RepoInfo {
	if p == nil || p.SourceRepo == nil {
		return nil
	}
	return p.SourceRepo
}

func (p *PullRequestEvent) GetTargetRepo() *RepoInfo {
	if p == nil || p.TargetRepo == nil {
		return nil
	}
	return p.TargetRepo
}

func (p *PullRequestEvent) GetUpdatedBy() *UserHook {
	if p == nil || p.UpdatedBy == nil {
		return nil
	}
	return p.UpdatedBy
}

func (p *PullRequestEvent) GetSender() *UserHook {
	if p == nil || p.Sender == nil {
		return nil
	}
	return p.Sender
}

func (p *PullRequestEvent) GetTargetUser() *UserHook {
	if p == nil || p.TargetUser == nil {
		return nil
	}
	return p.TargetUser
}

func (p *PullRequestEvent) GetEnterprise() *EnterpriseHook {
	if p == nil || p.Enterprise == nil {
		return nil
	}
	return p.Enterprise
}

func (p *PullRequestEvent) GetHookName() string{
	if p == nil || p.HookName == nil {
		return ""
	}
	return *p.HookName
}

func (p *PullRequestEvent) GetPassword() string{
	if p == nil || p.Password == nil {
		return ""
	}
	return *p.Password
}
