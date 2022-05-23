package didutil

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// DidUtilCmd returns did cobra Command.
func DidUtilCmd() *cobra.Command {
	// the answers will be written to this struct
	answers := struct {
		Goal          string // survey will match the question and field names
		FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
		Age           int    // if the types don't match, survey will convert it
	}{}

	// run the survey
	cmd := &cobra.Command{
		Use:   "did",
		Short: "Generate DID Documents and handle WebAuthn",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// run the survey
			selectOpt := &survey.Select{
				Message: "Choose an option:",
				Options: []string{"Create a new DID", "Update existing DID"},
				Default: "Create a new DID",
			}
			// perform the questions
			survey.AskOne(selectOpt, &answers.Goal, nil)

			if answers.Goal == "Create a new DID" {
				err = surveyNewDid()
				cobra.CheckErr(err)
			} else {
				err = surveyExistingDid()
				cobra.CheckErr(err)
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
