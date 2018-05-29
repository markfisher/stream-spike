/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fake

import (
	v1alpha1 "github.com/scothis/stream-spike/pkg/client/clientset/versioned/typed/spike.local/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSpikeV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSpikeV1alpha1) Brokers(namespace string) v1alpha1.BrokerInterface {
	return &FakeBrokers{c, namespace}
}

func (c *FakeSpikeV1alpha1) Streams(namespace string) v1alpha1.StreamInterface {
	return &FakeStreams{c, namespace}
}

func (c *FakeSpikeV1alpha1) Subscriptions(namespace string) v1alpha1.SubscriptionInterface {
	return &FakeSubscriptions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSpikeV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
