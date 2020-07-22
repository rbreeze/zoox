package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewAddCommand() *cobra.Command {
	var name string
	var link string

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a zoom meeting link",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				name = GetLine("Enter a meeting name:")
			}
			if link == "" {
				link = GetLine("Enter meeting link:")
			}

			viper.Set(name, link)
			clipboard.WriteAll(link)

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

    addCmd.Flags().StringVarP(&name, "name", "n", "", "meeting name") 
    addCmd.Flags().StringVarP(&link, "link", "l", "", "meeting link") 
	return addCmd
}
