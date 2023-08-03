// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shoot

//var _ = Describe("Shoot Tests", Label("Shoot", "high-availability", "upgrade-to-zone"), func() {
//	test := func(shoot *gardencorev1beta1.Shoot) {
//		f := defaultShootCreationFramework()
//		f.Shoot = shoot
//
//		f.Shoot.Spec.ControlPlane = nil
//
//		It("Create, Upgrade (non-HA to HA with failure tolerance type 'zone') and Delete Shoot", Offset(1), func() {
//			By("Create Shoot")
//			ctx, cancel := context.WithTimeout(parentCtx, 30*time.Minute)
//			defer cancel()
//
//			Expect(f.CreateShootAndWaitForCreation(ctx, false)).To(Succeed())
//			f.Verify()
//
//			By("Upgrade Shoot (non-HA to HA with failure tolerance type 'zone')")
//			ctx, cancel = context.WithTimeout(parentCtx, 30*time.Minute)
//			defer cancel()
//			highavailability.UpgradeAndVerify(ctx, f.ShootFramework, v1beta1.FailureToleranceTypeZone)
//
//			By("Delete Shoot")
//			ctx, cancel = context.WithTimeout(parentCtx, 20*time.Minute)
//			defer cancel()
//			Expect(f.DeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())
//		})
//	}
//
//	Context("Shoot with workers", func() {
//		test(e2e.DefaultShoot("e2e-upd-zone"))
//	})
//
//	Context("Workerless Shoot", Label("workerless"), func() {
//		test(e2e.DefaultWorkerlessShoot("e2e-upd-zone"))
//	})
//})
