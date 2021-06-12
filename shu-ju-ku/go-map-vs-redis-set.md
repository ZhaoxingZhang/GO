---
description: Quite similar
---

# Go Map VS Redis Set

## TL;DR

| struct | map | \[\]buckets | bucket | \#kv | load\_factor | 扩容 | 缩容 |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| GO Map | hmap | buckets, oldbuckets | bmap | 8 | 6.5 | 2/1.25 | - |
| Redis Set | dict | dictht\[0\], dictht\[1\] | dictEntry | 1 | 1/5/0.1 | 2 | 0.5 |

## [Go Map](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/)

![overall](https://img-blog.csdnimg.cn/20210612122201856.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXNtYW4y,size_16,color_FFFFFF,t_70)

![tophash](https://img-blog.csdnimg.cn/20210612122235294.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXNtYW4y,size_16,color_FFFFFF,t_70)

* 负载因子: load\_factor = 6.5
* 碰撞：key冲突 -&gt; overflow
* 增量扩容 ：load\_factor&gt;6.5，&lt;1024翻倍或者1/4
* 等量扩容：overflow&gt;max\(2^15, key\)，空洞太多
* 动态扩容：2个bucket，迁移old-new时rehash，标记key是否已经迁移
* map不执行缩容

## Redis Set

![dictEntry&#x540E;&#x9762;&#x53EF;&#x4EE5;&#x63A5;&#x4E0B;&#x4E00;&#x4E2A;&#x8282;&#x70B9;](https://img-blog.csdnimg.cn/20210612121849636.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXNtYW4y,size_16,color_FFFFFF,t_70)

* load\_factor =dictht\[0\].use/dictht\[1\].size 1\(没有AOF/BGSAVE\)/ 5
* 动态扩容：翻倍；迁移0-1时rehash，标记key是否已经迁移；完成后0，1标记交换
* 动态缩容：load\_factor&lt;0.1是减半

