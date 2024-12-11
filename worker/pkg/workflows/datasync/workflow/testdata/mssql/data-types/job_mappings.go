
// Code generated by Neosync jobmapping_generator. DO NOT EDIT.
// source: create-table.sql

package mssql_datatypes

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
)

func GetDefaultSyncJobMappings()[]*mgmtv1alpha1.JobMapping {
  return []*mgmtv1alpha1.JobMapping{
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_bigint",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_numeric",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_bit",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_smallint",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_decimal",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_smallmoney",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_int",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_tinyint",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_money",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_float",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_real",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_date",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_datetimeoffset",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_datetime2",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_smalldatetime",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_datetime",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_time",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_char",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_varchar",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_text",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_nchar",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_nvarchar",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_json",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_ntext",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_uniqueidentifier",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
		{
			Schema: "alltypes",
			Table:  "alldatatypes",
			Column: "col_xml",
			Transformer: &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_PASSTHROUGH,
			},
		},
	} 
}



func GetTableColumnTypeMap() map[string]map[string]string {
	return map[string]map[string]string{
		"alltypes.alldatatypes": {
			"col_bigint": "BIGINT",
			"col_numeric": "NUMERIC(18,0)",
			"col_bit": "BIT",
			"col_smallint": "SMALLINT",
			"col_decimal": "DECIMAL(18,0)",
			"col_smallmoney": "SMALLMONEY",
			"col_int": "INT",
			"col_tinyint": "TINYINT",
			"col_money": "MONEY",
			"col_float": "FLOAT",
			"col_real": "REAL",
			"col_date": "DATE",
			"col_datetimeoffset": "DATETIMEOFFSET",
			"col_datetime2": "DATETIME2",
			"col_smalldatetime": "SMALLDATETIME",
			"col_datetime": "DATETIME",
			"col_time": "TIME",
			"col_char": "CHAR(10)",
			"col_varchar": "VARCHAR(50)",
			"col_text": "TEXT",
			"col_nchar": "NCHAR(10)",
			"col_nvarchar": "NVARCHAR(50)",
			"col_json": "NVARCHAR(MAX)",
			"col_ntext": "NTEXT",
			"col_uniqueidentifier": "UNIQUEIDENTIFIER",
			"col_xml": "XML",
		},
	}
}
