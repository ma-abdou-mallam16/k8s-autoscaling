package mock

import (
	"crypto/rsa"
	"math/big"
	"net/http"
	"net/http/httptest"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/client/metadata"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/session"
)

// Session is a mock session which is used to hit the mock server
var Session = func() *session.Session {
	// server is the mock server that simply writes a 200 status back to the client
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	return session.Must(session.NewSession(&aws.Config{
		DisableSSL: aws.Bool(true),
		Endpoint:   aws.String(server.URL),
	}))
}()

// NewMockClient creates and initializes a client that will connect to the
// mock server
func NewMockClient(cfgs ...*aws.Config) *client.Client {
	c := Session.ClientConfig("Mock", cfgs...)

	svc := client.New(
		*c.Config,
		metadata.ClientInfo{
			ServiceName:   "Mock",
			SigningRegion: c.SigningRegion,
			Endpoint:      c.Endpoint,
			APIVersion:    "2015-12-08",
			JSONVersion:   "1.1",
			TargetPrefix:  "MockServer",
		},
		c.Handlers,
	)

	return svc
}

// RSAPrivateKey is used for testing functionality that requires some
// sort of private key. Taken from crypto/rsa/rsa_test.go
//
// Credit to golang 1.11
var RSAPrivateKey = &rsa.PrivateKey{
	PublicKey: rsa.PublicKey{
		N: fromBase10("14314132931241006650998084889274020608918049032671858325988396851334124245188214251956198731333464217832226406088020736932173064754214329009979944037640912127943488972644697423190955557435910767690712778463524983667852819010259499695177313115447116110358524558307947613422897787329221478860907963827160223559690523660574329011927531289655711860504630573766609239332569210831325633840174683944553667352219670930408593321661375473885147973879086994006440025257225431977751512374815915392249179976902953721486040787792801849818254465486633791826766873076617116727073077821584676715609985777563958286637185868165868520557"),
		E: 3,
	},
	D: fromBase10("9542755287494004433998723259516013739278699355114572217325597900889416163458809501304132487555642811888150937392013824621448709836142886006653296025093941418628992648429798282127303704957273845127141852309016655778568546006839666463451542076964744073572349705538631742281931858219480985907271975884773482372966847639853897890615456605598071088189838676728836833012254065983259638538107719766738032720239892094196108713378822882383694456030043492571063441943847195939549773271694647657549658603365629458610273821292232646334717612674519997533901052790334279661754176490593041941863932308687197618671528035670452762731"),
	Primes: []*big.Int{
		fromBase10("130903255182996722426771613606077755295583329135067340152947172868415809027537376306193179624298874215608270802054347609836776473930072411958753044562214537013874103802006369634761074377213995983876788718033850153719421695468704276694983032644416930879093914927146648402139231293035971427838068945045019075433"),
		fromBase10("109348945610485453577574767652527472924289229538286649661240938988020367005475727988253438647560958573506159449538793540472829815903949343191091817779240101054552748665267574271163617694640513549693841337820602726596756351006149518830932261246698766355347898158548465400674856021497190430791824869615170301029"),
	},
}

// Taken from crypto/rsa/rsa_test.go
//
// Credit to golang 1.11
func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}
