# Octopus

WIP

See help:

`go run main.go --help`

Example 1 (no parallel process):

`go run main.go -file sleep.csv -process sleep.sh`

Example 1 (with parallel process):

`go run main.go -file sleep.csv -max-parallel 5 -process sleep.sh`