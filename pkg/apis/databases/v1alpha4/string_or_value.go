package v1alpha4

import "fmt"

func (ir *InlineOrRef) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var inline string
	err := unmarshal(&inline)
	if err == nil {
		ir.Value = inline
		return nil
	}

	type otherType struct {
		ValueFrom *ValueFrom `yaml:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
	}
	var t otherType
	err = unmarshal(&t)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ir.copyValueFrom(t.ValueFrom)

	return nil
}

func (ir *InlineOrRef) copyValueFrom(valueFrom *ValueFrom) {
	ir.ValueFrom = &ValueFrom{}

	if valueFrom.SecretKeyRef != nil {
		ir.ValueFrom.SecretKeyRef = &SecretKeyRef{
			Name: valueFrom.SecretKeyRef.Name,
			Key:  valueFrom.SecretKeyRef.Key,
		}
	}
}

type InlineOrRef struct {
	Value     string     `json:"-" yaml:"-"`
	ValueFrom *ValueFrom `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}

type ValueFrom struct {
	SecretKeyRef *SecretKeyRef `json:"secretKeyRef,omitempty" yaml:"secretKeyRef,omitempty"`
	Vault        *Vault        `json:"vault,omitempty" yaml:"vault,omitempty"`
}

type SecretKeyRef struct {
	Name string `json:"name" yaml:"name"`
	Key  string `json:"key" yaml:"key"`
}

type Vault struct {
	AgentInject bool   `json:"agentInject,omitempty" yaml:"agentInject,omitempty"`
	Secret      string `json:"secret" yaml:"secret"`
	Role        string `json:"role" yaml:"role"`

	Endpoint                string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	ServiceAccount          string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
	ServiceAccountNamespace string `json:"serviceAccountNamespace,omitempty" yaml:"serviceAccountNamespace,omitempty"`
}
