package tfstateviz_test

import (
	tfstateviz "github.com/kishaningithub/terralens/pkg"
	"github.com/kishaningithub/terralens/pkg/parser"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestBurnDownDetailed(t *testing.T) {
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
| Module | Resource type | Count |
| --- | --- | ---:|
| root | aws_glue_catalog_database | 2 |
| root | aws_iam_policy | 1 |
| test_mwaa | aws_mwaa_environment | 3 |
| test_mwaa | aws_iam_policy | 1 |
| vpc | aws_iam_role | 1 |
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tfstateviz.BurnDownDetailed(tt.args.resources)
			require.Equal(t, tt.want, got)
		})
	}
}
