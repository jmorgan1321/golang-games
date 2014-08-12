package managers

type ResourceManager struct {
}

func (r *ResourceManager) GetFileData(filename string) (string, error) {
	mockFileSystem := map[string]string{
		"LevelSpaceFile": "{}",
		"GocFile":        "{}",
	}

	return mockFileSystem[filename], nil
}

func (r *ResourceManager) Construct() error {
	return nil
}
