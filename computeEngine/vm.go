package computeengine

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateComputeEngineInstance(ctx *pulumi.Context) error {
	// defaultAccount, err := serviceAccount.NewAccount(ctx, "defaultAccount", &serviceAccount.AccountArgs{
	// 	AccountId:   pulumi.String("service_account_id"),
	// 	DisplayName: pulumi.String("Service Account"),
	// })
	// if err != nil {
	// 	return err
	// }

	static, err := compute.NewAddress(ctx, "rocket-chat-static-ip", &compute.AddressArgs{
		Description: pulumi.String("A static IP for the rocket chat vm"),
		Name:        pulumi.String("rocket-chat-static-ip"),
		Region:      pulumi.String("europe-west1"),
	})
	if err != nil {
		return err
	}

	_, err = compute.NewInstance(ctx, "rocket-chat", &compute.InstanceArgs{
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
		DeletionProtection:    pulumi.Bool(false),
		Description:           pulumi.String("Just a test instance deployed using pulumi"),
		MachineType:           pulumi.String("n1-standard-1"),
		Metadata:              nil,
		MetadataStartupScript: nil,
		MinCpuPlatform:        nil,
		Name:                  pulumi.String("rocket-chat"),
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
		// ServiceAccount: &compute.InstanceServiceAccountArgs{
		// 	Email: defaultAccount.Email,
		// 	Scopes: pulumi.StringArray{
		// 		pulumi.String("cloud-platform"),
		// 	},
		// },
		ShieldedInstanceConfig: nil,
		Tags: pulumi.StringArray{
			pulumi.String("dev"),
		},
		Zone: pulumi.String("europe-west1-d"),
	})
	if err != nil {
		return err
	}
	return nil
}
