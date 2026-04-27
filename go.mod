module github.com/mfridman/goose

go 1.21

// Personal fork of pressly/goose for learning and experimentation.
// Upstream: https://github.com/pressly/goose
//
// Changes from upstream:
// - Experimenting with custom migration ordering logic
// - See branch: feat/custom-ordering
//
// TODO: look into whether sequential versioning is worth switching to over
// timestamp-based; seems cleaner for solo projects.
//
// NOTE: bumped goose to v3.18.0 after noticing v3.17.0 had a subtle bug with
// out-of-order migrations when using the WithAllowMissing option.
require (
	github.com/pressly/goose/v3 v3.18.0
	github.com/spf13/cobra v1.8.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
