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
		Short: "reset your zoom cli configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if !yes {
				yString := getLine("Are you sure? This will remove your zoom CLI configuration (type Y to confirm)")
				if yString == "Y" {
					yes = true
				} else {
					return
				}
			}
			if yes {
				var err = os.Remove(configFilePath + "/" + configFileName)
				if err != nil {
					log.Error(err)
				} else {
					fmt.Println("zoom CLI configuration reset")
				}
			}
		},
	}

	resetCmd.Flags().BoolVarP(&yes, "yes", "y", false, "use this flag to skip confirmation")
	return resetCmd
}
