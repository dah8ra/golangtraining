package system

type FileSystem interface {
	GetFiles(userInfo *map[string]string) ([]map[string]string, error)
}

type FileManager struct {
	FileSystem
}

type DefaultFileSystem struct {
}

func (dfs DefaultFileSystem) GetFiles(userInfo *map[string]string) ([]map[string]string, error) {
	files := make([]map[string]string, 0)
	for i := 0; i < 5; i++ {
		file := make(map[string]string)
		file["size"] = "90210"
		file["name"] = "ftp.txt"
		files = append(files, file)
	}
	return files, nil
}

func NewDefaultFileSystem() *FileManager {
	fm := FileManager{}
	fm.FileSystem = DefaultFileSystem{}
	return &fm
}