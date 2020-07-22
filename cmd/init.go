package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewInitCommand() *cobra.Command {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize default Zoom link",
		Run: func(cmd *cobra.Command, args []string) {
			var l string
			if len(args) > 0 {
				l = args[0]
			} else {
				l = GetLine("Enter default meeting link:")
			} 
			viper.Set("default", l)

			if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
				os.Mkdir(configFilePath, 0700)
			}

			_, err := os.OpenFile(configFilePath+"/"+configFileName, os.O_RDONLY|os.O_CREATE, 0700)
			if err != nil {
				log.WithFields(log.Fields{ "path": configFilePath, "name": configFileName }).WithError(err).Errorf("Error creating config file")
			}

			err = viper.WriteConfig()
			if err != nil {
				log.WithFields(log.Fields{ "path": configFilePath, "name": configFileName }).WithError(err).Errorf("Error saving config file")
			}
		},
	}
	return initCmd
}
