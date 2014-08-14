package managers

/*
import (
    "code.google.com/p/go.exp/fsnotify"
    "fmt"
)

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        fmt.Println("error:", err)
        return
    }

    go func() {
        err = watcher.Watch("D:/work/fbl_grfx_dev_p/windows/sandbox/go_tools/rsrc")
        if err != nil {
            fmt.Println("watch error:", err)
            return
        }
    }()

    for {
        select {
        case ev := <-watcher.Event:
            fmt.Println("event:", ev)
        case err := <-watcher.Error:
            fmt.Println("error:", err)
        }
    }
}
*/

/*



*/

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
