# Harborctl
A CLI interact to Harbor Registry.

```bash
# harborctl help

Usage:
  harborctl [flags]
  harborctl [command]

Available Commands:
  get         Display one or many resources
  health      Check status of Harbor's component services
  help        Help about any command
  version     Print the version number of Harbor CLI

Flags:
      --config string   config file (default is $HOME/harbor.yaml)
  -h, --help            help for harborctl

Use "harborctl [command] --help" for more information about a command.
```

# How to use

### Check Harbor's components status
```bash
# harborctl health
```

```
┌───┬─────────────┬─────────┐
│   │ SERVICE     │ STATUS  │
├───┼─────────────┼─────────┤
│ 1 │ jobservice  │ healthy │
├───┼─────────────┼─────────┤
│ 2 │ registry    │ healthy │
├───┼─────────────┼─────────┤
│ 3 │ core        │ healthy │
├───┼─────────────┼─────────┤
│ 4 │ portal      │ healthy │
├───┼─────────────┼─────────┤
│ 5 │ registryctl │ healthy │
├───┼─────────────┼─────────┤
│ 6 │ redis       │ healthy │
├───┼─────────────┼─────────┤
│ 7 │ database    │ healthy │
└───┴─────────────┴─────────┘
```

### Get project

```bash
# harborctl get project
```

```
┌───┬────────────────────┬───────────────┬────────┬──────────────────┐
│   │ PROJECT NAME       │ OWNER         │ PUBLIC │ TOTAL REPOSITORY │
├───┼────────────────────┼───────────────┼────────┼──────────────────┤
│ 1 │ hub (id: 12)       │ admin (id: 1) │ false  │              530 │
├───┼────────────────────┼───────────────┼────────┼──────────────────┤
│ 2 │ data (id: 59)      │ admin (id: 1) │ true   │               17 │
├───┼────────────────────┼───────────────┼────────┼──────────────────┤
│ 3 │ hub_cache (id: 49) │ admin (id: 1) │ true   │               80 │
├───┼────────────────────┼───────────────┼────────┼──────────────────┤
│ 4 │ toy (id: 24)       │ admin (id: 1) │ true   │                1 │
└───┴────────────────────┴───────────────┴────────┴──────────────────┘
```

### Get informations of a user
```
# harborctl get user bob
```

```
┌───┬──────────────┬──────────────────────┬──────────────────┐
│   │ USER NAME    │ EMAIL                │ ADMIN PERMISSION │
├───┼──────────────┼──────────────────────┼──────────────────┤
│ 1 │ bob          │ bod@company.org      │ true             │
└───┴──────────────┴──────────────────────┴──────────────────┘
```

### Get all users of registry
```
# harborctl get users
```

```
┌───┬──────────────┬──────────────────────┬──────────────────┐
│   │ USER NAME    │ EMAIL                │ ADMIN PERMISSION │
├───┼──────────────┼──────────────────────┼──────────────────┤
│ 1 │ bob          │ bod@company.org      │ true             │
├───┼──────────────┼──────────────────────┼──────────────────┤
│ 2 │ jane         │ jane@company.org     │ false            │
├───┼──────────────┼──────────────────────┼──────────────────┤
│ 3 │ foo          │ foo@company.org      │ false            │
├───┼──────────────┼──────────────────────┼──────────────────┤
│ 4 │ bar          │ bar@company.org      │ true             │
├───┼──────────────┼──────────────────────┼──────────────────┤
│ 5 │ jack         │ jack@company.org     │ false            │
└───┴──────────────┴──────────────────────┴──────────────────┘
```
