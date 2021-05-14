package helm

import (
	"fmt"

	"github.com/prateekgogia/kit/cli/pkg/util/logging"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/storage/driver"
)

var log = logging.NewNamedLogger("Helm")

type Client struct {
	Configuration *action.Configuration
	Settings      *cli.EnvSettings
}

func NewClient() (*Client, error) {
	settings := cli.New()
	configuration := &action.Configuration{}
	if err := configuration.Init(settings.RESTClientGetter(), settings.Namespace(), "", log.Debugf); err != nil {
		return nil, fmt.Errorf("initializing helm, %w", err)
	}
	return &Client{
		Settings:      settings,
		Configuration: configuration,
	}, nil
}

type ApplyOptions struct {
	ChartName   string
	ReleaseName string
	Namespace   string
	Values      map[string]interface{}
}

// Apply a helm chart idempotently
func (h *Client) Apply(options ApplyOptions) error {
	history := action.NewHistory(h.Configuration)
	history.Max = 1
	releases, err := history.Run(options.ReleaseName);
	if  err == driver.ErrReleaseNotFound {
		log.Infof("Creating helm release %s", options.ReleaseName)
		if err := h.create(options); err != nil {
			return fmt.Errorf("failed to create helm release %s, %w", options.ReleaseName, err)
		}
		return nil
	}

	log.Infof("Updating helm release %s", options.ReleaseName)
	upgrade := action.NewUpgrade(h.Configuration)
	upgrade.Namespace = options.Namespace
	upgrade.Atomic = true

	if _, err := upgrade.Run(options.ReleaseName, releases[0].Chart, options.Values); err != nil {
		return fmt.Errorf("failed to upgrade helm release %s, %w", options.ReleaseName, err)
	}
	return nil
}

func (h *Client) create(options ApplyOptions) error {
	install := action.NewInstall(h.Configuration)
	install.ReleaseName = options.ReleaseName
	install.Namespace = options.Namespace
	install.CreateNamespace = true
	install.IsUpgrade = true
	install.Atomic = true
	install.IncludeCRDs = true
	path, err := install.LocateChart(options.ChartName, h.Settings)
	if err != nil {
		return fmt.Errorf("locating helm chart %s, %w", options.ChartName, err)
	}
	chart, err := loader.Load(path)
	if err != nil {
		return fmt.Errorf("loading helm chart %s, %w", options.ChartName, err)
	}
	if _, err := install.Run(chart, options.Values); err != nil {
		return fmt.Errorf("installing helm release %s, %w", options.ReleaseName, err)
	}
	return nil
}
