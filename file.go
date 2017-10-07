package main

type file interface {
	Name() string
	URL() string
}

type alp struct {
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

func ptQueryDigestNew() file {
	return &ptQueryDigest{
		name: "pt-query-digest",
		url:  "https://www.percona.com/get/pt-query-digest",
	}
}

func (p *ptQueryDigest) Name() string { return p.name }
func (p *ptQueryDigest) URL() string  { return p.url }
