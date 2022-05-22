# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2022-05-21

- remove seed return chance during harvest
- add recipe for enhanced seed crafting by hand
- rename vanilla->researcher, crystal-hell->standard
- fix file generation
- fix schematic recipe unlocks, update hotbox recipe
- fix standard recipe unlocks, update hotbox recipe
- hide fully grown plant types
- hide growing plant types
- fix gracecorn stage 2 model
- set hotbox repair resources
- resolve issue with explosive trait collision
- update readme
- make Thorny and Explosive traits incompatible
- fix perk level descriptions, add [MOD] indicator
- fix bonus trait; support for higher lotl levels

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
