/*
 * update_status_test.go
 *
 * This source file is part of the FoundationDB open source project
 *
 * Copyright 2020 Apple Inc. and the FoundationDB project authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controllers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("update_status", func() {
	Context("when env var is set with 1", func() {
		It("should return 1", func() {
			instance := FdbInstance{
				Pod: &corev1.Pod{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{{
							Env: []corev1.EnvVar{
								{
									Name:  "STORAGE_SERVERS_PER_POD",
									Value: "1",
								},
							},
						}},
					},
				},
			}

			storageServersPerPod, err := getStorageServersPerPodForInstance(&instance)
			Expect(err).NotTo(HaveOccurred())
			Expect(storageServersPerPod).To(Equal(1))
		})
	})

	Context("when env var is set with 2", func() {
		It("should return 2", func() {
			instance := FdbInstance{
				Pod: &corev1.Pod{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{{
							Env: []corev1.EnvVar{
								{
									Name:  "STORAGE_SERVERS_PER_POD",
									Value: "2",
								},
							},
						}},
					},
				},
			}

			storageServersPerPod, err := getStorageServersPerPodForInstance(&instance)
			Expect(err).NotTo(HaveOccurred())
			Expect(storageServersPerPod).To(Equal(2))
		})
	})

	Context("when env var is unset", func() {
		It("should return 1", func() {
			instance := FdbInstance{
				Pod: &corev1.Pod{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{{
							Env: []corev1.EnvVar{},
						}},
					},
				},
			}

			storageServersPerPod, err := getStorageServersPerPodForInstance(&instance)
			Expect(err).NotTo(HaveOccurred())
			Expect(storageServersPerPod).To(Equal(1))
		})
	})
})
