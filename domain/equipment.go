package domain

import "herkansing/onion/dndapi"

type Equipment struct {
	MainHand Weapon `json:"main_hand"`
	OffHand  Weapon `json:"off_hand"`
	Armor    Armor  `json:"armor"`
	Shield   Shield `json:"shield"`
	Gear     Gear   `json:"gear"`
}

type Weapon struct {
	APIEquipment dndapi.APIEquipment `json:"api_equipment"`
}

type Armor struct {
	APIEquipment dndapi.APIEquipment `json:"api_equipment"`
}

type Shield struct {
	APIEquipment dndapi.APIEquipment `json:"api_equipment"`
}

type Gear struct {
	APIEquipment dndapi.APIEquipment `json:"api_equipment"`
}
