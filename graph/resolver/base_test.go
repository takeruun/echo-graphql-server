package resolver_test

import (
	"app/graph/resolver"
	"app/test_utils/mock_usecases"
)

var mockAuthUsecase *mock_usecases.MockAuthUsecase
var resolvers resolver.Resolver
