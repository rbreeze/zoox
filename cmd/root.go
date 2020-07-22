package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    configFilePath = os.Getenv("HOME") + "/.zoox"
    configFileType = "yaml"
    configFileName = "config.yaml"
    defaultKey     = "default"
)

type ErrorMessage string

const (
    LinkNotSet ErrorMessage = "Link not set"
)

func GetLine(s string) string {
    fmt.Printf("%s ", s)
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Error(err)
    }
    return strings.Replace(input, "\n", "", -1)
}

func GetArg(args []string, n int, defaultVal string) string {
    if len(args) > n {
        return args[n]
    } else {
        return defaultVal
    }
}

func NewCommand() *cobra.Command {
    var command = &cobra.Command{
        Use:   "zoox",
        Short: "some useful commands for using Zoom",
        Run: func(cmd *cobra.Command, args []string) {
            name := GetArg(args, 0, defaultKey)
            link := viper.GetString(name)
            if link != "" {
                fmt.Println(link)
            } else {
                log.Error(LinkNotSet)
            }
        },
    }

    command.AddCommand(NewInitCommand())
    command.AddCommand(NewAddCommand())
    command.AddCommand(NewResetCommand())
    command.AddCommand(NewCopyCommand())
    command.AddCommand(NewGetCommand())
    command.AddCommand(NewOpenCommand())

    command.PersistentPreRun = func(cmd *cobra.Command, args []string) {
        viper.AddConfigPath(configFilePath)
        viper.SetConfigName(configFileName)
        viper.SetConfigType(configFileType)
        if err := viper.ReadInConfig(); err != nil {
            if _, ok := err.(viper.ConfigFileNotFoundError); ok {
                if cmd.Name() == "init" {
                    return
                } else if cmd.Name() == "reset" {
                    log.Warn("Configuration not reset because it doesn't exist.")
                } else {
                    log.Warn("No links found. Run zoox init to fix this.")
                }
                os.Exit(0)
            }
        }
    }

    return command
}

var rootCmd = &cobra.Command{
    Use:   "zoox",
    Short: "some useful commands for using Zoom",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
