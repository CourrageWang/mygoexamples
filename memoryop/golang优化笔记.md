### golang 优化笔记
>> https://segmentfault.com/a/1190000005020158  参考雨痕大神博客

* slice 代替rray做法有待商榷，对于一些短小的对象，复制成本远远小于在堆上分配回收内存
array函数会在栈上操作，而slice则需要makeslice 继而在堆上分配内存。

* map 优化 map会按需扩张，但是必须付出数据拷贝重新哈希成本，应该尽可能预设足够容量空间，对于小对象直接将数据交由map存储，远比用指针(value 为指针类型)高效，不但减少了堆内存的分配，垃圾回收器不会扫描非指针类型的key/value，int key 要比string key更快，map不会收缩不再使用的空间，就算把健值删除了，仍然保留内存空间待以后使用，可以如下释放
```go
   dctmap := make(map[int]string, capacity)
   dctmap = nil // 释放当前map对象
```
* 即便是panic后 defer依然会执行

* 