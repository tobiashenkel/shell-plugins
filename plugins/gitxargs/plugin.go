package gitxargs

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "gitxargs",
		Platform: schema.PlatformInfo{
			Name:     "git-xargs",
			Homepage: sdk.URL("https://github.com/gruntwork-io/git-xargs"),
		},
		Credentials: []schema.CredentialType{
			PersonalAccessToken(),
		},
		Executables: []schema.Executable{
			gitxargsCLI(),
		},
	}
}
