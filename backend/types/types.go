// package types Defines types used throughout ZaneCluster
package types

import (
	"crypto/x509"
	"time"
)

type ClusterMember struct {
	ip       string
	port     int
	id       string
	lastseen time.Time
}

type ClusterMembers struct {
	count   int
	members []ClusterMember
}

type Encryption int

type EncryptedClusterPassword struct {
	encryption Encryption
	pwd        string
}

type Network struct {
	net  string
	mask string
	port int
}

type SecureClusterConfig struct {
	trustedCertPool *x509.CertPool
	serverCert      string
	serverKey       string
	clusterPassword EncryptedClusterPassword
	clusterNet      Network
}

type Cluster struct {
	members      ClusterMembers
	secure       bool
	secureConfig *SecureClusterConfig
	clusterNet   Network
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
