package cmd

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mrShow(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command("../lab_bin", "mr", "1")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Log(string(b))
		t.Error(err)
	}

	require.Contains(t, string(b), `
#1 Test MR for lab list
===================================
This MR is to remain open for testing the `+"`lab mr list`"+` functionality
-----------------------------------
Project: alerque/test
Branches: mrtest->master
Status: Open
Assignee: alerque
Author: alerque
Milestone: 1.0
Labels: documentation
WebURL: https://gitlab.com/alerque/test/merge_requests/1
`)
}
