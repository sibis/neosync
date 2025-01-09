
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_email.go

package transformers

import (
	"strings"
	"fmt"
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateEmail struct{}

type GenerateEmailOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
	emailType string
}

func NewGenerateEmail() *GenerateEmail {
	return &GenerateEmail{}
}

func NewGenerateEmailOpts(
	maxLengthArg *int64,
	emailTypeArg *string,
  seedArg *int64,
) (*GenerateEmailOpts, error) {
	maxLength := int64(100000)
	if maxLengthArg != nil {
		maxLength = *maxLengthArg
	}
	
	emailType := string(GenerateEmailType_UuidV4.String())
	if emailTypeArg != nil {
		emailType = *emailTypeArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateEmailOpts{
		maxLength: maxLength,
		emailType: emailType,
		randomizer: rng.New(seed),
	}, nil
}

func (o *GenerateEmailOpts) BuildBloblangString(
) string {
	fnStr := []string{
		"max_length:%v",
		"email_type:%q",
	}

	params := []any{
	 	o.maxLength,
	 	o.emailType,
	}

	

	template := fmt.Sprintf("generate_email(%s)", strings.Join(fnStr, ","))
	return fmt.Sprintf(template, params...)
}

func (t *GenerateEmail) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateEmail",
		Description: "Generates a new randomized email address.",
		Example: "",
	}, nil
}

func (t *GenerateEmail) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateEmailOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 100000
	}
	transformerOpts.maxLength = maxLength

	emailType, ok := opts["emailType"].(string)
	if !ok {
		emailType = GenerateEmailType_UuidV4.String()
	}
	transformerOpts.emailType = emailType

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
