package gitee

func (pj *ProjectHook) GetNameSpace() string {
	if pj == nil {
		return ""
	}

	return pj.Namespace
}

func (pj *ProjectHook) GetPath() string {
	if pj == nil {
		return ""
	}

	return pj.Path
}