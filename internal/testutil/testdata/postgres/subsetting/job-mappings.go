
// Code generated by Neosync jobmapping_generator. DO NOT EDIT.
// source: create-tables.sql

package postgres_subsetting

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

func GetDefaultSyncJobMappings(schema string)[]*mgmtv1alpha1.JobMapping {
  return []*mgmtv1alpha1.JobMapping{
		{
			Schema: schema,
			Table:  "test_2_x",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_x",
			Column: "name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_x",
			Column: "created",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_b",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_b",
			Column: "name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_b",
			Column: "created",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_a",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_a",
			Column: "x_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_c",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_c",
			Column: "name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_c",
			Column: "created",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_c",
			Column: "a_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_c",
			Column: "b_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_d",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_d",
			Column: "c_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_e",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_2_e",
			Column: "c_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_a",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_b",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_b",
			Column: "a_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_c",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_c",
			Column: "b_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_d",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_d",
			Column: "c_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_e",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_e",
			Column: "d_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_f",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_g",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_g",
			Column: "f_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_h",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_h",
			Column: "g_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_i",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "test_3_i",
			Column: "h_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "addresses",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "addresses",
			Column: "order_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "customers",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "customers",
			Column: "address_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "orders",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "orders",
			Column: "customer_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "payments",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "payments",
			Column: "customer_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "division",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "division",
			Column: "division_name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "division",
			Column: "location",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "employees",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "employees",
			Column: "division_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "employees",
			Column: "first_name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "employees",
			Column: "last_name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "employees",
			Column: "email",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "projects",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "projects",
			Column: "project_name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "projects",
			Column: "start_date",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "projects",
			Column: "end_date",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "projects",
			Column: "responsible_employee_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "projects",
			Column: "responsible_division_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "bosses",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "bosses",
			Column: "manager_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "bosses",
			Column: "big_boss_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "minions",
			Column: "id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "minions",
			Column: "boss_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "users",
			Column: "user_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "users",
			Column: "name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "users",
			Column: "email",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "users",
			Column: "manager_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "users",
			Column: "mentor_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "initiatives",
			Column: "initiative_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "initiatives",
			Column: "name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "initiatives",
			Column: "description",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "initiatives",
			Column: "lead_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "initiatives",
			Column: "client_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "task_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "title",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "description",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "status",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "initiative_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "assignee_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "tasks",
			Column: "reviewer_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "skills",
			Column: "skill_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "skills",
			Column: "name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "skills",
			Column: "category",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "user_skills",
			Column: "user_skill_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "user_skills",
			Column: "user_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "user_skills",
			Column: "skill_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "user_skills",
			Column: "proficiency_level",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "comment_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "content",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "created_at",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "user_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "task_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "initiative_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "comments",
			Column: "parent_comment_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "attachment_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "file_name",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "file_path",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "uploaded_by",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "task_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "initiative_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
		{
			Schema: schema,
			Table:  "attachments",
			Column: "comment_id",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Config:
					&mgmtv1alpha1.TransformerConfig{
						Config: &mgmtv1alpha1.TransformerConfig_PassthroughConfig{
							PassthroughConfig: &mgmtv1alpha1.Passthrough{},
						},
					},
			},
		},
	} 
}

