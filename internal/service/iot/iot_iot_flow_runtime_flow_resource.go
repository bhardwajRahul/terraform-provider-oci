// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotFlowRuntimeFlowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotFlowRuntimeFlowWithContext,
		ReadContext:   readIotIotFlowRuntimeFlowWithContext,
		UpdateContext: updateIotIotFlowRuntimeFlowWithContext,
		DeleteContext: deleteIotIotFlowRuntimeFlowWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"flows_document": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
			},
			"iot_flow_runtime_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createIotIotFlowRuntimeFlowWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeFlowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotFlowRuntimeFlowWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeFlowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotIotFlowRuntimeFlowWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotFlowRuntimeFlowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotIotFlowRuntimeFlowWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type IotIotFlowRuntimeFlowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    map[string]interface{}
	DisableNotFoundRetries bool
}

func (s *IotIotFlowRuntimeFlowResourceCrud) ID() string {
	return GetIotFlowRuntimeFlowCompositeId(s.D.Get("iot_flow_runtime_id").(string))
}

func (s *IotIotFlowRuntimeFlowResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.UpdateIotFlowRuntimeFlowsRequest{}

	if flowsDocument, ok := s.D.GetOkExists("flows_document"); ok {
		flowsDocumentMap, err := JsonStringToMap(flowsDocument.(string))
		if err != nil {
			return err
		}
		request.FlowsDocument = flowsDocumentMap
	}

	if iotFlowRuntimeId, ok := s.D.GetOkExists("iot_flow_runtime_id"); ok {
		tmp := iotFlowRuntimeId.(string)
		request.IotFlowRuntimeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateIotFlowRuntimeFlows(ctx, request)
	if err != nil {
		return err
	}

	s.Res = response.Object
	return nil
}

func (s *IotIotFlowRuntimeFlowResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotFlowRuntimeFlowsRequest{}

	if iotFlowRuntimeId, ok := s.D.GetOkExists("iot_flow_runtime_id"); ok {
		tmp := iotFlowRuntimeId.(string)
		request.IotFlowRuntimeId = &tmp
	}

	iotFlowRuntimeId, err := parseIotFlowRuntimeFlowCompositeId(s.D.Id())
	if err == nil {
		request.IotFlowRuntimeId = &iotFlowRuntimeId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotFlowRuntimeFlows(ctx, request)
	if err != nil {
		return err
	}

	s.Res = response.Object
	return nil
}

func (s *IotIotFlowRuntimeFlowResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_iot.UpdateIotFlowRuntimeFlowsRequest{}

	if flowsDocument, ok := s.D.GetOkExists("flows_document"); ok {
		flowsDocumentMap, err := JsonStringToMap(flowsDocument.(string))
		if err != nil {
			return err
		}
		request.FlowsDocument = flowsDocumentMap
	}

	if iotFlowRuntimeId, ok := s.D.GetOkExists("iot_flow_runtime_id"); ok {
		tmp := iotFlowRuntimeId.(string)
		request.IotFlowRuntimeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateIotFlowRuntimeFlows(ctx, request)
	if err != nil {
		return err
	}

	s.Res = response.Object
	return nil
}

func (s *IotIotFlowRuntimeFlowResourceCrud) SetData() error {
	iotFlowRuntimeId, err := parseIotFlowRuntimeFlowCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("iot_flow_runtime_id", &iotFlowRuntimeId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	return nil
}

func GetIotFlowRuntimeFlowCompositeId(iotFlowRuntimeId string) string {
	iotFlowRuntimeId = url.PathEscape(iotFlowRuntimeId)
	compositeId := "iotFlowRuntimes/" + iotFlowRuntimeId + "/flows"
	return compositeId
}

func parseIotFlowRuntimeFlowCompositeId(compositeId string) (iotFlowRuntimeId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("iotFlowRuntimes/.*/flows", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	iotFlowRuntimeId, _ = url.PathUnescape(parts[1])

	return
}
