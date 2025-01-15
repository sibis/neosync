# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: mgmt/v1alpha1/metrics.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'mgmt/v1alpha1/metrics.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bmgmt/v1alpha1/metrics.proto\x12\rmgmt.v1alpha1\x1a\x1b\x62uf/validate/validate.proto\"d\n\x04\x44\x61te\x12\x1e\n\x04year\x18\x01 \x01(\rB\n\xbaH\x07*\x05\x18\x8fN(\x00R\x04year\x12\x1f\n\x05month\x18\x02 \x01(\rB\t\xbaH\x06*\x04\x18\x1f(\x00R\x05month\x12\x1b\n\x03\x64\x61y\x18\x03 \x01(\rB\t\xbaH\x06*\x04\x18\x1f(\x00R\x03\x64\x61y\"\xa6\x02\n\x1aGetDailyMetricCountRequest\x12)\n\x05start\x18\x01 \x01(\x0b\x32\x13.mgmt.v1alpha1.DateR\x05start\x12%\n\x03\x65nd\x18\x02 \x01(\x0b\x32\x13.mgmt.v1alpha1.DateR\x03\x65nd\x12\x37\n\x06metric\x18\x03 \x01(\x0e\x32\x1f.mgmt.v1alpha1.RangedMetricNameR\x06metric\x12)\n\naccount_id\x18\x04 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\taccountId\x12!\n\x06job_id\x18\x05 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\x05jobId\x12!\n\x06run_id\x18\x06 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\x05runIdB\x0c\n\nidentifier\"Q\n\x1bGetDailyMetricCountResponse\x12\x32\n\x07results\x18\x01 \x03(\x0b\x32\x18.mgmt.v1alpha1.DayResultR\x07results\"J\n\tDayResult\x12\'\n\x04\x64\x61te\x18\x01 \x01(\x0b\x32\x13.mgmt.v1alpha1.DateR\x04\x64\x61te\x12\x14\n\x05\x63ount\x18\x02 \x01(\x04R\x05\x63ount\"\xbb\x02\n\x15GetMetricCountRequest\x12\x37\n\x06metric\x18\x03 \x01(\x0e\x32\x1f.mgmt.v1alpha1.RangedMetricNameR\x06metric\x12)\n\naccount_id\x18\x04 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\taccountId\x12!\n\x06job_id\x18\x05 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\x05jobId\x12!\n\x06run_id\x18\x06 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\x05runId\x12\x30\n\tstart_day\x18\x07 \x01(\x0b\x32\x13.mgmt.v1alpha1.DateR\x08startDay\x12,\n\x07\x65nd_day\x18\x08 \x01(\x0b\x32\x13.mgmt.v1alpha1.DateR\x06\x65ndDayB\x0c\n\nidentifierJ\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03\".\n\x16GetMetricCountResponse\x12\x14\n\x05\x63ount\x18\x01 \x01(\x04R\x05\x63ount*]\n\x10RangedMetricName\x12\"\n\x1eRANGED_METRIC_NAME_UNSPECIFIED\x10\x00\x12%\n!RANGED_METRIC_NAME_INPUT_RECEIVED\x10\x01\x32\xe1\x01\n\x0eMetricsService\x12n\n\x13GetDailyMetricCount\x12).mgmt.v1alpha1.GetDailyMetricCountRequest\x1a*.mgmt.v1alpha1.GetDailyMetricCountResponse\"\x00\x12_\n\x0eGetMetricCount\x12$.mgmt.v1alpha1.GetMetricCountRequest\x1a%.mgmt.v1alpha1.GetMetricCountResponse\"\x00\x42\xc8\x01\n\x11\x63om.mgmt.v1alpha1B\x0cMetricsProtoP\x01ZPgithub.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1;mgmtv1alpha1\xa2\x02\x03MXX\xaa\x02\rMgmt.V1alpha1\xca\x02\rMgmt\\V1alpha1\xe2\x02\x19Mgmt\\V1alpha1\\GPBMetadata\xea\x02\x0eMgmt::V1alpha1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'mgmt.v1alpha1.metrics_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\021com.mgmt.v1alpha1B\014MetricsProtoP\001ZPgithub.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1;mgmtv1alpha1\242\002\003MXX\252\002\rMgmt.V1alpha1\312\002\rMgmt\\V1alpha1\342\002\031Mgmt\\V1alpha1\\GPBMetadata\352\002\016Mgmt::V1alpha1'
  _globals['_DATE'].fields_by_name['year']._loaded_options = None
  _globals['_DATE'].fields_by_name['year']._serialized_options = b'\272H\007*\005\030\217N(\000'
  _globals['_DATE'].fields_by_name['month']._loaded_options = None
  _globals['_DATE'].fields_by_name['month']._serialized_options = b'\272H\006*\004\030\037(\000'
  _globals['_DATE'].fields_by_name['day']._loaded_options = None
  _globals['_DATE'].fields_by_name['day']._serialized_options = b'\272H\006*\004\030\037(\000'
  _globals['_GETDAILYMETRICCOUNTREQUEST'].fields_by_name['account_id']._loaded_options = None
  _globals['_GETDAILYMETRICCOUNTREQUEST'].fields_by_name['account_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETDAILYMETRICCOUNTREQUEST'].fields_by_name['job_id']._loaded_options = None
  _globals['_GETDAILYMETRICCOUNTREQUEST'].fields_by_name['job_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETDAILYMETRICCOUNTREQUEST'].fields_by_name['run_id']._loaded_options = None
  _globals['_GETDAILYMETRICCOUNTREQUEST'].fields_by_name['run_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETMETRICCOUNTREQUEST'].fields_by_name['account_id']._loaded_options = None
  _globals['_GETMETRICCOUNTREQUEST'].fields_by_name['account_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETMETRICCOUNTREQUEST'].fields_by_name['job_id']._loaded_options = None
  _globals['_GETMETRICCOUNTREQUEST'].fields_by_name['job_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETMETRICCOUNTREQUEST'].fields_by_name['run_id']._loaded_options = None
  _globals['_GETMETRICCOUNTREQUEST'].fields_by_name['run_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_RANGEDMETRICNAME']._serialized_start=999
  _globals['_RANGEDMETRICNAME']._serialized_end=1092
  _globals['_DATE']._serialized_start=75
  _globals['_DATE']._serialized_end=175
  _globals['_GETDAILYMETRICCOUNTREQUEST']._serialized_start=178
  _globals['_GETDAILYMETRICCOUNTREQUEST']._serialized_end=472
  _globals['_GETDAILYMETRICCOUNTRESPONSE']._serialized_start=474
  _globals['_GETDAILYMETRICCOUNTRESPONSE']._serialized_end=555
  _globals['_DAYRESULT']._serialized_start=557
  _globals['_DAYRESULT']._serialized_end=631
  _globals['_GETMETRICCOUNTREQUEST']._serialized_start=634
  _globals['_GETMETRICCOUNTREQUEST']._serialized_end=949
  _globals['_GETMETRICCOUNTRESPONSE']._serialized_start=951
  _globals['_GETMETRICCOUNTRESPONSE']._serialized_end=997
  _globals['_METRICSSERVICE']._serialized_start=1095
  _globals['_METRICSSERVICE']._serialized_end=1320
# @@protoc_insertion_point(module_scope)
