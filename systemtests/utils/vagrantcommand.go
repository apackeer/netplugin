/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"os"
	"os/exec"
)

type TestCommand interface {
	Run(cmd string, args ...string) error
	RunWithOutput(cmd string, args ...string) ([]byte, error)
}

type VagrantCommand struct {
	ContivNodes int
	ContivEnv   string
}

func (c *VagrantCommand) getCmd(cmd string, args ...string) *exec.Cmd {
	newArgs := append([]string{cmd}, args...)
	osCmd := exec.Command("vagrant", newArgs...)
	osCmd.Env = os.Environ()
	if c.ContivNodes != 0 {
		osCmd.Env = append(osCmd.Env, fmt.Sprintf("CONTIV_NODES=%d", c.ContivNodes))
	}
	if c.ContivEnv != "" {
		osCmd.Env = append(osCmd.Env, fmt.Sprintf("CONTIV_ENV=%s", c.ContivEnv))
	}

	return osCmd
}

func (c *VagrantCommand) Run(cmd string, args ...string) error {
	return c.getCmd(cmd, args...).Run()
}

func (c *VagrantCommand) RunWithOutput(cmd string, args ...string) ([]byte, error) {
	return c.getCmd(cmd, args...).Output()
}
