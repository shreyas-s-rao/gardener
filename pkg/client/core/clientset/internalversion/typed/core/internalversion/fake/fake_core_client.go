// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	internalversion "github.com/gardener/gardener/pkg/client/core/clientset/internalversion/typed/core/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCore struct {
	*testing.Fake
}

func (c *FakeCore) BackupBuckets() internalversion.BackupBucketInterface {
	return &FakeBackupBuckets{c}
}

func (c *FakeCore) BackupEntries(namespace string) internalversion.BackupEntryInterface {
	return &FakeBackupEntries{c, namespace}
}

func (c *FakeCore) CloudProfiles() internalversion.CloudProfileInterface {
	return &FakeCloudProfiles{c}
}

func (c *FakeCore) ControllerDeployments() internalversion.ControllerDeploymentInterface {
	return &FakeControllerDeployments{c}
}

func (c *FakeCore) ControllerInstallations() internalversion.ControllerInstallationInterface {
	return &FakeControllerInstallations{c}
}

func (c *FakeCore) ControllerRegistrations() internalversion.ControllerRegistrationInterface {
	return &FakeControllerRegistrations{c}
}

func (c *FakeCore) ExposureClasses() internalversion.ExposureClassInterface {
	return &FakeExposureClasses{c}
}

func (c *FakeCore) InternalSecrets(namespace string) internalversion.InternalSecretInterface {
	return &FakeInternalSecrets{c, namespace}
}

func (c *FakeCore) Projects() internalversion.ProjectInterface {
	return &FakeProjects{c}
}

func (c *FakeCore) Quotas(namespace string) internalversion.QuotaInterface {
	return &FakeQuotas{c, namespace}
}

func (c *FakeCore) SecretBindings(namespace string) internalversion.SecretBindingInterface {
	return &FakeSecretBindings{c, namespace}
}

func (c *FakeCore) Seeds() internalversion.SeedInterface {
	return &FakeSeeds{c}
}

func (c *FakeCore) Shoots(namespace string) internalversion.ShootInterface {
	return &FakeShoots{c, namespace}
}

func (c *FakeCore) ShootStates(namespace string) internalversion.ShootStateInterface {
	return &FakeShootStates{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCore) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
