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
Government: gov_science_directorate (auth_oligarchic)
Origin: origin_default
Ethics: ethic_authoritarian, ethic_fanatic_materialist
Civics: civic_technocracy, civic_crafters

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
  Physics: tech_administrative_ai (progress: 151.3)
  Society: tech_planetary_unification (progress: 134.6)
  Engineering: tech_powered_exoskeletons (progress: 134.6)

Available Research Options:
  Physics: tech_physics_1, tech_shields_2, tech_administrative_ai, tech_power_plant_2
  Society: tech_genome_mapping, tech_doctrine_navy_size_1, tech_eco_simulation, tech_society_1, tech_planetary_unification
  Engineering: tech_ship_armor_2, tech_powered_exoskeletons, tech_afterburners_1, tech_torpedoes_1

Completed Technologies (31):
  - tech_solar_panel_network
  - tech_space_exploration
  - tech_corvettes
  - tech_starbase_1
  - tech_starbase_2
  - tech_assault_armies
  - tech_ship_armor_1
  - tech_thrusters_1
  - tech_space_defense_station_1
  - tech_basic_industry
  - tech_mechanized_mining
  - tech_space_construction
  - tech_mass_drivers_1
  - tech_flak_batteries_1
  - tech_missiles_1
  - tech_basic_science_lab_1
  - tech_fission_power
  - tech_reactor_boosters_1
  - tech_shields_1
  - tech_power_plant_1
  - tech_hyper_drive_1
  - tech_lasers_1
  - tech_pd_tracking_1
  - tech_planetary_defenses
  - tech_interplanetary_commerce
  - tech_industrial_farming
  - tech_hydroponics
  - tech_colonization_1
  - tech_basic_health
  - tech_planetary_government
  - tech_holo_entertainment
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
11     SPEC_Bebaki_planet        pc_continental     19  5219    71 col_capital  4/19 [city:1 energy:1 farming:1 mining:1]
```

## get_planets (detail: planet_id=11)

```
=== Planet 11: SPEC_Bebaki_planet ===
Class: pc_continental, Size: 19
Designation: col_capital
Owner: 0, Controller: 0

Population: 5219 pops
Stability: 70.7
Crime: 0.0
Amenities: 4656 (used: 3780, free: 876)
Housing: 5600 (used: 5216, free: 384)

District Slots: 4 / 19 used
  district_city: 1
  district_generator: 1
  district_mining: 1
  district_farming: 1

Districts Detail (4):
  district_city (lvl 3) — slots: zone_research_unity, zone_industrial
    [zone_default] building_capital
    [zone_default] building_precinct_house
    [zone_research_unity] building_research_lab_1
    [zone_industrial] building_factory_1
  district_generator (lvl 3)
  district_mining (lvl 2)
  district_farming (lvl 4)

Planetary Features (10):
  - d_hot_springs
  - d_rushing_waterfalls
  - d_tempestous_mountain
  - d_veiny_cliffs
  - d_prosperous_mesa
  - d_rich_mountain
  - d_rugged_woods
  - d_fertile_lands
  - d_rugged_woods
  - d_black_soil

Blockers (3):
  - d_decrepit_dwellings
  - d_failing_infrastructure -> d_prosperous_mesa
  - d_failing_infrastructure -> d_prosperous_mesa

Modifiers:
  - prosp_uni_mod (7020 days remaining)

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
2107   Ispyria III               pc_desert            25      yes 14
3224   Camthrin II               pc_continental       25      yes 12
269    Uldor III                 pc_savannah          24      yes 10
294    Istora III                pc_desert            23      yes 14
3537   Rathadore III             pc_desert            23      yes 15
3895   Hocry II                  pc_gaia              23      yes 13
287    Perqim III                pc_alpine            23      yes 14
2287   Tystra III                pc_tropical          22      yes 16
2986   Zastea III                pc_arid              22      yes 12
3063   Cowhig III                pc_tundra            22      yes 14
778    Dossel I                  pc_desert            21      yes 15
1776   Kodracca III              pc_arid              20      yes 12
863    NAME_Unique_System_2_Pla… pc_relic             20      yes 5
1211   Imoleto I                 pc_continental       20      yes 13
3357   Irthius II                pc_tropical          20      yes 15
580    Saidainope III            pc_alpine            20      yes 14
261    Bailleulus III            pc_continental       20      yes 13
2370   Areysak II                pc_continental       20      yes 14
1024   Vrittaka II               pc_tropical          19      yes 13
1829   NAME_wenkwort_prime       pc_gaia              19      yes 9
228    Haedus III                pc_desert            18      yes 11
2980   Atmir II                  pc_savannah          18      yes 12
2952   Owhziea II                pc_ocean             18      yes 10
3408   Ruinam II                 pc_tropical          18      yes 13
152    Jhurope III               pc_arid              18      yes 13
277    Taggallion III            pc_desert            18      yes 11
396    Emara III                 pc_arid              18      yes 13
3586   Ophele III                pc_desert            18      yes 12
3274   Uproth II b               pc_arctic            17      yes 15
1213   Imoleto III               pc_gaia              17      yes 9
1805   Toiubos III               pc_arid              17      yes 11
328    Theta_Scorpii II          pc_alpine            16      yes 14
578    Saidainope II             pc_ocean             16      yes 13
213    Dodonnam III              pc_desert            16      yes 12
318    Alassia III               pc_desert            16      yes 8
1031   Unur II                   pc_tropical          15      yes 13
3620   Betria III                pc_savannah          14      yes 11
3736   Avishek III a             pc_arctic            14      yes 13
774    Ay'arolee III c           pc_alpine            14      yes 13
2063   Uxfriri III               pc_arctic            14      yes 11
251    Lazon I                   pc_desert            14      yes 9
775    Ay'arolee III d           pc_desert            13      yes 11
1257   NAME_HillosC              pc_tundra            13      yes 8
189    Hydrobius II              pc_continental       13      yes 11
2660   Seb III                   pc_tropical          13      yes 12
244    Fomalhaut I               pc_tropical          13      yes 10
1156   Aytoun III a              pc_tropical          13      yes 11
3774   Sjoberg III a             pc_arid              13      yes 13
3782   Sjoberg V a               pc_alpine            12      yes 11
200    Terzam II                 pc_savannah          12      yes 11
1806   Toiubos III a             pc_alpine            12      yes 8
1212   Imoleto II                pc_continental       12      yes 11
2709   Offe'ei III a             pc_arctic            11      yes 11
1855   Fidhilam III a            pc_continental       11      yes 12
1944   Bazzanac IV a             pc_arid              10      yes 7
558    NAME_UbogleeltD b         pc_gaia               5      yes 5
221    Iolam III                 pc_alpine            16       no 13
```

## get_leaders (summary)

```
=== Leaders (SPEC_Bebaki) — 16 leaders ===
Use get_leaders with leader_id for full detail.

ID           Name                      Class        Lvl Age Assignment           Traits
--------------------------------------------------------------------------------------------------------------
120          HUM1_CHR_Jonnara          commander      1   0 farmer               leader_trait_trickster
805306373    HUM1_CHR_Falatir          commander      1  38 Council #16777236    leader_trait_fleet_organizer
121          HUM1_CHR_Darmull          commander      1   0 bureaucrat           leader_trait_adaptable
119          HUM1_CHR_Jonnara          commander      1   0 politician           leader_trait_aggressive
822083587    HUM1_CHR_Shibbala         envoy          1  33 politician           -
754974722    HUM1_CHR_Bemalona         envoy          1  34 bureaucrat           -
117          HUM1_CHR_Corrona          official       1   0 engineer             trait_ruler_eye_for_talent
118          HUM1_CHR_Campramara       official       1   0 miner                leader_trait_resilient
126          HUM1_CHR_Karba            official       1  32 Council #16777238    leader_trait_politician
116          HUM1_CHR_Haghonona        official       1   0 politician           leader_trait_adaptable
124          HUM1_CHR_Kashnak          scientist      1   0 technician           leader_trait_expertise_propulsion, leader_trait_politician, leader_trait_army_veteran, leader_trait_destructive
822083588    HUM1_CHR_Thaloth          scientist      1  36 Fleet: HUM1_SHIP_Un… leader_trait_expertise_biology, leader_trait_archaeologist
122          HUM1_CHR_Jesslur          scientist      1   0 foundry              leader_trait_expertise_industry, leader_trait_spark_of_genius
125          HUM1_CHR_Jesslur          scientist      1  40 Council #16777237    leader_trait_expertise_military_theory, leader_trait_architectural_interest
123          HUM1_CHR_Haghonon         scientist      1   0 foundry              leader_trait_expertise_new_worlds, leader_trait_carefree
520093696    HUM1_CHR_Thalotha         scientist      1  28 technician           leader_trait_expertise_particles, leader_trait_spark_of_genius
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
Ethic: ethic_materialist
Job: bureaucrat
Recruited: 2200.01.01

Assignment: Council #16777236
Location: ship (id=3)
Council: council_position (id=16777236, position=0)

Traits:
  - leader_trait_fleet_organizer
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

