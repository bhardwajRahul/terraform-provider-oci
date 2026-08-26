// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlBlueGreenDeploymentRequiredOnlyResource = MysqlBlueGreenDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Required, acctest.Create, MysqlBlueGreenDeploymentRepresentation)

	MysqlBlueGreenDeploymentResourceConfig = MysqlBlueGreenDeploymentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Update, MysqlBlueGreenDeploymentRepresentation)

	MysqlBlueGreenDeploymentSingularDataSourceRepresentation = map[string]interface{}{
		"blue_green_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_blue_green_deployment.test_blue_green_deployment.id}`},
	}

	MysqlBlueGreenDeploymentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"source_db_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"target_db_system_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_blue_green_deployment.test_blue_green_deployment.target_db_system_id}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentDataSourceFilterRepresentation}}
	MysqlBlueGreenDeploymentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_mysql_blue_green_deployment.test_blue_green_deployment.id}`}},
	}

	MysqlBlueGreenDeploymentRepresentation = map[string]interface{}{
		"channel_details":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentChannelDetailsRepresentation},
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                      acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"source_db_system_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"target_db_system_details":          acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentTargetDbSystemDetailsRepresentation},
		"delete_target_db_system_on_delete": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"defined_tags":                      acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"switchover_trigger":                acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}
	MysqlBlueGreenDeploymentChannelDetailsRepresentation = map[string]interface{}{
		"source_password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"source_username": acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"ssl_mode":        acctest.Representation{RepType: acctest.Required, Create: `REQUIRED`},
	}
	MysqlBlueGreenDeploymentTargetDbSystemDetailsRepresentation = map[string]interface{}{
		"mysql_version":           acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.mysql_version}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_configuration.test_mysql_configuration.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `50`},
		"shape_name":              acctest.Representation{RepType: acctest.Optional, Create: `MySQL.2`},
	}
	MysqlBlueGreenDeploymentImportedTargetDbSystemRepresentation = acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
		"shape_name": acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"admin_password", "admin_username"}},
		}},
	})
	MysqlBlueGreenDeploymentSecurityListRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"egress_security_rules":  []acctest.RepresentationGroup{{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentEgressSecurityRuleRepresentation}},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentIngressSecurityRuleRepresentation}},
	}
	MysqlBlueGreenDeploymentEgressSecurityRuleRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `all`},
	}
	MysqlBlueGreenDeploymentIngressSecurityRuleRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.cidr_block}`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentTcpOptionsRepresentation},
	}
	MysqlBlueGreenDeploymentTcpOptionsRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `3306`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `3306`},
	}
	MysqlBlueGreenDeploymentBackupPolicyRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
	MysqlBlueGreenDeploymentResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, MysqlBlueGreenDeploymentSecurityListRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
				"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_security_list.id}`}},
			})) +
		DefinedTagsDependencies +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(MysqlMysqlConfigurationRepresentation, map[string]interface{}{
				"shape_name": acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(MysqlMysqlDbSystemRepresentation, map[string]interface{}{
				"backup_policy": acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlBlueGreenDeploymentBackupPolicyRepresentation},
				"shape_name":    acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
			}))
)

// issue-routing-tag: mysql/default
func TestMysqlBlueGreenDeploymentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlBlueGreenDeploymentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() + `
provider "oci" {
	ignore_defined_tags = [
		"Oracle-Recommended-Tags.ResourceType",
		"Oracle-Tags.CreatedBy",
		"Oracle-Tags.CreatedOn",
	]
}
`

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_mysql_blue_green_deployment.test_blue_green_deployment"
	sourceDbSystemResourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"
	targetDbSystemResourceName := "oci_mysql_mysql_db_system.test_target_db_system"
	datasourceName := "data.oci_mysql_blue_green_deployments.test_blue_green_deployments"
	singularDatasourceName := "data.oci_mysql_blue_green_deployment.test_blue_green_deployment"
	preserveTargetRepresentation := acctest.RepresentationCopyWithNewProperties(MysqlBlueGreenDeploymentRepresentation, map[string]interface{}{
		"delete_target_db_system_on_delete": acctest.Representation{RepType: acctest.Required, Create: `false`},
	})
	preserveTargetResourceConfig := acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Required, acctest.Create, preserveTargetRepresentation)
	importedTargetResourceConfig := acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_target_db_system", acctest.Required, acctest.Create, MysqlBlueGreenDeploymentImportedTargetDbSystemRepresentation)
	updateWithoutSwitchoverRepresentation := acctest.RepresentationCopyWithNewProperties(MysqlBlueGreenDeploymentRepresentation, map[string]interface{}{
		"switchover_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
	})

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MysqlBlueGreenDeploymentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Create, MysqlBlueGreenDeploymentRepresentation), "mysql", "blueGreenDeployment", t)

	acctest.ResourceTest(t, testAccCheckMysqlBlueGreenDeploymentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies + preserveTargetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_username", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "delete_target_db_system_on_delete", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "source_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.#", "1"),
				resource.TestCheckResourceAttrPair(resourceName, "target_db_system_details.0.mysql_version", sourceDbSystemResourceName, "mysql_version"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Import the generated target DB System so Terraform can manage it independently.
		{
			Config:             config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies + preserveTargetResourceConfig + importedTargetResourceConfig,
			ImportState:        true,
			ImportStateIdFunc:  getMysqlBlueGreenDeploymentTargetDbSystemOcid(resourceName),
			ImportStatePersist: true,
			ResourceName:       targetDbSystemResourceName,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrPair(resourceName, "target_db_system_id", targetDbSystemResourceName, "id"),
			),
		},

		// Delete the Blue/Green deployment and verify that its imported target remains.
		{
			Config: config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies + importedTargetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(targetDbSystemResourceName, "id"),
				func(s *terraform.State) error {
					if _, ok := s.RootModule().Resources[resourceName]; ok {
						return fmt.Errorf("expected %s to be removed from state", resourceName)
					}
					return nil
				},
			),
		},

		// Delete the imported target DB System before the next Create.
		{
			Config: config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Create, MysqlBlueGreenDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_username", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_details.0.configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.0.data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttrPair(resourceName, "target_db_system_details.0.mysql_version", sourceDbSystemResourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_details.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						exportId := "oci_mysql_blue_green_deployment:" + resId
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&exportId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MysqlBlueGreenDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(MysqlBlueGreenDeploymentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_username", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_details.0.configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.0.data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttrPair(resourceName, "target_db_system_details.0.mysql_version", sourceDbSystemResourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_details.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Update, updateWithoutSwitchoverRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.source_username", "adminUser"),
				resource.TestCheckResourceAttr(resourceName, "channel_details.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_details.0.configuration_id"),
				resource.TestCheckResourceAttr(resourceName, "target_db_system_details.0.data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttrPair(resourceName, "target_db_system_details.0.mysql_version", sourceDbSystemResourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(resourceName, "target_db_system_details.0.shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify switchover after the preceding updates have completed
		{
			Config: config + compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Update, MysqlBlueGreenDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "switchover_trigger", "1"),
				resource.TestCheckResourceAttr(resourceName, "switchover_status", "SWITCHOVER_COMPLETED"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_blue_green_deployments", "test_blue_green_deployments", acctest.Optional, acctest.Update, MysqlBlueGreenDeploymentDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Optional, acctest.Update, MysqlBlueGreenDeploymentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_db_system_id"),

				resource.TestCheckResourceAttr(datasourceName, "blue_green_deployment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "blue_green_deployment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_blue_green_deployment", "test_blue_green_deployment", acctest.Required, acctest.Create, MysqlBlueGreenDeploymentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlBlueGreenDeploymentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "blue_green_deployment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_db_system_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replication_channel_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ssl_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "switchover_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_db_system_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_db_system_details.0.data_storage_size_in_gb", "50"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "target_db_system_details.0.mysql_version", sourceDbSystemResourceName, "mysql_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MysqlBlueGreenDeploymentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"channel_details",
				"delete_target_db_system_on_delete",
				"switchover_trigger",
			},
			ResourceName: resourceName,
		},
	})
}

func getMysqlBlueGreenDeploymentTargetDbSystemOcid(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok || rs.Primary == nil {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		targetDbSystemId := rs.Primary.Attributes["target_db_system_id"]
		if targetDbSystemId == "" {
			return "", fmt.Errorf("target_db_system_id not found: %s", resourceName)
		}
		return targetDbSystemId, nil
	}
}

func testAccCheckMysqlBlueGreenDeploymentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BlueGreenDeploymentsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_blue_green_deployment" {
			noResourceFound = false
			request := oci_mysql.GetBlueGreenDeploymentRequest{}

			tmp := rs.Primary.ID
			request.BlueGreenDeploymentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetBlueGreenDeployment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.BlueGreenDeploymentLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MysqlBlueGreenDeployment") {
		resource.AddTestSweepers("MysqlBlueGreenDeployment", &resource.Sweeper{
			Name:         "MysqlBlueGreenDeployment",
			Dependencies: acctest.DependencyGraph["blueGreenDeployment"],
			F:            sweepMysqlBlueGreenDeploymentResource,
		})
	}
}

func sweepMysqlBlueGreenDeploymentResource(compartment string) error {
	blueGreenDeploymentsClient := acctest.GetTestClients(&schema.ResourceData{}).BlueGreenDeploymentsClient()
	blueGreenDeploymentIds, err := getMysqlBlueGreenDeploymentIds(compartment)
	if err != nil {
		return err
	}
	for _, blueGreenDeploymentId := range blueGreenDeploymentIds {
		if ok := acctest.SweeperDefaultResourceId[blueGreenDeploymentId]; !ok {
			deleteBlueGreenDeploymentRequest := oci_mysql.DeleteBlueGreenDeploymentRequest{}

			deleteBlueGreenDeploymentRequest.BlueGreenDeploymentId = &blueGreenDeploymentId

			deleteBlueGreenDeploymentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := blueGreenDeploymentsClient.DeleteBlueGreenDeployment(context.Background(), deleteBlueGreenDeploymentRequest)
			if error != nil {
				fmt.Printf("Error deleting BlueGreenDeployment %s %s, It is possible that the resource is already deleted. Please verify manually \n", blueGreenDeploymentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &blueGreenDeploymentId, MysqlBlueGreenDeploymentSweepWaitCondition, time.Duration(3*time.Minute),
				MysqlBlueGreenDeploymentSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlBlueGreenDeploymentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BlueGreenDeploymentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	blueGreenDeploymentsClient := acctest.GetTestClients(&schema.ResourceData{}).BlueGreenDeploymentsClient()

	listBlueGreenDeploymentsRequest := oci_mysql.ListBlueGreenDeploymentsRequest{}
	listBlueGreenDeploymentsRequest.CompartmentId = &compartmentId
	listBlueGreenDeploymentsRequest.LifecycleState = oci_mysql.BlueGreenDeploymentSummaryLifecycleStateActive
	listBlueGreenDeploymentsResponse, err := blueGreenDeploymentsClient.ListBlueGreenDeployments(context.Background(), listBlueGreenDeploymentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BlueGreenDeployment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, blueGreenDeployment := range listBlueGreenDeploymentsResponse.Items {
		id := *blueGreenDeployment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BlueGreenDeploymentId", id)
	}
	return resourceIds, nil
}

func MysqlBlueGreenDeploymentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if blueGreenDeploymentResponse, ok := response.Response.(oci_mysql.GetBlueGreenDeploymentResponse); ok {
		return blueGreenDeploymentResponse.LifecycleState != oci_mysql.BlueGreenDeploymentLifecycleStateDeleted
	}
	return false
}

func MysqlBlueGreenDeploymentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BlueGreenDeploymentsClient().GetBlueGreenDeployment(context.Background(), oci_mysql.GetBlueGreenDeploymentRequest{
		BlueGreenDeploymentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
