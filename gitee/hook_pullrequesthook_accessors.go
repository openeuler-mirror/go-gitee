package gitee

func (ph *PullRequestHook) GetUser() *UserHook {
	if ph == nil {
		return nil
	}

	return ph.User
}

func (ph *PullRequestHook) GetAssignee() *UserHook {
	if ph == nil {
		return nil
	}

	return ph.Assignee
}

func (ph *PullRequestHook) GetMilestone() *MilestoneHook {
	if ph == nil {
		return nil
	}

	return ph.Milestone
}

func (ph *PullRequestHook) GetHead() *BranchHook {
	if ph == nil {
		return nil
	}

	return ph.Head
}

func (ph *PullRequestHook) GetBase() *BranchHook {
	if ph == nil {
		return nil
	}

	return ph.Base
}
func (ph *PullRequestHook) GetUpdatedBy() *UserHook {
	if ph == nil {
		return nil
	}

	return ph.UpdatedBy
}