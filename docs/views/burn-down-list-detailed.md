# Burn down list detailed

## Command

```bash
terraform show -json | terralens -d burn_down_list
```

## Result

|   | Module        | Resource type                     | No. of resources |
|--:|---------------|-----------------------------------|-----------------:|
| 1 | lambda_module | aws_iam_role_policy_attachment    |                2 |
| 2 | lambda_module | aws_iam_policy                    |                1 |
| 3 | lambda_module | aws_lambda_function               |                1 |
| 4 | lambda_module | aws_iam_role                      |                1 |
| 5 | root          | aws_s3_bucket_public_access_block |                6 |
| 6 | root          | aws_s3_bucket                     |                6 |
| 7 | root          | aws_s3_bucket_policy              |                4 |

## Computation logic

The modules are first alphabetically sorted and within the module resources are sorted by the descending order of the
"No. of resources".