package setup

import (
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

const bin = "/usr/bin"

type file interface {
	Name() string
	URL() string
	Callback() error
}

type alp struct {
	filename  string
	name, url string
}

type ptQueryDigest struct {
	name, url string
}

func alpNew() file {
	return &alp{
		name: "alp_linux_amd64.zip",
		url:  "https://github.com/tkuchiki/alp/releases/download/v0.3.1/alp_linux_amd64.zip",
	}
}
func (a *alp) Name() string { return a.name }
func (a *alp) URL() string  { return a.url }
func (a *alp) Callback() error {
	if err := archiver.Zip.Open(a.name, a.filename); err != nil {
		return err
	}
	dirname := a.filename
	if err := os.Rename(
		filepath.Join(dirname, a.filename),
		filepath.Join(bin, a.filename),
	); err != nil {
		return err
	}
	if err := os.Remove(a.name); err != nil {
		return err
	}
	if err := os.RemoveAll(a.filename); err != nil {
		return err
	}
	return nil
}

func ptQueryDigestNew() file {
	return &ptQueryDigest{
		name: "pt-query-digest",
		url:  "https://www.percona.com/get/pt-query-digest",
	}
}

func (p *ptQueryDigest) Name() string { return p.name }
func (p *ptQueryDigest) URL() string  { return p.url }
func (p *ptQueryDigest) Callback() error {
	if err := os.Rename(p.name, filepath.Join(bin, p.name)); err != nil {
		return err
	}
	return nil
}
