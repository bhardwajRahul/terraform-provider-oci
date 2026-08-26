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

func MysqlBlueGreenDeploymentsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readMysqlBlueGreenDeploymentsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"blue_green_deployment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MysqlBlueGreenDeploymentResource()),
						},
					},
				},
			},
		},
	}
}

func readMysqlBlueGreenDeploymentsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlBlueGreenDeploymentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlueGreenDeploymentsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MysqlBlueGreenDeploymentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.BlueGreenDeploymentsClient
	Res    *oci_mysql.ListBlueGreenDeploymentsResponse
}

func (s *MysqlBlueGreenDeploymentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlBlueGreenDeploymentsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_mysql.ListBlueGreenDeploymentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sourceDbSystemId, ok := s.D.GetOkExists("source_db_system_id"); ok {
		tmp := sourceDbSystemId.(string)
		request.SourceDbSystemId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.BlueGreenDeploymentSummaryLifecycleStateEnum(state.(string))
	}

	if targetDbSystemId, ok := s.D.GetOkExists("target_db_system_id"); ok {
		tmp := targetDbSystemId.(string)
		request.TargetDbSystemId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListBlueGreenDeployments(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBlueGreenDeployments(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlBlueGreenDeploymentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlBlueGreenDeploymentsDataSource-", MysqlBlueGreenDeploymentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	blueGreenDeployment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BlueGreenDeploymentSummaryToMap(item))
	}
	blueGreenDeployment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MysqlBlueGreenDeploymentsDataSource().Schema["blue_green_deployment_collection"].Elem.(*schema.Resource).Schema)
		blueGreenDeployment["items"] = items
	}

	resources = append(resources, blueGreenDeployment)
	if err := s.D.Set("blue_green_deployment_collection", resources); err != nil {
		return err
	}

	return nil
}
