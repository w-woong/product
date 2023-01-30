# Product API
## Find Product
```base
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
https://localhost:49002/v1/product/product1-f12e-4fa0-bf4f-ba002c11a671
```

## Find Products by group id
```base
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X GET \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
'https://localhost:49002/v1/product/group/group001-f12e-4fa0-bf4f-ba002c11a671'
```