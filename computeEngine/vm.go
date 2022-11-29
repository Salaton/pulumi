package computeengine

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dns"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateComputeEngineInstance(ctx *pulumi.Context) error {
	static, err := compute.NewAddress(ctx, "casdoor-static-ip", &compute.AddressArgs{
		Description: pulumi.String("A static IP for the rocket chat vm"),
		Name:        pulumi.String("casdoor-static-ip"),
		Region:      pulumi.String("europe-west1"),
	})
	if err != nil {
		return err
	}

	casdoorInstance, err := compute.NewInstance(ctx, "casdoor", &compute.InstanceArgs{
		AllowStoppingForUpdate: pulumi.Bool(true),
		BootDisk: &compute.InstanceBootDiskArgs{
			AutoDelete:              pulumi.Bool(true),
			DeviceName:              nil,
			DiskEncryptionKeyRaw:    nil,
			DiskEncryptionKeySha256: nil,
			InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
				Image:  pulumi.String("ubuntu-os-cloud/ubuntu-2004-lts"),
				Labels: nil,
				Size:   pulumi.IntPtr(20),
			},
			KmsKeySelfLink: nil,
			Mode:           nil,
			Source:         nil,
		},
		DeletionProtection: pulumi.Bool(false),
		Description:        pulumi.String("Just a test instance deployed using pulumi"),
		MachineType:        pulumi.String("n1-standard-1"),
		Metadata:           nil,
		MetadataStartupScript: pulumi.String(`
		#! /bin/bash
		sudo ufw allow 22
		sudo ufw allow 80/tcp
		sudo ufw allow 8080/tcp
		sudo ufw allow 443/tcp`,
		),
		MinCpuPlatform: nil,
		Name:           pulumi.String("casdoor"),
		NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
			&compute.InstanceNetworkInterfaceArgs{
				Network: pulumi.String("default"),
				AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
					&compute.InstanceNetworkInterfaceAccessConfigArgs{
						NatIp: static.Address,
					},
				},
			},
		},
		NetworkPerformanceConfig: nil,
		Project:                  nil,
		ReservationAffinity:      nil,
		ResourcePolicies:         nil,
		Scheduling:               nil,
		ScratchDisks: compute.InstanceScratchDiskArray{
			&compute.InstanceScratchDiskArgs{
				Interface: pulumi.String("SCSI"),
			},
		},
		ShieldedInstanceConfig: nil,
		Tags: pulumi.StringArray{
			pulumi.String("dev"),
			pulumi.String("https-server"),
			pulumi.String("http-server"),
		},
		Zone: pulumi.String("europe-west1-d"),
	})
	if err != nil {
		return err
	}

	// Create A record set
	managedZone, err := dns.LookupManagedZone(ctx, &dns.LookupManagedZoneArgs{Name: "skofel"})
	if err != nil {
		return err
	}

	_, err = dns.NewRecordSet(ctx, "recordSet", &dns.RecordSetArgs{
		Name:        pulumi.Sprintf("casdoor-testing.%v", managedZone.DnsName),
		ManagedZone: pulumi.String(managedZone.Name),
		Type:        pulumi.String("A"),
		Ttl:         pulumi.Int(300),
		Rrdatas: pulumi.StringArray{
			casdoorInstance.NetworkInterfaces.ApplyT(func(networkInterfaces []compute.InstanceNetworkInterface) (string, error) {
				return *networkInterfaces[0].AccessConfigs[0].NatIp, nil
			}).(pulumi.StringOutput),
		},
	}, pulumi.DependsOn([]pulumi.Resource{static}))
	// TODO: should also depend on creation of the DNS managed zone
	if err != nil {
		return err
	}

	return nil
}
