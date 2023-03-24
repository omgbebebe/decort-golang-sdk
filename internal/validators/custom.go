package validators

import "github.com/go-playground/validator/v10"

// protoValidator is used to validate Proto fields.
func protoValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, protoValues)
}

// accessTypeValidator is used to validate AccessType fields.
func accessTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, accessTypeValues)
}

// resTypesValidator is used to validate ResTypes fields.
func resTypesValidator(fe validator.FieldLevel) bool {
	fieldSlice, ok := fe.Field().Interface().([]string)
	if !ok {
		return false
	}

	for _, value := range fieldSlice {
		if !StringInSlice(value, resTypesValues) {
			return false
		}
	}

	return true
}

// driverValidator is used to validate Driver fields.
func driverValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, driverValues)
}

// accountCUTypeValidator is used to validate CUType field.
func accountCUTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, accountCUTypeValues)
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

// computeOrderValidator is used to validate Order field.
func computeOrderValidator(fe validator.FieldLevel) bool {
	fieldSlice, ok := fe.Field().Interface().([]string)
	if !ok {
		return false
	}

	for _, value := range fieldSlice {
		if !StringInSlice(value, computeOrderValues) {
			return false
		}
	}

	return true
}

// computeDataDisksValidator is used to validate DataDisks field.
func computeDataDisksValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, computeDataDisksValues)
}

// diskTypeValidator is used to validate Type field.
func diskTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, diskTypeValues)
}

// flipgroupClientTypeValidator is used to validate ClientType field.
func flipgroupClientTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, flipgroupClientTypeValues)
}

// kvmNetTypeValidator is used to validate NetType field.
func kvmNetTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, kvmNetTypeValues)
}

// lbAlgorithmValidator is used to validate Algorithm field.
func lbAlgorithmValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, lbAlgorithmValues)
}

// rgDefNetValidator is used to validate DefNet field.
func rgDefNetValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, rgDefNetValues)
}

// rgNetTypeValidator is used to validate NetType field.
func rgNetTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, rgNetTypeValues)
}

// vinsTypeValidator is used to validate Type field.
func vinsTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, vinsTypeValues)
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

// sepFieldTypeValidator is used to validate FieldType field.
func sepFieldTypeValidator(fe validator.FieldLevel) bool {
	fieldValue := fe.Field().String()

	return StringInSlice(fieldValue, sepFieldTypeValues)
}
