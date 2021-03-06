// Copyright 2017 Raymond Jelierse
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"google.golang.org/api/admin/directory/v1"
	"log"
)

type organizations []*admin.UserOrganization

func parseOrganizations(o interface{}) organizations {
	bytes, err := json.Marshal(o)
	if err != nil {
		log.Fatalf("Could not encode organizations: %v", err)
	}

	var orgs organizations

	err = json.Unmarshal(bytes, &orgs)
	if err != nil {
		log.Fatalf("Cound not decode organizations: %v", err)
	}

	return orgs
}

func (orgs organizations) Primary() *admin.UserOrganization {
	for _, org := range orgs {
		if org.Primary {
			return org
		}
	}

	return nil
}
