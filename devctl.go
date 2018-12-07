// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

package main

import "github.com/bcaldwell/devctl/cmd"

// Version sets command version
var Version = "dev"

// BuildDate sets the date that the current build was built
var BuildDate = "n/a"

func main() {
	// pass down build date and version number set at build time
	cmd.Version = Version
	cmd.BuildDate = BuildDate
	cmd.Execute()
}
