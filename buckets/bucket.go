package bucket

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateStorageBucket(ctx *pulumi.Context) error {
	// Create a GCP resource (Storage Bucket)
	bucket, err := storage.NewBucket(ctx, "my-bucket", &storage.BucketArgs{
		Location: pulumi.String("US"),
	})
	if err != nil {
		return err
	}

	bucketObject, err := CreateBucketObject(ctx, bucket, "index.html")
	if err != nil {
		return err
	}
	bucketEndpoint := pulumi.Sprintf("http://storage.googleapis.com/%s/%s", bucket.Name, bucketObject.Name)

	// Export the DNS name of the bucket
	ctx.Export("bucketName", bucket.Url)
	ctx.Export("bucketEndpoint", bucketEndpoint)
	return nil
}

func CreateBucketObject(ctx *pulumi.Context, bucket *storage.Bucket, filename string) (*storage.BucketObject, error) {
	bucketObject, err := storage.NewBucketObject(ctx, filename, &storage.BucketObjectArgs{
		Bucket: bucket.Name,
		Source: pulumi.NewFileAsset(filename),
	})
	if err != nil {
		return nil, err
	}

	return bucketObject, nil
}
