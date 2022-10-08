Trie Source Code
===

* 介绍|Introduction  
  * trie 树源码，目前支持完全匹配查找，前缀匹配查找，前缀包含匹配，子串包含匹配四种操作。  
  * trie source code,support perfect match,prefix start match,prefix include match and subword include match. 

-------

* 编译方式|Build  
 ```bash
 g++ char_trie.cpp 
 ```
-------

* 使用方式|Tutorial  
  * 查看main函数  
  * check the main() function

-------
* 查看树|View the tree
```cpp
Trie T;
T.Print();
```

```cpp
(root)
    └---(B)
    |   └---(r)
    |   |   └---(a)
    |   |   |   └---(d)
    |   |   |       └---(f)
    |   |   |           └---(o)
    |   |   |               └---(r)
    |   |   |                   └---(d,1)
    |   |   └---(y)
    |   |       └---(a)
    |   |           └---(n,1)
    |   └---(e)
    |   |   └---(n,1)
    |   └---(a)
    |       └---(r)
    |           └---(t)
    |               └---(h)
    |                   └---(o)
    |                       └---(l)
    |                           └---(o)
    |                               └---(m)
    |                                   └---(e)
    |                                       └---(w,1)
    └---(A)
    |   └---(l)
    |   |   └---(b)
    |   |       └---(i)
    |   |       |   └---(n,1)
    |   |       └---(e)
    |   |           └---(r)
    |   |               └---(t,1)
    |   └---(b)
    |       └---(n)
    |           └---(e)
    |               └---(r,1)
    └---(Y)
    |   └---(u)
    |       └---(a)
    |           └---(n)
    |               └---(y)
    |                   └---(e,1)
    └---(J)
        └---(a)
            └---(c)
            |   └---(k,1)
            └---(y,1)
```

