package main

import (
	bucket "quickstart/buckets"
	computeengine "quickstart/computeEngine"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP resource (Storage Bucket) - Details are hardcoded for now
		err := bucket.CreateStorageBucket(ctx)
		if err != nil {
			return err
		}

		err = computeengine.CreateComputeEngineInstance(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
