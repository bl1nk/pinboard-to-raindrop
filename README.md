# Pinboard to Raindrop

Converts a Pinboard JSON export to a CSV file to be imported on Raindrop.

## How to

1. Install this script

   No precompiled binaries, sorry.

   ```shell
   go install github.com/bl1nk/pinboard-to-raindrop@latest
   ```

1. Visit [the Pinboard JSON export URL](https://pinboard.in/export/format:json/).

1. Convert the Pinboard JSON to a CSV file to be imported to Raindrop.

   ```shell
   pinboard-to-raindrop -input pinboard_export.2022.09.20_15.57.json -output to-be-imported.csv
   ```

1. Visit [the Raindrop import page](https://app.raindrop.io/settings/import) and choose `to-be-imported.csv`.
