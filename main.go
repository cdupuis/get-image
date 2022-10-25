/*
 * Copyright © 2022 Docker, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
)

type ImageId struct {
	name string
}

func (i ImageId) Context() name.Repository {
	return name.Repository{}
}

func (i ImageId) Identifier() string {
	return i.name
}

func (i ImageId) Name() string {
	return i.name
}

func (i ImageId) Scope(s string) string {
	return ""
}

func (i ImageId) String() string {
	return i.name
}

func main() {
	image := os.Args[1]
	img, err := daemon.Image(ImageId{name: image})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(img.ConfigName())
}
