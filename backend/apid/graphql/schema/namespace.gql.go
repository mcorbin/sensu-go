// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	fmt "fmt"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// NamespaceEnvironmentFieldResolver implement to resolve requests for the Namespace's environment field.
type NamespaceEnvironmentFieldResolver interface {
	// Environment implements response to request for environment field.
	Environment(p graphql.ResolveParams) (string, error)
}

// NamespaceOrganizationFieldResolver implement to resolve requests for the Namespace's organization field.
type NamespaceOrganizationFieldResolver interface {
	// Organization implements response to request for organization field.
	Organization(p graphql.ResolveParams) (string, error)
}

//
// NamespaceFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Namespace' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type NamespaceFieldResolvers interface {
	NamespaceEnvironmentFieldResolver
	NamespaceOrganizationFieldResolver
}

// NamespaceAliases implements all methods on NamespaceFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type NamespaceAliases struct{}

// Environment implements response to request for 'environment' field.
func (_ NamespaceAliases) Environment(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := fmt.Sprint(val)
	return ret, err
}

// Organization implements response to request for 'organization' field.
func (_ NamespaceAliases) Organization(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := fmt.Sprint(val)
	return ret, err
}

// NamespaceType Namespace represents the unique details describing where a resource is located.
var NamespaceType = graphql.NewType("Namespace", graphql.ObjectKind)

// RegisterNamespace registers Namespace object type with given service.
func RegisterNamespace(svc *graphql.Service, impl NamespaceFieldResolvers) {
	svc.RegisterObject(_ObjectTypeNamespaceDesc, impl)
}
func _ObjTypeNamespaceEnvironmentHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(NamespaceEnvironmentFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Environment(p)
	}
}

func _ObjTypeNamespaceOrganizationHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(NamespaceOrganizationFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Organization(p)
	}
}

func _ObjectTypeNamespaceConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Namespace represents the unique details describing where a resource is located.",
		Fields: graphql1.Fields{
			"environment": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "environment indicates to which env a check belongs to.",
				Name:              "environment",
				Type:              graphql1.String,
			},
			"organization": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "organization indicates to which org a check belongs to.",
				Name:              "organization",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see NamespaceFieldResolvers.")
		},
		Name: "Namespace",
	}
}

// describe Namespace's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeNamespaceDesc = graphql.ObjectDesc{
	Config: _ObjectTypeNamespaceConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"environment":  _ObjTypeNamespaceEnvironmentHandler,
		"organization": _ObjTypeNamespaceOrganizationHandler,
	},
}
