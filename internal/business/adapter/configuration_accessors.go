// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"fmt"
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/core"
)

func (c *Configuration) GetKeycloak() *core.Keycloak {
	if c.Auth != nil && c.Auth.Keycloak != nil {
		return c.Auth.Keycloak
	}
	return nil
}

func (c *Configuration) GetKeycloakReactApp() *core.KeycloakReactApp {
	if c.Auth != nil && c.Auth.Keycloak != nil && c.Auth.Keycloak.KeycloakReactApp != nil {
		return c.Auth.Keycloak.KeycloakReactApp
	}
	return nil
}

func (c *Configuration) GetAccountURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.Manager.AccountEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Account URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAnalyticURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.Manager.AnalyticEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Analytic URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAPIURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.Manager.APIEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAuthURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.Manager.AuthEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Auth URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetManagerURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.Manager.ManagerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Manager URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAuthType() string {
	if c.Auth == nil || c.Auth.Type == "horusec" {
		return ""
	}

	return c.Auth.Type
}
