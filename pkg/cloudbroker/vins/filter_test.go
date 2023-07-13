package vins

import "testing"

var vinsItems = ListVINS{
	Data: []ItemVINS{
		{
			AccountID:   1,
			AccountName: "std",
			CreatedBy:   "sample_user_1@decs3o",
			CreatedTime: 1676898844,
			DefaultGW:   "",
			DefaultQOS: QOS{
				ERate:   0,
				GUID:    "",
				InBurst: 0,
				InRate:  0,
			},
			DeletedBy:          "",
			DeletedTime:        0,
			Description:        "",
			ExternalIP:         "",
			GID:                212,
			GUID:               1,
			ID:                 1,
			LockStatus:         "UNLOCKED",
			ManagerID:          0,
			ManagerType:        "",
			Milestones:         363485,
			Name:               "vins01",
			NetMask:            24,
			Network:            "192.168.1.0/24",
			PreReservationsNum: 32,
			PriVNFDevID:        29557,
			Redundant:          false,
			RGID:               7971,
			RGName:             "rg_01",
			SecVNFDevID:        0,
			Status:             "ENABLED",
			UpdatedBy:          "",
			UpdatedTime:        0,
			UserManaged:        true,
			VNFs: ItemVNFs{
				DHCP: 51997,
				DNS:  0,
				FW:   0,
				GW:   0,
				NAT:  0,
				VPN:  0,
			},
			VXLANID: 3544,
		},
		{
			AccountID:   2,
			AccountName: "std2",
			CreatedBy:   "sample_user_1@decs3o",
			CreatedTime: 1676898948,
			DefaultGW:   "",
			DefaultQOS: QOS{
				ERate:   0,
				GUID:    "",
				InBurst: 0,
				InRate:  0,
			},
			DeletedBy:          "",
			DeletedTime:        0,
			Description:        "",
			ExternalIP:         "",
			GID:                212,
			GUID:               2,
			ID:                 2,
			LockStatus:         "LOCKED",
			ManagerID:          0,
			ManagerType:        "",
			Milestones:         363508,
			Name:               "vins02",
			NetMask:            24,
			Network:            "192.168.2.0/24",
			PreReservationsNum: 32,
			PriVNFDevID:        29558,
			Redundant:          false,
			RGID:               7972,
			RGName:             "rg_02",
			SecVNFDevID:        0,
			Status:             "ENABLED",
			UpdatedBy:          "",
			UpdatedTime:        0,
			UserManaged:        true,
			VNFs: ItemVNFs{
				DHCP: 51998,
				DNS:  0,
				FW:   0,
				GW:   0,
				NAT:  0,
				VPN:  0,
			},
			VXLANID: 3545,
		},
		{
			AccountID:   3,
			AccountName: "std3",
			CreatedBy:   "sample_user_2@decs3o",
			CreatedTime: 1676899026,
			DefaultGW:   "",
			DefaultQOS: QOS{
				ERate:   0,
				GUID:    "",
				InBurst: 0,
				InRate:  0,
			},
			DeletedBy:          "",
			DeletedTime:        0,
			Description:        "",
			ExternalIP:         "",
			GID:                212,
			GUID:               3,
			ID:                 3,
			LockStatus:         "UNLOCKED",
			ManagerID:          0,
			ManagerType:        "",
			Milestones:         363549,
			Name:               "vins03",
			NetMask:            24,
			Network:            "192.168.3.0/24",
			PreReservationsNum: 32,
			PriVNFDevID:        29559,
			Redundant:          false,
			RGID:               7973,
			RGName:             "rg_03",
			SecVNFDevID:        0,
			Status:             "DISABLED",
			UpdatedBy:          "",
			UpdatedTime:        0,
			UserManaged:        true,
			VNFs: ItemVNFs{
				DHCP: 51999,
				DNS:  0,
				FW:   0,
				GW:   0,
				NAT:  0,
				VPN:  0,
			},
			VXLANID: 3546,
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := vinsItems.FilterByID(2).FindOne()

	if actual.ID != 2 {
		t.Fatal("expected ID 2, found: ", actual.ID)
	}
}

func TestFilterByName(t *testing.T) {
	actual := vinsItems.FilterByName("vins01").FindOne()

	if actual.Name != "vins01" {
		t.Fatal("expected Name 'vins01', found: ", actual.Name)
	}
}

func TestFilterByAccountID(t *testing.T) {
	actual := vinsItems.FilterByAccountID(3).FindOne()

	if actual.AccountID != 3 {
		t.Fatal("expected AccountID 3, found: ", actual.AccountID)
	}
}

func TestFilterByCreatedBy(t *testing.T) {
	actual := vinsItems.FilterByCreatedBy("sample_user_1@decs3o")

	if len(actual.Data) != 2 {
		t.Fatal("expected 2 found, actual: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.CreatedBy != "sample_user_1@decs3o" {
			t.Fatal("expected CreatedBy 'sample_user_1@decs3o', found: ", item.CreatedBy)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := vinsItems.FilterFunc(func(iv ItemVINS) bool {
		return iv.RGID == 7971
	}).
		FindOne()

	if actual.RGID != 7971 {
		t.Fatal("expected RGID 7971, found: ", actual.RGID)
	}
}

func TestSortByCreatedTime(t *testing.T) {
	actual := vinsItems.SortByCreatedTime(false)

	if actual.Data[0].CreatedTime != 1676898844 || actual.Data[2].CreatedTime != 1676899026 {
		t.Fatal("expected ascending order, found descending")
	}
}
