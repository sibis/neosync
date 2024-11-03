
// Code generated by Neosync neosync_transformer_list_generator.go. DO NOT EDIT.

package transformers

func GetNeosyncTransformers() []NeosyncTransformer {
	return []NeosyncTransformer{
		NewTransformCharacterScramble(),
		NewTransformE164PhoneNumber(),
		NewTransformEmail(),
		NewTransformFirstName(),
		NewTransformFloat64(),
		NewTransformFullName(),
		NewTransformInt64(),
		NewTransformInt64PhoneNumber(),
		NewTransformLastName(),
		NewTransformString(),
		NewTransformStringPhoneNumber(),
	}
}

func GetNeosyncGenerators() []NeosyncGenerator {
	return []NeosyncGenerator{
			NewGenerateBool(),
			NewGenerateBusinessName(),
			NewGenerateCardNumber(),
			NewGenerateCategorical(),
			NewGenerateCity(),
			NewGenerateCountry(),
			NewGenerateEmail(),
			NewGenerateFirstName(),
			NewGenerateFloat64(),
			NewGenerateFullAddress(),
			NewGenerateFullName(),
			NewGenerateGender(),
			NewGenerateInt64(),
			NewGenerateInt64PhoneNumber(),
			NewGenerateInternationalPhoneNumber(),
			NewGenerateLastName(),
			NewGenerateRandomString(),
			NewGenerateSHA256Hash(),
			NewGenerateSSN(),
			NewGenerateState(),
			NewGenerateStreetAddress(),
			NewGenerateStringPhoneNumber(),
			NewGenerateUnixTimestamp(),
			NewGenerateUsername(),
			NewGenerateUTCTimestamp(),
			NewGenerateUUID(),
			NewGenerateZipcode(),
	}
}
