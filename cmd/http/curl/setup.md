# Setup Curl
## Find Product
```bash
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"product1-f12e-4fa0-bf4f-dd002c11a671","img_url":"https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80","name":"Product A", "price":20000}}' \
https://localhost:49002/v1/product

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"product2-f12e-4fa0-bf4f-dd002c11a671","img_url":"https://images.unsplash.com/photo-1520342868574-5fa3804e551c?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=6ff92caffcdd63681a35134a6770ed3b&auto=format&fit=crop&w=1951&q=80","name":"Product A", "price":20000}}' \
https://localhost:49002/v1/product


curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"group001-f12e-4fa0-bf4f-ba002c11a670","name":"Group A", "description":"Group A description","products":[{"id":"product1-f12e-4fa0-bf4f-dd002c11a671"},{"id":"product2-f12e-4fa0-bf4f-dd002c11a671"}]}}' \
https://localhost:49002/v1/product/group
```
