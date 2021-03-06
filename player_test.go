package radiko

import (
	"path"
	"testing"
)

func TestDownloadPlayer(t *testing.T) {
	dir, removeDir := createTestTempDir(t)
	defer removeDir() // clean up

	playerPath := path.Join(dir, "myplayer.swf")
	err := DownloadPlayer(playerPath)
	if err != nil {
		t.Errorf("Failed to download player.swf: %s", err)
	}
}
