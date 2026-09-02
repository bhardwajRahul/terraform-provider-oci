package cluster_health

import (
	oci_cluster_health "github.com/oracle/oci-go-sdk/v65/clusterhealth"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("cluster_health", clusterHealthResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportClusterHealthDiagnosisStoreHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cluster_health_diagnosis_store",
	DatasourceClass:        "oci_cluster_health_diagnosis_stores",
	DatasourceItemsAttr:    "diagnosis_store_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "diagnosis_store",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cluster_health.DiagnosisStoreLifecycleStateActive),
	},
}

var clusterHealthResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportClusterHealthDiagnosisStoreHints},
	},
}
