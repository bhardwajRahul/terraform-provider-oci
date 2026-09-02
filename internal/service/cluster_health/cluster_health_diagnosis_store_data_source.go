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

func ClusterHealthDiagnosisStoreDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["diagnosis_store_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(ClusterHealthDiagnosisStoreResource(), fieldMap, readSingularClusterHealthDiagnosisStoreWithContext)
}

func readSingularClusterHealthDiagnosisStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ClusterHealthDiagnosisStoreDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosisClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ClusterHealthDiagnosisStoreDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cluster_health.DiagnosisClient
	Res    *oci_cluster_health.GetDiagnosisStoreResponse
}

func (s *ClusterHealthDiagnosisStoreDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ClusterHealthDiagnosisStoreDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_cluster_health.GetDiagnosisStoreRequest{}

	if diagnosisStoreId, ok := s.D.GetOkExists("diagnosis_store_id"); ok {
		tmp := diagnosisStoreId.(string)
		request.DiagnosisStoreId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cluster_health")

	response, err := s.Client.GetDiagnosisStore(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ClusterHealthDiagnosisStoreDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastAcceptedRequestToken != nil {
		s.D.Set("last_accepted_request_token", *s.Res.LastAcceptedRequestToken)
	}

	if s.Res.LastCompletedRequestToken != nil {
		s.D.Set("last_completed_request_token", *s.Res.LastCompletedRequestToken)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectStoreBucket != nil {
		s.D.Set("object_store_bucket", *s.Res.ObjectStoreBucket)
	}

	if s.Res.ObjectStoreNamespace != nil {
		s.D.Set("object_store_namespace", *s.Res.ObjectStoreNamespace)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
