package common

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*gql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*gql.Field

var CoordinateGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Coordinate",
	Fields: gql.Fields{
		"lat": &gql.Field{
			Type: gql.NewNonNull(gql.Float),
		},
		"lng": &gql.Field{
			Type: gql.NewNonNull(gql.Float),
		},
	},
})

var CoordinateGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "CoordinateInput",
	Fields: gql.InputObjectConfigFieldMap{
		"lat": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Float),
		},
		"lng": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Float),
		},
	},
})

var CoordinateGraphqlArgs = gql.FieldConfigArgument{
	"lat": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Float),
	},
	"lng": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Float),
	},
}

func CoordinateFromArgs(args map[string]interface{}) *Coordinate {
	return CoordinateInstanceFromArgs(&Coordinate{}, args)
}

func CoordinateInstanceFromArgs(objectFromArgs *Coordinate, args map[string]interface{}) *Coordinate {
	if args["lat"] != nil {
		val := args["lat"]
		objectFromArgs.Lat = float64(val.(float64))
	}
	if args["lng"] != nil {
		val := args["lng"]
		objectFromArgs.Lng = float64(val.(float64))
	}
	return objectFromArgs
}

func (objectFromArgs *Coordinate) FromArgs(args map[string]interface{}) {
	CoordinateInstanceFromArgs(objectFromArgs, args)
}

func (msg *Coordinate) XXX_GraphqlType() *gql.Object {
	return CoordinateGraphqlType
}

func (msg *Coordinate) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return CoordinateGraphqlArgs
}

func (msg *Coordinate) XXX_Package() string {
	return "common"
}

var Int32RangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Int32Range",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"max": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var Int32RangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "Int32RangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"max": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var Int32RangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"max": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func Int32RangeFromArgs(args map[string]interface{}) *Int32Range {
	return Int32RangeInstanceFromArgs(&Int32Range{}, args)
}

func Int32RangeInstanceFromArgs(objectFromArgs *Int32Range, args map[string]interface{}) *Int32Range {
	if args["min"] != nil {
		val := args["min"]
		objectFromArgs.Min = int32(val.(int))
	}
	if args["max"] != nil {
		val := args["max"]
		objectFromArgs.Max = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *Int32Range) FromArgs(args map[string]interface{}) {
	Int32RangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *Int32Range) XXX_GraphqlType() *gql.Object {
	return Int32RangeGraphqlType
}

func (msg *Int32Range) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return Int32RangeGraphqlArgs
}

func (msg *Int32Range) XXX_Package() string {
	return "common"
}

var Int32OptionalRangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Int32OptionalRange",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: gql.Int,
		},
		"max": &gql.Field{
			Type: gql.Int,
		},
	},
})

var Int32OptionalRangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "Int32OptionalRangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: gql.Int,
		},
		"max": &gql.InputObjectFieldConfig{
			Type: gql.Int,
		},
	},
})

var Int32OptionalRangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: gql.Int,
	},
	"max": &gql.ArgumentConfig{
		Type: gql.Int,
	},
}

func Int32OptionalRangeFromArgs(args map[string]interface{}) *Int32OptionalRange {
	return Int32OptionalRangeInstanceFromArgs(&Int32OptionalRange{}, args)
}

func Int32OptionalRangeInstanceFromArgs(objectFromArgs *Int32OptionalRange, args map[string]interface{}) *Int32OptionalRange {
	if args["min"] != nil {
		val := args["min"]
		ptr := int32(val.(int))
		objectFromArgs.Min = &ptr
	}
	if args["max"] != nil {
		val := args["max"]
		ptr := int32(val.(int))
		objectFromArgs.Max = &ptr
	}
	return objectFromArgs
}

func (objectFromArgs *Int32OptionalRange) FromArgs(args map[string]interface{}) {
	Int32OptionalRangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *Int32OptionalRange) XXX_GraphqlType() *gql.Object {
	return Int32OptionalRangeGraphqlType
}

func (msg *Int32OptionalRange) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return Int32OptionalRangeGraphqlArgs
}

func (msg *Int32OptionalRange) XXX_Package() string {
	return "common"
}

var StringRangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "StringRange",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"max": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var StringRangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "StringRangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"max": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var StringRangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"max": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func StringRangeFromArgs(args map[string]interface{}) *StringRange {
	return StringRangeInstanceFromArgs(&StringRange{}, args)
}

func StringRangeInstanceFromArgs(objectFromArgs *StringRange, args map[string]interface{}) *StringRange {
	if args["min"] != nil {
		val := args["min"]
		objectFromArgs.Min = string(val.(string))
	}
	if args["max"] != nil {
		val := args["max"]
		objectFromArgs.Max = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *StringRange) FromArgs(args map[string]interface{}) {
	StringRangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *StringRange) XXX_GraphqlType() *gql.Object {
	return StringRangeGraphqlType
}

func (msg *StringRange) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return StringRangeGraphqlArgs
}

func (msg *StringRange) XXX_Package() string {
	return "common"
}

var MoneyRangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "MoneyRange",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: MoneyGraphqlType,
		},
		"max": &gql.Field{
			Type: MoneyGraphqlType,
		},
	},
})

var MoneyRangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "MoneyRangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: MoneyGraphqlInputType,
		},
		"max": &gql.InputObjectFieldConfig{
			Type: MoneyGraphqlInputType,
		},
	},
})

var MoneyRangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: MoneyGraphqlInputType,
	},
	"max": &gql.ArgumentConfig{
		Type: MoneyGraphqlInputType,
	},
}

func MoneyRangeFromArgs(args map[string]interface{}) *MoneyRange {
	return MoneyRangeInstanceFromArgs(&MoneyRange{}, args)
}

func MoneyRangeInstanceFromArgs(objectFromArgs *MoneyRange, args map[string]interface{}) *MoneyRange {
	if args["min"] != nil {
		val := args["min"]
		objectFromArgs.Min = MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["max"] != nil {
		val := args["max"]
		objectFromArgs.Max = MoneyFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *MoneyRange) FromArgs(args map[string]interface{}) {
	MoneyRangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *MoneyRange) XXX_GraphqlType() *gql.Object {
	return MoneyRangeGraphqlType
}

func (msg *MoneyRange) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return MoneyRangeGraphqlArgs
}

func (msg *MoneyRange) XXX_Package() string {
	return "common"
}

var DateRangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "DateRange",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: DateGraphqlType,
		},
		"max": &gql.Field{
			Type: DateGraphqlType,
		},
	},
})

var DateRangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "DateRangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: DateGraphqlInputType,
		},
		"max": &gql.InputObjectFieldConfig{
			Type: DateGraphqlInputType,
		},
	},
})

var DateRangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: DateGraphqlInputType,
	},
	"max": &gql.ArgumentConfig{
		Type: DateGraphqlInputType,
	},
}

func DateRangeFromArgs(args map[string]interface{}) *DateRange {
	return DateRangeInstanceFromArgs(&DateRange{}, args)
}

func DateRangeInstanceFromArgs(objectFromArgs *DateRange, args map[string]interface{}) *DateRange {
	if args["min"] != nil {
		val := args["min"]
		objectFromArgs.Min = DateFromArgs(val.(map[string]interface{}))
	}
	if args["max"] != nil {
		val := args["max"]
		objectFromArgs.Max = DateFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *DateRange) FromArgs(args map[string]interface{}) {
	DateRangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *DateRange) XXX_GraphqlType() *gql.Object {
	return DateRangeGraphqlType
}

func (msg *DateRange) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return DateRangeGraphqlArgs
}

func (msg *DateRange) XXX_Package() string {
	return "common"
}

var DateGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Date",
	Fields: gql.Fields{
		"year": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"month": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"day": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var DateGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "DateInput",
	Fields: gql.InputObjectConfigFieldMap{
		"year": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"month": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"day": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var DateGraphqlArgs = gql.FieldConfigArgument{
	"year": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"month": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"day": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func DateFromArgs(args map[string]interface{}) *Date {
	return DateInstanceFromArgs(&Date{}, args)
}

func DateInstanceFromArgs(objectFromArgs *Date, args map[string]interface{}) *Date {
	if args["year"] != nil {
		val := args["year"]
		objectFromArgs.Year = int32(val.(int))
	}
	if args["month"] != nil {
		val := args["month"]
		objectFromArgs.Month = int32(val.(int))
	}
	if args["day"] != nil {
		val := args["day"]
		objectFromArgs.Day = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *Date) FromArgs(args map[string]interface{}) {
	DateInstanceFromArgs(objectFromArgs, args)
}

func (msg *Date) XXX_GraphqlType() *gql.Object {
	return DateGraphqlType
}

func (msg *Date) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return DateGraphqlArgs
}

func (msg *Date) XXX_Package() string {
	return "common"
}

var AddressGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Address",
	Fields: gql.Fields{
		"streetAddress": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"town": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"country": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"postalCode": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var AddressGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AddressInput",
	Fields: gql.InputObjectConfigFieldMap{
		"streetAddress": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"town": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"country": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"postalCode": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var AddressGraphqlArgs = gql.FieldConfigArgument{
	"streetAddress": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"town": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"country": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"postalCode": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func AddressFromArgs(args map[string]interface{}) *Address {
	return AddressInstanceFromArgs(&Address{}, args)
}

func AddressInstanceFromArgs(objectFromArgs *Address, args map[string]interface{}) *Address {
	if args["streetAddress"] != nil {
		val := args["streetAddress"]
		objectFromArgs.StreetAddress = string(val.(string))
	}
	if args["town"] != nil {
		val := args["town"]
		objectFromArgs.Town = string(val.(string))
	}
	if args["country"] != nil {
		val := args["country"]
		objectFromArgs.Country = string(val.(string))
	}
	if args["postalCode"] != nil {
		val := args["postalCode"]
		objectFromArgs.PostalCode = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *Address) FromArgs(args map[string]interface{}) {
	AddressInstanceFromArgs(objectFromArgs, args)
}

func (msg *Address) XXX_GraphqlType() *gql.Object {
	return AddressGraphqlType
}

func (msg *Address) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AddressGraphqlArgs
}

func (msg *Address) XXX_Package() string {
	return "common"
}
