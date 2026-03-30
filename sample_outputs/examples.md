# Sample Outputs

Generated from a real Stellaris save (Cetus v4.3.2, early game 2200.07.01).

## ping

```
pong
```

## get_empire_overview

```
=== Empire Overview (2200.07.01) ===
Game: mp_Bebakian League (version Cetus v4.3.2)

Name: SPEC_Bebaki
Government: Science Directorate (Oligarchic)
Origin: Default
Ethics: Authoritarian, Fanatic Materialist
Civics: Technocracy, Crafters

Power Ratings:
  Military: 176.6
  Economy:  388.1
  Tech:     252.5
  Victory Rank: 7

Empire Stats:
  Pops: 5216
  Planets: 1
  Fleet Size: 15
  Empire Size: 56

Monthly Net Balance:
  alloys: 11.0, consumer_goods: 16.9, energy: 74.4, engineering_research: 18.7, food: 35.4, influence: 4.6, minerals: 22.9, physics_research: 21.1, society_research: 18.7, trade: 22.2, unity: 15.6
```

## get_research_status

```
=== Research Status (SPEC_Bebaki) ===

Current Research:
  Physics: Administrative Ai (progress: 151.3)
  Society: Planetary Unification (progress: 134.6)
  Engineering: Powered Exoskeletons (progress: 134.6)

Available Research Options:
  Physics: Physics 1, Shields 2, Administrative Ai, Power Plant 2
  Society: Genome Mapping, Doctrine Navy Size 1, Eco Simulation, Society 1, Planetary Unification
  Engineering: Ship Armor 2, Powered Exoskeletons, Afterburners 1, Torpedoes 1

Completed Technologies (31):
  - Solar Panel Network
  - Space Exploration
  - Corvettes
  - Starbase 1
  - Starbase 2
  - Assault Armies
  - Ship Armor 1
  - Thrusters 1
  - Space Defense Station 1
  - Basic Industry
  - Mechanized Mining
  - Space Construction
  - Mass Drivers 1
  - Flak Batteries 1
  - Missiles 1
  - Basic Science Lab 1
  - Fission Power
  - Reactor Boosters 1
  - Shields 1
  - Power Plant 1
  - Hyper Drive 1
  - Lasers 1
  - Pd Tracking 1
  - Planetary Defenses
  - Interplanetary Commerce
  - Industrial Farming
  - Hydroponics
  - Colonization 1
  - Basic Health
  - Planetary Government
  - Holo Entertainment
```

## get_economy_breakdown

```
=== Economy Breakdown (SPEC_Bebaki) ===

--- INCOME ---
  country_base: alloys: 5.0, consumer_goods: 15.0, energy: 20.0, engineering_research: 10.0, food: 20.0, influence: 3.0, minerals: 20.0, physics_research: 10.0, society_research: 10.0, unity: 10.0
  country_ethic: influence: 0.5
  country_power_projection: influence: 1.1
  orbital_mining_deposits: energy: 10.0, minerals: 13.0
  orbital_research_deposits: physics_research: 3.0
  planet_artisans: consumer_goods: 30.0, trade: 9.4
  planet_biologists: society_research: 5.8
  planet_bureaucrats: unity: 4.8
  planet_civilians: trade: 26.1
  planet_engineers: engineering_research: 6.0
  planet_farmers: food: 67.6
  planet_metallurgists: alloys: 7.3
  planet_miners: minerals: 22.5
  planet_physicists: physics_research: 5.4
  planet_politicians: engineering_research: 2.7, physics_research: 2.7, society_research: 2.9, unity: 10.8
  planet_technician: energy: 50.7
  pop_category_civilians: trade: 1.9
  pop_category_rulers: trade: 1.0
  pop_category_specialists: trade: 4.2
  pop_category_workers: trade: 1.8
  starbase_modules: energy: 6.0
  trade_policy: energy: 22.2, trade: 22.2
  TOTAL: alloys: 12.3, consumer_goods: 45.0, energy: 108.9, engineering_research: 18.7, food: 87.6, influence: 4.6, minerals: 55.5, physics_research: 21.1, society_research: 18.7, trade: 66.7, unity: 25.6

--- EXPENSES ---
  leader_commanders: unity: 2.0
  leader_officials: unity: 2.0
  leader_scientists: unity: 6.0
  planet_artisans: minerals: 22.9
  planet_biologists: consumer_goods: 2.2
  planet_buildings: energy: 8.8
  planet_bureaucrats: consumer_goods: 2.6
  planet_districts_cities: energy: 4.8
  planet_districts_farming: energy: 3.2
  planet_districts_generator: energy: 2.4
  planet_districts_mining: energy: 1.6
  planet_engineers: consumer_goods: 2.4
  planet_metallurgists: minerals: 9.8
  planet_physicists: consumer_goods: 2.2
  planet_politicians: consumer_goods: 6.5
  planet_pops: food: 52.2
  pop_category_civilians: consumer_goods: 1.9
  pop_category_rulers: consumer_goods: 2.0
  pop_category_specialists: consumer_goods: 6.4
  pop_category_workers: consumer_goods: 1.8
  ship_components: alloys: 0.2, energy: 1.3
  ships: alloys: 1.1, energy: 4.7
  starbase_buildings: energy: 0.9
  starbase_modules: energy: 0.9
  starbases: energy: 1.9
  station_gatherers: energy: 3.0
  station_researchers: energy: 1.0
  trade_policy: trade: 44.4
  TOTAL: alloys: 1.3, consumer_goods: 28.1, energy: 34.5, food: 52.2, minerals: 32.7, trade: 44.4, unity: 10.0

--- NET BALANCE ---
  alloys: +11.0
  consumer_goods: +16.9
  energy: +74.4
  engineering_research: +18.7
  food: +35.4
  influence: +4.6
  minerals: +22.9
  physics_research: +21.1
  society_research: +18.7
  trade: +22.2
  unity: +15.6
```

## get_neighbors

```
=== Known Empires (SPEC_Bebaki) ===
Your Military Power: 176.6

No known empires yet.
```

## get_fleet_power (summary)

```
=== Fleet Power (SPEC_Bebaki) ===
Total Military Power: 176.6
Fleet Size: 15
Use get_fleet_power with fleet_id for ship details.

ID       Name                           Class                        Power Ships
--------------------------------------------------------------------------------
0        SPEC_Bebaki_system             shipclass_starbase (station)    770.3     1
1        HUM1_SHIP_UntabbtheGrim        shipclass_science_ship (civilian)      0.0     1
2        HUM1_SHIP_NehdabantheHoly      shipclass_constructor (civilian)      0.0     1
3        HUM1_FLEET_PereksArmada        shipclass_military           176.6     3
204      SPEC_Bebaki_system             shipclass_mining_station (station)      0.0     1
205      HUM1_PLANET_KojoggsKeep        shipclass_mining_station (station)      0.0     1
206      HUM1_PLANET_KarbasLanding      shipclass_mining_station (station)      0.0     1
207      HUM1_PLANET_UltraksPoint       shipclass_mining_station (station)      0.0     1
208      HUM1_PLANET_HadriggasBastion   shipclass_research_station (station)      0.0     1
243      HUM1_PLANET_FindirbansOutpost  shipclass_mining_station (station)      0.0     1
16777286 HUM1_SHIP_YhlattheSavior       shipclass_science_ship (civilian)      0.0     1

Military Fleet Power (excl. stations/civilian): 176.6

Comparison with Known Empires:
  No known empires.
```

## get_fleet_power (detail: fleet_id=3)

```
=== Fleet 3: HUM1_FLEET_PereksArmada ===
Class: shipclass_military
Military Power: 176.6
Ships: 3

  --- Ship 3: HUM1_SHIP_BoroktheAffable ---
  Design: HUM1_CLASS_Jesslur (corvette)
  Hull:   200 / 200
  Shield: 200 / 200
  Armor:  100 / 100
  Section: CORVETTE_MID_S3
  Weapons:
    - SMALL_MASS_DRIVER_1 [SMALL_GUN_01]
    - SMALL_MASS_DRIVER_1 [SMALL_GUN_02]
    - MISSILE_1 [SMALL_GUN_03]

  --- Ship 4: HUM1_SHIP_RussutheSour ---
  Design: HUM1_CLASS_Jesslur (corvette)
  Hull:   200 / 200
  Shield: 200 / 200
  Armor:  100 / 100
  Section: CORVETTE_MID_S3
  Weapons:
    - SMALL_MASS_DRIVER_1 [SMALL_GUN_01]
    - SMALL_MASS_DRIVER_1 [SMALL_GUN_02]
    - MISSILE_1 [SMALL_GUN_03]

  --- Ship 5: HUM1_SHIP_LurtheUnpredictable ---
  Design: HUM1_CLASS_Jesslur (corvette)
  Hull:   200 / 200
  Shield: 200 / 200
  Armor:  100 / 100
  Section: CORVETTE_MID_S3
  Weapons:
    - SMALL_MASS_DRIVER_1 [SMALL_GUN_01]
    - SMALL_MASS_DRIVER_1 [SMALL_GUN_02]
    - MISSILE_1 [SMALL_GUN_03]
```

## get_planets (summary)

```
=== Planets (SPEC_Bebaki) — 1 planets ===
Use get_planets with planet_id for details.

ID     Name                      Class            Size  Pops  Stab Design.      Districts
--------------------------------------------------------------------------------------------------------------
11     SPEC_Bebaki_planet        Continental        19  5219    71 col_capital  4/19 [City:1 Farming:1 Generator:1 Mining:1]
```

## get_planets (detail: planet_id=11)

```
=== Planet 11: SPEC_Bebaki_planet ===
Class: Continental, Size: 19
Designation: Capital
Owner: 0, Controller: 0

Population: 5219 pops
Stability: 70.7
Crime: 0.0
Amenities: 4656 (used: 3780, free: 876)
Housing: 5600 (used: 5216, free: 384)

District Slots: 4 / 19 used
  Generator: 1
  Mining: 1
  Farming: 1
  City: 1

Districts Detail (4):
  City (lvl 3) — slots: Research Unity, Industrial
    [Default] Capital
    [Default] Precinct House
    [Research Unity] Research Lab 1
    [Industrial] Factory 1
  Generator (lvl 3)
  Mining (lvl 2)
  Farming (lvl 4)

Planetary Features (10):
  - Hot Springs
  - Rushing Waterfalls
  - Tempestous Mountain
  - Veiny Cliffs
  - Prosperous Mesa
  - Rich Mountain
  - Rugged Woods
  - Fertile Lands
  - Rugged Woods
  - Black Soil

Blockers (3):
  - Decrepit Dwellings
  - Failing Infrastructure -> Prosperous Mesa
  - Failing Infrastructure -> Prosperous Mesa

Modifiers:
  - Prosp Uni Mod (7020 days remaining)

Monthly Production:
  alloys: 7.3, consumer_goods: 30.0, energy: 50.7, engineering_research: 8.7, food: 67.6, minerals: 22.5, physics_research: 8.1, society_research: 8.7, trade: 44.4, unity: 15.6
Monthly Upkeep:
  consumer_goods: 28.1, energy: 20.8, food: 52.2, minerals: 32.7
Monthly Profit:
  alloys: 7.3, consumer_goods: 1.9, energy: 29.9, engineering_research: 8.7, food: 15.4, minerals: -10.1, physics_research: 8.1, society_research: 8.7, trade: 44.4, unity: 15.6
```

## get_planets (available=true)

```
=== Available Habitable Planets (SPEC_Bebaki) — 57 planets ===
Use get_planets with planet_id for details.

ID     Name                      Class              Size Surveyed Deposits
-------------------------------------------------------------------------------------
3224   Camthrin II               Continental          25      yes 12
2107   Ispyria III               Desert               25      yes 14
269    Uldor III                 Savannah             24      yes 10
294    Istora III                Desert               23      yes 14
3537   Rathadore III             Desert               23      yes 15
3895   Hocry II                  Gaia                 23      yes 13
287    Perqim III                Alpine               23      yes 14
3063   Cowhig III                Tundra               22      yes 14
2986   Zastea III                Arid                 22      yes 12
2287   Tystra III                Tropical             22      yes 16
778    Dossel I                  Desert               21      yes 15
863    NAME_Unique_System_2_Pla… Relic                20      yes 5
261    Bailleulus III            Continental          20      yes 13
2370   Areysak II                Continental          20      yes 14
1776   Kodracca III              Arid                 20      yes 12
580    Saidainope III            Alpine               20      yes 14
1211   Imoleto I                 Continental          20      yes 13
3357   Irthius II                Tropical             20      yes 15
1024   Vrittaka II               Tropical             19      yes 13
1829   NAME_wenkwort_prime       Gaia                 19      yes 9
3408   Ruinam II                 Tropical             18      yes 13
277    Taggallion III            Desert               18      yes 11
2952   Owhziea II                Ocean                18      yes 10
2980   Atmir II                  Savannah             18      yes 12
3586   Ophele III                Desert               18      yes 12
152    Jhurope III               Arid                 18      yes 13
396    Emara III                 Arid                 18      yes 13
228    Haedus III                Desert               18      yes 11
1213   Imoleto III               Gaia                 17      yes 9
1805   Toiubos III               Arid                 17      yes 11
3274   Uproth II b               Arctic               17      yes 15
318    Alassia III               Desert               16      yes 8
213    Dodonnam III              Desert               16      yes 12
328    Theta_Scorpii II          Alpine               16      yes 14
578    Saidainope II             Ocean                16      yes 13
1031   Unur II                   Tropical             15      yes 13
251    Lazon I                   Desert               14      yes 9
774    Ay'arolee III c           Alpine               14      yes 13
2063   Uxfriri III               Arctic               14      yes 11
3620   Betria III                Savannah             14      yes 11
3736   Avishek III a             Arctic               14      yes 13
1257   NAME_HillosC              Tundra               13      yes 8
3774   Sjoberg III a             Arid                 13      yes 13
775    Ay'arolee III d           Desert               13      yes 11
244    Fomalhaut I               Tropical             13      yes 10
1156   Aytoun III a              Tropical             13      yes 11
189    Hydrobius II              Continental          13      yes 11
2660   Seb III                   Tropical             13      yes 12
1212   Imoleto II                Continental          12      yes 11
3782   Sjoberg V a               Alpine               12      yes 11
1806   Toiubos III a             Alpine               12      yes 8
200    Terzam II                 Savannah             12      yes 11
1855   Fidhilam III a            Continental          11      yes 12
2709   Offe'ei III a             Arctic               11      yes 11
1944   Bazzanac IV a             Arid                 10      yes 7
558    NAME_UbogleeltD b         Gaia                  5      yes 5
221    Iolam III                 Alpine               16       no 13
```

## get_leaders (summary)

```
=== Leaders (SPEC_Bebaki) — 16 leaders ===
Use get_leaders with leader_id for full detail.

ID           Name                      Class        Lvl Age Assignment           Traits
--------------------------------------------------------------------------------------------------------------
119          HUM1_CHR_Jonnara          commander      1   0 politician           Aggressive
121          HUM1_CHR_Darmull          commander      1   0 bureaucrat           Adaptable
805306373    HUM1_CHR_Falatir          commander      1  38 Council #16777236    Fleet Organizer
120          HUM1_CHR_Jonnara          commander      1   0 farmer               Trickster
754974722    HUM1_CHR_Bemalona         envoy          1  34 bureaucrat           -
822083587    HUM1_CHR_Shibbala         envoy          1  33 politician           -
116          HUM1_CHR_Haghonona        official       1   0 politician           Adaptable
118          HUM1_CHR_Campramara       official       1   0 miner                Resilient
126          HUM1_CHR_Karba            official       1  32 Council #16777238    Politician
117          HUM1_CHR_Corrona          official       1   0 engineer             Ruler Eye For Talent
122          HUM1_CHR_Jesslur          scientist      1   0 foundry              Expertise Industry, Spark Of Genius
822083588    HUM1_CHR_Thaloth          scientist      1  36 Fleet: HUM1_SHIP_Un… Expertise Biology, Archaeologist
520093696    HUM1_CHR_Thalotha         scientist      1  28 technician           Expertise Particles, Spark Of Genius
124          HUM1_CHR_Kashnak          scientist      1   0 technician           Expertise Propulsion, Politician, Army Veteran, Destructive
123          HUM1_CHR_Haghonon         scientist      1   0 foundry              Expertise New Worlds, Carefree
125          HUM1_CHR_Jesslur          scientist      1  40 Council #16777237    Expertise Military Theory, Architectural Interest
```

## get_leaders (detail: leader_id=805306373)

```
=== Leader 805306373: HUM1_CHR_Falatir ===
Class: commander
Tier: leader_tier_default
Level: 1 (bonus: 0)
Experience: 69.0
Age: 38
Gender: male
Ethic: Materialist
Job: bureaucrat
Recruited: 2200.01.01

Assignment: Council #16777236
Location: ship (id=3)
Council: council_position (id=16777236, position=0)

Traits:
  - Fleet Organizer
```

## get_traditions_ascension

```
=== Traditions & Ascension Perks (SPEC_Bebaki) ===

Traditions:
  (none adopted)

Ascension Perks:
  (none adopted)

Note: Early game — traditions require Unity to adopt.
```

## get_notifications

```
=== Game Status (SPEC_Bebaki — 2200.07.01) ===

Players:
  - Froupix (Country 0)
  - Bidoof (Country 1)

Active Wars:
  No active wars.

Federations:
  No federations.

Diplomacy:
  Known empires: 0
  Victory rank: 7
```

