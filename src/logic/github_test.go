// Copyright 2017 The StudyGolang Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package logic_test

import (
	"logic"
	"testing"

	"github.com/polaris1119/config"
	"github.com/polaris1119/logger"
)

func TestPullPR(t *testing.T) {
	logger.Init(config.ROOT+"/log", config.ConfigFile.MustValue("global", "log_level", "DEBUG"))

	err := logic.DefaultGithub.PullPR("studygolang/GCTT", false)
	if err != nil {
		t.Error("pull request error:", err)
	}
}
