package controllers

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	hoopsv1alpha1 "github.com/pauldotyu/hoops-operator/api/v1alpha1"
)

// when creating a standings resource it should create a new configmap
var _ = Describe("Standings Controller", func() {
	const (
		standingsName      = "test-standings"
		standingsNamespace = "default"
		dataSourceUrl      = "https://www.basketball-reference.com/leagues/NBA_2021_standings.html"
		configMapName      = "test-standings-configmap"
		configMapKey       = "standings"
		timeout            = time.Second * 10
		interval           = time.Millisecond * 250
	)

	Context("When creating a new Standings resource", func() {
		BeforeEach(func() {
			// create a new Standings resource
			standings := &hoopsv1alpha1.Standings{
				ObjectMeta: metav1.ObjectMeta{
					Name:      standingsName,
					Namespace: standingsNamespace,
				},
				Spec: hoopsv1alpha1.StandingsSpec{
					DataSourceUrl: dataSourceUrl,
					ConfigMapName: configMapName,
					ConfigMapKey:  configMapKey,
				},
			}
			Expect(k8sClient.Create(ctx, standings)).Should(Succeed())
		})
		It("should create a new configmap", func() {
			// create a new configmap
			configMap := &corev1.ConfigMap{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, client.ObjectKey{Name: configMapName, Namespace: standingsNamespace}, configMap)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
