// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

const (
	appType  = "app"
	appLabel = "app"
)

var appSecrets = map[string]string{
	"onf.cacrt":  caCert,
	"client.crt": clientCert,
	"client.key": clientKey,
}

func newApp(name string, client *client) *App {
	labels := getLabels(appType)
	labels[appLabel] = name
	return &App{
		Service: newService(name, 0, labels, "", appSecrets, []string{}, client),
	}
}

// App provides methods for adding and modifying applications
type App struct {
	*Service
}
