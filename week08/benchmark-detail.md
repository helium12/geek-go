1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

Redis Benchmark环境:
VMware® Workstation 15 Pro
Redis版本: 6.2.5
处理器： 1 core 
内存： 4.0 GB
操作系统： CentOS Linux release 7.7.1908

----------------1、redis get、set性能数据
----------1.1 redis get性能
 redis-benchmark -h 127.0.0.1 -p 6379 -t get -n 100000 -d 10
// 10万请求   10字节
Summary:
  throughput summary: 33456.00 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.013     0.336     0.663     1.031    14.095    26.623
// 10万请求   20字节
Summary:
  throughput summary: 25859.84 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.377     0.384     0.783     4.807    15.175    32.831
// 10万请求   50字节
Summary:
  throughput summary: 24850.89 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.356     0.424     0.911     1.783    14.911    22.511
// 10万请求   100字节
Summary:
  throughput summary: 37425.15 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.920     0.392     0.855     1.439     2.623    20.031
// 10万请求   200字节
Summary:
  throughput summary: 30543.68 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.107     0.392     0.831     1.335    13.063    26.975
// 10万请求   1k字节
Summary:
  throughput summary: 17793.59 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.958     0.400     0.943    11.959    19.423    31.151
// 10万请求   5k字节
Summary:
  throughput summary: 20929.26 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.719     0.480     1.151     2.687    19.407    44.639


----------1.2 redis set性能
 redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 10
// 10万请求   10字节
Summary:
  throughput summary: 38565.37 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.905     0.344     0.663     0.991    12.943    23.775
// 10万请求   20字节
Summary:
  throughput summary: 44622.94 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.780     0.352     0.695     1.047     2.759    21.231
// 10万请求   50字节
Summary:
  throughput summary: 21877.05 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.588     0.352     0.815     5.679    21.807    31.871
// 10万请求   100字节
Summary:
  throughput summary: 37341.30 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.912     0.376     0.823     1.351     2.975    21.695
// 10万请求   200字节
Summary:
  throughput summary: 46446.82 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.751     0.352     0.719     1.047     1.839     6.967
// 10万请求   1k字节
Summary:
  throughput summary: 23239.60 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.482     0.368     0.863     1.919    20.735    32.575
// 10万请求   5k字节
Summary:
  throughput summary: 22055.58 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.574     0.400     1.103     2.647    16.543    40.287

----------------2、
127.0.0.1:6379> flushdb
OK
127.0.0.1:6379> info memory
* Memory
used_memory:874792
used_memory_human:854.29K
used_memory_rss:6017024
used_memory_rss_human:5.74M
used_memory_peak:22114176
used_memory_peak_human:21.09M
used_memory_peak_perc:3.96%
used_memory_overhead:830344
used_memory_startup:809848
used_memory_dataset:44448
used_memory_dataset_perc:68.44%
allocator_allocated:869512
allocator_active:1798144
allocator_resident:4255744
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.07
allocator_frag_bytes:928632
allocator_rss_ratio:2.37
allocator_rss_bytes:2457600
rss_overhead_ratio:1.41
rss_overhead_bytes:1761280
mem_fragmentation_ratio:7.22
mem_fragmentation_bytes:5183256
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

该条件下数据基本维持在97500左右
//10万数据, 10字节 9.12M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 10 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97586
127.0.0.1:6379> info memory
* Memory
used_memory:9730712
used_memory_human:9.28M
used_memory_rss:14848000
used_memory_rss_human:14.16M
used_memory_peak:22114176
used_memory_peak_human:21.09M
used_memory_peak_perc:44.00%
used_memory_overhead:5782360
used_memory_startup:809848
used_memory_dataset:3948352
used_memory_dataset_perc:44.26%
allocator_allocated:10162688
allocator_active:10543104
allocator_resident:15908864
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.04
allocator_frag_bytes:380416
allocator_rss_ratio:1.51
allocator_rss_bytes:5365760
rss_overhead_ratio:0.93
rss_overhead_bytes:-1060864
mem_fragmentation_ratio:1.53
mem_fragmentation_bytes:5158312
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0


//10万数据, 20字节  9.36M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 20 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97510
127.0.0.1:6379> info memory
* Memory
used_memory:10506856
used_memory_human:10.02M
used_memory_rss:16408576
used_memory_rss_human:15.65M
used_memory_peak:37169328
used_memory_peak_human:35.45M
used_memory_peak_perc:28.27%
used_memory_overhead:5779320
used_memory_startup:809848
used_memory_dataset:4727536
used_memory_dataset_perc:48.75%
allocator_allocated:10937696
allocator_active:11702272
allocator_resident:17698816
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.07
allocator_frag_bytes:764576
allocator_rss_ratio:1.51
allocator_rss_bytes:5996544
rss_overhead_ratio:0.93
rss_overhead_bytes:-1290240
mem_fragmentation_ratio:1.57
mem_fragmentation_bytes:5942744
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

//10万数据, 50字节 11.1M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 50 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97510
127.0.0.1:6379> info memory
* Memory
used_memory:13627600
used_memory_human:13.00M
used_memory_rss:18919424
used_memory_rss_human:18.04M
used_memory_peak:37169328
used_memory_peak_human:35.45M
used_memory_peak_perc:36.66%
used_memory_overhead:5779320
used_memory_startup:809848
used_memory_dataset:7848280
used_memory_dataset_perc:61.23%
allocator_allocated:14061176
allocator_active:14585856
allocator_resident:20119552
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.04
allocator_frag_bytes:524680
allocator_rss_ratio:1.38
allocator_rss_bytes:5533696
rss_overhead_ratio:0.94
rss_overhead_bytes:-1200128
mem_fragmentation_ratio:1.39
mem_fragmentation_bytes:5332848
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

//10万数据, 100字节 17.9M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 100 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97585
127.0.0.1:6379> info memory
* Memory
used_memory:19101512
used_memory_human:18.22M
used_memory_rss:24866816
used_memory_rss_human:23.71M
used_memory_peak:37169328
used_memory_peak_human:35.45M
used_memory_peak_perc:51.39%
used_memory_overhead:5782320
used_memory_startup:809848
used_memory_dataset:13319192
used_memory_dataset_perc:72.82%
allocator_allocated:19531096
allocator_active:20299776
allocator_resident:25833472
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.04
allocator_frag_bytes:768680
allocator_rss_ratio:1.27
allocator_rss_bytes:5533696
rss_overhead_ratio:0.96
rss_overhead_bytes:-966656
mem_fragmentation_ratio:1.30
mem_fragmentation_bytes:5806328
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

//10万数据, 200字节 28.28M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 200 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97586
127.0.0.1:6379> info memory
* Memory
used_memory:30031472
used_memory_human:28.64M
used_memory_rss:35848192
used_memory_rss_human:34.19M
used_memory_peak:37169328
used_memory_peak_human:35.45M
used_memory_peak_perc:80.80%
used_memory_overhead:5782360
used_memory_startup:809848
used_memory_dataset:24249112
used_memory_dataset_perc:82.98%
allocator_allocated:30470464
allocator_active:31223808
allocator_resident:36712448
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.02
allocator_frag_bytes:753344
allocator_rss_ratio:1.18
allocator_rss_bytes:5488640
rss_overhead_ratio:0.98
rss_overhead_bytes:-864256
mem_fragmentation_ratio:1.20
mem_fragmentation_bytes:5857744
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
  
//10万数据, 1k字节 91.5M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 1000 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97527
127.0.0.1:6379> info memory
* Memory
used_memory:108036608
used_memory_human:103.03M
used_memory_rss:115023872
used_memory_rss_human:109.70M
used_memory_peak:111092960
used_memory_peak_human:105.95M
used_memory_peak_perc:97.25%
used_memory_overhead:5780000
used_memory_startup:809848
used_memory_dataset:102256608
used_memory_dataset_perc:95.36%
allocator_allocated:108474048
allocator_active:109252608
allocator_resident:117440512
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.01
allocator_frag_bytes:778560
allocator_rss_ratio:1.07
allocator_rss_bytes:8187904
rss_overhead_ratio:0.98
rss_overhead_bytes:-2416640
mem_fragmentation_ratio:1.07
mem_fragmentation_bytes:7028288
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

//10万数据, 5k字节 432.3M
redis-benchmark -h 127.0.0.1 -p 6379 -t set -n 100000 -d 5000 -r 2000000
127.0.0.1:6379> dbsize
(integer) 97497
127.0.0.1:6379> info memory
* Memory
used_memory:507352960
used_memory_human:483.85M
used_memory_rss:514568192
used_memory_rss_human:490.73M
used_memory_peak:510409160
used_memory_peak_human:486.76M
used_memory_peak_perc:99.40%
used_memory_overhead:5778800
used_memory_startup:809848
used_memory_dataset:501574160
used_memory_dataset_perc:99.02%
allocator_allocated:507807672
allocator_active:508555264
allocator_resident:519118848
total_system_memory:3953991680
total_system_memory_human:3.68G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.00
allocator_frag_bytes:747592
allocator_rss_ratio:1.02
allocator_rss_bytes:10563584
rss_overhead_ratio:0.99
rss_overhead_bytes:-4550656
mem_fragmentation_ratio:1.01
mem_fragmentation_bytes:7256256
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0