# Reference

```bash
$ terralens -h
See your Terraform state world with clarity and precision - TerraLens, your visual command center for infrastructure insight!

Usage:
  terralens [flags]

Examples:

## Default usage (shows burn_down_list)
terraform show -json | terralens

## Show detailed burn down list
terraform show -json | terralens --display burn_down_list_detailed


Flags:
  -d, --display string   must be one of burn_down_list,burn_down_list_detailed (default "burn_down_list")
  -h, --help             help for terralens
  -v, --version          version for terralens
```