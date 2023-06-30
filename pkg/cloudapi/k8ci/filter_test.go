package k8ci

import "testing"

var k8ciItems = ListK8CI{
	Data: []ItemK8CI{
		{
			CreatedTime: 123902139,
			Status:      "ENABLED",
			Description: "",
			ID:          1,
			Name:        "purple_snake",
			Version:     "1",
			LBImageID:   654,
			NetworkPlugins: []string{
				"flannel",
				"calico",
				"weavenet",
			},
		},
		{
			CreatedTime: 123902232,
			Status:      "ENABLED",
			Description: "",
			ID:          2,
			Name:        "green_giant",
			Version:     "2",
			LBImageID:   654,
			NetworkPlugins: []string{
				"flannel",
				"calico",
				"weavenet",
			},
		},
		{
			CreatedTime: 123902335,
			Status:      "DISABLED",
			Description: "",
			ID:          3,
			Name:        "magenta_cloud",
			Version:     "3",
			NetworkPlugins: []string{
				"flannel",
				"calico",
				"weavenet",
			},
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := k8ciItems.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := k8ciItems.FilterByName("magenta_cloud").FindOne()

	if actual.Name != "magenta_cloud" {
		t.Fatal("expected Name 'magenta_cloud', found: ", actual.Name)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := k8ciItems.FilterFunc(func(ikc ItemK8CI) bool {
		return ikc.CreatedTime > 123902139
	})

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.CreatedTime < 123902139 {
			t.Fatal("expected CreatedTime greater than 123902139, found: ", item.CreatedTime)
		}
	}
}

func TestSortingByCreatedTime(t *testing.T) {
	actual := k8ciItems.SortByCreatedTime(true)

	if actual.Data[0].CreatedTime != 123902335 && actual.Data[2].CreatedTime != 123902139 {
		t.Fatal("expected inverse sort, found normal")
	}
}
