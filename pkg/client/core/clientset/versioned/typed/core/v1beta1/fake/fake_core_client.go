// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/gardener/gardener/pkg/client/core/clientset/versioned/typed/core/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1beta1 struct {
	*testing.Fake
}

func (c *FakeCoreV1beta1) BackupBuckets() v1beta1.BackupBucketInterface {
	return &FakeBackupBuckets{c}
}

func (c *FakeCoreV1beta1) BackupEntries(namespace string) v1beta1.BackupEntryInterface {
	return &FakeBackupEntries{c, namespace}
}

func (c *FakeCoreV1beta1) CloudProfiles() v1beta1.CloudProfileInterface {
	return &FakeCloudProfiles{c}
}

func (c *FakeCoreV1beta1) ControllerDeployments() v1beta1.ControllerDeploymentInterface {
	return &FakeControllerDeployments{c}
}

func (c *FakeCoreV1beta1) ControllerInstallations() v1beta1.ControllerInstallationInterface {
	return &FakeControllerInstallations{c}
}

func (c *FakeCoreV1beta1) ControllerRegistrations() v1beta1.ControllerRegistrationInterface {
	return &FakeControllerRegistrations{c}
}

func (c *FakeCoreV1beta1) ExposureClasses() v1beta1.ExposureClassInterface {
	return &FakeExposureClasses{c}
}

func (c *FakeCoreV1beta1) InternalSecrets(namespace string) v1beta1.InternalSecretInterface {
	return &FakeInternalSecrets{c, namespace}
}

func (c *FakeCoreV1beta1) Projects() v1beta1.ProjectInterface {
	return &FakeProjects{c}
}

func (c *FakeCoreV1beta1) Quotas(namespace string) v1beta1.QuotaInterface {
	return &FakeQuotas{c, namespace}
}

func (c *FakeCoreV1beta1) SecretBindings(namespace string) v1beta1.SecretBindingInterface {
	return &FakeSecretBindings{c, namespace}
}

func (c *FakeCoreV1beta1) Seeds() v1beta1.SeedInterface {
	return &FakeSeeds{c}
}

func (c *FakeCoreV1beta1) Shoots(namespace string) v1beta1.ShootInterface {
	return &FakeShoots{c, namespace}
}

func (c *FakeCoreV1beta1) ShootStates(namespace string) v1beta1.ShootStateInterface {
	return &FakeShootStates{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
