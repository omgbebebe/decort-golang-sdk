package grid

import "testing"

var grids = ListGrids{
	Data: []ItemGridList{
		{
			Resources: Resources{
				Current: RecordResource{
					CPU:         84,
					DiskSize:    976,
					DiskSizeMax: 1200,
					ExtIPs:      132,
					ExtTraffic:  0,
					GPU:         79500,
					RAM:         0,
					SEPs:        map[string]map[string]DiskUsage{},
				},
				Reserved: RecordResource{
					CPU:         123,
					DiskSize:    976,
					DiskSizeMax: 1200,
					ExtIPs:      132,
					ExtTraffic:  0,
					GPU:         0,
					RAM:         152600,
					SEPs:        map[string]map[string]DiskUsage{},
				},
			},
			Flag:         "",
			GID:          212,
			GUID:         1,
			ID:           1,
			LocationCode: "alfa",
			Name:         "alfa",
		},
		{
			Resources: Resources{
				Current: RecordResource{
					CPU:         84,
					DiskSize:    976,
					DiskSizeMax: 1200,
					ExtIPs:      132,
					ExtTraffic:  0,
					GPU:         79500,
					RAM:         0,
					SEPs:        map[string]map[string]DiskUsage{},
				},
				Reserved: RecordResource{
					CPU:         123,
					DiskSize:    976,
					DiskSizeMax: 1200,
					ExtIPs:      132,
					ExtTraffic:  0,
					GPU:         0,
					RAM:         152600,
					SEPs:        map[string]map[string]DiskUsage{},
				},
			},
			Flag:         "",
			GID:          666,
			GUID:         2,
			ID:           2,
			LocationCode: "beta",
			Name:         "beta",
		},
		{
			Resources: Resources{
				Current: RecordResource{
					CPU:         84,
					DiskSize:    976,
					DiskSizeMax: 1200,
					ExtIPs:      132,
					ExtTraffic:  0,
					GPU:         79500,
					RAM:         0,
					SEPs:        map[string]map[string]DiskUsage{},
				},
				Reserved: RecordResource{
					CPU:         123,
					DiskSize:    976,
					DiskSizeMax: 1200,
					ExtIPs:      132,
					ExtTraffic:  0,
					GPU:         0,
					RAM:         152600,
					SEPs:        map[string]map[string]DiskUsage{},
				},
			},
			Flag:         "",
			GID:          777,
			GUID:         3,
			ID:           3,
			LocationCode: "gamma",
			Name:         "gamma",
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := grids.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByGID(t *testing.T) {
	actual := grids.FilterByGID(777).FindOne()

	if actual.GID != 777 {
		t.Fatal("expected ID 777, found: ", actual.GID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := grids.FilterByName("gamma").FindOne()

	if actual.Name != "gamma" {
		t.Fatal("expected Name 'gamma', found: ", actual.Name)
	}
}

func TestFilterByLocationCode(t *testing.T) {
	actual := grids.FilterByLocationCode("alfa").FindOne()

	if actual.LocationCode != "alfa" {
		t.Fatal("expected LocationCode 'alfa', found: ", actual.LocationCode)
	}
}

func TestFilterFunc(t *testing.T) {
	actual := grids.FilterFunc(func(rg ItemGridList) bool {
		return rg.GID == 777
	}).
		FindOne()

	if actual.GID != 777 {
		t.Fatal("expected GID 777, found: ", actual.GID)
	}
}
