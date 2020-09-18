package events

type SecretRotationStep string

const (
	SecretRotationStepCreateSecret = "createSecret"
	SecretRotationStepSetSecret    = "setSecret"
	SecretRotationStepTestSecret   = "testSecret"
	SecretRotationStepFinishSecret = "finishSecret"
)

type SecretManagerRotation struct {
	Step               SecretRotationStep `json:"Step"`
	SecretID           string             `json:"SecretId"`
	ClientRequestToken string             `json:"ClientRequestToken"`
}
