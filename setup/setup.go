package setup

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/Code-Hex/chair/setup/generate"
	"github.com/pkg/errors"
)

const (
	indent             = "    "
	perm   os.FileMode = 0755
)

func CommandNew() *cobra.Command {
	return &cobra.Command{
		Use:   "setup",
		Short: "Run initial setup for isucon environment",
		Long: `Run initial setup for isucon environment.

What will install? => alp v0.3.1, pt-query-digest
What will generate script? => "restart.sh" into current directory
Where will install path? => ` + bin,
		Example: indent + "sudo chair setup",
		RunE:    run,
	}
}

func run(cmd *cobra.Command, args []string) error {
	if os.Geteuid() > 0 {
		return errors.New("If you want to initialize, you should run as a superuser")
	}
	if err := generate.RestartScript(); err != nil {
		return errors.Wrap(err, "Failed to create restart script")
	}

	files := []file{
		alpNew(),
		ptQueryDigestNew(),
	}

	var wg sync.WaitGroup
	for _, v := range files {
		fmt.Println("Start:", v.URL())
		wg.Add(1)
		go func(fi file) {
			defer wg.Done()
			if err := download(fi); err != nil {
				log.Println(err)
			}
			fmt.Println("Done:", fi.URL())
			if err := fi.Callback(); err != nil {
				log.Println(err)
			}
		}(v)
	}
	wg.Wait()
	return nil
}

func download(file file) error {
	req, err := http.NewRequest("GET", file.URL(), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to make request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "Failed to download: %s", file.URL())
	}
	defer resp.Body.Close()

	f, err := os.Create(file.Name())
	if err != nil {
		return errors.Wrapf(err, "Failed to create %s", file.Name())
	}
	io.Copy(f, resp.Body)
	f.Chmod(perm)
	f.Close()

	return nil
}
