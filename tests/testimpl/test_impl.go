package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	awsClient := GetAWSElasticacheClient(t)

	t.Run("TestIsDeployed", func(t *testing.T) {
		subnetGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "elasticache_subnet_group_name")
		out, err := awsClient.DescribeCacheSubnetGroups(context.TODO(), &elasticache.DescribeCacheSubnetGroupsInput{
			CacheSubnetGroupName: aws.String(subnetGroupName),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCacheSubnetGroups: %v", err)
		}

		assert.Equal(t, len(out.CacheSubnetGroups), 1, "Subnet group does not exists!")
	})

	t.Run("TestSubnetsAttached", func(t *testing.T) {
		subnetGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "elasticache_subnet_group_name")
		subnetsIds := terraform.OutputList(t, ctx.TerratestTerraformOptions(), "elasticache_subnet_subnets_ids")
		out, err := awsClient.DescribeCacheSubnetGroups(context.TODO(), &elasticache.DescribeCacheSubnetGroupsInput{
			CacheSubnetGroupName: aws.String(subnetGroupName),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCacheSubnetGroups: %v", err)
		}

		found := 0

		for _, v := range subnetsIds {
			for _, v2 := range out.CacheSubnetGroups[0].Subnets {
				if v == *v2.SubnetIdentifier {
					found++
				}
			}
		}

		assert.Equal(t, len(subnetsIds), found, "Expected assigned subnets to subnet group does not match current assigned subnets to subnet group!")
	})
}

func GetAWSElasticacheClient(t *testing.T) *elasticache.Client {
	awsElasticacheClient := elasticache.NewFromConfig(GetAWSConfig(t))
	return awsElasticacheClient
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
