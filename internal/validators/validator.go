package validators

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once
	instance *DecortValidator
)

type DecortValidator struct {
	decortValidator *validator.Validate
}

// getDecortValidator returns singleton instance of DecortValidator.
func getDecortValidator() *validator.Validate {
	if instance == nil {
		once.Do(func() {
			instance = new(DecortValidator)
			instance.decortValidator = validator.New()

			err := registerAllValidators(instance.decortValidator)
			if err != nil {
				panic(err)
			}
		})
	}

	return instance.decortValidator
}

// registerAllValidators registers all custom validators in DecortValidator.
func registerAllValidators(validate *validator.Validate) error {
	err := validate.RegisterValidation("proto", protoValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("accessType", accessTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("resTypes", resTypesValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("driver", driverValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("imageBootType", imageBootTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("imageType", imageTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("imageDrivers", imageDriversValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("imageArchitecture", imageArchitectureValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("accountCUType", accountCUTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("bserviceMode", bserviceModeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computeTopology", computeTopologyValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computePolicy", computePolicyValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computeMode", computeModeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computeDiskType", computeDiskTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computeNetType", computeNetTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computeOrder", computeOrderValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("computeDataDisks", computeDataDisksValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("diskType", diskTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("flipgroupClientType", flipgroupClientTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("kvmNetType", kvmNetTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("lbAlgorithm", lbAlgorithmValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("rgDefNet", rgDefNetValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("rgNetType", rgNetTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("vinsType", vinsTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("sepFieldType", sepFieldTypeValidator)
	if err != nil {
		return err
	}

	err = validate.RegisterValidation("hwPath", hwPathValidator)
	if err != nil {
		return err
	}

	return nil
}
