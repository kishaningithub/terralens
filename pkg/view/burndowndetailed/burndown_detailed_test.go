package burndowndetailed_test

import (
	"github.com/kishaningithub/terralens/pkg/parser"
	"github.com/kishaningithub/terralens/pkg/view/burndowndetailed"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestBurnDownDetailedView(t *testing.T) {
	type args struct {
		resources []parser.TerraformResource
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should list all resource counts in descending order",
			args: args{
				resources: []parser.TerraformResource{
					{
						Address: "module.test_mwaa.nested1.nested2.aws_iam_policy.test_mwaa_permissions",
						Type:    "aws_iam_policy",
					},
					{
						Address: "aws_iam_policy.test_policy",
						Type:    "aws_iam_policy",
					},
					{
						Address: "aws_glue_catalog_database.test_db",
						Type:    "aws_glue_catalog_database",
					},
					{
						Address: "aws_glue_catalog_database.test_db1",
						Type:    "aws_glue_catalog_database",
					},
					{
						Address: "module.test_mwaa.aws_mwaa_environment.test_airflow_env",
						Type:    "aws_mwaa_environment",
					},
					{
						Address: "module.test_mwaa.aws_mwaa_environment.test_airflow_env2",
						Type:    "aws_mwaa_environment",
					},
					{
						Address: "module.test_mwaa.aws_mwaa_environment.test_airflow_env3",
						Type:    "aws_mwaa_environment",
					},
					{
						Address: "module.vpc.aws_iam_role.test_role",
						Type:    "aws_iam_role",
					},
				},
			},
			want: strings.TrimSpace(`
| | Module | Resource type | No. of resources |
| ---:| --- | --- | ---:|
| 1 | root | aws_glue_catalog_database | 2 |
| 2 | root | aws_iam_policy | 1 |
| 3 | test_mwaa | aws_mwaa_environment | 3 |
| 4 | test_mwaa | aws_iam_policy | 1 |
| 5 | vpc | aws_iam_role | 1 |
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := burndowndetailed.BurnDownDetailedView(tt.args.resources)
			require.Equal(t, tt.want, got)
		})
	}
}
