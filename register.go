package validate

import "reflect"

var (
	// global validators. contains built-in and user custom
	validators map[string]int
	// all validators func meta information
	validatorMetas map[string]*funcMeta
)

// init: register all built-in validators
func init() {
	validators = make(map[string]int)
	validatorMetas = make(map[string]*funcMeta)

	for n, fv := range validatorValues {
		validators[n] = 1 // built in
		validatorMetas[n] = newFuncMeta(n, true, fv)
	}
}

// validator func reflect.Value
var validatorValues = map[string]reflect.Value{
	// int value
	"lt":  reflect.ValueOf(Lt),
	"gt":  reflect.ValueOf(Gt),
	"min": reflect.ValueOf(Min),
	"max": reflect.ValueOf(Max),
	// value check
	"enum":     reflect.ValueOf(Enum),
	"notIn":    reflect.ValueOf(NotIn),
	"between":  reflect.ValueOf(Between),
	"regexp":   reflect.ValueOf(Regexp),
	"isEqual":  reflect.ValueOf(IsEqual),
	"intEqual": reflect.ValueOf(IntEqual),
	"notEqual": reflect.ValueOf(NotEqual),
	// contains
	"contains":    reflect.ValueOf(Contains),
	"notContains": reflect.ValueOf(NotContains),
	// string contains
	"stringContains": reflect.ValueOf(StringContains),
	"startsWith":     reflect.ValueOf(StartsWith),
	"endsWith":       reflect.ValueOf(EndsWith),
	// data type check
	"isInt":     reflect.ValueOf(IsInt),
	"isMap":     reflect.ValueOf(IsMap),
	"isUint":    reflect.ValueOf(IsUint),
	"isBool":    reflect.ValueOf(IsBool),
	"isFloat":   reflect.ValueOf(IsFloat),
	"isInts":    reflect.ValueOf(IsInts),
	"isArray":   reflect.ValueOf(IsArray),
	"isSlice":   reflect.ValueOf(IsSlice),
	"isString":  reflect.ValueOf(IsString),
	"isStrings": reflect.ValueOf(IsStrings),
	// length
	"length":       reflect.ValueOf(Length),
	"minLength":    reflect.ValueOf(MinLength),
	"maxLength":    reflect.ValueOf(MaxLength),
	"stringLength": reflect.ValueOf(StringLength),
	// string
	"isIntString": reflect.ValueOf(IsIntString),
	// ip
	"isIP":        reflect.ValueOf(IsIP),
	"isIPv4":      reflect.ValueOf(IsIPv4),
	"isIPv6":      reflect.ValueOf(IsIPv6),
	"isEmail":     reflect.ValueOf(IsEmail),
	"isASCII":     reflect.ValueOf(IsASCII),
	"isAlpha":     reflect.ValueOf(IsAlpha),
	"isAlphaNum":  reflect.ValueOf(IsAlphaNum),
	"isAlphaDash": reflect.ValueOf(IsAlphaDash),
	"isBase64":    reflect.ValueOf(IsBase64),
	"isCIDR":      reflect.ValueOf(IsCIDR),
	"isCIDRv4":    reflect.ValueOf(IsCIDRv4),
	"isCIDRv6":    reflect.ValueOf(IsCIDRv6),
	"isDNSName":   reflect.ValueOf(IsDNSName),
	"isDataURI":   reflect.ValueOf(IsDataURI),
	"isEmpty":     reflect.ValueOf(IsEmpty),
	"isHexColor":  reflect.ValueOf(IsHexColor),
	"isISBN10":    reflect.ValueOf(IsISBN10),
	"isISBN13":    reflect.ValueOf(IsISBN13),
	"isJSON":      reflect.ValueOf(IsJSON),
	"isLatitude":  reflect.ValueOf(IsLatitude),
	"isLongitude": reflect.ValueOf(IsLongitude),
	"isMAC":       reflect.ValueOf(IsMAC),
	"isMultiByte": reflect.ValueOf(IsMultiByte),
	"isNumber":    reflect.ValueOf(IsNumber),
	"isNumeric":   reflect.ValueOf(IsNumeric),
	"isCnMobile":  reflect.ValueOf(IsCnMobile),
	//
	"isStringNumber":   reflect.ValueOf(IsStringNumber),
	"hasWhitespace":    reflect.ValueOf(HasWhitespace),
	"isHexadecimal":    reflect.ValueOf(IsHexadecimal),
	"isPrintableASCII": reflect.ValueOf(IsPrintableASCII),
	//
	"isRGBColor": reflect.ValueOf(IsRGBColor),
	"isURL":      reflect.ValueOf(IsURL),
	"isFullURL":  reflect.ValueOf(IsFullURL),
	"isUUID":     reflect.ValueOf(IsUUID),
	"isUUID3":    reflect.ValueOf(IsUUID3),
	"isUUID4":    reflect.ValueOf(IsUUID4),
	"isUUID5":    reflect.ValueOf(IsUUID5),
	// file system
	"isPath":     reflect.ValueOf(IsPath),
	"isDirPath":  reflect.ValueOf(IsDirPath),
	"isFilePath": reflect.ValueOf(IsFilePath),
	"isUnixPath": reflect.ValueOf(IsUnixPath),
	"isWinPath":  reflect.ValueOf(IsWinPath),
	// date check
	"isDate":     reflect.ValueOf(IsDate),
	"afterDate":  reflect.ValueOf(AfterDate),
	"beforeDate": reflect.ValueOf(BeforeDate),
	//
	"afterOrEqualDate":  reflect.ValueOf(AfterOrEqualDate),
	"beforeOrEqualDate": reflect.ValueOf(BeforeOrEqualDate),
}

// define validator alias name mapping
var validatorAliases = map[string]string{
	// alias -> real name
	"in":    "enum",
	"range": "between",
	// type
	"int":     "isInt",
	"integer": "isInt",
	"uint":    "isUint",
	"bool":    "isBool",
	"float":   "isFloat",
	"map":     "isMap",
	"ints":    "isInts", // []int
	"str":     "isString",
	"string":  "isString",
	"strings": "isStrings", // []string
	"arr":     "isArray",
	"array":   "isArray",
	"slice":   "isSlice",
	// val
	"regex":  "regexp",
	"eq":     "isEqual",
	"equal":  "isEqual",
	"intEq":  "intEqual",
	"int_eq": "intEqual",
	"ne":     "notEqual",
	"notEq":  "notEqual",
	"not_eq": "notEqual",
	// int compare
	"lte":          "max",
	"gte":          "min",
	"lessThan":     "lt",
	"less_than":    "lt",
	"greaterThan":  "gt",
	"greater_than": "gt",
	// len
	"len":      "length",
	"lenEq":    "length",
	"len_eq":   "length",
	"lengthEq": "length",
	"minLen":   "minLength",
	"maxLen":   "maxLength",
	"minSize":  "minLength",
	"min_size": "minLength",
	"maxSize":  "maxLength",
	"max_size": "maxLength",
	// string rune length
	"strlen":     "stringLength",
	"strLen":     "stringLength",
	"str_len":    "stringLength",
	"strLength":  "stringLength",
	"runeLen":    "stringLength",
	"rune_len":   "stringLength",
	"runeLength": "stringLength",
	// string contains
	"string_contains": "stringContains",
	"str_contains":    "stringContains",
	"startWith":       "startsWith",
	"start_with":      "startsWith",
	"starts_with":     "startsWith",
	"endWith":         "endsWith",
	"end_with":        "endsWith",
	"ends_with":       "endsWith",
	// string
	"ip":        "isIP",
	"ipv4":      "isIPv4",
	"ipv6":      "isIPv6",
	"email":     "isEmail",
	"intStr":    "isIntString",
	"int_str":   "isIntString",
	"strInt":    "isIntString",
	"str_int":   "isIntString",
	"intString": "isIntString",
	//
	"stringNum":      "isStringNumber",
	"string_num":     "isStringNumber",
	"strNumber":      "isStringNumber",
	"str_number":     "isStringNumber",
	"strNum":         "isStringNumber",
	"str_num":        "isStringNumber",
	"stringNumber":   "isStringNumber",
	"hexadecimal":    "isHexadecimal",
	"hasWhitespace":  "hasWhitespace",
	"has_whitespace": "hasWhitespace",
	"has_wp":         "hasWhitespace",
	"printableASCII": "isPrintableASCII",
	//
	"ascii":      "isASCII",
	"ASCII":      "isASCII",
	"alpha":      "isAlpha",
	"alphaNum":   "isAlphaNum",
	"alpha_num":  "isAlphaNum",
	"alphaDash":  "isAlphaDash",
	"alpha_dash": "isAlphaDash",
	"base64":     "isBase64",
	"cidr":       "isCIDR",
	"CIDR":       "isCIDR",
	"CIDRv4":     "isCIDRv4",
	"CIDRv6":     "isCIDRv6",
	"dnsName":    "isDNSName",
	"dns_name":   "isDNSName",
	"DNSName":    "isDNSName",
	"dataURI":    "isDataURI",
	"data_URI":   "isDataURI",
	"data_uri":   "isDataURI",
	"empty":      "isEmpty",
	"filePath":   "isFilePath",
	"filepath":   "isFilePath",
	"hexColor":   "isHexColor",
	"isbn10":     "isISBN10",
	"ISBN10":     "isISBN10",
	"isbn13":     "isISBN13",
	"ISBN13":     "isISBN13",
	"json":       "isJSON",
	"JSON":       "isJSON",
	"lat":        "isLatitude",
	"latitude":   "isLatitude",
	"lon":        "isLongitude",
	"longitude":  "isLongitude",
	"mac":        "isMAC",
	"multiByte":  "isMultiByte",
	"num":        "isNumber",
	"number":     "isNumber",
	"numeric":    "isNumeric",
	"rgbColor":   "isRGBColor",
	"RGBColor":   "isRGBColor",
	"url":        "isURL",
	"URL":        "isURL",
	"fullURL":    "isFullURL",
	"fullUrl":    "isFullURL",
	"uuid":       "isUUID",
	"uuid3":      "isUUID3",
	"uuid4":      "isUUID4",
	"uuid5":      "isUUID5",
	"UUID":       "isUUID",
	"UUID3":      "isUUID3",
	"UUID4":      "isUUID4",
	"UUID5":      "isUUID5",
	"unixPath":   "isUnixPath",
	"winPath":    "isWinPath",
	"cnMobile":   "isCnMobile",
	// date
	"date":    "isDate",
	"gtDate":  "afterDate",
	"ltDate":  "beforeDate",
	"gteDate": "afterOrEqualDate",
	"lteDate": "beforeOrEqualDate",
	// uploaded file
	"img":       "isImage",
	"file":      "isFile",
	"image":     "isImage",
	"mime":      "inMimeTypes",
	"mimes":     "inMimeTypes",
	"mimeType":  "inMimeTypes",
	"mimeTypes": "inMimeTypes",
	// requiredXXX
	"requiredIf":         "required_if",
	"requiredUnless":     "required_unless",
	"requiredWith":       "required_with",
	"requiredWithAll":    "required_with_all",
	"requiredWithout":    "required_without",
	"requiredWithoutAll": "required_without_all",
}
