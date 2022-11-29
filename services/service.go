package service

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func EnableAPI(ctx *pulumi.Context, projectID, service string, disableDependentServices bool) (*projects.Service, error) {
	newService, err := projects.NewService(ctx, "project", &projects.ServiceArgs{
		DisableDependentServices: pulumi.Bool(disableDependentServices),
		Project:                  pulumi.String(projectID),
		Service:                  pulumi.String(service),
	})
	if err != nil {
		return nil, err
	}

	return newService, nil
}

func EnableService(ctx *pulumi.Context) (*projects.Service, error) {
	availableServices := []string{"dns.googleapis.com"}
	var newService *projects.Service
	for _, service := range availableServices {
		newService, _ = EnableAPI(ctx, "intricate-idiom-364006", service, true)
		// if err != nil {
		// 	return nil, err
		// }
	}
	return newService, nil
}
