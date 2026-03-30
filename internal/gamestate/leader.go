package gamestate

type Leader struct {
	Name            LeaderName      `clausewitz:"name"`
	Species         int             `clausewitz:"species"`
	Gender          string          `clausewitz:"gender"`
	Country         int             `clausewitz:"country"`
	Class           string          `clausewitz:"class"`
	Tier            string          `clausewitz:"tier"`
	Experience      float64         `clausewitz:"experience"`
	Age             int             `clausewitz:"age"`
	Level           int             `clausewitz:"level"`
	BonusSkillLevel int             `clausewitz:"bonus_skill_level"`
	Job             string          `clausewitz:"job"`
	Ethic           string          `clausewitz:"ethic"`
	Planet          int             `clausewitz:"planet"`
	Traits          []string        `clausewitz:"traits"`
	Location        LeaderLocation  `clausewitz:"location"`
	CouncilLocation LeaderLocation  `clausewitz:"council_location"`
	RecruitmentDate string          `clausewitz:"recruitment_date"`
}

type LeaderName struct {
	FullNames LocalizedName `clausewitz:"full_names"`
}

type LeaderLocation struct {
	Type       string `clausewitz:"type"`
	Assignment string `clausewitz:"assignment"`
	ID         int    `clausewitz:"id"`
	Position   int    `clausewitz:"position"`
}
