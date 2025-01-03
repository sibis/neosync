// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/anonymization.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb.js";
import type { TransformerConfig } from "./transformer_pb.js";
import { file_mgmt_v1alpha1_transformer } from "./transformer_pb.js";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file mgmt/v1alpha1/anonymization.proto.
 */
export const file_mgmt_v1alpha1_anonymization: GenFile = /*@__PURE__*/
  fileDesc("CiFtZ210L3YxYWxwaGExL2Fub255bWl6YXRpb24ucHJvdG8SDW1nbXQudjFhbHBoYTEimgIKFEFub255bWl6ZU1hbnlSZXF1ZXN0EiQKCmlucHV0X2RhdGEYASADKAlCELpIDZIBCggBEBkiBHICEAESPwoUdHJhbnNmb3JtZXJfbWFwcGluZ3MYAiADKAsyIS5tZ210LnYxYWxwaGExLlRyYW5zZm9ybWVyTWFwcGluZxJLChRkZWZhdWx0X3RyYW5zZm9ybWVycxgDIAEoCzIoLm1nbXQudjFhbHBoYTEuRGVmYXVsdFRyYW5zZm9ybWVyc0NvbmZpZ0gAiAEBEhcKD2hhbHRfb25fZmFpbHVyZRgEIAEoCBIcCgphY2NvdW50X2lkGAUgASgJQgi6SAVyA7ABAUIXChVfZGVmYXVsdF90cmFuc2Zvcm1lcnMiYAoVQW5vbnltaXplTWFueVJlc3BvbnNlEhMKC291dHB1dF9kYXRhGAEgAygJEjIKBmVycm9ycxgCIAMoCzIiLm1nbXQudjFhbHBoYTEuQW5vbnltaXplTWFueUVycm9ycyJvChJUcmFuc2Zvcm1lck1hcHBpbmcSGgoKZXhwcmVzc2lvbhgBIAEoCUIGukgDyAEBEj0KC3RyYW5zZm9ybWVyGAIgASgLMiAubWdtdC52MWFscGhhMS5UcmFuc2Zvcm1lckNvbmZpZ0IGukgDyAEBIqgBChlEZWZhdWx0VHJhbnNmb3JtZXJzQ29uZmlnEjEKB2Jvb2xlYW4YAiABKAsyIC5tZ210LnYxYWxwaGExLlRyYW5zZm9ybWVyQ29uZmlnEisKAW4YAyABKAsyIC5tZ210LnYxYWxwaGExLlRyYW5zZm9ybWVyQ29uZmlnEisKAXMYBCABKAsyIC5tZ210LnYxYWxwaGExLlRyYW5zZm9ybWVyQ29uZmlnIkEKE0Fub255bWl6ZU1hbnlFcnJvcnMSEwoLaW5wdXRfaW5kZXgYASABKAMSFQoNZXJyb3JfbWVzc2FnZRgCIAEoCSLxAQoWQW5vbnltaXplU2luZ2xlUmVxdWVzdBISCgppbnB1dF9kYXRhGAEgASgJEj8KFHRyYW5zZm9ybWVyX21hcHBpbmdzGAIgAygLMiEubWdtdC52MWFscGhhMS5UcmFuc2Zvcm1lck1hcHBpbmcSSwoUZGVmYXVsdF90cmFuc2Zvcm1lcnMYAyABKAsyKC5tZ210LnYxYWxwaGExLkRlZmF1bHRUcmFuc2Zvcm1lcnNDb25maWdIAIgBARIcCgphY2NvdW50X2lkGAQgASgJQgi6SAVyA7ABAUIXChVfZGVmYXVsdF90cmFuc2Zvcm1lcnMiLgoXQW5vbnltaXplU2luZ2xlUmVzcG9uc2USEwoLb3V0cHV0X2RhdGEYASABKAky2AEKFEFub255bWl6YXRpb25TZXJ2aWNlElwKDUFub255bWl6ZU1hbnkSIy5tZ210LnYxYWxwaGExLkFub255bWl6ZU1hbnlSZXF1ZXN0GiQubWdtdC52MWFscGhhMS5Bbm9ueW1pemVNYW55UmVzcG9uc2UiABJiCg9Bbm9ueW1pemVTaW5nbGUSJS5tZ210LnYxYWxwaGExLkFub255bWl6ZVNpbmdsZVJlcXVlc3QaJi5tZ210LnYxYWxwaGExLkFub255bWl6ZVNpbmdsZVJlc3BvbnNlIgBCzgEKEWNvbS5tZ210LnYxYWxwaGExQhJBbm9ueW1pemF0aW9uUHJvdG9QAVpQZ2l0aHViLmNvbS9udWNsZXVzY2xvdWQvbmVvc3luYy9iYWNrZW5kL2dlbi9nby9wcm90b3MvbWdtdC92MWFscGhhMTttZ210djFhbHBoYTGiAgNNWFiqAg1NZ210LlYxYWxwaGExygINTWdtdFxWMWFscGhhMeICGU1nbXRcVjFhbHBoYTFcR1BCTWV0YWRhdGHqAg5NZ210OjpWMWFscGhhMWIGcHJvdG8z", [file_buf_validate_validate, file_mgmt_v1alpha1_transformer]);

/**
 * @generated from message mgmt.v1alpha1.AnonymizeManyRequest
 */
export type AnonymizeManyRequest = Message<"mgmt.v1alpha1.AnonymizeManyRequest"> & {
  /**
   * Array of stringified JSON data to be anonymized (up to 25 items)
   *
   * @generated from field: repeated string input_data = 1;
   */
  inputData: string[];

  /**
   * Array of Transformer mappings
   *
   * @generated from field: repeated mgmt.v1alpha1.TransformerMapping transformer_mappings = 2;
   */
  transformerMappings: TransformerMapping[];

  /**
   * Optional default transformations for any unmapped keys
   *
   * @generated from field: optional mgmt.v1alpha1.DefaultTransformersConfig default_transformers = 3;
   */
  defaultTransformers?: DefaultTransformersConfig;

  /**
   * Flag to indicate whether to stop processing when an error occurs
   * true: stops on first error encounter
   *
   * @generated from field: bool halt_on_failure = 4;
   */
  haltOnFailure: boolean;

  /**
   * The unique account identifier
   *
   * @generated from field: string account_id = 5;
   */
  accountId: string;
};

/**
 * Describes the message mgmt.v1alpha1.AnonymizeManyRequest.
 * Use `create(AnonymizeManyRequestSchema)` to create a new message.
 */
export const AnonymizeManyRequestSchema: GenMessage<AnonymizeManyRequest> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 0);

/**
 * @generated from message mgmt.v1alpha1.AnonymizeManyResponse
 */
export type AnonymizeManyResponse = Message<"mgmt.v1alpha1.AnonymizeManyResponse"> & {
  /**
   * Array of anonymized JSON data
   *
   * @generated from field: repeated string output_data = 1;
   */
  outputData: string[];

  /**
   * Array of errors that occured during anonymization
   *
   * @generated from field: repeated mgmt.v1alpha1.AnonymizeManyErrors errors = 2;
   */
  errors: AnonymizeManyErrors[];
};

/**
 * Describes the message mgmt.v1alpha1.AnonymizeManyResponse.
 * Use `create(AnonymizeManyResponseSchema)` to create a new message.
 */
export const AnonymizeManyResponseSchema: GenMessage<AnonymizeManyResponse> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 1);

/**
 * @generated from message mgmt.v1alpha1.TransformerMapping
 */
export type TransformerMapping = Message<"mgmt.v1alpha1.TransformerMapping"> & {
  /**
   * JQ Expression or Field Path to apply the transformation to
   *
   * @generated from field: string expression = 1;
   */
  expression: string;

  /**
   * Configuration of Transformer to apply
   *
   * @generated from field: mgmt.v1alpha1.TransformerConfig transformer = 2;
   */
  transformer?: TransformerConfig;
};

/**
 * Describes the message mgmt.v1alpha1.TransformerMapping.
 * Use `create(TransformerMappingSchema)` to create a new message.
 */
export const TransformerMappingSchema: GenMessage<TransformerMapping> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 2);

/**
 * @generated from message mgmt.v1alpha1.DefaultTransformersConfig
 */
export type DefaultTransformersConfig = Message<"mgmt.v1alpha1.DefaultTransformersConfig"> & {
  /**
   * Boolean
   *
   * @generated from field: mgmt.v1alpha1.TransformerConfig boolean = 2;
   */
  boolean?: TransformerConfig;

  /**
   * Number
   *
   * @generated from field: mgmt.v1alpha1.TransformerConfig n = 3;
   */
  n?: TransformerConfig;

  /**
   * String
   *
   * @generated from field: mgmt.v1alpha1.TransformerConfig s = 4;
   */
  s?: TransformerConfig;
};

/**
 * Describes the message mgmt.v1alpha1.DefaultTransformersConfig.
 * Use `create(DefaultTransformersConfigSchema)` to create a new message.
 */
export const DefaultTransformersConfigSchema: GenMessage<DefaultTransformersConfig> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 3);

/**
 * @generated from message mgmt.v1alpha1.AnonymizeManyErrors
 */
export type AnonymizeManyErrors = Message<"mgmt.v1alpha1.AnonymizeManyErrors"> & {
  /**
   * Index of input data that caused error
   *
   * @generated from field: int64 input_index = 1;
   */
  inputIndex: bigint;

  /**
   * Error message
   *
   * @generated from field: string error_message = 2;
   */
  errorMessage: string;
};

/**
 * Describes the message mgmt.v1alpha1.AnonymizeManyErrors.
 * Use `create(AnonymizeManyErrorsSchema)` to create a new message.
 */
export const AnonymizeManyErrorsSchema: GenMessage<AnonymizeManyErrors> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 4);

/**
 * @generated from message mgmt.v1alpha1.AnonymizeSingleRequest
 */
export type AnonymizeSingleRequest = Message<"mgmt.v1alpha1.AnonymizeSingleRequest"> & {
  /**
   * Stringified JSON data to be anonymized
   *
   * @generated from field: string input_data = 1;
   */
  inputData: string;

  /**
   * Array of Transformer mappings
   *
   * @generated from field: repeated mgmt.v1alpha1.TransformerMapping transformer_mappings = 2;
   */
  transformerMappings: TransformerMapping[];

  /**
   * Optional default transformations for any unmapped keys
   *
   * @generated from field: optional mgmt.v1alpha1.DefaultTransformersConfig default_transformers = 3;
   */
  defaultTransformers?: DefaultTransformersConfig;

  /**
   * The unique account identifier
   *
   * @generated from field: string account_id = 4;
   */
  accountId: string;
};

/**
 * Describes the message mgmt.v1alpha1.AnonymizeSingleRequest.
 * Use `create(AnonymizeSingleRequestSchema)` to create a new message.
 */
export const AnonymizeSingleRequestSchema: GenMessage<AnonymizeSingleRequest> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 5);

/**
 * @generated from message mgmt.v1alpha1.AnonymizeSingleResponse
 */
export type AnonymizeSingleResponse = Message<"mgmt.v1alpha1.AnonymizeSingleResponse"> & {
  /**
   * Anonymized JSON data
   *
   * @generated from field: string output_data = 1;
   */
  outputData: string;
};

/**
 * Describes the message mgmt.v1alpha1.AnonymizeSingleResponse.
 * Use `create(AnonymizeSingleResponseSchema)` to create a new message.
 */
export const AnonymizeSingleResponseSchema: GenMessage<AnonymizeSingleResponse> = /*@__PURE__*/
  messageDesc(file_mgmt_v1alpha1_anonymization, 6);

/**
 * @generated from service mgmt.v1alpha1.AnonymizationService
 */
export const AnonymizationService: GenService<{
  /**
   * Anonymizes many JSON strings by applying specified transformation mappings.
   *
   * @generated from rpc mgmt.v1alpha1.AnonymizationService.AnonymizeMany
   */
  anonymizeMany: {
    methodKind: "unary";
    input: typeof AnonymizeManyRequestSchema;
    output: typeof AnonymizeManyResponseSchema;
  },
  /**
   * Anonymizes a single JSON strings by applying specified transformation mappings.
   *
   * @generated from rpc mgmt.v1alpha1.AnonymizationService.AnonymizeSingle
   */
  anonymizeSingle: {
    methodKind: "unary";
    input: typeof AnonymizeSingleRequestSchema;
    output: typeof AnonymizeSingleResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_mgmt_v1alpha1_anonymization, 0);

