package main

import (
	"github.com/google/go-github/v41/github"
)

type OrganizationService struct {
	*BaseService
}

type Organization struct {
	Name         *string      `json:"organization,omitempty" yaml:"organization,omitempty"`
	Repositories []Repository `json:"repositories,omitempty" yaml:"repositories,omitempty"`
}

func NewOrganizationService(p *Provisioner) *OrganizationService {
	return &OrganizationService{
		BaseService: &BaseService{
			provisioner: p,
		},
	}
}

func (o *OrganizationService) Get(orgName string) (*github.Organization, *github.Response, error) {
	return o.provisioner.Client.Organizations.Get(getContext(), orgName)
}
