// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package vault_akv_plugin

import (
	"github.com/hashicorp/go-hclog"
	"testing"
)

var (
	akvClient *keyvaultClient
)

const (
	VaultName  = "anjuna-keyvault"
	SecretName = "hello"
)

func TestInitAkvClient(t *testing.T) {
	logger := hclog.New(&hclog.LoggerOptions{})
	akvClientRet, err := InitKeyvaultClient(&logger)
	if err != nil {
		t.Errorf("Failed initializing Azure Key Vault client")
	}

	akvClient = akvClientRet
}

func TestListSecrets(t *testing.T) {
	secrets, err := akvClient.ListSecrets(VaultName)
	if err != nil {
		t.Errorf("Failed listing secrets")
	}

	t.Logf("%v", secrets)
}

func TestSetSecret(t *testing.T) {
	err := akvClient.SetSecret(VaultName, SecretName, "world")
	if err != nil {
		t.Errorf("Failed setting secret")
	}
}

func TestGetSecret(t *testing.T) {
	value, err := akvClient.GetSecret(VaultName, SecretName)
	if err != nil {
		t.Errorf("Failed getting secret")
	}

	t.Logf("%s=%s", SecretName, value)
}

func TestDeleteSecret(t *testing.T) {
	err := akvClient.DeleteSecret(VaultName, SecretName)
	if err != nil {
		t.Errorf("Failed deleting secret")
	}
}
