## Hello


### 目前项目架构说明

> 后端    

|类别|名称|功能
|-|-|-
|框架|beego|负责提供静态访问,路由,逻辑处理,与数据库的交互
|数据库|mysql|数据持久化
|缓存|rethinkdb|


> 前端

|类别|名称|功能
|-|-|-
|包管理|npm|负责第三方控件的管理
|框架|react(+redux)|实时页面的展示
|工具|babel|
|工具|webpack|


> Install

`npm install --save react`

`npm install --save react-dom`

`npm install --save-dev webpack`

`npm install --save-dev babel-loader`

`npm install --save-dev babel-core`

`npm install --save-dev babel-preset-es2015`

`npm install --save-dev babel-preset-react`
