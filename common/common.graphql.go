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

var DateGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Date",
	Fields: gql.Fields{
		"year": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"month": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"date": &gql.Field{
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
		"date": &gql.InputObjectFieldConfig{
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
	"date": &gql.ArgumentConfig{
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
	if args["date"] != nil {
		val := args["date"]
		objectFromArgs.Date = int32(val.(int))
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
