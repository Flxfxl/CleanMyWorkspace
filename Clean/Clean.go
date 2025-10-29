package Clean

func CleanWorkSpace(workspace *[][]*string) *[][]*string {
	if workspace == nil {
		return workspace
	}

	a := *workspace
	for i := range a {
		for j := range a[i] {
			a[i][j] = nil
		}
	}
	return &a
}
