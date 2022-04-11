package cmd

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v3"

	"github.com/lspserver/proxy/config"
)

var (
	app        = kingpin.New("proxy", "lspserver proxy").Version(config.Version + "-build-" + config.Build)
	configFile = app.Flag("config-file", "Config file (.yml)").Required().String()
	listenPort = app.Flag("listen-address", "Listening to address").Default("127.0.0.1:49093").String()
)

func Run() error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	_, err := initConfig(*configFile)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	return nil
}

func initConfig(name string) (*config.Config, error) {
	c := config.New()

	fi, err := os.Open(name)
	if err != nil {
		return c, errors.Wrap(err, "failed to open")
	}

	defer func() {
		_ = fi.Close()
	}()

	buf, _ := io.ReadAll(fi)

	if err := yaml.Unmarshal(buf, c); err != nil {
		return c, errors.Wrap(err, "failed to unmarshal")
	}

	return c, nil
}
