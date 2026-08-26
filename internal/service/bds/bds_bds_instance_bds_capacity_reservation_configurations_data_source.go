// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceBdsCapacityReservationConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readBdsBdsInstanceBdsCapacityReservationConfigurationsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bds_capacity_reservation_configuration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(BdsBdsInstanceBdsCapacityReservationConfigurationResource()),
						},
					},
				},
			},
		},
	}
}

func readBdsBdsInstanceBdsCapacityReservationConfigurationsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCapacityReservationConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BdsBdsInstanceBdsCapacityReservationConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.ListBdsCapacityReservationConfigurationsResponse
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.ListBdsCapacityReservationConfigurationsRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_bds.BdsCapacityReservationConfigurationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.ListBdsCapacityReservationConfigurations(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBdsCapacityReservationConfigurations(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BdsBdsInstanceBdsCapacityReservationConfigurationsDataSource-", BdsBdsInstanceBdsCapacityReservationConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	bdsInstanceBdsCapacityReservationConfiguration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BdsCapacityReservationConfigurationSummaryToMap(item))
	}
	bdsInstanceBdsCapacityReservationConfiguration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BdsBdsInstanceBdsCapacityReservationConfigurationsDataSource().Schema["bds_capacity_reservation_configuration_collection"].Elem.(*schema.Resource).Schema)
		bdsInstanceBdsCapacityReservationConfiguration["items"] = items
	}

	resources = append(resources, bdsInstanceBdsCapacityReservationConfiguration)
	if err := s.D.Set("bds_capacity_reservation_configuration_collection", resources); err != nil {
		return err
	}

	return nil
}
