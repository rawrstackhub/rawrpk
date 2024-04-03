package common

const (
	Github = iota + 1
)

type RepoFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"download_url"`
}

type PkgData struct {
	Title   string   //EX: lsf
	Mainsrc int8     //EX: Github
	Gblsrc  []string //EX: raw.github.com/rawrstackhub/lsf
	Lcldir  string   //EX: C:/users/username/rawrpk/lsf
}
