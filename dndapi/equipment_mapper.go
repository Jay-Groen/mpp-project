package dndapi

import "herkansing/onion/domain"

// ToDomainEquipment converts an APIEquipment object from the D&D API
// into a domain.EquipmentSpecific used by the business logic.
func ToDomainEquipment(apiEq APIEquipment) domain.EquipmentSpecific {
	return domain.EquipmentSpecific{
		Index:               apiEq.Index,
		Name:                apiEq.Name,
		EquipmentCategory:   toDomainAPIReference(apiEq.EquipmentCategory),
		Desc:                apiEq.Desc,
		Special:             apiEq.Special,
		WeaponCategory:      apiEq.WeaponCategory,
		WeaponRange:         apiEq.WeaponRange,
		CategoryRange:       apiEq.CategoryRange,
		ArmorCategory:       apiEq.ArmorCategory,
		ArmorClass:          toDomainArmorClass(apiEq.ArmorClass),
		StrMinimum:          apiEq.StrMinimum,
		StealthDisadvantage: apiEq.StealthDisadvantage,
		Cost:                toDomainCost(apiEq.Cost),
		Damage:              toDomainEquipmentDamage(apiEq.Damage),
		TwoHandedDamage:     toDomainEquipmentDamage(apiEq.TwoHandedDamage),
		Range:               toDomainEquipmentRange(apiEq.Range),
		ThrowRange:          toDomainThrowRange(apiEq.ThrowRange),
		Weight:              apiEq.Weight,
		Properties:          toDomainReferences(apiEq.Properties),
		Contents:            toDomainReferences(apiEq.Contents),
		VehicleCategory:     apiEq.VehicleCategory,
		Speed:               toDomainSpeed(apiEq.Speed),
		Capacity:            apiEq.Capacity,
	}
}

// --- Helper conversion functions ---

func toDomainAPIReference(ref APIReference) domain.APIReference {
	return domain.APIReference{
		Index: ref.Index,
		Name:  ref.Name,
		URL:   ref.URL,
	}
}

func toDomainArmorClass(ac *ArmorClass) *domain.ArmorClass {
	if ac == nil {
		return nil
	}
	return &domain.ArmorClass{
		Base:     ac.Base,
		DexBonus: ac.DexBonus,
		MaxBonus: ac.MaxBonus,
	}
}

func toDomainCost(cost *Cost) *domain.Cost {
	if cost == nil {
		return nil
	}
	return &domain.Cost{
		Quantity: cost.Quantity,
		Unit:     cost.Unit,
	}
}

func toDomainEquipmentDamage(dmg *EquipmentDamage) *domain.EquipmentDamage {
	if dmg == nil {
		return nil
	}
	return &domain.EquipmentDamage{
		DamageDice: dmg.DamageDice,
		DamageType: toDomainAPIReference(dmg.DamageType),
	}
}

func toDomainEquipmentRange(rng *EquipmentRange) *domain.EquipmentRange {
	if rng == nil {
		return nil
	}
	return &domain.EquipmentRange{
		Normal: rng.Normal,
		Long:   rng.Long,
	}
}

func toDomainThrowRange(rng *ThrowRange) *domain.ThrowRange {
	if rng == nil {
		return nil
	}
	return &domain.ThrowRange{
		Normal: rng.Normal,
		Long:   rng.Long,
	}
}

func toDomainSpeed(sp *Speed) *domain.Speed {
	if sp == nil {
		return nil
	}
	return &domain.Speed{
		Quantity: sp.Quantity,
		Unit:     sp.Unit,
	}
}

func toDomainReferences(refs []APIReference) []domain.APIReference {
	out := make([]domain.APIReference, len(refs))
	for i, r := range refs {
		out[i] = toDomainAPIReference(r)
	}
	return out
}
