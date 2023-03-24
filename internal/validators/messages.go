package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func errorMessage(fe validator.FieldError) string {
	prefix := "validation-error:"

	switch fe.Tag() {

	// Common Validators
	case "required":
		return fmt.Sprintf("%s %s is required", prefix, fe.Field())
	case "gt":
		return fmt.Sprintf("%s %s can't be less or equal to zero", prefix, fe.Field())
	case "min":
		return fmt.Sprintf("%s %s: not enough elements", prefix, fe.Field())
	case "max":
		return fmt.Sprintf("%s %s: too many elements", prefix, fe.Field())
	case "url":
		return fmt.Sprintf("%s %s: unexpected URL format", prefix, fe.Field())
	case "email":
		return fmt.Sprintf("%s %s: unexpected E-Mail format", prefix, fe.Field())

	case "driver":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(driverValues))

	case "accessType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(accessTypeValues))

	case "resTypes":
		return fmt.Sprintf("%s %s can contain only the following values: %s",
			prefix,
			fe.Field(),
			joinValues(resTypesValues))

	case "proto":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(protoValues))

	// Account Validators
	case "accountCUType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(accountCUTypeValues))

	// BService Validators
	case "bserviceMode":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(bserviceModeValues))

	// Compute Validators
	case "computeTopology":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computeTopologyValues))

	case "computePolicy":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computePolicyValues))

	case "computeMode":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computeModeValues))

	case "computeDiskType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computeDiskTypeValues))

	case "computeNetType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computeNetTypeValues))

	case "computeOrder":
		return fmt.Sprintf("%s %s can contain only the following values: %s",
			prefix,
			fe.Field(),
			joinValues(computeOrderValues))

	case "computeDataDisks":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computeDataDisksValues))

	// Disk Validators
	case "diskType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(diskTypeValues))

	// Flipgroup Validators
	case "flipgroupClientType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(flipgroupClientTypeValues))

	// KVM_X86/KVM_PPC Validators
	case "kvmNetType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(kvmNetTypeValues))

	// LB Validators
	case "lbAlgorithm":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(lbAlgorithmValues))

	// RG Validators
	case "rgDefNet":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(rgDefNetValues))

	case "rgNetType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(rgNetTypeValues))

	// ViNS Validators
	case "vinsType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(vinsTypeValues))

	// Image Validators
	case "imageBootType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(imageBootTypeValues))

	case "imageType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(imageTypeValues))

	case "imageDrivers":
		return fmt.Sprintf("%s %s must contain only the following: %s",
			prefix,
			fe.Field(),
			joinValues(imageDriversValues))

	case "imageArchitecture":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(imageArchitectureValues))

	// SEP Validators
	case "sepFieldType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(sepFieldTypeValues))

	}

	return fe.Error()
}

func joinValues(values []string) string {
	return strings.Join(values, ", ")
}
