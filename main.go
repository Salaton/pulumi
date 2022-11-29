package main

import (
	bucket "quickstart/buckets"
	computeengine "quickstart/computeEngine"
	clouddns "quickstart/dns"
	service "quickstart/services"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// First, enable any services
		_, err := service.EnableService(ctx)
		if err != nil {
			return err
		}

		// Create a GCP resource (Storage Bucket) - Details are hardcoded for now
		err = bucket.CreateStorageBucket(ctx)
		if err != nil {
			return err
		}

		err = computeengine.CreateComputeEngineInstance(ctx)
		if err != nil {
			return err
		}

		err = clouddns.CreateManagedZone(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
