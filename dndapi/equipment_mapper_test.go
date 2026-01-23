package dndapi

import (
    "herkansing/onion/domain"
    "reflect"
    "testing"
)

func TestToDomainEquipment_MapsCoreFields(t *testing.T) {
    api := APIEquipment{
        Index: "leather-armor",
        Name:  "Leather Armor",
        EquipmentCategory: APIReference{
            Index: "armor",
            Name:  "Armor",
            URL:   "/api/equipment-categories/armor",
        },
        Desc: []string{"A simple suit of leather armor."},
        Special: []string{"No special rules."},
        WeaponCategory: "simple",
        WeaponRange:    "melee",
        CategoryRange:  "martial",
        ArmorCategory:  "Light",
        StrMinimum:     0,
        StealthDisadvantage: false,
        Weight: 10,
    }

    got := ToDomainEquipment(api)

    if got.Index != api.Index || got.Name != api.Name {
        t.Fatalf("Index/Name not mapped: got (%q,%q), want (%q,%q)", got.Index, got.Name, api.Index, api.Name)
    }

    wantCat := domain.APIReference{
        Index: api.EquipmentCategory.Index,
        Name:  api.EquipmentCategory.Name,
        URL:   api.EquipmentCategory.URL,
    }
    if !reflect.DeepEqual(got.EquipmentCategory, wantCat) {
        t.Fatalf("EquipmentCategory = %#v, want %#v", got.EquipmentCategory, wantCat)
    }

    if !reflect.DeepEqual(got.Desc, api.Desc) {
        t.Fatalf("Desc = %#v, want %#v", got.Desc, api.Desc)
    }
    if !reflect.DeepEqual(got.Special, api.Special) {
        t.Fatalf("Special = %#v, want %#v", got.Special, api.Special)
    }
}
