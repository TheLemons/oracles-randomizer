# Owl statue hints

In the randomizer, owl statue messages are replaced with information about the
seed. The messages are of the form "[location] holds [item].", and follow these
rules:

- Checks that are already required to reach the owl statue in logic are not
  hinted at.
- Each dungeon is a location.
- Groups of map tiles sharing a name are locations.
- Single map tiles with unique names are not locations, and use the name of the
  adjacent/containing group instead.
- Subrosia is partitioned into North Subrosia, South Subrosia, and the Temple
  of Seasons.
- No two owls give the same hint.

Hints and their corresponding owls appear in the log file.
