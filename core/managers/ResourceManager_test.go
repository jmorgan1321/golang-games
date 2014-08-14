package managers

import (
    "github.com/jmorgan1321/golang-games/core/debug"
	"github.com/jmorgan1321/golang-games/core/types"
	"strings"
	"testing"
)

type ResourceManagerInitComponent struct {
    restartOnChange bool
}
func (*ResourceManagerInitComponent) isComponent() {}


func ResourceManagerAcceptanceTest(t *testing.T) {
    // config := setUpMockConfig()
	config := &BasicSpace{}
    config.AddComponent()

    config.CoreData.RestartChan := make(chan bool)

    rsrcMgr := NewResourceManager()
	rsrcMgr.Init(config)

	// // init := config.GetComponent(ResourceManagerInitComponent)
	// debug.ExpectEQ(t, true, rsrcMgr.restartOnChange, "Bad init (restartOnChange).")
	// // debug.ExpectEQ(t, "", rsrcMgr.rootWatchDir, "Bad init (rootWatchDir).")

	// pathFromRsrcDir := "test/json/objects/testFile.jrm"
	// fileData, err := rscrMgr.GetFile(pathFromRsrcDir)
	// debug.ExpectOK(err)

	// expData := `{"test1":["item1", "item2"], "test2": true}`
	// debug.ExpectEQ(t, strings.TrimSpace(expData), strings.TrimSpace(fileData), "serialziation failed.")

	// ModifyTestFile()
	// debug.ExpectEQ(t, 0, len(rsrcMgr.FileChanges), "changes updated too soon.")

	// rsrcMgr.BeginFrame()
	// debug.ExpectEQ(t, 1, len(rsrcMgr.FileChanges), "changes didn't update.")
	// expChanges := []string{"test/json/objects/testFile.jrm"}
	// debug.ExpectEQ(t, expChanges, rsrcMgr.FileChanges, "file not stored correctly.")

	// debug.ExpectEQ(t, false, core.Restart, "core shouldn't be notified yet.")
	// rsrcMgr.EndFrame()
	// debug.ExpectEQ(t, true, core.Restart, "core wasn't notified of restart.")
}
