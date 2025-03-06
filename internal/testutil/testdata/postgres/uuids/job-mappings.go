
// Code generated by Neosync jobmapping_generator. DO NOT EDIT.
// source: create-tables.sql

package postgres_uuids

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

func GetDefaultSyncJobMappings(schema string)[]*mgmtv1alpha1.JobMapping {
  return []*mgmtv1alpha1.JobMapping{
		{
			Schema: schema,
			Table:  "store_notifications",
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
			Table:  "stores",
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
			Table:  "stores",
			Column: "notifications_id",
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
			Table:  "store_customers",
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
			Table:  "store_customers",
			Column: "store_id",
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
			Table:  "store_customers",
			Column: "referred_by_code",
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
			Table:  "referral_codes",
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
			Table:  "referral_codes",
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
	}
}

