// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2021 Renesas Inc.
// Copyright 2021 EPAM Systems Inc.
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

package cryptutils_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	log "github.com/sirupsen/logrus"

	"gitpct.epam.com/epmd-aepr/aos_common/utils/cryptutils"
)

/*******************************************************************************
 * Vars
 ******************************************************************************/

var tmpDir string

/*******************************************************************************
 * Init
 ******************************************************************************/

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		FullTimestamp:    true})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

/*******************************************************************************
 * Main
 ******************************************************************************/

func TestMain(m *testing.M) {
	var err error

	tmpDir, err = ioutil.TempDir("", "fileutils_")
	if err != nil {
		log.Fatalf("Error create temporary dir: %s", err)
	}

	ret := m.Run()

	if err := os.RemoveAll(tmpDir); err != nil {
		log.Fatalf("Error removing tmp dir: %s", err)
	}

	os.Exit(ret)
}

/*******************************************************************************
 * Tests
 ******************************************************************************/

func TestGetCert(t *testing.T) {
	crtName := path.Join(tmpDir, "certFile1.crt")

	crt1, err := os.Create(crtName)
	if err != nil {
		t.Fatal("Can't create file: ", err)
	}
	crt1.Close()

	crt2, err := os.Create(path.Join(tmpDir, "certFile2.crt"))
	if err != nil {
		t.Fatal("Can't create file: ", err)
	}
	crt2.Close()

	if err := os.MkdirAll(path.Join(tmpDir, "testCrtDir"), 0755); err != nil {
		t.Fatal("Can't create folder: ", err)
	}

	crtFile, err := cryptutils.GetCertFileFromDir(tmpDir)
	if err != nil {
		t.Error("Error GetCertFile: ", err)
	}

	if crtFile != crtName {
		t.Error("Incorrect crt file: ", crtFile)
	}
}

func TestGetKey(t *testing.T) {
	keyName := path.Join(tmpDir, "keyFile1.key")

	key1, err := os.Create(keyName)
	if err != nil {
		t.Fatal("Can't create file: ", err)
	}
	key1.Close()

	key2, err := os.Create(path.Join(tmpDir, "keyFile2.key"))
	if err != nil {
		t.Fatal("Can't create file: ", err)
	}
	key2.Close()

	keyFile, err := cryptutils.GetKeyFileFromDir(tmpDir)
	if err != nil {
		t.Error("Error GetKeyFile: ", err)
	}

	if keyFile != keyName {
		t.Error("Incorrect crt file: ", keyFile)
	}
}
