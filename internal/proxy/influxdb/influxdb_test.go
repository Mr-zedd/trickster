/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package influxdb

import (
	"testing"

	cr "github.com/Comcast/trickster/internal/cache/registration"
	"github.com/Comcast/trickster/internal/config"
)

func TestConfiguration(t *testing.T) {
	oc := config.OriginConfig{Type: "TEST"}
	client := Client{Config: oc}
	c := client.Configuration()
	if c.Type != "TEST" {
		t.Errorf("expected %s got %s", "TEST", c.Type)
	}
}

func TestCacheInstance(t *testing.T) {

	err := config.Load("trickster", "test", nil)
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}

	cr.LoadCachesFromConfig()
	cache, err := cr.GetCache("default")
	if err != nil {
		t.Error(err)
	}
	client := Client{Cache: cache}
	c := client.CacheInstance()

	if c.Configuration().Type != "memory" {
		t.Errorf("expected %s got %s", "memory", c.Configuration().Type)
	}
}

func TestOriginName(t *testing.T) {

	client := Client{Name: "TEST"}
	c := client.OriginName()

	if c != "TEST" {
		t.Errorf("expected %s got %s", "TEST", c)
	}

}