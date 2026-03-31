package gamestate

type Fleet struct {
	Name          LocalizedName `clausewitz:"name"`
	Ships         []int         `clausewitz:"ships"`
	MilitaryPower float64       `clausewitz:"military_power"`
	ShipClass     string        `clausewitz:"ship_class"`
	Station       bool          `clausewitz:"station"`
	Civilian      bool          `clausewitz:"civilian"`
}

type Ship struct {
	Fleet                    int            `clausewitz:"fleet"`
	Name                     LocalizedName  `clausewitz:"name"`
	ShipDesignImplementation ShipDesignImpl `clausewitz:"ship_design_implementation"`
	Section                  ShipSection    `clausewitz:"section"`
	Hitpoints                float64        `clausewitz:"hitpoints"`
	ShieldHitpoints          float64        `clausewitz:"shield_hitpoints"`
	ArmorHitpoints           float64        `clausewitz:"armor_hitpoints"`
	MaxHitpoints             float64        `clausewitz:"max_hitpoints"`
	MaxShieldHitpoints       float64        `clausewitz:"max_shield_hitpoints"`
	MaxArmorHitpoints        float64        `clausewitz:"max_armor_hitpoints"`
}

type ShipDesignImpl struct {
	Design int `clausewitz:"design"`
}

type ShipSection struct {
	Design string       `clausewitz:"design"`
	Slot   string       `clausewitz:"slot"`
	Weapon []ShipWeapon `clausewitz:"weapon"`
}

type ShipWeapon struct {
	Template      string `clausewitz:"template"`
	ComponentSlot string `clausewitz:"component_slot"`
}

type ShipDesign struct {
	Name         LocalizedName     `clausewitz:"name"`
	GrowthStages []ShipGrowthStage `clausewitz:"growth_stages"`
}

type ShipGrowthStage struct {
	ShipSize          string             `clausewitz:"ship_size"`
	Section           DesignSection      `clausewitz:"section"`
	RequiredComponent []string           `clausewitz:"required_component"`
}

type DesignSection struct {
	Template  string            `clausewitz:"template"`
	Slot      string            `clausewitz:"slot"`
	Component []DesignComponent `clausewitz:"component"`
}

type DesignComponent struct {
	Slot     string `clausewitz:"slot"`
	Template string `clausewitz:"template"`
}
