package browser

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/KeisukeYamashita/biko/providers"
	"github.com/urfave/cli"
)

// Open wraps the browser opener
func Open(c *cli.Context, provider providers.Provider) error {
	if err := provider.Init(c); err != nil {
		return err
	}
	url, err := provider.GetTargetURL()
	if err != nil {
		return err
	}
	fmt.Print(url)
	return openbrowser(url)
}

func openbrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return err
	}

	return nil
}
