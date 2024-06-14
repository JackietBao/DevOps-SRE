### 加载外部配置文件

#### Input config

```yaml
filebeat.config.inputs:
  enabled: true
  path: inputs.d/*.yml
```

`inputs.d`目录下的配置文件示例：

```yaml
- type: log
  paths:
    - /var/log/mysql.log
  scan_frequency: 10s

- type: log
  paths:
    - /var/log/apache.log
  scan_frequency: 5s
```

#### Module config

```bash
filebeat.config.modules:
  enabled: true
  path: ${path.config}/modules.d/*.yml
```



外部配置文件示例：

```yaml
- module: apache
  access:
    enabled: true
    var.paths: [/var/log/apache2/access.log*]
  error:
    enabled: true
    var.paths: [/var/log/apache2/error.log*]
```



### Elasticsearch output

#### 配置模板加载

默认情况下，如果启用了Elasticsearch输出，Filebeat会自动加载推荐的模板文件`fields.yml`，如果要使用默认索引模板，则不需要其他配置，否则，你可以将`filebeat.yml`配置文件中的默认值更改为：



-  **加载不同的模板** 

```plain
setup.template.name: "your_template_name"
setup.template.fields: "path/to/fields.yml"
```


如果模板已存在，则除非你配置Filebeat，否则不会覆盖该模板。 

-  **覆盖现有模板** 

```plain
setup.template.overwrite: true
```

 

-  **禁用自动模板加载** 

```plain
setup.template.enabled: false
```


如果禁用自动模板加载，则需要手动加载模板。 

-  **更改索引名称**
   如果要将事件发送到支持索引生命周期管理的集群，请参阅配置索引生命周期管理以了解如何更改索引名称。
   默认情况下，当禁用或不支持索引生命周期管理时，Filebeat使用时间系列索引，索引名为`filebeat-7.3.0-yyyy.MM.dd`，其中`yyyy.MM.dd`是事件索引的日期，要使用其他名称，请在Elasticsearch输出中设置`index`选项。你指定的值应包括索引的根名称以及版本和日期信息，你还需要配置`setup.template.name`和`setup.template.pattern`选项以匹配新名称，例如： 

```plain
output.elasticsearch.index: "customname-%{[agent.version]}-%{+yyyy.MM.dd}"
setup.template.name: "customname"
setup.template.pattern: "customname-*"
```


如果你使用的是预先构建的Kibana仪表板，请同时设置`setup.dashboards.index`选项，例如：  



#### 自定义索引名



##### 配置filebeat



如果我们不配置则默认会生成 ，如下类格式的索引，且如果检测到有的话，会默认一直使用这个日期



```bash
# ---------------------------- Elasticsearch Output ----------------------------
output.elasticsearch:
  hosts: ["https://myEShost:9200"]
  
# 输出到es的索引格式
filebeat-8.0.0-rc1-2022-01-20
```



从上面的配置中我们可以看到，数据会送到ES里去，但是只会送到ES的filebeat-*的索引里去，这显然不是我们想要的。



使用默认索引时要写入事件的索引名称。默认为`“filebeat-%{[agent.version]}-%{+yyyy.MM.dd}”`，例如“filebeat-8.0.0-rc1-2022-01-20”。如果更改此设置，还需要配置 `setup.template.name` 和 `setup.template.pattern` 选项。



要想送到指定索引，我们对filebeat.yml的更改配置如下：



```bash
# ============================== Filebeat inputs ===============================
filebeat.inputs:
- type: log
  paths:
    - /var/log/nginx/access*
  fields:
    source: access
- type: log
  paths:
    - /var/log/nginx/error*
  fields:
    source: error
# ============================== Filebeat modules ==============================
filebeat.config.modules:
  enabled: true
  path: ${path.config}/modules.d/*.yml
setup.template.enabled: false
setup.template.name: "nginx"
setup.template.pattern: "nginx-*"
setup.template.overwrite: true
setup.ilm.enabled: false
#  ============================== Elasticsearch Output ==============================
output.elasticsearch:
  hosts: ["192.168.197.130:9200"]
  # 这里的index前缀nginx与模板的pattern匹配，中间这一串设置为field.source变量，如果下面indices格式没有匹配上，则使用该index格式
  index: "nginx-%{[fields.source]}-*"  
  indices:
  # 这里的前缀nginx同为与模板的pattern匹配，中间为field.source具体的值，当前面的input的field.source值与这里的匹配时，则index设置为定义的格式
    - index: "nginx-access-%{[agent.version]}-%{+yyyy.MM.dd}" 
      when.equals:                                            
      fields:
          source: "access"
    - index: "nginx-error-%{[agent.version]}-%{+yyyy.MM.dd}"
      when.equals:
        fields:
          source: "error"
# ============================== Processors ==============================
processors:
  - add_host_metadata: ~
  - add_cloud_metadata: ~
```



```bash
when.contains: 包含
when.equals: 等于
```



```bash
# 相关模板字段意义：
setup.template.name: “nginx”      # 设置一个新的模板，模板的名称
setup.template.pattern: “nginx-*” # 模板匹配那些索引，这里表示以nginx开头的所有的索引
setup.template.enabled: false     # 关掉默认的模板配置
setup.template.overwrite: true    # 开启新设置的模板
setup.ilm.enabled: false      # 索引生命周期管理ilm功能默认开启，开启的情况下索引名称只能为filebeat-*， 通过setup.ilm.enabled: false进行关闭；如果要使用自定义的索引名称，同时又需要启用ilm，可以修改filebeat的模板
```



##### 查看es是否增加了新的索引



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125101754863.png)



##### 在kibana上关联es索引



点击【Managerment】—【 Stack Management】



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125102933262.png)



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125113301351.png)



索引名使用通配符的形式，这样以后索引的 `nginx-access` 开头的索引库收集来的日志可以聚合展示了，否则每个月甚至每天都需要创建索引



可以看到在选择字段的时候也比原来的要少，因为是自定义的模板，因此字段只有我们定义的内容，默认的模板把所有支持的字段都会加上



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125113456292.png)



创建成功



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125113614934.png)



进行查看



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125113659346.png)



![img](https://gitee.com/salted-egg-yolk-pie/ima/raw/master/imgs/image-20220125113712249.png)