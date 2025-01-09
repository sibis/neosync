
/* prettier-ignore */
/**
 * Code generated by Neosync neosync_transformer_typescript_declaration_generator.go. DO NOT EDIT.
 */

/* eslint @typescript-eslint/no-explicit-any: 0 */

declare namespace neosync {

  /**
	 * Transformers
   */
	
	export interface TransformCharacterScrambleOptions {
		/** A custom regular expression. This regex is used to manipulate input data during the transformation process. */
		userProvidedRegex?: string;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * 
   */
	declare function transformCharacterScramble(value: any, options: TransformCharacterScrambleOptions): any;

	
	export interface TransformE164PhoneNumberOptions {
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** Specifies the maximum length for the transformed data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Transforms an existing E164 formatted phone number.
   */
	declare function transformE164PhoneNumber(value: any, options: TransformE164PhoneNumberOptions): any;

	
	export interface TransformEmailOptions {
		/** Specifies the maximum length for the transformed data. This field ensures that the output does not exceed a certain number of characters. */
		preserveLength?: boolean;
		/** A boolean indicating whether the domain part of the email should be preserved. */
		preserveDomain?: boolean;
		/** A list of domains that should be excluded from the transformation */
		excludedDomains?: any[];
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		maxLength?: number;
		/** An optional seed value used for generating deterministic transformations. */
		seed?: number;
		/** Specifies the type of email to transform, with options including `uuidv4`, `fullname`, or `any`. */
		emailType?: string;
		/** Specifies the action to take when an invalid email is encountered, with options including `reject`, `passthrough`, `null`, or `generate`. */
		invalidEmailAction?: string;
	}

  /**
   * Anonymizes and transforms an existing email address.
   */
	declare function transformEmail(value: any, options: TransformEmailOptions): any;

	
	export interface TransformFirstNameOptions {
		/** Specifies the maximum length for the transformed data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** An optional seed value used for generating deterministic transformations. */
		seed?: number;
	}

  /**
   * Transforms an existing first name
   */
	declare function transformFirstName(value: any, options: TransformFirstNameOptions): any;

	
	export interface TransformFloat64Options {
		/** Specifies the minimum value for the range of the float. */
		randomizationRangeMin?: number;
		/** Specifies the maximum value for the randomization range of the float. */
		randomizationRangeMax?: number;
		/** An optional parameter that defines the number of significant digits for the float. */
		precision?: number;
		/** An optional parameter that defines the number of decimal places for the float. */
		scale?: number;
		/** An optional seed value used for generating deterministic transformations. */
		seed?: number;
	}

  /**
   * Anonymizes and transforms an existing float value.
   */
	declare function transformFloat64(value: any, options: TransformFloat64Options): any;

	
	export interface TransformFullNameOptions {
		/** Specifies the maximum length for the transformed data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** An optional seed value used for generating deterministic transformations. */
		seed?: number;
	}

  /**
   * Transforms an existing full name.
   */
	declare function transformFullName(value: any, options: TransformFullNameOptions): any;

	
	export interface TransformInt64Options {
		/** Specifies the minimum value for the range of the int. */
		randomizationRangeMin?: number;
		/** Specifies the maximum value for the range of the int. */
		randomizationRangeMax?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Anonymizes and transforms an existing int64 value.
   */
	declare function transformInt64(value: any, options: TransformInt64Options): any;

	
	export interface TransformInt64PhoneNumberOptions {
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Anonymizes and transforms an existing int64 phone number.
   */
	declare function transformInt64PhoneNumber(value: any, options: TransformInt64PhoneNumberOptions): any;

	
	export interface TransformLastNameOptions {
		/** Specifies the maximum length for the transformed data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** An optional seed value used for generating deterministic transformations. */
		seed?: number;
	}

  /**
   * Anonymizes and transforms an existing last name.
   */
	declare function transformLastName(value: any, options: TransformLastNameOptions): any;

	
	export interface TransformStringOptions {
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** Specifies the minimum length of the transformed value. */
		minLength?: number;
		/** Specifies the maximum length of the transformed value. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Anonymizes and transforms an existing string value.
   */
	declare function transformString(value: any, options: TransformStringOptions): any;

	
	export interface TransformStringPhoneNumberOptions {
		/** Whether the original length of the input data should be preserved during transformation. If set to true, the transformation logic will ensure that the output data has the same length as the input data. */
		preserveLength?: boolean;
		/** Specifies the maximum length for the transformed data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Anonymizes and transforms an existing phone number that is typed as a string.
   */
	declare function transformStringPhoneNumber(value: any, options: TransformStringPhoneNumberOptions): any;

	
	export interface TransformUuidOptions {
		/** An optional seed value used for generating deterministic transformations. */
		seed?: number;
	}

  /**
   * Transforms an existing UUID to a UUID v5
   */
	declare function transformUuid(value: any, options: TransformUuidOptions): any;

	


  /**
	 * Generators
   */
	
	export interface GenerateBoolOptions {
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random boolean value.
   */
	declare function generateBool(options: GenerateBoolOptions): any;

	
	export interface GenerateBusinessNameOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random business name between 2 and 35 characters long.
   */
	declare function generateBusinessName(options: GenerateBusinessNameOptions): any;

	
	export interface GenerateCardNumberOptions {
		/** A boolean indicating whether the generated value should pass the Luhn algorithm check. */
		validLuhn?: boolean;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a 16 digit card number that is valid by Luhn valid by default.
   */
	declare function generateCardNumber(options: GenerateCardNumberOptions): any;

	
	export interface GenerateCategoricalOptions {
		/** A list of comma-separated string values to randomly select from. */
		categories?: string;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly selects a value from a defined set of categorical values.
   */
	declare function generateCategorical(options: GenerateCategoricalOptions): any;

	
	export interface GenerateCityOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly selects a city from a list of predefined US cities.
   */
	declare function generateCity(options: GenerateCityOptions): any;

	
	export interface GenerateCountryOptions {
		/** If true returns the full country name instead of the two character country code. */
		generateFullName?: boolean;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly selects a country and by default, returns it as a 2-letter country code.
   */
	declare function generateCountry(options: GenerateCountryOptions): any;

	
	export interface GenerateEmailOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** Specifies the type of email type to generate, with options including `uuidv4`, `fullname`, or `any`. */
		emailType?: string;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a new randomized email address.
   */
	declare function generateEmail(options: GenerateEmailOptions): any;

	
	export interface GenerateFirstNameOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random first name between 2 and 12 characters long.
   */
	declare function generateFirstName(options: GenerateFirstNameOptions): any;

	
	export interface GenerateFloat64Options {
		/** A boolean indicating whether the sign of the float should be randomized. */
		randomizeSign?: boolean;
		/** Specifies the minimum value for the generated float. */
		min?: number;
		/** Specifies the maximum value for the generated float */
		max?: number;
		/** An optional parameter that defines the number of significant digits for the generated float. */
		precision?: number;
		/** An optional parameter that defines the number of decimal places for the generated float. */
		scale?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random floating point number with a max precision of 17. Go float64 adheres to the IEEE 754 standard for double-precision floating-point numbers.
   */
	declare function generateFloat64(options: GenerateFloat64Options): any;

	
	export interface GenerateFullAddressOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a randomly selected real full address that exists in the United States.
   */
	declare function generateFullAddress(options: GenerateFullAddressOptions): any;

	
	export interface GenerateFullNameOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a new full name consisting of a first and last name.
   */
	declare function generateFullName(options: GenerateFullNameOptions): any;

	
	export interface GenerateGenderOptions {
		/** Shortens length of generated value to 1. */
		abbreviate?: boolean;
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly generates one of the following genders: female (f), male (m), undefined (u), nonbinary (n).
   */
	declare function generateGender(options: GenerateGenderOptions): any;

	
	export interface GenerateInt64Options {
		/** A boolean indicating whether the sign of the float should be randomized. */
		randomizeSign?: boolean;
		/** Specifies the minimum value for the generated int. */
		min?: number;
		/** Specifies the maximum value for the generated int. */
		max?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random int64 value with a default length of 4.
   */
	declare function generateInt64(options: GenerateInt64Options): any;

	
	export interface GenerateInt64PhoneNumberOptions {
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a new int64 phone number with a default length of 10.
   */
	declare function generateInt64PhoneNumber(options: GenerateInt64PhoneNumberOptions): any;

	
	export interface GenerateInternationalPhoneNumberOptions {
		/** Specifies the minimum value for the generated phone number. */
		min?: number;
		/** Specifies the maximum value for the generated phone number. */
		max?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a new random international phone number including the + sign and no hyphens.
   */
	declare function generateInternationalPhoneNumber(options: GenerateInternationalPhoneNumberOptions): any;

	
	export interface GenerateIpAddressOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** IP type to generate. */
		ipType?: string;
		/** Optional seed for deterministic generation */
		seed?: number;
	}

  /**
   * Generates IPv4 or IPv6 addresses with support for different network classes.
   */
	declare function generateIpAddress(options: GenerateIpAddressOptions): any;

	
	export interface GenerateLastNameOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random last name.
   */
	declare function generateLastName(options: GenerateLastNameOptions): any;

	
	export interface GenerateRandomStringOptions {
		/** Specifies the minimum length for the generated string. */
		min?: number;
		/** Specifies the maximum length for the generated string. */
		max?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random string of alphanumeric characters..
   */
	declare function generateRandomString(options: GenerateRandomStringOptions): any;

	
	export interface GenerateSHA256HashOptions {
	}

  /**
   * Generates a random SHA256 hash and returns it as a string.
   */
	declare function generateSHA256Hash(options: GenerateSHA256HashOptions): any;

	
	export interface GenerateSSNOptions {
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random social security numbers including the hyphens in the format xxx-xx-xxxx.
   */
	declare function generateSSN(options: GenerateSSNOptions): any;

	
	export interface GenerateStateOptions {
		/** If true returns the full state name instead of the two character state code. */
		generateFullName?: boolean;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly selects a US state and by default, returns it as a 2-letter state code.
   */
	declare function generateState(options: GenerateStateOptions): any;

	
	export interface GenerateStreetAddressOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly generates a street address.
   */
	declare function generateStreetAddress(options: GenerateStreetAddressOptions): any;

	
	export interface GenerateStringPhoneNumberOptions {
		/** Specifies the minimum length for the generated phone number. */
		min?: number;
		/** Specifies the maximum length for the generated phone number. */
		max?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a random 10 digit phone number and returns it as a string with no hyphens.
   */
	declare function generateStringPhoneNumber(options: GenerateStringPhoneNumberOptions): any;

	
	export interface GenerateUnixTimestampOptions {
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly generates a Unix timestamp that is in the past.
   */
	declare function generateUnixTimestamp(options: GenerateUnixTimestampOptions): any;

	
	export interface GenerateUsernameOptions {
		/** Specifies the maximum length for the generated data. This field ensures that the output does not exceed a certain number of characters. */
		maxLength?: number;
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly generates a username
   */
	declare function generateUsername(options: GenerateUsernameOptions): any;

	
	export interface GenerateUTCTimestampOptions {
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Randomly generates a UTC timestamp.
   */
	declare function generateUTCTimestamp(options: GenerateUTCTimestampOptions): any;

	
	export interface GenerateUUIDOptions {
		/** Determines whether the generated UUID should include hyphens. If set to true, the UUID will be formatted with hyphens (e.g., d853d251-e135-4fe4-a4eb-0aea6bfaf645). If set to false, the hyphens will be omitted (e.g., d853d251e1354fe4a4eb0aea6bfaf645). */
		includeHyphens?: boolean;
	}

  /**
   * Generates a new UUIDv4 id.
   */
	declare function generateUUID(options: GenerateUUIDOptions): any;

	
	export interface GenerateZipcodeOptions {
		/** An optional seed value used to generate deterministic outputs. */
		seed?: number;
	}

  /**
   * Generates a randomly selected US zipcode.
   */
	declare function generateZipcode(options: GenerateZipcodeOptions): any;

	
}