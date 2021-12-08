/*
Copyright 2020 NDDO.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/yndd/nddo-grpc/ndd"
	resource "github.com/yndd/nddo-grpc/resource/resourcepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultTimeout = 30 * time.Second
	maxMsgSize     = 512 * 1024 * 1024
)

func NewClient(ctx context.Context, c *ndd.Config) (resource.ResourceClient, error) {
	var opts []grpc.DialOption
	fmt.Printf("grpc client config: %v\n", *c)
	//if c.Insecure {
	//	opts = append(opts, grpc.WithInsecure())
	//} else {
	tlsConfig, err := newTLS(*c)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	//}
	timeoutCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	conn, err := grpc.DialContext(timeoutCtx, c.Address, opts...)
	if err != nil {
		return nil, err
	}
	//defer conn.Close()
	client := resource.NewResourceClient(conn)

	return client, nil
}

// newTLS sets up a new TLS profile
func newTLS(c ndd.Config) (*tls.Config, error) {
	tlsConfig := &tls.Config{
		Renegotiation:      tls.RenegotiateNever,
		InsecureSkipVerify: c.SkipVerify,
	}
	err := loadCerts(tlsConfig)
	if err != nil {
		return nil, err
	}
	return tlsConfig, nil
}

func loadCerts(tlscfg *tls.Config) error {
	/*
		if *c.TLSCert != "" && *c.TLSKey != "" {
			certificate, err := tls.LoadX509KeyPair(*c.TLSCert, *c.TLSKey)
			if err != nil {
				return err
			}
			tlscfg.Certificates = []tls.Certificate{certificate}
			tlscfg.BuildNameToCertificate()
		}
		if c.TLSCA != nil && *c.TLSCA != "" {
			certPool := x509.NewCertPool()
			caFile, err := ioutil.ReadFile(*c.TLSCA)
			if err != nil {
				return err
			}
			if ok := certPool.AppendCertsFromPEM(caFile); !ok {
				return errors.New("failed to append certificate")
			}
			tlscfg.RootCAs = certPool
		}
	*/
	return nil
}
