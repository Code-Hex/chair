package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/Code-Hex/chair/gen"
	"github.com/pkg/errors"
)

func runInit() error {
	if os.Geteuid() > 0 {
		return errors.New("If you want to initialize, you should run as a superuser")
	}
	restartScript := gen.GenerateRestartScript()
	f, err := os.Create("restart.sh")
	if err != nil {
		return errors.Wrap(err, "Failed to create restart.sh")
	}
	f.WriteString(restartScript)
	f.Chmod(perm)
	f.Close()

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
