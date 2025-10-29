package domain

type Equipment struct {
	MainHand Weapon `json:"main_hand"`
	OffHand  Weapon `json:"off_hand"`
	Armor    Armor  `json:"armor"`
	Shield   Shield `json:"shield"`
	Gear     Gear   `json:"gear"`
}

type Weapon struct {
	EquipmentSpecific EquipmentSpecific `json:"api_equipment"`
}

type Armor struct {
	EquipmentSpecific EquipmentSpecific `json:"api_equipment"`
}

type Shield struct {
	EquipmentSpecific EquipmentSpecific `json:"api_equipment"`
}

type Gear struct {
	EquipmentSpecific EquipmentSpecific `json:"api_equipment"`
}

// APIEquipment represents any equipment item from the D&D 5e API.
type EquipmentSpecific struct {
	Index               string           `json:"index,omitempty"`
	Name                string           `json:"name"`
	EquipmentCategory   APIReference     `json:"equipment_category"`
	Desc                []string         `json:"desc"`
	Special             []string         `json:"special"`
	WeaponCategory      string           `json:"weapon_category,omitempty"`
	WeaponRange         string           `json:"weapon_range,omitempty"`
	CategoryRange       string           `json:"category_range,omitempty"`
	ArmorCategory       string           `json:"armor_category,omitempty"`
	ArmorClass          *ArmorClass      `json:"armor_class,omitempty"`
	StrMinimum          int              `json:"str_minimum,omitempty"`
	StealthDisadvantage bool             `json:"stealth_disadvantage,omitempty"`
	Cost                *Cost            `json:"cost,omitempty"`
	Damage              *EquipmentDamage `json:"damage,omitempty"`
	TwoHandedDamage     *EquipmentDamage `json:"2h_damage,omitempty"`
	Range               *EquipmentRange  `json:"range,omitempty"`
	ThrowRange          *ThrowRange      `json:"throw_range,omitempty"`
	Weight              float64          `json:"weight,omitempty"`
	Properties          []APIReference   `json:"properties,omitempty"`
	Contents            []APIReference   `json:"contents,omitempty"`
	VehicleCategory     string           `json:"vehicle_category,omitempty"`
	Speed               *Speed           `json:"speed,omitempty"`
	Capacity            string           `json:"capacity,omitempty"`
}

type APIReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Cost struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type EquipmentDamage struct {
	DamageDice string       `json:"damage_dice"`
	DamageType APIReference `json:"damage_type"`
}

type EquipmentRange struct {
	Normal int  `json:"normal"`
	Long   *int `json:"long,omitempty"`
}

type ThrowRange struct {
	Normal int `json:"normal"`
	Long   int `json:"long"`
}

type ArmorClass struct {
	Base     int  `json:"base"`
	DexBonus bool `json:"dex_bonus"`
	MaxBonus *int `json:"max_bonus,omitempty"`
}

type Speed struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
