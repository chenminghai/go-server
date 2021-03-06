### 添加收货地址

[POST] /v1/user/address

| 参数          | 类型     | 说明                       | 必选 |
| ------------- | -------- | -------------------------- | ---- |
| name          | `string` | 收件人                     | \*   |
| phone         | `string` | 收件人手机号               | \*   |
| province_code | `string` | 省份代码，6 位数           | \*   |
| city_code     | `string` | 城市代码，6 位数           | \*   |
| area_code     | `string` | 县城代码，6 位数           | \*   |
| address       | `string` | 详细地址，具体的街道门牌号 | \*   |
| is_default    | `bool`   | 是否设置为默认地址         | \*   |

### 更新收货地址

[PUT] /v1/user/address/a/:address_id

| 参数          | 类型     | 说明                       | 必选 |
| ------------- | -------- | -------------------------- | ---- |
| name          | `string` | 收件人                     |      |
| phone         | `string` | 收件人手机号               |      |
| province_code | `string` | 省份代码，6 位数           |      |
| city_code     | `string` | 城市代码，6 位数           |      |
| area_code     | `string` | 县城代码，6 位数           |      |
| address       | `string` | 详细地址，具体的街道门牌号 |      |
| is_default    | `bool`   | 是否设置为默认地址         |      |

### 删除收货地址

[DELETE] /v1/user/address/a/:address_id

删除收货地址

### 收货地址列表

[GET] /v1/user/address

获取我的收货地址列表

### 获取默认收货地址

[GET] /v1/user/address/default

获取我的默认收货地址

### 获取地址详情

[GET] /v1/user/address/a/:address_id

获取某一个地址的详细信息

### 获取全国地区码列表

[GET] /v1/area

获取全国地区码列表

### 获取指定地区码的详情

[GET] /v1/area/:area_code

获取指定地区码的详情, `area_code` 是最详细的地区码
