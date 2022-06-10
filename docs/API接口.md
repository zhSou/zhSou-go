## 文档查询
前端提供查询表达式，

后端返回查询结果

### 请求

POST /query

| 字段          | 类型                               | 描述      |
|-------------|----------------------------------|---------|
| query       | string                           | 待查询的表达式 |
| page        | int                              | 页码编号，从1开始 |
| limit       | int                              | 单页限制的数目 |
| filterWord  | string[]                         | 过滤词，也可空数组 |
| highLight | {preTag: string, postTag:string} | 高亮显示信息  |
#### 示例

```json
{
  "query": "游戏",
  "page":1,
  "limit":10,
  "filterWord":["手机"],
  "highLight": {
    "preTag": "<span style='color:red'>",
    "postTag": "</span>"
  }
}
```

### 响应

| 字段        | 类型       | 描述        |
|-----------|----------|-----------|
| useTime   | float    | 搜索用时，单位s  |
| total     | int      | 搜索结果数     |
| totalPage | int      | 总页数       |
| page      | int      | 当前页       |
| records   | Record[] | 当前页的所有记录项 |


对于Record格式如下

| 字段       | 类型                         | 描述                   |
| -------- |----------------------------|----------------------|
| id       | int                        | 文档编号                 |
| text     | string                     | 该文档的索引建立文本区域，会拼接高亮信息 |
| document | {url:string, text: string} | 原始文档对象               |

#### 示例
```json
{
	"code": 0,
	"msg": "Success",
	"data": {
		"useTime": 0.0062808,
		"total": 1870,
		"page": 1,
		"totalPage": 187,
		"records": [
			{
				"id": 124,
				"text": "电商平台销售无限制<span style='color:red'>游戏</span>账号就算家长不给孩子开设<span style='color:red'>游戏</span>账号,万能的某",
				"document": {
					"url": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fpic3.zhimg.com%2Fv2-cec8bd65270b88d933d2a33172eb2b3a_r.jpg&refer=http%3A%2F%2Fpic3.zhimg.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1632564339&t=41f92622009306972603b0280faa8c4c",
					"text": "电商平台销售无限制游戏账号就算家长不给孩子开设游戏账号,万能的某"
				}
			},
			{
				"id": 977,
				"text": "<span style='color:red'>游戏</span>本体关于宇宙竞技场",
				"document": {
					"url": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg.nga.178.com%2Fattachments%2Fmon_202108%2F25%2F-39t2Qd0gp-c6moZdT1kShs-13i.jpg&refer=http%3A%2F%2Fimg.nga.178.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1632584688&t=7a7eb1adb81d6a7b93ce21da14d4acc7",
					"text": "游戏本体关于宇宙竞技场"
				}
			},
            ...
        ]
    }
}
```

## 读取文档
### 请求

POST /getDocuments

给定文档编号json数组，返回文档获取结果，若部分id有误，则忽略错误id
#### 请求示例
```json
[1,4,7,-1]
```

### 响应
```json
{
  "code": 0,
  "msg": "Success",
  "data": {
    "documents": {
      "1": {
        "url": "https://gss0.baidu.com/70cFfyinKgQFm2e88IuM_a/forum/w=580/sign=ef25202540a7d933bfa8e47b9d4ad194/0b1ba6efce1b9d161d6c8050f1deb48f8d5464b1.jpg",
        "text": "13/14赛季 英超第5轮 曼城 vs 曼联 13.09.22"
      },
      "4": {
        "url": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fnimg.ws.126.net%2F%3Furl%3Dhttp%253A%252F%252Fdingyue.ws.126.net%252F2021%252F0825%252F36524b12j00qye3c7001pc000ku009ig.jpg%26thumbnail%3D650x2147483647%26quality%3D80%26type%3Djpg&refer=http%3A%2F%2Fnimg.ws.126.net&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1632584681&t=a649dc4d1af9bd6d539aad5a75e93810",
        "text": "一是持续加大\"双一流\"建设支持力度."
      },
      "7": {
        "url": "https://gimg2.baidu.com/image_search/src=http%3A%2F%2F5b0988e595225.cdn.sohucs.com%2Fimages%2F20180309%2F7e0cb3c963b24803bbd19abe8667e62f.jpg&refer=http%3A%2F%2F5b0988e595225.cdn.sohucs.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1632565715&t=de3752c650cb9a87ceaaaa9ecb737f62",
        "text": "今秋起湖北高一新生不分文理科,高考录取总成绩由\"3 3"
      }
    }
  }
}
```