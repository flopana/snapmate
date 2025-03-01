# Snapmate

Snapmate is an automatic snapshot utility for timeshift on Arch Linux.<br>
Inspired by [timeshift-autosnap](https://aur.archlinux.org/packages/timeshift-autosnap) and [snap-pac](https://archlinux.org/packages/extra/any/snap-pac/)

## Installation
### AUR
Snapmate is available on the AUR as [snapmate](https://aur.archlinux.org/packages/snapmate/)

## Usage
Snapmate gets triggered by a libalpm hook and creates a snapshot of the system before upgrading packages.
Snapshots can be skipped by setting the `SKIP_SNAPSHOT` environment variable to anything other than an empty string.

## Configuration
### Snapshots
- `maxSnapshots` - Maximum number of snapshots to keep
- `deleteSnapshots` - Delete old snapshots if the number of snapshots exceeds `maxSnapshots`
- `minTimeBetween` - Minimum time between snapshots in minutes, if the last snapshot is older than this value, a new snapshot will be created

### Logging
- `debugLog` - Enable debug logging

### Database
- `path` - Path to the database file