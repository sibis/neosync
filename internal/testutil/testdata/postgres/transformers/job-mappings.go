
// Code generated by Neosync jobmapping_generator. DO NOT EDIT.
// source: create-tables.sql

package postgres_transformers

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

func GetDefaultSyncJobMappings(schema string)[]*mgmtv1alpha1.JobMapping {
  return []*mgmtv1alpha1.JobMapping{
		{
			Schema: schema,
			Table:  "transformers",
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
			Table:  "transformers",
			Column: "e164_phone_number",
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
			Table:  "transformers",
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
			Table:  "transformers",
			Column: "measurement",
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
			Table:  "transformers",
			Column: "int64",
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
			Table:  "transformers",
			Column: "int64_phone_number",
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
			Table:  "transformers",
			Column: "string_phone_number",
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
			Table:  "transformers",
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
			Table:  "transformers",
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
			Table:  "transformers",
			Column: "full_name",
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
			Table:  "transformers",
			Column: "str",
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
			Table:  "transformers",
			Column: "character_scramble",
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
			Table:  "transformers",
			Column: "bool",
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
			Table:  "transformers",
			Column: "card_number",
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
			Table:  "transformers",
			Column: "categorical",
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
			Table:  "transformers",
			Column: "city",
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
			Table:  "transformers",
			Column: "full_address",
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
			Table:  "transformers",
			Column: "gender",
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
			Table:  "transformers",
			Column: "international_phone",
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
			Table:  "transformers",
			Column: "sha256",
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
			Table:  "transformers",
			Column: "ssn",
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
			Table:  "transformers",
			Column: "state",
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
			Table:  "transformers",
			Column: "street_address",
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
			Table:  "transformers",
			Column: "unix_time",
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
			Table:  "transformers",
			Column: "username",
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
			Table:  "transformers",
			Column: "utc_timestamp",
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
			Table:  "transformers",
			Column: "uuid",
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
			Table:  "transformers",
			Column: "zipcode",
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

