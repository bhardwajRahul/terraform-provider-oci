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

func IotIotFlowRuntimeFlowDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["iot_flow_runtime_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(IotIotFlowRuntimeFlowResource(), fieldMap, readSingularIotIotFlowRuntimeFlowWithContext)
}

func readSingularIotIotFlowRuntimeFlowWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeFlowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotIotFlowRuntimeFlowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.GetIotFlowRuntimeFlowsResponse
}

func (s *IotIotFlowRuntimeFlowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotIotFlowRuntimeFlowDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotFlowRuntimeFlowsRequest{}

	if iotFlowRuntimeId, ok := s.D.GetOkExists("iot_flow_runtime_id"); ok {
		tmp := iotFlowRuntimeId.(string)
		request.IotFlowRuntimeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.GetIotFlowRuntimeFlows(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotIotFlowRuntimeFlowDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IotIotFlowRuntimeFlowDataSource-", IotIotFlowRuntimeFlowDataSource(), s.D))

	return nil
}
