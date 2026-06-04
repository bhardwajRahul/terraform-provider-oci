// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cluster_health

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cluster_health "github.com/oracle/oci-go-sdk/v65/clusterhealth"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ClusterHealthDiagnosisStoresDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readClusterHealthDiagnosisStoresWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"diagnosis_store_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"diagnosis_store_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ClusterHealthDiagnosisStoreResource()),
						},
					},
				},
			},
		},
	}
}

func readClusterHealthDiagnosisStoresWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ClusterHealthDiagnosisStoresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosisClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ClusterHealthDiagnosisStoresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cluster_health.DiagnosisClient
	Res    *oci_cluster_health.ListDiagnosisStoresResponse
}

func (s *ClusterHealthDiagnosisStoresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ClusterHealthDiagnosisStoresDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_cluster_health.ListDiagnosisStoresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if diagnosisStoreId, ok := s.D.GetOkExists("diagnosis_store_id"); ok {
		tmp := diagnosisStoreId.(string)
		request.DiagnosisStoreId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cluster_health.DiagnosisStoreLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cluster_health")

	response, err := s.Client.ListDiagnosisStores(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDiagnosisStores(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ClusterHealthDiagnosisStoresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ClusterHealthDiagnosisStoresDataSource-", ClusterHealthDiagnosisStoresDataSource(), s.D))
	resources := []map[string]interface{}{}
	diagnosisStore := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiagnosisStoreSummaryToMap(item))
	}
	diagnosisStore["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ClusterHealthDiagnosisStoresDataSource().Schema["diagnosis_store_collection"].Elem.(*schema.Resource).Schema)
		diagnosisStore["items"] = items
	}

	resources = append(resources, diagnosisStore)
	if err := s.D.Set("diagnosis_store_collection", resources); err != nil {
		return err
	}

	return nil
}
