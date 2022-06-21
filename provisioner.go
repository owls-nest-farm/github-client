package main

import (
	"sync"

	"github.com/google/go-github/v41/github"
)

type BaseService struct {
	provisioner *Provisioner
}

type Provisioner struct {
	Client  *github.Client
	Configs []Organization

	Organizations *OrganizationService
	Repositories  *RepositoryService
}

func NewProvisioner(configs []Organization) *Provisioner {
	p := &Provisioner{
		Client:  getClient(),
		Configs: configs,
	}

	p.Organizations = NewOrganizationService(p)
	p.Repositories = NewRepositoryService(p)

	return p
}

func (p *Provisioner) ProcessConfig(organization Organization, destroy bool) {
	var wg sync.WaitGroup
	wg.Add(len(organization.Repositories))

	// TODO: Ensure org name was given.
	org, _, err := p.Organizations.Get(*organization.Name)
	if err != nil {
		panic(err)
	}

	for _, repository := range organization.Repositories {
		rc := RepositoryCtx{
			Organization: org,
			Repository:   repository,
		}
		go func(rc RepositoryCtx) {
			if !destroy {
				p.Repositories.Create(&rc)
			} else {
				p.Repositories.Delete(&rc)
			}
			wg.Done()
		}(rc)
	}

	wg.Wait()
}

func (p *Provisioner) ProcessConfigs(destroy bool) {
	for _, organization := range p.Configs {
		p.ProcessConfig(organization, destroy)
	}
}
