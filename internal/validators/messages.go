package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func errorMessage(fe validator.FieldError) string {
	prefix := "validation-error:"

	switch fe.Tag() {

	// Default Validators
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

	// Account Validators
	case "accountAccessType":
		return fmt.Sprintf("%s %s must be one of the followng: %s",
			prefix,
			fe.Field(),
			joinValues(accountAccessTypeValues))

	case "accountCUType":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(accountCUTypeValues))

	// BService Validators
	case "bserviceDriver":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(bserviceDriverValues))

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

	case "computeProto":
		return fmt.Sprintf("%s %s must be one of the following: %s",
			prefix,
			fe.Field(),
			joinValues(computeProtoValues))

	// Image Validators
	case "bootType":
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

	}

	return fe.Error()
}

func joinValues(values []string) string {
	return strings.Join(values, ", ")
}
