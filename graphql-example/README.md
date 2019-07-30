# GraphQL-example







## 使用步骤

**前提**

- Go 1.11 以上
- Docker

**下载安装依赖**
``` 
cd ${PROJECT} && go mod download
```

**创建数据库**

``` 
cd ${PROJECT} && bash db.sh
```

**迁移数据库表**

``` 
cd ${PROJECT} && make sync
```

**启动**

``` 
cd ${PROJECT} && make run
```

## 接口示例

```text
query {
    ping{
        data
        code
    }   
}

# 结果
{
	"data": {
		"ping": {
			"code": "200",
			"data": "pong"
		}
	}
}


```

```text

query HeartBeat($Data: String!){
    pingWithData(data: $Data){
        data
        code
    }
}
# 变量
{
	"Data":"GraphQL"
}
# 结果

{
	"data": {
		"pingWithData": {
			"code": "200",
			"data": "GraphQL"
		}
	}
}

```

```text
mutation {
    createVote(
        title: "去哪玩？",
        description:"本次团建去哪玩？",
        options:[
            {
                name: "杭州西湖"
            },{
                name:"安徽黄山"
            },{
                name:"香港九龙"
            }
            ],
        deadline: "2019-08-01 00:00:00",
        class: SINGLE
        ) {
            id
            title
            deadline
            description
            createdAt
            updatedAt
            options{
                name
            }
            class
            classString
        }
}

# 结果

{
	"data": {
		"vote": {
			"class": "SINGLE",
			"classString": "单选",
			"createdAt": "2019-07-30T19:33:27+08:00",
			"deadline": "2019-08-01T00:00:00+08:00",
			"description": "本次团建去哪玩？",
			"id": "1",
			"options": [
				{
					"name": "杭州西湖"
				},
				{
					"name": "安徽黄山"
				},
				{
					"name": "香港九龙"
				}
			],
			"title": "去哪玩？",
			"updatedAt": "2019-07-30T19:33:27+08:00"
		}
	}
}

```

```text

query{
    vote(id:1){
            id
            title
            deadline
            description
            createdAt
            updatedAt
            options{
                name
            }
            class
            classString
    }
}

# 结果

query{
    vote(id:1){
            id
            title
            deadline
            description
            createdAt
            updatedAt
            options{
                name
            }
            class
            classString
    }
}

```

```text
mutation {
    updateVote(
        id:1,
        title:"本次团建上哪玩？",
        description: "希望大家踊跃参与"
        ){
            id
            title
            deadline
            description
            createdAt
            updatedAt
            options{
                name
            }
            class
            classString
        }
}

# 结果

{
	"data": {
		"updateVote": {
			"class": "SINGLE",
			"classString": "单选",
			"createdAt": "2019-07-30T19:33:27+08:00",
			"deadline": "2019-08-01T00:00:00+08:00",
			"description": "希望大家踊跃参与",
			"id": "1",
			"options": [
				{
					"name": "杭州西湖"
				},
				{
					"name": "安徽黄山"
				},
				{
					"name": "香港九龙"
				}
			],
			"title": "本次团建上哪玩？",
			"updatedAt": "2019-07-30T20:18:19+08:00"
		}
	}
}

```