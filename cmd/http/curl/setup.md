# Setup Curl

## Find Product

```bash
# Add Product
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"product1-f12e-4fa0-bf4f-dd002c11a671","img_url":"https://localhost:49006/v1/resource/image/product/41_detail_094.jpg","name":"ALL NEW 골드브이랑 갈바닉 LED마사지기 LED VRANG", "price":198000}}' \
https://localhost:49002/v1/product

curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"product2-f12e-4fa0-bf4f-dd002c11a671","img_url":"https://localhost:49006/v1/resource/image/product/41_detail_175.jpg","name":"ALL NEW 레드브이랑 갈바닉 LED마사지기 LED VRANG", "price":198000}}' \
https://localhost:49002/v1/product

# Add Product group
curl --insecure -H "Content-Type: application/json; charset=utf-8" \
-X POST \
-H 'Authorization: Bearer ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b' \
-d '{"document":{"id":"group001-f12e-4fa0-bf4f-ba002c11a670","name":"베스트", "description":"제이뷰티코리아의 베스트 상품을 만나보세요","products":[{"id":"product1-f12e-4fa0-bf4f-dd002c11a671"},{"id":"product2-f12e-4fa0-bf4f-dd002c11a671"}]}}' \
https://localhost:49002/v1/product/group
```
