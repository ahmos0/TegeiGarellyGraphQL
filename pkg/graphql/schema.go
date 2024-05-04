package graphql

import (
	"github.com/graphql-go/graphql"
)

var itemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Item",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"worksName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"imageUrl": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"other": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"allItems": &graphql.Field{
				Type: graphql.NewList(itemType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return getAllItems()
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"putItem": &graphql.Field{
				Type: itemType,
				Args: graphql.FieldConfigArgument{
					"uuid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"worksName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"imageUrl": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"other": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					uuid, _ := params.Args["uuid"].(string)
					worksName, _ := params.Args["worksName"].(string)
					author, _ := params.Args["author"].(string)
					imageUrl, _ := params.Args["imageUrl"].(string)
					other, _ := params.Args["other"].(string)
					return saveItem(uuid, worksName, author, imageUrl, other)
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
