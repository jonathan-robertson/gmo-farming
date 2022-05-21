# Contributing to the GMO Farming Modlet Project

This modlet is truly massive due to all the plant traits and supported combinations of them.

Becuase of this, we relied on Go to generate the necessary xpath files to make this modlet work as envisioned.

## Code Structure

This project generates 2 separate themes of this mod that the admin can choose between: Standard and Researcher (find out more about them in the [README](README.md)).

Folder | Purpose
--- | ---
Config-Shared | pure, non-generated XML/XPath. Changes made here will show up in both themes
Config-Researcher | ignored by git, generated
Config-Standard | ignored by git, generated
Config | ignored by git, filled by build script
data | go [data](./data) package responsible for specifics related to traits and plants
gen | go [gen](./gen) package responsible for generating xml files

- `build-researcher-locally.sh` is a shell script that allows one to build and prepare a local copy of this mod with the Researcher Theme (helpful for local testing).
- `build-standard-locally.sh` is the same kind of shell script, but for the Standard Theme rather than Researcher.

## Learn Go

If you don't already know Go, you can learn it [over here](https://go.dev/learn/).

I highly recommend the [Tour of Go](https://go.dev/tour/), which can usually get someone up to speed in about 30mins to 1hr.

## More Questions?

Feel free to reach out to me: <https://github.com/jonathan-robertson/gmo-farming/discussions/53>
