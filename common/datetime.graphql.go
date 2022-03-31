package common

import (
	"errors"
	gql "github.com/graphql-go/graphql"
	"time"
)

var DateTimeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "DateTimeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"dateTime": &gql.InputObjectFieldConfig{
			Type: gql.NewInputObject(gql.InputObjectConfig{
				Name: "DateTimeFullObjectInput",
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
					"hours": &gql.InputObjectFieldConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"minutes": &gql.InputObjectFieldConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"seconds": &gql.InputObjectFieldConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"nanos": &gql.InputObjectFieldConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"utcOffsetSeconds": &gql.InputObjectFieldConfig{
						Type: gql.Int,
					},
					"timeZone": &gql.InputObjectFieldConfig{
						Type: gql.String,
					},
				},
			}),
		},
		"ISOString": &gql.InputObjectFieldConfig{
			Type: gql.String,
		},
	},
})

var DateTimeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "DateTime",
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
		"hours": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"minutes": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"seconds": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"nanos": &gql.Field{
			Type: gql.Int,
		},
		"utcOffsetSeconds": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"timeZone": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(*DateTime); ok && t.TimeZone != "" {
					return t.TimeZone, nil
				}
				return "UTC", nil
			},
		},
		"ISOString": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(*DateTime); ok && t != nil {
					dateTime := DateTimeToTime(*t)
					if dateTime.IsZero() {
						return nil, errors.New("cannot convert zero dateTime to iso string")
					}
					return dateTime.Format(time.RFC3339), nil
				}
				return nil, errors.New("invalid dateTime object")
			},
		},
		"unix": &gql.Field{
			Type: gql.Int,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(*DateTime); ok && t != nil {
					dateTime := DateTimeToTime(*t)
					if dateTime.IsZero() {
						return nil, errors.New("cannot convert zero dateTime to unix")
					}
					return dateTime.Unix(), nil
				}
				return nil, errors.New("invalid dateTime object")
			},
		},
		"format": &gql.Field{
			Description: `https://golang.org/pkg/time/#Time.Format Use Format() from Go's time package to format dates and times easily using the reference time "Mon Jan 2 15:04:05 -0700 MST 2006" (https://gotime.agardner.me/)`,
			Args: gql.FieldConfigArgument{
				"layout": &gql.ArgumentConfig{
					Description: "Mon Jan 2 15:04:05 -0700 MST 2006",
					Type:        gql.String,
				},
			},
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(*DateTime); ok && t != nil {
					dateTime := DateTimeToTime(*t)
					if dateTime.IsZero() {
						return nil, errors.New("cannot convert zero dateTime to formatable string")
					}
					return dateTime.Format(p.Args["layout"].(string)), nil
				}
				return nil, errors.New("invalid dateTime object")
			},
		},
	},
})

func DateTimeFromArgs(args map[string]interface{}) *DateTime {
	if args["dateTime"] != nil {
		obj, err := parseObject(args["dateTime"].(map[string]interface{}))
		if err != nil {
			return nil
		}
		return obj
	} else if args["ISOString"] != nil {
		obj, err := parseIsoString(args["ISOString"].(string))
		if err != nil {
			return nil
		}
		return obj
	}
	return nil
}

func parseIsoString(isoString string) (*DateTime, error) {
	time, err := time.Parse(time.RFC3339, isoString)
	if err != nil {
		return nil, err
	}
	return DateTimeFromTime(time), nil
}

func parseObject(args map[string]interface{}) (*DateTime, error) {
	objectFromArgs := &DateTime{}
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
	if args["hours"] != nil {
		val := args["hours"]
		objectFromArgs.Hours = int32(val.(int))
	}
	if args["minutes"] != nil {
		val := args["minutes"]
		objectFromArgs.Minutes = int32(val.(int))
	}
	if args["seconds"] != nil {
		val := args["seconds"]
		objectFromArgs.Seconds = int32(val.(int))
	}
	if args["nanos"] != nil {
		val := args["nanos"]
		objectFromArgs.Nanos = int32(val.(int))
	}
	if args["utcOffsetSeconds"] != nil {
		val := args["utcOffsetSeconds"]
		objectFromArgs.UtcOffsetSeconds = int64(val.(int))
	}
	if args["timeZone"] != nil {
		val := args["timeZone"]
		objectFromArgs.TimeZone = string(val.(string))
	}
	return objectFromArgs, nil
}

func DateTimeToTime(t DateTime) time.Time {
	loc := time.UTC
	if t.TimeZone != "" {
		loc_, err := time.LoadLocation(t.TimeZone)
		if err == nil {
			loc = loc_
		}
	} else if t.UtcOffsetSeconds != 0 {
		loc = time.FixedZone("Fixed", int(t.UtcOffsetSeconds))
	}
	return time.Date(int(t.Year), time.Month(t.Month), int(t.Day), int(t.Hours), int(t.Minutes), int(t.Seconds), int(t.Nanos), loc)
}

func DateTimeFromTime(t time.Time) *DateTime {
	_, offset := t.Zone()
	return &DateTime{
		Year:             int32(t.Year()),
		Month:            int32(t.Month()),
		Day:              int32(t.Day()),
		Hours:            int32(t.Hour()),
		Minutes:          int32(t.Minute()),
		Seconds:          int32(t.Second()),
		Nanos:            int32(t.Nanosecond()),
		UtcOffsetSeconds: int64(offset),
		TimeZone:         t.Location().String(),
	}
}
