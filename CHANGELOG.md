# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.1] - 2022-08-03

- update readme badges and github actions
  - otherwise identical to 1.2.0

## [1.2.0] - 2022-07-17

- fix loss of bonus crop after LivingOffTheLand 3
- add 50% chance to recover enhanced/trait seeds
- disable seed return chance for renewable crops

## [1.1.0] - 2022-06-07

- update hotbox enhancement category icon to upgrade
- add icon glyphs to each seed type

## [1.0.1] - 2022-05-28

- fix bug with remembering seed unlocks after lvl 3

## [1.0.0] - 2022-05-22

- add recipe for enhanced seed crafting by hand
- fix bonus trait; support for higher lotl levels
- fix file generation
- fix gracecorn stage 2 model
- fix perk level descriptions, add [MOD] indicator
- fix researcher 2nd trait seed ingredient count
- fix schematic recipe unlocks, update hotbox recipe
- fix standard recipe unlocks, update hotbox recipe
- fix xml file clean/build process
- hide fully grown plant types
- hide growing plant types
- make Thorny and Explosive traits incompatible
- remove seed return chance during harvest
- rename vanilla->researcher, crystal-hell->standard
- resolve issue with explosive trait collision
- set hotbox repair resources
- update readme

## [0.6.0] - 2022-05-19

- remove 25% decrease to crop yield for renewable
- automate release build pipeline
- add rotting flesh to enhanced seed recipe

## [0.5.0] - 2022-05-16

- adjust localization for lvl 3 Living off the Land
- add schematics for unlocking plant trait recipes
- fix potato, pumpkin, and yucca seed descriptions
- update stage 3 icons to reflect stage 1 plant
- remove traces of Sweet trait; dropped
- add recipes to make schematics
- fix aloe display name
- add localization for schematics
- add unlock source for each schematic item/recipe
- adjust crafting categories
- restructure generator's architecture

## [0.4.1] - 2022-05-06

- remove unused file in CrystalHell folder

## [0.4.0] - 2022-05-06

- split file generation into separate folders
- support division of Vanilla and CrystalHell

## [0.3.0] - 2022-05-05

- move plant-trait compatibility check to trait
- fix issue with corn going from stage 2 -> 3
- add aloe
- add blueberry
- add chrysanthemum
- add coffee
- add cotton
- add goldenrod
- add gracecorn
- add hop
- add potato
- add pumpkin
- add yucca

## [0.2.0] - 2022-05-01

- add Thorny trait, along with all necessary functionality

## [0.1.0] - 2022-05-01

- initial proof of concept
- generation of:
  - blocks.xml
  - recipes.xml
  - Localization.txt (naming)
- custom 'hotbox' workstation and 5 trait combinations
- planning to add:
  - 2 more traits: thorny and sweet
    - thorny: damage and/or bleed while walking through plant
    - sweet: attract animals - chance to spawn an animal upon plant maturity
  - craftable schematics for each seed type and trait type, meant to slow progression, but also increase the feeling of reward
