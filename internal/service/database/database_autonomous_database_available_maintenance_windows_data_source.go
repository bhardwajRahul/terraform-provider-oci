// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatabaseAutonomousDatabaseAvailableMaintenanceWindowsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_database_maintenance_window_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"availability_domain": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"day_of_week": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"is_maintenance_window_change_scheduled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"maintenance_end_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"maintenance_start_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabaseAvailableMaintenanceWindowsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAvailableMaintenanceWindowsResponse
}

func (s *DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.ListAvailableMaintenanceWindowsRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAvailableMaintenanceWindows(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAvailableMaintenanceWindows(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSource-", DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSource(), s.D))
	resources := []map[string]interface{}{}
	autonomousDatabaseAvailableMaintenanceWindow := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AutonomousDatabaseMaintenanceWindowSummaryToMap(&item))
	}
	autonomousDatabaseAvailableMaintenanceWindow["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSource().Schema["autonomous_database_maintenance_window_collection"].Elem.(*schema.Resource).Schema)
		autonomousDatabaseAvailableMaintenanceWindow["items"] = items
	}

	resources = append(resources, autonomousDatabaseAvailableMaintenanceWindow)
	if err := s.D.Set("autonomous_database_maintenance_window_collection", resources); err != nil {
		return err
	}

	return nil
}
