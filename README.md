## 知识罗列
### 1、外键引起的死锁。
    
> accounts.id 被 B 表作为外键使用，
> B 表操作（如insert）时候，由于要使用 accounts.id 外键，
> 导致 accounts 该行加锁，
> 如果此时 B 表要 select account 时候发现了死锁。

> 解决办法：B 表 insert 后面的操作，如果不影响 accounts 的 id 的话，
> 可以 通知 accounts 我不需要改变你的 id，你不需要加锁。
> 但是：如果 2个事务 操作 insert accounts 还是会发生死锁。 
    
    ```sql
    SELECT * FROM accounts
    WHERE id = $1 LIMIT 1
    FOR NO KEY UPDATE;
    ```

### 2、并发的问题。
