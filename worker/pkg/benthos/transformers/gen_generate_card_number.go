
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_card_number.go

package transformers

import (
	"strings"
	"fmt"
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateCardNumber struct{}

type GenerateCardNumberOpts struct {
	randomizer     rng.Rand
	
	validLuhn bool
}

func NewGenerateCardNumber() *GenerateCardNumber {
	return &GenerateCardNumber{}
}

func NewGenerateCardNumberOpts(
	validLuhnArg *bool,
  seedArg *int64,
) (*GenerateCardNumberOpts, error) {
	validLuhn := bool(false)
	if validLuhnArg != nil {
		validLuhn = *validLuhnArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateCardNumberOpts{
		validLuhn: validLuhn,
		randomizer: rng.New(seed),
	}, nil
}

func (o *GenerateCardNumberOpts) BuildBloblangString(
) string {
	fnStr := []string{
		"valid_luhn:%v",
	}

	params := []any{
	 	o.validLuhn,
	}

	

	template := fmt.Sprintf("generate_card_number(%s)", strings.Join(fnStr, ","))
	return fmt.Sprintf(template, params...)
}

func (t *GenerateCardNumber) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateCardNumber",
		Description: "Generates a 16 digit card number that is valid by Luhn valid by default.",
		Example: "",
	}, nil
}

func (t *GenerateCardNumber) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateCardNumberOpts{}

	validLuhn, ok := opts["validLuhn"].(bool)
	if !ok {
		validLuhn = false
	}
	transformerOpts.validLuhn = validLuhn

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
