package cmd

import (
	"net/url"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ZoomData struct {
	host string
	confNo   string
	password string
}

func parseZoomURL(raw string) *ZoomData {
	parsed, err := url.Parse(raw)
	if err != nil {
		log.WithFields(log.Fields{ "url": raw }).WithError(err).Error("Error parsing zoom URL")
	}
	vals := parsed.Query()
	cn := strings.Replace(parsed.Path, "/j/", "", 1)
	return &ZoomData{
		confNo: cn,
		password: vals.Get("pwd"),
		host: parsed.Host,
	}
}

func constructZoomClientURL(data *ZoomData) string {
	return "zoommtg://" + data.host + "/join?confNo=" + data.confNo + "&pwd=" + data.password
}

func NewOpenCommand() *cobra.Command {
	openCmd := &cobra.Command{
		Use:   "open",
		Short: "open a zoom meeting",
		Run: func(cmd *cobra.Command, args []string) {
			key := GetArg(args, 0, defaultKey)
			link := viper.GetString(key)
			parsed := parseZoomURL(link)
			open := exec.Command("open", constructZoomClientURL(parsed))
			err := open.Run()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
		},
	}
	return openCmd
}
