
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: transform_lastname.go

package transformers

import (
	"strings"
	"fmt"
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type TransformLastName struct{}

type TransformLastNameOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
	preserveLength bool
}

func NewTransformLastName() *TransformLastName {
	return &TransformLastName{}
}

func NewTransformLastNameOpts(
	maxLengthArg *int64,
	preserveLengthArg *bool,
  seedArg *int64,
) (*TransformLastNameOpts, error) {
	maxLength := int64(100)
	if maxLengthArg != nil {
		maxLength = *maxLengthArg
	}
	
	preserveLength := bool(false)
	if preserveLengthArg != nil {
		preserveLength = *preserveLengthArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &TransformLastNameOpts{
		maxLength: maxLength,
		preserveLength: preserveLength,
		randomizer: rng.New(seed),	
	}, nil
}

func (o *TransformLastNameOpts) BuildBloblangString(
	valuePath string,	
) string {
	fnStr := []string{ 
		"max_length:%v",
		"value:this.%s", 
		"preserve_length:%v",
	}

	params := []any{
	 	o.maxLength,
		valuePath,
	 	o.preserveLength,
	}

	

	template := fmt.Sprintf("transform_last_name(%s)", strings.Join(fnStr, ","))
	return fmt.Sprintf(template, params...)
}

func (t *TransformLastName) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "transformLastName",
		Description: "Anonymizes and transforms an existing last name.",
		Example: "",
	}, nil
}

func (t *TransformLastName) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &TransformLastNameOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 100
	}
	transformerOpts.maxLength = maxLength

	preserveLength, ok := opts["preserveLength"].(bool)
	if !ok {
		preserveLength = false
	}
	transformerOpts.preserveLength = preserveLength

	var seedArg *int64
	if seedValue, ok := opts["seed"].(int64); ok {
			seedArg = &seedValue
	}
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
	if err != nil {
		return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	transformerOpts.randomizer = rng.New(seed)

	return transformerOpts, nil
}
