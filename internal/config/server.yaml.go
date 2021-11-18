package config

type Server struct {
	NodeID int32  `yaml:"node_id"`
	SvrApp string `yaml:"svr_app"`
}
