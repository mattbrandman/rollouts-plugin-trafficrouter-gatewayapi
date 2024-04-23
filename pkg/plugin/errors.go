package plugin

const (
	GatewayAPIUpdateError                    = "error updating Gateway API %q: %s"
	GatewayAPIManifestError                  = "No routes configured. One of 'tcpRoutes', 'httpRoutes', 'tcpRoute' or 'httpRoute' should be set"
	HTTPRouteFieldIsEmptyError               = "httpRoute field is empty. It has to be set to remove managed routes"
	InvalidHeaderMatchTypeError              = "invalid header match type"
	BackendRefWasNotFoundInHTTPRouteError    = "backendRef was not found in httpRoute"
	BackendRefWasNotFoundInTCPRouteError     = "backendRef was not found in tcpRoute"
	BackendRefListWasNotFoundInTCPRouteError = "backendRef list was not found in tcpRoute"
	ManagedRouteMapEntryDeleteError          = "can't delete key %q from managedRouteMap. The key %q is not in the managedRouteMap "
)
