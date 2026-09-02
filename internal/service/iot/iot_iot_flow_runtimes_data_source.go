// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotFlowRuntimesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readIotIotFlowRuntimesWithContext,
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
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iot_domain_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iot_flow_runtime_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(IotIotFlowRuntimeResource()),
						},
					},
				},
			},
		},
	}
}

func readIotIotFlowRuntimesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotIotFlowRuntimesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.ListIotFlowRuntimesResponse
}

func (s *IotIotFlowRuntimesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotIotFlowRuntimesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.ListIotFlowRuntimesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_iot.IotFlowRuntimeLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.ListIotFlowRuntimes(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIotFlowRuntimes(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IotIotFlowRuntimesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotIotFlowRuntimesDataSource-", IotIotFlowRuntimesDataSource(), s.D))
	resources := []map[string]interface{}{}
	iotFlowRuntime := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, IotFlowRuntimeSummaryToMap(item))
	}
	iotFlowRuntime["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, IotIotFlowRuntimesDataSource().Schema["iot_flow_runtime_collection"].Elem.(*schema.Resource).Schema)
		iotFlowRuntime["items"] = items
	}

	resources = append(resources, iotFlowRuntime)
	if err := s.D.Set("iot_flow_runtime_collection", resources); err != nil {
		return err
	}

	return nil
}
