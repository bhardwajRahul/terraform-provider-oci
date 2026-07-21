// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlBlueGreenDeploymentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["blue_green_deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(MysqlBlueGreenDeploymentResource(), fieldMap, readSingularMysqlBlueGreenDeploymentWithContext)
}

func readSingularMysqlBlueGreenDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlBlueGreenDeploymentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlueGreenDeploymentsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MysqlBlueGreenDeploymentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.BlueGreenDeploymentsClient
	Res    *oci_mysql.GetBlueGreenDeploymentResponse
}

func (s *MysqlBlueGreenDeploymentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlBlueGreenDeploymentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_mysql.GetBlueGreenDeploymentRequest{}

	if blueGreenDeploymentId, ok := s.D.GetOkExists("blue_green_deployment_id"); ok {
		tmp := blueGreenDeploymentId.(string)
		request.BlueGreenDeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetBlueGreenDeployment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlBlueGreenDeploymentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActiveDbSystemId != nil {
		s.D.Set("active_db_system_id", *s.Res.ActiveDbSystemId)
	}

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ReplicationChannelId != nil {
		s.D.Set("replication_channel_id", *s.Res.ReplicationChannelId)
	}

	if s.Res.SourceDbSystemId != nil {
		s.D.Set("source_db_system_id", *s.Res.SourceDbSystemId)
	}

	s.D.Set("ssl_mode", s.Res.SslMode)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("switchover_status", s.Res.SwitchoverStatus)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDbSystemDetails != nil {
		s.D.Set("target_db_system_details", []interface{}{BlueGreenDeploymentTargetDbSystemDetailsToMap(s.Res.TargetDbSystemDetails)})
	} else {
		s.D.Set("target_db_system_details", nil)
	}

	if s.Res.TargetDbSystemId != nil {
		s.D.Set("target_db_system_id", *s.Res.TargetDbSystemId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
