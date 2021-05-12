/*
 *
 * Copyright 2021 The Vitess Authors.
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
 * /
 */

package server

import (
	"errors"
	"io/ioutil"
	"path"

	"github.com/go-git/go-git/v5"
)

// setupLocalVitess is used to setup the local clone of vitess
func (s *Server) setupLocalVitess() error {
	files, err := ioutil.ReadDir(s.localVitessPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Name() == "vitess" && file.IsDir() {
			return nil
		}
	}
	_, err = git.PlainClone(s.getVitessPath(), false, &git.CloneOptions{
		URL: "https://github.com/vitessio/vitess",
	})
	return err
}

// getVitessPath is used to find the path of the directory where the local vitess clone should exits
func (s *Server) getVitessPath() string {
	return path.Join(s.localVitessPath, "vitess")
}

// fetchLocalVitess is used to execute
func (s *Server) fetchLocalVitess() error {
	r, err := git.PlainOpen(s.getVitessPath())
	if err != nil {
		return err
	}
	err = r.Fetch(&git.FetchOptions{
		Tags: git.AllTags,
	})
	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			return nil
		}
		return err
	}
	return nil
}
