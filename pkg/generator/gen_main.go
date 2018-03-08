// Copyright 2018 The Operator-SDK Authors
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

package generator

import (
	"io"
	"path/filepath"
	"text/template"
)

const (
	// sdkImport is the operator-sdk import path.
	sdkImport = "github.com/coreos/operator-sdk/pkg/sdk"
)

// Main contains all the customized data needed to generate cmd/<projectName>/main.go for a new operator
// when pairing with mainTmpl template.
type Main struct {
	// imports
	OperatorSDKImport string
	StubImport        string
}

// renderMainFile generates the cmd/<projectName>/main.go file given a repo path ("github.com/coreos/play")
func renderMainFile(w io.Writer, repo string) error {
	t := template.New("cmd/<projectName>/main.go")
	t, err := t.Parse(mainTmpl)
	if err != nil {
		return err
	}

	m := Main{
		OperatorSDKImport: sdkImport,
		StubImport:        filepath.Join(repo, stubDir),
	}
	return t.Execute(w, m)
}
