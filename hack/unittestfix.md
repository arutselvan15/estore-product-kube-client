# Unit test fix
    
    Auto generate code test will fail due to version conflict. Add below fix to get make test working.
    ```
    pkg/client/clientset/versioned/typed/estore/vi/estore_client.go
    
    // import below pkg
    "github.com/arutselvan15/estore-product-kube-client/pkg/client/clientset/versioned/scheme"
   	"k8s.io/apimachinery/pkg/runtime/serializer"
    
    // comment this line
    //config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
   
    // add this line
    config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
    ```
    Run ```make test``` to verify the unit tests.