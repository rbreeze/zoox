package cmd

import (
	"bufio"
	"fmt"
	"os"

	clipboard "github.com/atotto/clipboard"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getLine(s string) string {
	fmt.Printf("%s: ", s)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Info("Something went wrong getting input from stdin")
	}
	return input
}

func NewInitCommand() *cobra.Command {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize your zoom config",
		Run: func(cmd *cobra.Command, args []string) {
			myLink := getLine("Enter your personal Zoom meeting link")
			viper.Set("myLink", myLink)
			clipboard.WriteAll(myLink)

			viper.AddConfigPath(configFilePath)
			viper.SetConfigName(configFileName)
			viper.SetConfigType(configFileType)

			err := viper.WriteConfig()
			if err != nil {
				log.Warnf("Error saving config: %s", err)
			}
		},
	}
	return initCmd
}
