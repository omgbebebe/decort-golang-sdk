package computeci

import "testing"

var computeciItems = ListComputeCI{
	{
		CustomFields: map[string]interface{}{},
		Description:  "",
		Drivers: []string{
			"KVM_X86",
		},
		GUID:     1,
		ID:       1,
		Name:     "computeci_1",
		Status:   "ENABLED",
		Template: "",
	},
	{
		CustomFields: map[string]interface{}{},
		Description:  "",
		Drivers: []string{
			"KVM_X86",
		},
		GUID:     2,
		ID:       2,
		Name:     "computeci_2",
		Status:   "ENABLED",
		Template: "",
	},
	{
		CustomFields: map[string]interface{}{},
		Description:  "",
		Drivers: []string{
			"SVA_KVM_X86",
		},
		GUID:     3,
		ID:       3,
		Name:     "computeci_3",
		Status:   "DISABLED",
		Template: "",
	},
}

func TestFilterByID(t *testing.T) {
	actual := computeciItems.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := computeciItems.FilterByName("computeci_3").FindOne()

	if actual.Name != "computeci_3" {
		t.Fatal("expected Name 'computeci_2', found: ", actual.Name)
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := computeciItems.FilterByStatus("ENABLED")

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.Status != "ENABLED" {
			t.Fatal("expected Status 'ENABLED', found: ", item.Status)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := computeciItems.FilterFunc(func(icc ItemComputeCI) bool {
		for _, item := range icc.Drivers {
			if item == "KVM_X86" {
				return true
			}
		}
		return false
	})

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		for _, driver := range item.Drivers {
			if driver != "KVM_X86" {
				t.Fatal("expected 'KVM_X86' Driver, found: ", driver)
			}
		}
	}
}
