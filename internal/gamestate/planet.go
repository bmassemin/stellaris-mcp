package gamestate

type Planet struct {
	Name             LocalizedName      `clausewitz:"name"`
	PlanetClass      string             `clausewitz:"planet_class"`
	PlanetSize       int                `clausewitz:"planet_size"`
	Owner            int                `clausewitz:"owner"`
	Controller       int                `clausewitz:"controller"`
	Stability        float64            `clausewitz:"stability"`
	Crime            float64            `clausewitz:"crime"`
	Amenities        float64            `clausewitz:"amenities"`
	AmenitiesUsage   float64            `clausewitz:"amenities_usage"`
	FreeAmenities    float64            `clausewitz:"free_amenities"`
	TotalHousing     float64            `clausewitz:"total_housing"`
	HousingUsage     float64            `clausewitz:"housing_usage"`
	FreeHousing      float64            `clausewitz:"free_housing"`
	NumPops          int                `clausewitz:"num_sapient_pops"`
	FinalDesignation string             `clausewitz:"final_designation"`
	AscensionTier    int                `clausewitz:"ascension_tier"`
	Districts        []int              `clausewitz:"districts"`
	BuildingsCache   []int              `clausewitz:"buildings_cache"`
	SurveyedBy       int                `clausewitz:"surveyed_by"`
	Deposits         []int              `clausewitz:"deposits"`
	TimedModifier    TimedModifierList  `clausewitz:"timed_modifier"`
	Produces         map[string]float64 `clausewitz:"produces"`
	Upkeep           map[string]float64 `clausewitz:"upkeep"`
	Profits          map[string]float64 `clausewitz:"profits"`
}

type TimedModifierList struct {
	Items []TimedModifier `clausewitz:"items"`
}

type TimedModifier struct {
	Modifier string `clausewitz:"modifier"`
	Days     int    `clausewitz:"days"`
}
