package config

type SonrConfig struct {
}

type IdentityConfig struct {
}

type ServiceConfig struct {
}

type DatastoreConfig struct {
}

type MatrixConfig struct {
}

type Config struct {
    Sonr     SonrConfig
    Identity IdentityConfig
    Service  ServiceConfig
    Datastore DatastoreConfig
    Matrix   MatrixConfig
}
