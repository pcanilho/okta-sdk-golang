/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import ()

type Protocol struct {
	Algorithms  *ProtocolAlgorithms          `json:"algorithms,omitempty"`
	Credentials *IdentityProviderCredentials `json:"credentials,omitempty"`
	Endpoints   *ProtocolEndpoints           `json:"endpoints,omitempty"`
	Issuer      *ProtocolEndpoint            `json:"issuer,omitempty"`
	RelayState  *ProtocolRelayState          `json:"relayState,omitempty"`
	Scopes      []string                     `json:"scopes,omitempty"`
	Settings    *ProtocolSettings            `json:"settings,omitempty"`
	Type        string                       `json:"type,omitempty"`
}
