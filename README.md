# go-study
Collect the test codes for daily learning go language.

# MD学习
## 基础语法
现在大家开始练习**着重**符号的练习
熟悉一下*斜体*的格式
突出 __*加粗和斜体*__ 的方式

> 注意：这里是需要特别关注的地方——块引用
>> 1. 第一项要求
>> 1. 第三项要求
>> 8. 第二项要求

无序排列示例
- 无序345
- 无序990

![雨绮弟弟](https://tse3-mm.cn.bing.net/th/id/OIP-C.C2dJajHedhhpvh9TFZCPHQHaLG?w=200&h=300&c=7&o=5&pid=1.7)

![花千骨](https://tse4-mm.cn.bing.net/th/id/OIP-C.DTqVEY3Llm1MSHsCLnL6mAHaLH?w=200&h=300&c=7&o=5&pid=1.7)

![本地图片](typora.png)


---

## 超链接演示
- 我的GitHub地址:*https://github.com/shanyongsy*

- [我的源代码](https://github.com/shanyongsy)

## 扩展语法-表格
| 任务 | 开始时间 | 结束时间 | 责任人 |
| --- | :---: | ---: | --- |
| 项目开发 | 2021年01月 | 2021年2月20日 | 张三 |
| 项目开发 | *2021* | **2021** | ![雨绮弟弟](https://tse3-mm.cn.bing.net/th/id/OIP-C.C2dJajHedhhpvh9TFZCPHQHaLG?w=200&h=300&c=7&o=5&pid=1.7) |

## 代码显示
```c++
bool
CivetHandler::handleGet(CivetServer *server, struct mg_connection *conn)
{
	UNUSED_PARAMETER(server);
	UNUSED_PARAMETER(conn);
	return false;
}

bool
CivetHandler::handleGet(CivetServer *server,
                        struct mg_connection *conn,
                        int *status_code)
{
	UNUSED_PARAMETER(server);
	UNUSED_PARAMETER(conn);
	if (status_code) {
		*status_code = -1;
	}
	return false;
}
```

## 脚注
刘[^1]关[^2]张[^3]三兄弟来到徐州。
[^1]: 刘备.
[^2]: 关羽.
[^3]: 张飞.

## 定义清单
四大古国
: 具有悠久历史和文化，以及对周边进行长时期意识影响的国家。

计划任务
: ~~基本语法.~~
: 扩展语法

## 任务清单
- [x] Write the press release
- [ ] Update the website
- [ ] Contact the media

## 禁用URL链接
`https://github.com/shanyongsy`
