package Clean

func CleanWorkspace(workspace *[][]*string) *[][]*string {

	for i := range *workspace {
		for j := range (*workspace)[i] {
			if (*workspace)[i][j] != nil {
				(*workspace)[i][j] = nil
			}
		}
	}
	return workspace
}
