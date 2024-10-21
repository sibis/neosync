
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_first_name.go

package transformers

import (
	"strings"
	"fmt"
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateFirstName struct{}

type GenerateFirstNameOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
}

func NewGenerateFirstName() *GenerateFirstName {
	return &GenerateFirstName{}
}

func NewGenerateFirstNameOpts(
	maxLengthArg *int64,
  seedArg *int64,
) (*GenerateFirstNameOpts, error) {
	maxLength := int64(100)
	if maxLengthArg != nil {
		maxLength = *maxLengthArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateFirstNameOpts{
		maxLength: maxLength,
		randomizer: rng.New(seed),	
	}, nil
}

func (o *GenerateFirstNameOpts) BuildBloblangString(	
) string {
	fnStr := []string{ 
		"max_length:%v",
	}

	params := []any{
	 	o.maxLength,
	}

	

	template := fmt.Sprintf("generate_first_name(%s)", strings.Join(fnStr, ","))
	return fmt.Sprintf(template, params...)
}

func (t *GenerateFirstName) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateFirstName",
		Description: "Generates a random first name between 2 and 12 characters long.",
		Example: "",
	}, nil
}

func (t *GenerateFirstName) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateFirstNameOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 100
	}
	transformerOpts.maxLength = maxLength

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
