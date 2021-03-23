package network

// ^ Router Protocol ID Option ^ //
type protocolRouterOption struct {
	local     bool
	group     bool
	groupName string
}

// ! Option to Set Protocol ID for Local
func SetIDForLocal() *protocolRouterOption {
	return &protocolRouterOption{
		local:     true,
		group:     false,
		groupName: "",
	}
}

// ! Option to Set Protocol ID for a Group, TODO
func SetIDForGroup(name string) *protocolRouterOption {
	return &protocolRouterOption{
		local:     false,
		group:     true,
		groupName: name,
	}
}
