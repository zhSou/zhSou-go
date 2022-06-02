前端提供查询表达式，

后端返回查询结果

### 请求

POST /query

| 字段        | 类型     | 是否必填 | 描述        |
| --------- | ------ |------|-----------|
| query     | string | 是    | 待查询的表达式   |
| page      | int    | 是    | 页码编号，从1开始 |
| limit     | int    | 是    | 单页限制的数目   |
#### 实例

```json
{
    "query":"手机游戏",
    "page": 1,
    "limit": 20
}
```

### 响应

| 字段        | 类型            | 描述        |
|-----------|---------------|-----------|
| useTime   | float         | 搜索用时，单位s  |
| total     | int           | 搜索结果数     |
| totalPage | int           | 总页数       |
| page      | int           | 当前页       |
| records   | array<Record> | 当前页的所有记录项 |
对于document格式如下

| 字段       | 类型     | 描述           |
| -------- | ------ | ------------ |
| id       | int    | 文档编号         |
| text     | string | 该文档的索引建立文本区域 |
| document | object | 任意对象         |