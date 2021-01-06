package syncthing

import (
	"io/ioutil"
	"testing"
)

// normally download with version
func TestDownLoad(t *testing.T) {
	versionToDownload := "v0.0.1"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath: tmpDir,
		Version: versionToDownload,
	}

	// download v0.0.1
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost") != versionToDownload {
		t.Error("Download version did not match the version: " + versionToDownload)
	}
}

// normally download with version and repeat 'InstallIfNeeded()'
func TestRepeatDownload(t *testing.T) {
	versionToDownload := "v0.0.1"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath: tmpDir,
		Version: versionToDownload,
	}

	// download v0.0.1
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost") != versionToDownload {
		t.Error("Download version did not match the version: " + versionToDownload)
	}

	for i := 0; i < 10; i++ {
		// download v0.0.1
		downloadAgain, err := mockedSyncthingInstaller.InstallIfNeeded()
		if err != nil {
			t.Fatal(err)
		}
		if downloadAgain {
			t.Fatal()
		}
	}
}

// though specify the commit id, download the version of `VERSION` is high priority
func TestDownLoadWithCommitId(t *testing.T) {
	versionToDownload := "v0.0.1"
	commitId := "forunittestwithmockedcommitid"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:  tmpDir,
		Version:  versionToDownload,
		CommitId: commitId,
	}

	// download v0.0.1
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost") != versionToDownload {
		t.Error("Download version did not match the version: " + versionToDownload)
	}
}

// though specify the commit id, download the version of `VERSION` is high priority and repeat 'InstallIfNeeded()'
func TestRepeatDownLoadWithCommitId(t *testing.T) {
	versionToDownload := "v0.0.1"
	commitId := "forunittestwithmockedcommitid"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:  tmpDir,
		Version:  versionToDownload,
		CommitId: commitId,
	}

	// download v0.0.1
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost") != versionToDownload {
		t.Error("Download version did not match the version: " + versionToDownload)
	}

	for i := 0; i < 10; i++ {
		// download v0.0.1
		downloadAgain, err := mockedSyncthingInstaller.InstallIfNeeded()
		if err != nil {
			t.Fatal(err)
		}
		if downloadAgain {
			t.Fatal()
		}
	}
}

// specify the commit id, and the VERSION is empty
func TestRepeatDownLoadWithCommitIdAndEmptyVersion(t *testing.T) {
	versionToDownload := ""
	commitId := "forunittestwithmockedcommitid"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:  tmpDir,
		Version:  versionToDownload,
		CommitId: commitId,
	}

	// download v0.0.1
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost-commit-id") != commitId {
		t.Error("Download version did not match the version: " + commitId)
	}

	for i := 0; i < 10; i++ {
		// download v0.0.1
		downloadAgain, err := mockedSyncthingInstaller.InstallIfNeeded()
		if err != nil {
			t.Fatal(err)
		}
		if downloadAgain {
			t.Fatal()
		}
	}
}

// specify the commit id, and the VERSION is empty
func TestDownLoadWithCommitIdAndEmptyVersion(t *testing.T) {
	versionToDownload := ""
	commitId := "forunittestwithmockedcommitid"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:  tmpDir,
		Version:  versionToDownload,
		CommitId: commitId,
	}

	// download forunittestwithmockedcommitid
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost-commit-id") != commitId {
		t.Error("Download version did not match the version: " + commitId)
	}
}

// specify the commit id, and the VERSION is empty
func TestDownLoadWithCommitIdAnInvalidVersion(t *testing.T) {
	versionToDownload := "invalidVersion"
	commitId := "forunittestwithmockedcommitid"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:  tmpDir,
		Version:  versionToDownload,
		CommitId: commitId,
	}

	// download forunittestwithmockedcommitid
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost-commit-id") != commitId {
		t.Error("Download version did not match the version: " + commitId)
	}
}

// specify the commit id, and the VERSION is empty
func TestRepeatDownLoadWithCommitIdAnInvalidVersion(t *testing.T) {
	versionToDownload := "invalidVersion"
	commitId := "forunittestwithmockedcommitid"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:  tmpDir,
		Version:  versionToDownload,
		CommitId: commitId,
	}

	// download forunittestwithmockedcommitid
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost-commit-id") != commitId {
		t.Error("Download version did not match the version: " + commitId)
	}

	for i := 0; i < 3; i++ {
		downloadAgain, err := mockedSyncthingInstaller.InstallIfNeeded()
		if err != nil {
			t.Fatal(err)
		}

		// this case will download again because the version 'invalidVersion' can not be downloaded
		if !downloadAgain {
			t.Fatal()
		}

		if mockedSyncthingInstaller.exec("-nocalhost-commit-id") != commitId {
			t.Error("Download version did not match the version: " + commitId)
		}
	}
}

// specify specifyVersion
func TestDownLoadWithSpecifyVersion(t *testing.T) {
	versionToDownload := "invalidVersion"
	commitId := "forunittestwithmockedcommitid"
	specifyVersion := "v0.2.0"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:        tmpDir,
		Version:        versionToDownload,
		CommitId:       commitId,
		SpecifyVersion: specifyVersion,
	}

	// download v0.2.0
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost") != specifyVersion {
		t.Error("Download version did not match the version: " + specifyVersion)
	}
}


// specify specifyVersion
func TestRepeatDownLoadWithSpecifyVersion(t *testing.T) {
	versionToDownload := "invalidVersion"
	commitId := "forunittestwithmockedcommitid"
	specifyVersion := "v0.2.0"
	tmpDir, _ := ioutil.TempDir("", "")

	mockedSyncthingInstaller := &SyncthingInstaller{
		BinPath:        tmpDir,
		Version:        versionToDownload,
		CommitId:       commitId,
		SpecifyVersion: specifyVersion,
	}

	// download v0.2.0
	_, err := mockedSyncthingInstaller.InstallIfNeeded()
	if err != nil {
		t.Fatal(err)
	}

	if mockedSyncthingInstaller.exec("-nocalhost") != specifyVersion {
		t.Error("Download version did not match the version: " + specifyVersion)
	}

	for i := 0; i < 3; i++ {
		// download v0.2.0
		downloadAgain, err := mockedSyncthingInstaller.InstallIfNeeded()
		if err != nil {
			t.Fatal(err)
		}

		if downloadAgain {
			t.Fatal()
		}
	}
}