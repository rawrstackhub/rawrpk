package common

const (
	Github = iota + 1
)

type RepoFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"download_url"`
}

type PkFile struct {
	Name       string
	FileURL    string
	InstallLoc string
}

var Pack PkFile
