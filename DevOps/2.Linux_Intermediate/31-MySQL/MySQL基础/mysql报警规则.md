```shell
groups:
- name: MySQLStatsAlert
  rules:
  # Mysql服务不可用
  - alert: Mysql服务不可用
    expr: mysql_up == 0
    for: 1m
    labels:
      severity: critical
    annotations:
      summary: "MySQL服务不可用"
      description: "MySQL数据库已经停止运行，请立即采取行动！"
  # MySQL CPU 使用率
  - alert: MySQL CPU 使用率过高
    expr: (sum by (instance) (rate(mysql_global_status_cpu_time_total{mode="user"}[2m])) * 100) > 80
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL CPU 使用率过高"
      description: "MySQL CPU 使用率过高，请关注数据库性能问题。"
  # MySQL 内存使用率
  - alert: MySQL 内存使用率过高
    expr: (sum(mysql_global_status_memory_used) / sum(mysql_global_variables_buffer_pool_size)) > 0.8
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL 内存使用率过高"
      description: "MySQL 内存使用率过高，请关注数据库性能问题。"
  # MySQL 磁盘空间使用率
  - alert: MySQL 磁盘空间使用率过高
    expr: (node_filesystem_avail{mountpoint="/var/lib/mysql"} * 100 / node_filesystem_size{mountpoint="/var/lib/mysql"}) < 10
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL 磁盘空间使用率过高"
      description: "MySQL 磁盘空间使用率过高，请关注磁盘空间问题。"
  # MySQL 占用连接数
  - alert: MySQL 占用连接数超过80%
    expr: mysql_global_status_max_used_connections > mysql_global_variables_max_connections * 0.8
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL占用连接数超过80%"
      description: "MySQL已使用的连接数已经超过最大连接数的80%。请考虑增加最大连接数。"
  # MySQL 打开文件数
  - alert: MySQL打开文件数超过阈值
    expr: mysql_global_status_innodb_num_open_files > (mysql_global_variables_open_files_limit) * 0.75
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL打开文件数超过阈值"
      description: "MySQL已打开的文件数已经超过阈值。请考虑增加open_files_limit参数的值。"
  # MySQL 读取缓冲区大小
  - alert: MySQL读取缓冲区大小超过最大允许包大小
    expr: mysql_global_variables_read_buffer_size > mysql_global_variables_slave_max_allowed_packet 
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL读取缓冲区大小超过最大允许包大小"
      description: "读取缓冲区大小（read_buffer_size）超过了最大允许包大小（max_allowed_packet）。这可能导致复制出现问题。"
  # MySQL 排序缓冲区大小
  - alert: MySQL排序缓冲区大小配置可能有误
    expr: mysql_global_variables_innodb_sort_buffer_size < 256*1024 or mysql_global_variables_read_buffer_size > 4*1024*1024 
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL排序缓冲区大小配置可能有误"
      description: "排序缓冲区大小可能过大或过小。推荐的sort_buffer_size值为256k到4M之间。"
  # MySQL 线程栈大小
  - alert: MySQL线程栈大小过小
    expr: mysql_global_variables_thread_stack < 196608
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL线程栈大小过小"
      description: "线程栈大小过小，可能会在使用存储语言结构时出现问题。通常的thread_stack_size值为256k。"
  # MySQL InnoDB 强制恢复模式
  - alert: MySQL InnoDB强制恢复模式已启用
    expr: mysql_global_variables_innodb_force_recovery != 0 
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL InnoDB强制恢复模式已启用"
      description: "MySQL InnoDB强制恢复模式已启用。该模式仅用于数据恢复目的，禁止写入数据。"
  # MySQL InnoDB 日志文件大小
  - alert: MySQL InnoDB日志文件大小过小
    expr: mysql_global_variables_innodb_log_file_size < 16777216 
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL InnoDB日志文件大小过小"
      description: "MySQL InnoDB日志文件大小可能过小。选择一个较小的InnoDB日志文件大小可能会对性能产生重大影响。"
  # MySQL InnoDB 刷新日志选项
  - alert: MySQL InnoDB在提交事务时刷新日志未启用
    expr: mysql_global_variables_innodb_flush_log_at_trx_commit != 1
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "MySQL InnoDB在提交事务时刷新日志未启用"
      description: "MySQL InnoDB在提交事务时未启用刷新日志选项。在默认情况下，该选项应该为1。"
  # MySQL 慢查询
  - alert: MySQL 慢查询
    expr: rate(mysql_slow_queries_total[1m]) > 5
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "MySQL 慢查询"
      description: "MySQL存在大量慢查询，请检查是否需要优化查询语句。"
#mysql服务挂掉时触发告警
  - alert: Mysql状态
    expr: mysql_up == 0
    for: 10s
    labels:
      severity: warning
    annotations:
      summary: ' {{ $labels.instance }} Mysql服务 '
      description: " {{ $labels.instance }} Mysql服务不可用  请检查"
#mysql主从IO线程停止时触发告警
  - alert: Mysql主从IO线程检测
    expr: mysql_slave_status_slave_io_running == 0
    for: 5s
    labels:
      severity: error
    annotations:
      summary: " {{ $labels.instance }} Mysql从节点IO线程"
      description: "Mysql主从IO线程故障，请检测"
#mysql主从sql线程停止时触发告警
  - alert: Mysql主从sql线程检测
    expr: mysql_slave_status_slave_sql_running == 0
    for: 5s
    labels:
      severity: error
    annotations:
      summary: "{{ $labels.instance }} Mysql从节点sql线程"
      description: "Mysql主从sql线程故障，请检测"
#MySQL主从复制延迟较高
  - alert: MySQL主从复制延迟较高
    expr: mysql_slave_seconds_behind_master > 1
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "MySQL主从复制延迟较高"
      description: "MySQL主从复制延迟为 {{ $value }} 秒"
```

