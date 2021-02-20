// This file is generated by "./lib/proto/generate"

package proto

/*

WebAuthn

This domain allows configuring virtual authenticators to test the WebAuthn
API.

*/

// WebAuthnAuthenticatorID ...
type WebAuthnAuthenticatorID string

// WebAuthnAuthenticatorProtocol ...
type WebAuthnAuthenticatorProtocol string

const (
	// WebAuthnAuthenticatorProtocolU2f enum const
	WebAuthnAuthenticatorProtocolU2f WebAuthnAuthenticatorProtocol = "u2f"

	// WebAuthnAuthenticatorProtocolCtap2 enum const
	WebAuthnAuthenticatorProtocolCtap2 WebAuthnAuthenticatorProtocol = "ctap2"
)

// WebAuthnCtap2Version ...
type WebAuthnCtap2Version string

const (
	// WebAuthnCtap2VersionCtap20 enum const
	WebAuthnCtap2VersionCtap20 WebAuthnCtap2Version = "ctap2_0"

	// WebAuthnCtap2VersionCtap21 enum const
	WebAuthnCtap2VersionCtap21 WebAuthnCtap2Version = "ctap2_1"
)

// WebAuthnAuthenticatorTransport ...
type WebAuthnAuthenticatorTransport string

const (
	// WebAuthnAuthenticatorTransportUsb enum const
	WebAuthnAuthenticatorTransportUsb WebAuthnAuthenticatorTransport = "usb"

	// WebAuthnAuthenticatorTransportNfc enum const
	WebAuthnAuthenticatorTransportNfc WebAuthnAuthenticatorTransport = "nfc"

	// WebAuthnAuthenticatorTransportBle enum const
	WebAuthnAuthenticatorTransportBle WebAuthnAuthenticatorTransport = "ble"

	// WebAuthnAuthenticatorTransportCable enum const
	WebAuthnAuthenticatorTransportCable WebAuthnAuthenticatorTransport = "cable"

	// WebAuthnAuthenticatorTransportInternal enum const
	WebAuthnAuthenticatorTransportInternal WebAuthnAuthenticatorTransport = "internal"
)

// WebAuthnVirtualAuthenticatorOptions ...
type WebAuthnVirtualAuthenticatorOptions struct {

	// Protocol ...
	Protocol WebAuthnAuthenticatorProtocol `json:"protocol"`

	// Ctap2Version (optional) Defaults to ctap2_0. Ignored if |protocol| == u2f.
	Ctap2Version WebAuthnCtap2Version `json:"ctap2Version,omitempty"`

	// Transport ...
	Transport WebAuthnAuthenticatorTransport `json:"transport"`

	// HasResidentKey (optional) Defaults to false.
	HasResidentKey bool `json:"hasResidentKey,omitempty"`

	// HasUserVerification (optional) Defaults to false.
	HasUserVerification bool `json:"hasUserVerification,omitempty"`

	// HasLargeBlob (optional) If set to true, the authenticator will support the largeBlob extension.
	// https://w3c.github.io/webauthn#largeBlob
	// Defaults to false.
	HasLargeBlob bool `json:"hasLargeBlob,omitempty"`

	// AutomaticPresenceSimulation (optional) If set to true, tests of user presence will succeed immediately.
	// Otherwise, they will not be resolved. Defaults to true.
	AutomaticPresenceSimulation bool `json:"automaticPresenceSimulation,omitempty"`

	// IsUserVerified (optional) Sets whether User Verification succeeds or fails for an authenticator.
	// Defaults to false.
	IsUserVerified bool `json:"isUserVerified,omitempty"`
}

// WebAuthnCredential ...
type WebAuthnCredential struct {

	// CredentialID ...
	CredentialID []byte `json:"credentialId"`

	// IsResidentCredential ...
	IsResidentCredential bool `json:"isResidentCredential"`

	// RpID (optional) Relying Party ID the credential is scoped to. Must be set when adding a
	// credential.
	RpID string `json:"rpId,omitempty"`

	// PrivateKey The ECDSA P-256 private key in PKCS#8 format.
	PrivateKey []byte `json:"privateKey"`

	// UserHandle (optional) An opaque byte sequence with a maximum size of 64 bytes mapping the
	// credential to a specific user.
	UserHandle []byte `json:"userHandle,omitempty"`

	// SignCount Signature counter. This is incremented by one for each successful
	// assertion.
	// See https://w3c.github.io/webauthn/#signature-counter
	SignCount int `json:"signCount"`

	// LargeBlob (optional) The large blob associated with the credential.
	// See https://w3c.github.io/webauthn/#sctn-large-blob-extension
	LargeBlob []byte `json:"largeBlob,omitempty"`
}

// WebAuthnEnable Enable the WebAuthn domain and start intercepting credential storage and
// retrieval with a virtual authenticator.
type WebAuthnEnable struct {
}

// ProtoReq of the command
func (m WebAuthnEnable) ProtoReq() string { return "WebAuthn.enable" }

// Call of the command, sessionID is optional.
func (m WebAuthnEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnDisable Disable the WebAuthn domain.
type WebAuthnDisable struct {
}

// ProtoReq of the command
func (m WebAuthnDisable) ProtoReq() string { return "WebAuthn.disable" }

// Call of the command, sessionID is optional.
func (m WebAuthnDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnAddVirtualAuthenticator Creates and adds a virtual authenticator.
type WebAuthnAddVirtualAuthenticator struct {

	// Options ...
	Options *WebAuthnVirtualAuthenticatorOptions `json:"options"`
}

// ProtoReq of the command
func (m WebAuthnAddVirtualAuthenticator) ProtoReq() string { return "WebAuthn.addVirtualAuthenticator" }

// Call of the command, sessionID is optional.
func (m WebAuthnAddVirtualAuthenticator) Call(c Client) (*WebAuthnAddVirtualAuthenticatorResult, error) {
	var res WebAuthnAddVirtualAuthenticatorResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// WebAuthnAddVirtualAuthenticatorResult Creates and adds a virtual authenticator.
type WebAuthnAddVirtualAuthenticatorResult struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`
}

// WebAuthnRemoveVirtualAuthenticator Removes the given authenticator.
type WebAuthnRemoveVirtualAuthenticator struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`
}

// ProtoReq of the command
func (m WebAuthnRemoveVirtualAuthenticator) ProtoReq() string {
	return "WebAuthn.removeVirtualAuthenticator"
}

// Call of the command, sessionID is optional.
func (m WebAuthnRemoveVirtualAuthenticator) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnAddCredential Adds the credential to the specified authenticator.
type WebAuthnAddCredential struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`

	// Credential ...
	Credential *WebAuthnCredential `json:"credential"`
}

// ProtoReq of the command
func (m WebAuthnAddCredential) ProtoReq() string { return "WebAuthn.addCredential" }

// Call of the command, sessionID is optional.
func (m WebAuthnAddCredential) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnGetCredential Returns a single credential stored in the given virtual authenticator that
// matches the credential ID.
type WebAuthnGetCredential struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`

	// CredentialID ...
	CredentialID []byte `json:"credentialId"`
}

// ProtoReq of the command
func (m WebAuthnGetCredential) ProtoReq() string { return "WebAuthn.getCredential" }

// Call of the command, sessionID is optional.
func (m WebAuthnGetCredential) Call(c Client) (*WebAuthnGetCredentialResult, error) {
	var res WebAuthnGetCredentialResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// WebAuthnGetCredentialResult Returns a single credential stored in the given virtual authenticator that
// matches the credential ID.
type WebAuthnGetCredentialResult struct {

	// Credential ...
	Credential *WebAuthnCredential `json:"credential"`
}

// WebAuthnGetCredentials Returns all the credentials stored in the given virtual authenticator.
type WebAuthnGetCredentials struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`
}

// ProtoReq of the command
func (m WebAuthnGetCredentials) ProtoReq() string { return "WebAuthn.getCredentials" }

// Call of the command, sessionID is optional.
func (m WebAuthnGetCredentials) Call(c Client) (*WebAuthnGetCredentialsResult, error) {
	var res WebAuthnGetCredentialsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// WebAuthnGetCredentialsResult Returns all the credentials stored in the given virtual authenticator.
type WebAuthnGetCredentialsResult struct {

	// Credentials ...
	Credentials []*WebAuthnCredential `json:"credentials"`
}

// WebAuthnRemoveCredential Removes a credential from the authenticator.
type WebAuthnRemoveCredential struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`

	// CredentialID ...
	CredentialID []byte `json:"credentialId"`
}

// ProtoReq of the command
func (m WebAuthnRemoveCredential) ProtoReq() string { return "WebAuthn.removeCredential" }

// Call of the command, sessionID is optional.
func (m WebAuthnRemoveCredential) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnClearCredentials Clears all the credentials from the specified device.
type WebAuthnClearCredentials struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`
}

// ProtoReq of the command
func (m WebAuthnClearCredentials) ProtoReq() string { return "WebAuthn.clearCredentials" }

// Call of the command, sessionID is optional.
func (m WebAuthnClearCredentials) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnSetUserVerified Sets whether User Verification succeeds or fails for an authenticator.
// The default is true.
type WebAuthnSetUserVerified struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`

	// IsUserVerified ...
	IsUserVerified bool `json:"isUserVerified"`
}

// ProtoReq of the command
func (m WebAuthnSetUserVerified) ProtoReq() string { return "WebAuthn.setUserVerified" }

// Call of the command, sessionID is optional.
func (m WebAuthnSetUserVerified) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// WebAuthnSetAutomaticPresenceSimulation Sets whether tests of user presence will succeed immediately (if true) or fail to resolve (if false) for an authenticator.
// The default is true.
type WebAuthnSetAutomaticPresenceSimulation struct {

	// AuthenticatorID ...
	AuthenticatorID WebAuthnAuthenticatorID `json:"authenticatorId"`

	// Enabled ...
	Enabled bool `json:"enabled"`
}

// ProtoReq of the command
func (m WebAuthnSetAutomaticPresenceSimulation) ProtoReq() string {
	return "WebAuthn.setAutomaticPresenceSimulation"
}

// Call of the command, sessionID is optional.
func (m WebAuthnSetAutomaticPresenceSimulation) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}