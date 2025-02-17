//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

// KeyVaultClientBackupKeyResponse contains the response from method KeyVaultClient.BackupKey.
type KeyVaultClientBackupKeyResponse struct {
	BackupKeyResult
}

// KeyVaultClientCreateKeyResponse contains the response from method KeyVaultClient.CreateKey.
type KeyVaultClientCreateKeyResponse struct {
	KeyBundle
}

// KeyVaultClientDecryptResponse contains the response from method KeyVaultClient.Decrypt.
type KeyVaultClientDecryptResponse struct {
	KeyOperationResult
}

// KeyVaultClientDeleteKeyResponse contains the response from method KeyVaultClient.DeleteKey.
type KeyVaultClientDeleteKeyResponse struct {
	DeletedKeyBundle
}

// KeyVaultClientEncryptResponse contains the response from method KeyVaultClient.Encrypt.
type KeyVaultClientEncryptResponse struct {
	KeyOperationResult
}

// KeyVaultClientGetDeletedKeyResponse contains the response from method KeyVaultClient.GetDeletedKey.
type KeyVaultClientGetDeletedKeyResponse struct {
	DeletedKeyBundle
}

// KeyVaultClientGetDeletedKeysResponse contains the response from method KeyVaultClient.GetDeletedKeys.
type KeyVaultClientGetDeletedKeysResponse struct {
	DeletedKeyListResult
}

// KeyVaultClientGetKeyResponse contains the response from method KeyVaultClient.GetKey.
type KeyVaultClientGetKeyResponse struct {
	KeyBundle
}

// KeyVaultClientGetKeyRotationPolicyResponse contains the response from method KeyVaultClient.GetKeyRotationPolicy.
type KeyVaultClientGetKeyRotationPolicyResponse struct {
	KeyRotationPolicy
}

// KeyVaultClientGetKeyVersionsResponse contains the response from method KeyVaultClient.GetKeyVersions.
type KeyVaultClientGetKeyVersionsResponse struct {
	KeyListResult
}

// KeyVaultClientGetKeysResponse contains the response from method KeyVaultClient.GetKeys.
type KeyVaultClientGetKeysResponse struct {
	KeyListResult
}

// KeyVaultClientGetRandomBytesResponse contains the response from method KeyVaultClient.GetRandomBytes.
type KeyVaultClientGetRandomBytesResponse struct {
	RandomBytes
}

// KeyVaultClientImportKeyResponse contains the response from method KeyVaultClient.ImportKey.
type KeyVaultClientImportKeyResponse struct {
	KeyBundle
}

// KeyVaultClientPurgeDeletedKeyResponse contains the response from method KeyVaultClient.PurgeDeletedKey.
type KeyVaultClientPurgeDeletedKeyResponse struct {
	// placeholder for future response values
}

// KeyVaultClientRecoverDeletedKeyResponse contains the response from method KeyVaultClient.RecoverDeletedKey.
type KeyVaultClientRecoverDeletedKeyResponse struct {
	KeyBundle
}

// KeyVaultClientReleaseResponse contains the response from method KeyVaultClient.Release.
type KeyVaultClientReleaseResponse struct {
	KeyReleaseResult
}

// KeyVaultClientRestoreKeyResponse contains the response from method KeyVaultClient.RestoreKey.
type KeyVaultClientRestoreKeyResponse struct {
	KeyBundle
}

// KeyVaultClientRotateKeyResponse contains the response from method KeyVaultClient.RotateKey.
type KeyVaultClientRotateKeyResponse struct {
	KeyBundle
}

// KeyVaultClientSignResponse contains the response from method KeyVaultClient.Sign.
type KeyVaultClientSignResponse struct {
	KeyOperationResult
}

// KeyVaultClientUnwrapKeyResponse contains the response from method KeyVaultClient.UnwrapKey.
type KeyVaultClientUnwrapKeyResponse struct {
	KeyOperationResult
}

// KeyVaultClientUpdateKeyResponse contains the response from method KeyVaultClient.UpdateKey.
type KeyVaultClientUpdateKeyResponse struct {
	KeyBundle
}

// KeyVaultClientUpdateKeyRotationPolicyResponse contains the response from method KeyVaultClient.UpdateKeyRotationPolicy.
type KeyVaultClientUpdateKeyRotationPolicyResponse struct {
	KeyRotationPolicy
}

// KeyVaultClientVerifyResponse contains the response from method KeyVaultClient.Verify.
type KeyVaultClientVerifyResponse struct {
	KeyVerifyResult
}

// KeyVaultClientWrapKeyResponse contains the response from method KeyVaultClient.WrapKey.
type KeyVaultClientWrapKeyResponse struct {
	KeyOperationResult
}
