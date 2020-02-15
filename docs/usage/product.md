
```yaml
apiVersion: estore/v1
kind: Product
metadata:
  finalizers:
  - product.estore.k8s.io/operator
  name: test-product
  namespace: ""
spec:
  displayName: TestProduct
  description: Test product
  brand: brandX
  price: 100.00
  categories:
  - health
status:
  conditions:
  - lastTransitionTime: 2017-12-27T07:00:46Z
    message: product available in inventory
    reason: ProductAvailableInInventory
    status: "True"
    type: InventoryCheck
  - lastTransitionTime: 2017-12-27T07:00:46Z
    message: product offers applied
    reason: ProductOffersApplied
    status: "True"
    type: ApplyOffer
  - lastTransitionTime: 2017-12-27T07:00:46Z
    message: product is posting ready status
    reason: ProductReady
    status: "True"
    type: Ready
  currentStatus:
    lastUpdateTime: 2017-12-27T07:00:06Z
    phase: Available
  lastOperation:
    description: Product is now ready
    lastUpdateTime: 2017-12-27T07:00:06Z
    state: Successful
    type: Create
```