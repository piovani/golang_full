package cmd

import "github.com/spf13/cobra"

type Cmd struct {
	cobra *cobra.Command
}

func NewCmd() *Cmd {
	return &Cmd{
		cobra: &cobra.Command{
			Use:     "go_full",
			Short:   "go full is complete aplication",
			Version: "1.0.0",
		},
	}
}
