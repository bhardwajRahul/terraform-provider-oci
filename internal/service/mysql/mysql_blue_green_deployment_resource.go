// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlBlueGreenDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.OneHour,
			Update: &tfresource.ThirtyMinutes,
			Delete: &tfresource.TwentyMinutes,
		},
		CreateContext: createMysqlBlueGreenDeploymentWithContext,
		ReadContext:   readMysqlBlueGreenDeploymentWithContext,
		UpdateContext: updateMysqlBlueGreenDeploymentWithContext,
		DeleteContext: deleteMysqlBlueGreenDeploymentWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"channel_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"source_password": {
							Type:      schema.TypeString,
							Required:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"source_username": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"ssl_mode": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"applier_username": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"ssl_ca_certificate": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"certificate_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PEM",
										}, true),
									},
									"contents": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_db_system_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"mysql_version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"configuration_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"data_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"delete_target_db_system_on_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to delete the target DB System when this Blue/Green deployment is deleted.",
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"switchover_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"active_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_channel_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"switchover_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_db_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMysqlBlueGreenDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlBlueGreenDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlueGreenDeploymentsClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}
	return nil

}

func readMysqlBlueGreenDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlBlueGreenDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlueGreenDeploymentsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateMysqlBlueGreenDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlBlueGreenDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlueGreenDeploymentsClient()

	if _, ok := sync.D.GetOkExists("switchover_trigger"); ok && sync.D.HasChange("switchover_trigger") {
		oldRaw, newRaw := sync.D.GetChange("switchover_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			sync.WorkRequestClient = m.(*client.OracleClients).MysqlWorkRequestsClient()
			err := sync.SwitchoverBlueGreenDeployment(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("switchover_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func deleteMysqlBlueGreenDeploymentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MysqlBlueGreenDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlueGreenDeploymentsClient()
	sync.DbSystemClient = m.(*client.OracleClients).DbSystemClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type MysqlBlueGreenDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.BlueGreenDeploymentsClient
	DbSystemClient         *oci_mysql.DbSystemClient
	WorkRequestClient      *oci_mysql.WorkRequestsClient
	Res                    *oci_mysql.BlueGreenDeployment
	DisableNotFoundRetries bool
}

func (s *MysqlBlueGreenDeploymentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlBlueGreenDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.BlueGreenDeploymentLifecycleStateCreating),
	}
}

func (s *MysqlBlueGreenDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.BlueGreenDeploymentLifecycleStateActive),
	}
}

func (s *MysqlBlueGreenDeploymentResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.BlueGreenDeploymentLifecycleStateUpdating),
	}
}

func (s *MysqlBlueGreenDeploymentResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.BlueGreenDeploymentLifecycleStateActive),
	}
}

func (s *MysqlBlueGreenDeploymentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.BlueGreenDeploymentLifecycleStateDeleting),
	}
}

func (s *MysqlBlueGreenDeploymentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.BlueGreenDeploymentLifecycleStateDeleted),
	}
}

func (s *MysqlBlueGreenDeploymentResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_mysql.CreateBlueGreenDeploymentRequest{}

	if channelDetails, ok := s.D.GetOkExists("channel_details"); ok {
		if tmpList := channelDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "channel_details", 0)
			tmp, err := s.mapToCreateBlueGreenDeploymentChannelDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ChannelDetails = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if sourceDbSystemId, ok := s.D.GetOkExists("source_db_system_id"); ok {
		tmp := sourceDbSystemId.(string)
		request.SourceDbSystemId = &tmp
	}

	if targetDbSystemDetails, ok := s.D.GetOkExists("target_db_system_details"); ok {
		if tmpList := targetDbSystemDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_db_system_details", 0)
			tmp, err := s.mapToCreateBlueGreenDeploymentTargetDbSystemDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetDbSystemDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateBlueGreenDeployment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BlueGreenDeployment
	return nil
}

func (s *MysqlBlueGreenDeploymentResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_mysql.GetBlueGreenDeploymentRequest{}

	tmp := s.D.Id()
	request.BlueGreenDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetBlueGreenDeployment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BlueGreenDeployment
	return nil
}

func (s *MysqlBlueGreenDeploymentResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_mysql.UpdateBlueGreenDeploymentRequest{}

	tmp := s.D.Id()
	request.BlueGreenDeploymentId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateBlueGreenDeployment(ctx, request)
	if err != nil {
		return err
	}

	return s.GetWithContext(ctx)
}

func (s *MysqlBlueGreenDeploymentResourceCrud) DeleteWithContext(ctx context.Context) error {
	deleteTargetDbSystem := false
	if value, ok := s.D.GetOkExists("delete_target_db_system_on_delete"); ok {
		deleteTargetDbSystem = value.(bool)
	}

	var targetDbSystemId string
	if deleteTargetDbSystem {
		if value, ok := s.D.GetOk("target_db_system_id"); ok {
			targetDbSystemId = value.(string)
		}
		if targetDbSystemId == "" {
			if err := s.GetWithContext(ctx); err != nil {
				return fmt.Errorf("failed to determine the target DB System before deleting the Blue/Green deployment: %w", err)
			}
			if s.Res.TargetDbSystemId != nil {
				targetDbSystemId = *s.Res.TargetDbSystemId
			}
		}
		if targetDbSystemId == "" {
			return fmt.Errorf("cannot delete the target DB System because its OCID is not available")
		}
	}

	deleteStarted := time.Now()
	deleteTimeout := s.D.Timeout(schema.TimeoutDelete)
	request := oci_mysql.DeleteBlueGreenDeploymentRequest{}

	tmp := s.D.Id()
	request.BlueGreenDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteBlueGreenDeployment(ctx, request)
	if err != nil && !isMysqlNotFoundError(err) {
		return err
	}

	if !deleteTargetDbSystem {
		return nil
	}

	if err == nil {
		remainingTimeout, timeoutErr := remainingMysqlBlueGreenDeleteTimeout(deleteStarted, deleteTimeout)
		if timeoutErr != nil {
			return timeoutErr
		}
		if waitErr := s.waitForBlueGreenDeploymentDeletion(ctx, remainingTimeout); waitErr != nil {
			return fmt.Errorf("failed waiting for Blue/Green deployment %s to be deleted: %w", s.D.Id(), waitErr)
		}
	}

	if err := s.deleteTargetDbSystem(ctx, targetDbSystemId); err != nil {
		return fmt.Errorf("failed to delete target DB System %s: %w", targetDbSystemId, err)
	}

	remainingTimeout, timeoutErr := remainingMysqlBlueGreenDeleteTimeout(deleteStarted, deleteTimeout)
	if timeoutErr != nil {
		return timeoutErr
	}
	if err := s.waitForTargetDbSystemDeletion(ctx, targetDbSystemId, remainingTimeout); err != nil {
		return fmt.Errorf("failed waiting for target DB System %s to be deleted: %w", targetDbSystemId, err)
	}

	return nil
}

func (s *MysqlBlueGreenDeploymentResourceCrud) waitForBlueGreenDeploymentDeletion(ctx context.Context, timeout time.Duration) error {
	request := oci_mysql.GetBlueGreenDeploymentRequest{
		BlueGreenDeploymentId: oci_common.String(s.D.Id()),
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	stateConf := &retry.StateChangeConf{
		Pending: s.DeletedPending(),
		Target:  s.DeletedTarget(),
		Refresh: func() (interface{}, string, error) {
			response, err := s.Client.GetBlueGreenDeployment(ctx, request)
			if isMysqlNotFoundError(err) {
				return s.D.Id(), string(oci_mysql.BlueGreenDeploymentLifecycleStateDeleted), nil
			}
			if err != nil {
				return nil, "", err
			}
			return &response.BlueGreenDeployment, string(response.LifecycleState), nil
		},
		Timeout: timeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func (s *MysqlBlueGreenDeploymentResourceCrud) deleteTargetDbSystem(ctx context.Context, targetDbSystemId string) error {
	request := oci_mysql.DeleteDbSystemRequest{
		DbSystemId: oci_common.String(targetDbSystemId),
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.DbSystemClient.DeleteDbSystem(ctx, request)
	if isMysqlNotFoundError(err) {
		return nil
	}
	return err
}

func (s *MysqlBlueGreenDeploymentResourceCrud) waitForTargetDbSystemDeletion(ctx context.Context, targetDbSystemId string, timeout time.Duration) error {
	request := oci_mysql.GetDbSystemRequest{
		DbSystemId: oci_common.String(targetDbSystemId),
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_mysql.DbSystemLifecycleStateDeleting),
		},
		Target: []string{
			string(oci_mysql.DbSystemLifecycleStateDeleted),
		},
		Refresh: func() (interface{}, string, error) {
			response, err := s.DbSystemClient.GetDbSystem(ctx, request)
			if isMysqlNotFoundError(err) {
				return targetDbSystemId, string(oci_mysql.DbSystemLifecycleStateDeleted), nil
			}
			if err != nil {
				return nil, "", err
			}
			return &response.DbSystem, string(response.LifecycleState), nil
		},
		Timeout: timeout,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}

func isMysqlNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	failure, ok := oci_common.IsServiceError(err)
	return ok && failure.GetHTTPStatusCode() == http.StatusNotFound
}

func remainingMysqlBlueGreenDeleteTimeout(started time.Time, timeout time.Duration) (time.Duration, error) {
	remaining := timeout - time.Since(started)
	if remaining <= 0 {
		return 0, fmt.Errorf("timed out deleting the Blue/Green deployment and its target DB System")
	}
	return remaining, nil
}

func (s *MysqlBlueGreenDeploymentResourceCrud) SetData() error {
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

func (s *MysqlBlueGreenDeploymentResourceCrud) SwitchoverBlueGreenDeployment(ctx context.Context) error {
	request := oci_mysql.SwitchoverBlueGreenDeploymentRequest{}

	idTmp := s.D.Id()
	request.BlueGreenDeploymentId = &idTmp

	if waitTimeInSeconds, ok := s.D.GetOkExists("wait_time_in_seconds"); ok {
		tmp := waitTimeInSeconds.(int)
		request.WaitTimeInSeconds = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.SwitchoverBlueGreenDeployment(ctx, request)
	if err != nil {
		return err
	}
	if response.OpcWorkRequestId != nil {
		if err := blueGreenDeploymentWaitForSwitchoverWorkRequest(
			ctx,
			response.OpcWorkRequestId,
			s.D.Timeout(schema.TimeoutUpdate),
			s.DisableNotFoundRetries,
			s.WorkRequestClient,
		); err != nil {
			return err
		}
	}

	const (
		switchoverPending   = "SWITCHOVER_PENDING"
		switchoverSucceeded = "SWITCHOVER_SUCCEEDED"
	)
	stateConf := &retry.StateChangeConf{
		Pending: []string{switchoverPending},
		Target:  []string{switchoverSucceeded},
		Refresh: func() (interface{}, string, error) {
			if err := s.GetWithContext(ctx); err != nil {
				return nil, "", err
			}

			if s.Res.LifecycleState == oci_mysql.BlueGreenDeploymentLifecycleStateFailed ||
				s.Res.SwitchoverStatus == oci_mysql.BlueGreenDeploymentSwitchoverStatusSwitchoverFailed {
				if s.Res.LifecycleDetails != nil {
					return s.Res, "", fmt.Errorf("blue/green deployment switchover failed: %s", *s.Res.LifecycleDetails)
				}
				return s.Res, "", fmt.Errorf("blue/green deployment switchover failed")
			}

			if s.Res.LifecycleState == oci_mysql.BlueGreenDeploymentLifecycleStateActive &&
				s.Res.SwitchoverStatus == oci_mysql.BlueGreenDeploymentSwitchoverStatusSwitchoverCompleted {
				return s.Res, switchoverSucceeded, nil
			}
			return s.Res, switchoverPending, nil
		},
		Timeout: s.D.Timeout(schema.TimeoutUpdate),
	}
	if _, waitErr := stateConf.WaitForStateContext(ctx); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("switchover_trigger")
	s.D.Set("switchover_trigger", val)

	return nil
}

func blueGreenDeploymentWaitForSwitchoverWorkRequest(
	ctx context.Context,
	workRequestId *string,
	timeout time.Duration,
	disableNotFoundRetries bool,
	workRequestClient *oci_mysql.WorkRequestsClient,
) error {
	retryPolicy := tfresource.GetRetryPolicy(disableNotFoundRetries, "mysql")
	response := oci_mysql.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_mysql.WorkRequestOperationStatusAccepted),
			string(oci_mysql.WorkRequestOperationStatusInProgress),
			string(oci_mysql.WorkRequestOperationStatusCanceling),
		},
		Target: []string{
			string(oci_mysql.WorkRequestOperationStatusSucceeded),
			string(oci_mysql.WorkRequestOperationStatusFailed),
			string(oci_mysql.WorkRequestOperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = workRequestClient.GetWorkRequest(ctx, oci_mysql.GetWorkRequestRequest{
				WorkRequestId: workRequestId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
			return &response.WorkRequest, string(response.Status), err
		},
		Timeout: timeout,
	}
	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return err
	}
	if response.Status == oci_mysql.WorkRequestOperationStatusSucceeded {
		return nil
	}

	return getErrorFromMysqlBlueGreenDeploymentWorkRequest(
		ctx,
		workRequestClient,
		workRequestId,
		retryPolicy,
	)
}

func getErrorFromMysqlBlueGreenDeploymentWorkRequest(
	ctx context.Context,
	workRequestClient *oci_mysql.WorkRequestsClient,
	workRequestId *string,
	retryPolicy *oci_common.RetryPolicy,
) error {
	response, err := workRequestClient.ListWorkRequestErrors(ctx, oci_mysql.ListWorkRequestErrorsRequest{
		WorkRequestId: workRequestId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return err
	}

	allErrors := make([]string, 0, len(response.Items))
	for _, workRequestError := range response.Items {
		if workRequestError.Code != nil && workRequestError.Message != nil {
			allErrors = append(allErrors, fmt.Sprintf("%s: %s", *workRequestError.Code, *workRequestError.Message))
		} else if workRequestError.Message != nil {
			allErrors = append(allErrors, *workRequestError.Message)
		}
	}
	if len(allErrors) == 0 {
		return fmt.Errorf("blue/green switchover work request %s did not succeed", *workRequestId)
	}

	return fmt.Errorf(
		"blue/green switchover work request %s did not succeed: %s",
		*workRequestId,
		strings.Join(allErrors, "\n"),
	)
}

func BlueGreenDeploymentSummaryToMap(obj oci_mysql.BlueGreenDeploymentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveDbSystemId != nil {
		result["active_db_system_id"] = string(*obj.ActiveDbSystemId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.SourceDbSystemId != nil {
		result["source_db_system_id"] = string(*obj.SourceDbSystemId)
	}

	result["state"] = string(obj.LifecycleState)

	result["switchover_status"] = string(obj.SwitchoverStatus)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDbSystemId != nil {
		result["target_db_system_id"] = string(*obj.TargetDbSystemId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *MysqlBlueGreenDeploymentResourceCrud) mapToCaCertificate(fieldKeyFormat string) (oci_mysql.CaCertificate, error) {
	var baseObject oci_mysql.CaCertificate
	//discriminator
	certificateTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "certificate_type"))
	var certificateType string
	if ok {
		certificateType = certificateTypeRaw.(string)
	} else {
		certificateType = "" // default value
	}
	switch strings.ToLower(certificateType) {
	case strings.ToLower("PEM"):
		details := oci_mysql.PemCaCertificate{}
		if contents, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "contents")); ok {
			tmp := contents.(string)
			details.Contents = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown certificate_type '%v' was specified", certificateType)
	}
	return baseObject, nil
}

func (s *MysqlBlueGreenDeploymentResourceCrud) mapToCreateBlueGreenDeploymentChannelDetails(fieldKeyFormat string) (oci_mysql.CreateBlueGreenDeploymentChannelDetails, error) {
	result := oci_mysql.CreateBlueGreenDeploymentChannelDetails{}

	if applierUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "applier_username")); ok {
		tmp := applierUsername.(string)
		result.ApplierUsername = &tmp
	}

	if sourcePassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_password")); ok {
		tmp := sourcePassword.(string)
		result.SourcePassword = &tmp
	}

	if sourceUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_username")); ok {
		tmp := sourceUsername.(string)
		result.SourceUsername = &tmp
	}

	if sslCaCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_ca_certificate")); ok {
		if tmpList := sslCaCertificate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "ssl_ca_certificate"), 0)
			tmp, err := s.mapToCaCertificate(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert ssl_ca_certificate, encountered error: %v", err)
			}
			result.SslCaCertificate = tmp
		}
	}

	if sslMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ssl_mode")); ok {
		result.SslMode = oci_mysql.SslModeEnum(sslMode.(string))
	}

	return result, nil
}

func CreateBlueGreenDeploymentChannelDetailsToMap(obj *oci_mysql.CreateBlueGreenDeploymentChannelDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplierUsername != nil {
		result["applier_username"] = string(*obj.ApplierUsername)
	}

	if obj.SourcePassword != nil {
		result["source_password"] = string(*obj.SourcePassword)
	}

	if obj.SourceUsername != nil {
		result["source_username"] = string(*obj.SourceUsername)
	}

	if obj.SslCaCertificate != nil {
		sslCaCertificateArray := []interface{}{}
		if sslCaCertificateMap := CaCertificateToMap(&obj.SslCaCertificate); sslCaCertificateMap != nil {
			sslCaCertificateArray = append(sslCaCertificateArray, sslCaCertificateMap)
		}
		result["ssl_ca_certificate"] = sslCaCertificateArray
	}

	result["ssl_mode"] = string(obj.SslMode)

	return result
}

func (s *MysqlBlueGreenDeploymentResourceCrud) mapToCreateBlueGreenDeploymentTargetDbSystemDetails(fieldKeyFormat string) (oci_mysql.CreateBlueGreenDeploymentTargetDbSystemDetails, error) {
	result := oci_mysql.CreateBlueGreenDeploymentTargetDbSystemDetails{}

	if configurationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration_id")); ok {
		tmp := configurationId.(string)
		result.ConfigurationId = &tmp
	}

	if dataStorageSizeInGB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_storage_size_in_gb")); ok {
		tmp := dataStorageSizeInGB.(int)
		result.DataStorageSizeInGBs = &tmp
	}

	if mysqlVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mysql_version")); ok {
		tmp := mysqlVersion.(string)
		result.MysqlVersion = &tmp
	}

	if shapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_name")); ok {
		tmp := shapeName.(string)
		result.ShapeName = &tmp
	}

	return result, nil
}

func BlueGreenDeploymentTargetDbSystemDetailsToMap(obj *oci_mysql.BlueGreenDeploymentTargetDbSystemDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigurationId != nil {
		result["configuration_id"] = string(*obj.ConfigurationId)
	}

	if obj.DataStorageSizeInGBs != nil {
		result["data_storage_size_in_gb"] = int(*obj.DataStorageSizeInGBs)
	}

	if obj.MysqlVersion != nil {
		result["mysql_version"] = string(*obj.MysqlVersion)
	}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	return result
}

func (s *MysqlBlueGreenDeploymentResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_mysql.ChangeBlueGreenDeploymentCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BlueGreenDeploymentId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.ChangeBlueGreenDeploymentCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	return tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s)
}
