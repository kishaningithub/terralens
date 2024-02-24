## Getting started

Run the following command inside your terraform repository

```bash
terraform show -json | terralens
```

this shows a burn down list in descending order of the "No. of resources"

|   | Resource name                                      | No. of resources |
|--:|----------------------------------------------------|-----------------:|
| 1 | module.test_vpc                                    |               22 |
| 2 | aws_s3_bucket                                      |                6 |
| 3 | aws_secretsmanager_secret                          |                5 |
| 4 | aws_instance                                       |                1 |

For more info on this view and other supported views checkout [this page](views)

If you have any idea on the view you need for your state file level analysis, feel free to raise an [issue](https://github.com/kishaningithub/terralens/issues)