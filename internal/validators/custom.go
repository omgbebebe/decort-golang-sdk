package validators

import "github.com/go-playground/validator/v10"

// accountCUTypeValidator is used to validate CUType field.
func accountCUTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, accountCUTypeValues)
}

// accountAccessTypeValidator is used to validate AccessType field.
func accountAccessTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, accountAccessTypeValues)
}

// bserviceDriverValidator is used to validate Driver field.
func bserviceDriverValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, bserviceDriverValues)
}

// bserviceModeValidator is used to validate Mode field.
func bserviceModeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, bserviceModeValues)
}

// computeTopologyValidator is used to validate Topology field.
func computeTopologyValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computeTopologyValues)
}

// computePolicyValidator is used to validate Policy field.
func computePolicyValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computePolicyValues)
}

// computeModeValidator is used to validate Mode field.
func computeModeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computeModeValues)
}

// computeDiskTypeValidator is used to validate DiskType field.
func computeDiskTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computeDiskTypeValues)
}

// computeNetTypeValidator is used to validate NetType field.
func computeNetTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computeNetTypeValues)
}

// computeProtoValidator is used to validate Proto field.
func computeProtoValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computeProtoValues)
}

// imageBootTypeValidator is used to validate BootType field.
func imageBootTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, imageBootTypeValues)
}

// imageTypeValidator is used to validate ImageType field.
func imageTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, imageTypeValues)
}

// imageDriversValidator is used to validate Drivers field.
func imageDriversValidator(fe validator.FieldLevel) bool {
	fieldSlice, ok := fe.Field().Interface().([]string)
	if !ok {
		return false
	}

	for _, item := range fieldSlice {
		if !StringInSlice(item, imageDriversValues) {
			return false
		}
	}

	return true
}

// imageArchitectureValidator is used to validate Architecture field.
func imageArchitectureValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, imageArchitectureValues)
}
