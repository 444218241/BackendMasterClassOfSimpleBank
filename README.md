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

### 2、Git&Dockerfile。

    git:
        - `git checkout -b ft/docker`
        - `git add .`
        - `git commit -m "add dockerfile"`
        - `git push origin ft/docker`
    docker:
        - build images: `docker build -t simplebank:latest .`
        - run container: `docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest`
            - 由于cotainer内调用localhost:5432会报错：的error：dial tcp 127.0.0.1:5432: connect: connection refused
            - 解决办法：
                - `docker run --name simplebank -p 8080:8080 -e GIN_MODE=release -e db_source="postgresql://root:secret@172.17.0.3:5432/simple_bank?sslmode=disable" simplebank:latest`
                - 设置network
                    1. `docker network ls` # 
                    2. `docker network inspect bridge` # 发现所有容器已经在 bridge 网络下了。
                    3. `docker network create bank-network` # 
                    4. `docker network connect bank-network 某个容器Name` # 
                    5. `docker network inspect bank-network` # 发现某个容器已经在 bank-network 下了。
                    6. `docker container inspect 某个容器Name` # 发现某个容器已经在 bank-network、bridge 两个网络下了。
                    7. `docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e db_source="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" simplebank:latest` # 重启指定network和要使用的容器Name。


### 3、并发的问题。
