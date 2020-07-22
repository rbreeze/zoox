package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewResetCommand() *cobra.Command {
	var yes bool

	var resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "reset zoox configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if !yes {
				yString := GetLine("Are you sure? (type Y to confirm):")
				if yString == "Y" {
					yes = true
				} else {
					fmt.Println("Reset cancelled.")
					return
				}
			}
			if yes {
				var err = os.Remove(configFilePath + "/" + configFileName)
				if err != nil {
					log.Error(err)
				} else {
					fmt.Println("Configuration reset.")
				}
			}
		},
	}

	resetCmd.Flags().BoolVarP(&yes, "yes", "y", false, "use this flag to skip confirmation")
	return resetCmd
}
