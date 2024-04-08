package common

const (
	Github = iota + 1
)

const (
	Install = iota + 1
	EnvVar
)

type RepoFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"download_url"`
}

type PkgData struct {
	Title      string   //EX: lsf
	Source     []string //EX: github / rawrstackhub
	InstallLoc string   //EX: C:\Users\user\rawrpk\lsf
	PkgURL     string   //EX: https://api.github.com/repos/rawrstackhub/lsf/contents/lsf.rawrpk
	DwnSrc     string   //EX: https://raw.githubusercontent.com/rawrstackhub/lsf/main/lsf.rawrpk

}
