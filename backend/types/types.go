// package types Defines types used throughout ZaneCluster
package types

import (
	"crypto/x509"
	"time"
)

type ClusterMember struct {
	IP       string
	Port     int
	ID       string
	LastSeen time.Time
}

type ClusterMembers struct {
	Count   int
	Members []ClusterMember
}

type Encryption int

type EncryptedClusterPassword struct {
	Encryption Encryption
	Pwd        string
}

type Network struct {
	Net  string
	Mask string
	Port int
}

type SecureClusterConfig struct {
	TrustedCertPool *x509.CertPool
	ServerCert      string
	ServerKey       string
	ClusterPassword EncryptedClusterPassword
	ClusterNet      Network
}

type Cluster struct {
	Members      ClusterMembers
	Secure       bool
	SecureConfig *SecureClusterConfig
	ClusterNet   Network
}

type ClusterConfig struct {
	Cluster struct {
		Network     string     `yaml:"network"`
		NetMask     string     `yaml:"netmask"`
		Port        int        `yaml:"port"`
		Encryption  Encryption `yaml:"encryption_protocol"`
		Password    string     `yaml:"encrypted_password"`
		SSLCertFile string     `yaml:"ssl_certificate_file"`
		SSLKeyFile  string     `yaml:"ssl_key_file"`
	} `yaml:"cluster"`
}
