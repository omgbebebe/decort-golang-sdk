package account

import (
	"testing"
)

var accounts = ListAccounts{
	Data: []ItemAccount{
		{
			ACL: []RecordACL{
				{
					IsExplicit: true,
					GUID:       "",
					Rights:     "CXDRAU",
					Status:     "CONFIRMED",
					Type:       "U",
					UgroupID:   "timofey_tkachev_1@decs3o",
				},
			},
			CreatedTime: 1676645275,
			DeletedTime: 0,
			ID:          132846,
			Name:        "std",
			Status:      "CONFIRMED",
			UpdatedTime: 1676645275,
		},
		{
			ACL: []RecordACL{
				{
					IsExplicit: true,
					GUID:       "",
					Rights:     "CXDRAU",
					Status:     "CONFIRMED",
					Type:       "U",
					UgroupID:   "not_really_timofey_tkachev_1@decs3o",
				},
			},
			CreatedTime: 1676878820,
			DeletedTime: 0,
			ID:          132847,
			Name:        "std_2",
			Status:      "CONFIRMED",
			UpdatedTime: 1676645275,
		},
		{
			ACL: []RecordACL{
				{
					IsExplicit: true,
					GUID:       "",
					Rights:     "CXDRAU",
					Status:     "CONFIRMED",
					Type:       "U",
					UgroupID:   "timofey_tkachev_1@decs3o",
				},
				{
					IsExplicit: true,
					GUID:       "",
					Rights:     "CXDRAU",
					Status:     "CONFIRMED",
					Type:       "U",
					UgroupID:   "second_account@decs3o",
				},
			},
			CreatedTime: 1676883850,
			DeletedTime: 1676883899,
			ID:          132848,
			Name:        "std_broker",
			Status:      "DELETED",
			UpdatedTime: 1676878820,
		},
	},
	EntryCount: 3,
}

func TestFilterByID(t *testing.T) {
	actual := accounts.FilterByID(132846).FindOne()

	if actual.ID != 132846 {
		t.Fatal("actual: ", actual.ID, " > expected: 132846")
	}
}

func TestFilterByUserGroupId(t *testing.T) {
	actual := accounts.FilterByUserGroupID("second_account@decs3o").FindOne()

	for _, item := range actual.ACL {
		if item.UgroupID == "second_account@decs3o" {
			return
		}
	}

	t.Fatal("second_account@decs3o has not been found. expected 1 found")
}

func TestFilterByName(t *testing.T) {
	actual := accounts.FilterByName("std_broker").FindOne()

	if actual.Name != "std_broker" {
		t.Fatal("actual: ", actual.Name, " >> expected: std_broker")
	}
}

func TestFilterByStatus(t *testing.T) {
	actual := accounts.FilterByStatus("CONFIRMED")

	if len(actual.Data) != 2 {
		t.Fatal("Expected 2 elements in slice, found: ", len(actual.Data))
	}

	for _, item := range actual.Data {
		if item.Status != "CONFIRMED" {
			t.Fatal("expected CONFIRMED, found: ", item.Status)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	actual := accounts.FilterFunc(func(ia ItemAccount) bool {
		return ia.DeletedTime == 0
	})

	for _, item := range actual.Data {
		if item.DeletedTime != 0 {
			t.Fatal("Expected DeletedTime = 0, found: ", item.DeletedTime)
		}
	}
}

func TestSortingByCreatedTime(t *testing.T) {
	actual := accounts.SortByCreatedTime(false)

	if actual.Data[0].Name != "std" {
		t.Fatal("Expected account std as earliest, found: ", actual.Data[0].Name)
	}

	actual = accounts.SortByCreatedTime(true)

	if actual.Data[0].Name != "std_broker" {
		t.Fatal("Expected account std_broker as latest, found: ", actual.Data[0].Name)
	}
}

func TestFilterEmpty(t *testing.T) {
	actual := accounts.FilterByID(0)

	if len(actual.Data) != 0 {
		t.Fatal("Expected 0 found, actual: ", len(actual.Data))
	}
}
