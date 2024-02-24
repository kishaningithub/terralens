# Burn down list

## Command

```bash
terraform show -json | terralens -d burn_down_list
```

## Result

It shows a burn down list in descending order of the "No. of resources"

|   | Resource name                                      | No. of resources |
|--:|----------------------------------------------------|-----------------:|
| 1 | module.test_vpc                                    |               22 |
| 2 | aws_s3_bucket                                      |                6 |
| 3 | aws_secretsmanager_secret                          |                5 |
| 4 | aws_instance                                       |                1 |

## Computation logic

For modules, the "No. of resources" column represents the number of resources created by that module. For example, if
module.test_vpc is shown as 22 it means this module has 22 resources within it.

For resource,the "No. of resources" column represents the number of resources. For example, if aws_s3_bucket is shown
as 6 it means that there are 6 s3 buckets.

