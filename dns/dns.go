package clouddns

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateManagedZone(ctx *pulumi.Context) error {
	zone, err := dns.NewManagedZone(ctx, "skofel", &dns.ManagedZoneArgs{
		Description: pulumi.String("A DNS zone that will be used by my services. Lol"),
		DnsName:     pulumi.String("skofel.com."),
		Name:        pulumi.String("skofel"),
	})
	if err != nil {
		return err
	}

	ctx.Export("managedZoneID", zone.ManagedZoneId)
	return nil
}
