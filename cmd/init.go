package cmd

import (
	"os"

	clipboard "github.com/atotto/clipboard"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewInitCommand() *cobra.Command {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize your zoom config",
		Run: func(cmd *cobra.Command, args []string) {
			myLink := getLine("Enter your personal Zoom meeting link")
			viper.Set("myLink", myLink)
			clipboard.WriteAll(myLink)

			if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
				os.Mkdir(configFilePath, 0700)
			}

			_, err := os.OpenFile(configFilePath+"/"+configFileName, os.O_RDONLY|os.O_CREATE, 0700)
			if err != nil {
				log.Errorf("Error creating config file: %s", err)
			}

			err = viper.WriteConfig()
			if err != nil {
				log.Warnf("Error saving config: %s", err)
			}
		},
	}
	return initCmd
}
