// Copyright © 2020 Uzhinskiy Boris <boris.ujinsky@gmail.com>
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

package pconf

import (
	"io/ioutil"
	"log"
	"strings"
)

type ConfigType map[string]string

func (cfg ConfigType) Parse(configfile string) (err error) {
	rawBytes, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}

	text := string(rawBytes)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, ";") {
			fields := strings.Split(line, "=")
			if len(fields) == 2 {
				tmp := strings.TrimSpace(fields[1])
				j1 := strings.Index(tmp, "\"")
				j2 := strings.LastIndex(tmp, "\"")
				if j1 == 0 && j2 > 0 {
					cfg[strings.TrimSpace(fields[0])] = tmp[j1+1 : j2]
				} else if j1 == -1 && j2 == -1 {
					cfg[strings.TrimSpace(fields[0])] = tmp
				} else {
					log.Println("invalid option")
				}
			}
		}
	}
	return
}
