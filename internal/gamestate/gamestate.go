package gamestate

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmassemin/stellaris-mcp/internal/clausewitz"
)

type GameState struct {
	Version    string              `clausewitz:"version"`
	Name       string              `clausewitz:"name"`
	Date       string              `clausewitz:"date"`
	Player     []Player            `clausewitz:"player"`
	Country    map[int]Country     `clausewitz:"country"`
	Fleet      map[int]Fleet       `clausewitz:"fleet"`
	Ships      map[int]Ship        `clausewitz:"ships"`
	ShipDesign map[int]ShipDesign  `clausewitz:"ship_design"`
	Planets    PlanetsWrapper      `clausewitz:"planets"`
	Buildings  map[int]Building    `clausewitz:"buildings"`
	Districts  map[int]District    `clausewitz:"districts"`
	Zones      map[int]Zone        `clausewitz:"zones"`
	Deposit    map[int]Deposit     `clausewitz:"deposit"`
	SpyNetworks map[int]SpyNetwork `clausewitz:"spy_networks"`
	PopJobs    map[int]PopJob     `clausewitz:"pop_jobs"`
	Leaders    map[int]Leader      `clausewitz:"leaders"`
	War        []War               `clausewitz:"war"`
	Federation []Federation        `clausewitz:"federation"`
}

type Building struct {
	Type     string `clausewitz:"type"`
	Position int    `clausewitz:"position"`
}

type SpyNetwork struct {
	Owner            int     `clausewitz:"owner"`
	Target           int     `clausewitz:"target"`
	Leader           int     `clausewitz:"leader"`
	Power            float64 `clausewitz:"power"`
	ActiveOperations []int   `clausewitz:"active_operations"`
	Formed           string  `clausewitz:"formed"`
}

type PopJob struct {
	Type      string `clausewitz:"type"`
	Planet    int    `clausewitz:"planet"`
	Workforce float64 `clausewitz:"workforce"`
}

type Deposit struct {
	Type     string `clausewitz:"type"`
	SwapType string `clausewitz:"swap_type"`
}

type District struct {
	Type  string `clausewitz:"type"`
	Level int    `clausewitz:"level"`
	Zones []int  `clausewitz:"zones"`
}

type Zone struct {
	Type      string `clausewitz:"type"`
	Buildings []int  `clausewitz:"buildings"`
}

type Player struct {
	Name    string `clausewitz:"name"`
	Country int    `clausewitz:"country"`
}

type PlanetsWrapper struct {
	Planet map[int]Planet `clausewitz:"planet"`
}

type War struct {
	Name       LocalizedName `clausewitz:"name"`
	Attackers  []WarParty    `clausewitz:"attackers"`
	Defenders  []WarParty    `clausewitz:"defenders"`
	StartDate  string        `clausewitz:"start_date"`
	EndDate    string        `clausewitz:"end_date"`
}

type WarParty struct {
	Country int `clausewitz:"country"`
}

type Federation struct {
	Name                  LocalizedName         `clausewitz:"name"`
	Leader                int                   `clausewitz:"leader"`
	Members               []int                 `clausewitz:"members"`
	FederationProgression FederationProgression `clausewitz:"federation_progression"`
}

type FederationProgression struct {
	FederationType string  `clausewitz:"federation_type"`
	Experience     float64 `clausewitz:"experience"`
	Cohesion       float64 `clausewitz:"cohesion"`
}

// LoadFromDir finds the most recent .sav file in dir,
// extracts the "gamestate" file from the ZIP, and parses it.
func LoadFromDir(dir string) (*GameState, error) {
	savPath, err := findLatestSave(dir)
	if err != nil {
		return nil, err
	}
	return LoadFromZip(savPath)
}

// LoadFromZip extracts "gamestate" from a .sav ZIP and parses it.
func LoadFromZip(path string) (*GameState, error) {
	data, err := extractGamestate(path)
	if err != nil {
		return nil, err
	}
	return Load(data)
}

// Load parses raw gamestate bytes into a GameState.
func Load(data []byte) (*GameState, error) {
	var gs GameState
	if err := clausewitz.Unmarshal(data, &gs); err != nil {
		return nil, fmt.Errorf("parse gamestate: %w", err)
	}
	return &gs, nil
}

func findLatestSave(dir string) (string, error) {
	var newest string
	var newestTime int64

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // skip unreadable dirs
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".sav") {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		if t := info.ModTime().UnixNano(); t > newestTime {
			newestTime = t
			newest = path
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("walk save directory %q: %w", dir, err)
	}

	if newest == "" {
		return "", fmt.Errorf("no .sav files found in %q", dir)
	}
	return newest, nil
}

func extractGamestate(zipPath string) ([]byte, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("open zip %q: %w", zipPath, err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == "gamestate" {
			rc, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("open gamestate in zip: %w", err)
			}
			defer rc.Close()
			return io.ReadAll(rc)
		}
	}
	return nil, fmt.Errorf("no 'gamestate' file found in %q", zipPath)
}
