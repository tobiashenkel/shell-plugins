package gitxargs

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func gitxargsCLI() schema.Executable {
	return schema.Executable{
		Name:    "git-xargs CLI",
		Runs:    []string{"git-xargs"},
		DocsURL: sdk.URL("https://github.com/gruntwork-io/git-xargs"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.PersonalAccessToken,
			},
		},
	}
}
