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

//var _ = Describe("Shoot Tests", Label("Shoot", "default"), func() {
//	test := func(shoot *gardencorev1beta1.Shoot) {
//		f := defaultShootCreationFramework()
//		f.Shoot = shoot
//
//		// explicitly use one version below the latest supported minor version so that Kubernetes version update test can be
//		// performed
//		f.Shoot.Spec.Kubernetes.Version = "1.26.0"
//
//		if !v1beta1helper.IsWorkerless(f.Shoot) {
//			// create two additional worker pools which explicitly specify the kubernetes version
//			pool1 := f.Shoot.Spec.Provider.Workers[0]
//			pool2, pool3 := pool1.DeepCopy(), pool1.DeepCopy()
//			pool2.Name += "2"
//			pool2.Kubernetes = &gardencorev1beta1.WorkerKubernetes{Version: &f.Shoot.Spec.Kubernetes.Version}
//			pool3.Name += "3"
//			pool3.Kubernetes = &gardencorev1beta1.WorkerKubernetes{Version: pointer.String("1.25.4")}
//			f.Shoot.Spec.Provider.Workers = append(f.Shoot.Spec.Provider.Workers, *pool2, *pool3)
//		}
//
//		It("Create, Update, Delete", Label("simple"), Offset(1), func() {
//			By("Create Shoot")
//			ctx, cancel := context.WithTimeout(parentCtx, 30*time.Minute)
//			defer cancel()
//			Expect(f.CreateShootAndWaitForCreation(ctx, false)).To(Succeed())
//			f.Verify()
//
//			var (
//				shootClient kubernetes.Interface
//				err         error
//			)
//			By("Verify shoot access using admin kubeconfig")
//			Eventually(func(g Gomega) {
//				shootClient, err = access.CreateShootClientFromAdminKubeconfig(ctx, f.GardenClient, f.Shoot)
//				g.Expect(err).NotTo(HaveOccurred())
//
//				g.Expect(shootClient.Client().List(ctx, &corev1.NamespaceList{})).To(Succeed())
//			}).Should(Succeed())
//
//			if !v1beta1helper.IsWorkerless(f.Shoot) {
//				By("Verify worker node labels")
//				commonNodeLabels := utils.MergeStringMaps(f.Shoot.Spec.Provider.Workers[0].Labels)
//				commonNodeLabels["networking.gardener.cloud/node-local-dns-enabled"] = "false"
//				commonNodeLabels["node.kubernetes.io/role"] = "node"
//
//				Eventually(func(g Gomega) {
//					for _, workerPool := range f.Shoot.Spec.Provider.Workers {
//						expectedNodeLabels := utils.MergeStringMaps(commonNodeLabels)
//						expectedNodeLabels["worker.gardener.cloud/pool"] = workerPool.Name
//						expectedNodeLabels["worker.gardener.cloud/cri-name"] = string(workerPool.CRI.Name)
//						expectedNodeLabels["worker.gardener.cloud/system-components"] = strconv.FormatBool(workerPool.SystemComponents.Allow)
//
//						kubernetesVersion := f.Shoot.Spec.Kubernetes.Version
//						if workerPool.Kubernetes != nil && workerPool.Kubernetes.Version != nil {
//							kubernetesVersion = *workerPool.Kubernetes.Version
//						}
//						expectedNodeLabels["worker.gardener.cloud/kubernetes-version"] = kubernetesVersion
//
//						nodeList := &corev1.NodeList{}
//						g.Expect(shootClient.Client().List(ctx, nodeList, client.MatchingLabels{
//							"worker.gardener.cloud/pool": workerPool.Name,
//						})).To(Succeed())
//						g.Expect(nodeList.Items).To(HaveLen(1), "worker pool %s should have exactly one Node", workerPool.Name)
//
//						for key, value := range expectedNodeLabels {
//							g.Expect(nodeList.Items[0].Labels).To(HaveKeyWithValue(key, value), "worker pool %s should have expected labels", workerPool.Name)
//						}
//					}
//				}).Should(Succeed())
//			}
//
//			By("Update Shoot")
//			ctx, cancel = context.WithTimeout(parentCtx, 20*time.Minute)
//			defer cancel()
//			shootupdatesuite.RunTest(ctx, &framework.ShootFramework{
//				GardenerFramework: f.GardenerFramework,
//				Shoot:             f.Shoot,
//			}, nil, nil)
//
//			By("Add skip readiness annotation")
//			ctx, cancel = context.WithTimeout(parentCtx, 10*time.Minute)
//			defer cancel()
//			Expect(f.ShootFramework.UpdateShoot(ctx, func(shoot *gardencorev1beta1.Shoot) error {
//				metav1.SetMetaDataAnnotation(&shoot.ObjectMeta, "shoot.gardener.cloud/skip-readiness", "")
//				// Use maintain operation to also execute tasks in the reconcile flow which are only performed during maintenance.
//				metav1.SetMetaDataAnnotation(&shoot.ObjectMeta, "gardener.cloud/operation", "maintain")
//				return nil
//			})).To(Succeed())
//
//			By("Wait for operation annotation to be gone (meaning controller picked up reconciliation request)")
//			Eventually(func(g Gomega) {
//				shoot := &gardencorev1beta1.Shoot{
//					ObjectMeta: metav1.ObjectMeta{
//						Name:      f.Shoot.Name,
//						Namespace: f.Shoot.Namespace,
//					},
//				}
//
//				g.Expect(f.GetShoot(ctx, shoot)).To(Succeed())
//				g.Expect(shoot.Annotations).ToNot(HaveKey("gardener.cloud/operation"))
//			}).Should(Succeed())
//
//			Expect(f.WaitForShootToBeReconciled(ctx, f.Shoot)).To(Succeed())
//			Expect(f.Shoot.Annotations).ToNot(HaveKey("shoot.gardener.cloud/skip-readiness"))
//
//			By("Delete Shoot")
//			ctx, cancel = context.WithTimeout(parentCtx, 20*time.Minute)
//			defer cancel()
//			Expect(f.DeleteShootAndWaitForDeletion(ctx, f.Shoot)).To(Succeed())
//		})
//	}
//
//	Context("Shoot with workers", Label("basic"), func() {
//		test(e2e.DefaultShoot("e2e-default"))
//	})
//
//	Context("Workerless Shoot", Label("workerless"), func() {
//		test(e2e.DefaultWorkerlessShoot("e2e-default"))
//	})
//})
