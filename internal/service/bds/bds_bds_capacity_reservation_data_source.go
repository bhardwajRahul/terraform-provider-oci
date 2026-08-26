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

func BdsBdsCapacityReservationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_capacity_reservation_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BdsBdsCapacityReservationResource(), fieldMap, readSingularBdsBdsCapacityReservationWithContext)
}

func readSingularBdsBdsCapacityReservationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsCapacityReservationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BdsBdsCapacityReservationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetBdsCapacityReservationResponse
}

func (s *BdsBdsCapacityReservationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsCapacityReservationDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.GetBdsCapacityReservationRequest{}

	if bdsCapacityReservationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationId.(string)
		request.BdsCapacityReservationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetBdsCapacityReservation(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsCapacityReservationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCapacityReservations != nil {
		s.D.Set("compute_capacity_reservations", []interface{}{ComputeCapacityReservationsToMap(s.Res.ComputeCapacityReservations)})
	} else {
		s.D.Set("compute_capacity_reservations", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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
