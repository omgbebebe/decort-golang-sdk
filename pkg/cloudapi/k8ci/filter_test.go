package k8ci

import "testing"

var k8ciItems = ListK8CI{
	ItemK8CI{
		CreatedTime: 123902139,
		RecordK8CI: RecordK8CI{
			Description: "",
			ID:          1,
			Name:        "purple_snake",
			Version:     "1",
		},
	},
	ItemK8CI{
		CreatedTime: 123902232,
		RecordK8CI: RecordK8CI{
			Description: "",
			ID:          2,
			Name:        "green_giant",
			Version:     "2",
		},
	},
	ItemK8CI{
		CreatedTime: 123902335,
		RecordK8CI: RecordK8CI{
			Description: "",
			ID:          3,
			Name:        "magenta_cloud",
			Version:     "3",
		},
	},
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

	if len(actual) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual))
	}

	for _, item := range actual {
		if item.CreatedTime < 123902139 {
			t.Fatal("expected CreatedTime greater than 123902139, found: ", item.CreatedTime)
		}
	}
}

func TestSortingByCreatedTime(t *testing.T) {
	actual := k8ciItems.SortByCreatedTime(true)

	if actual[0].CreatedTime != 123902335 && actual[2].CreatedTime != 123902139 {
		t.Fatal("expected inverse sort, found normal")
	}
}
