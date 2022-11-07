# redis-examples-golang

Referências:

https://redis.io/docs
https://redis.com/ebook/redis-in-action/


## STRING

SET chave valor
GET chave

SET chave valor2
GET chav

SETNX chave valor3
GET chave

SETNX chave2 valor4

MGET chave chaave2

INCR count-views
GET count-views

INCR count-views
GET count-views

INCRBY count-views 10
GET count-views

INCRBY count-views -5
GET count-views

## LISTS

RPUSH lista-teste valor1
RPUSH lista-teste valor2
RPUSH lista-teste valor3
LPOP lista-teste
RPOP lista-teste

LLEN lista-teste


RPUSH lista-teste a b c d e f g
LTRIM lista-teste 0 2
LLEN lista-teste

BLPOP lista-teste 5

## SETS

SADD user:123:favorites 0001
SADD user:123:favorites 0002
SADD user:123:favorites 0003
SADD user:456:favorites 0001

SISMEMBER user:123:favorites 0001
SISMEMBER user:456:favorites 0002

SINTER user:123:favorites user:456:favorites

SCARD user:123:favorites

## HASHES

HSET user:123 name Miguel email lmiguel@redventures.com role back-end
HGET user:123 name
HGETALL user:123
HMGET name email

HINCRBY user:123 deploys 1
HINCRBY user:123 deploys 1
HGET user:123 deploys

## SORTED SETS

ZADD placar 25 user:123
ZADD placar 30 user:456
ZADD placar 80 user:789

ZRANGE placar 0 1 REV WITHSCORES

ZRANK placar user:123
ZRANK placar user:789
ZRANK placar user:456

ZREVRANK placar user:123

## STREAMS

XADD create-lead * lead_id 001 source www.google.com
XADD create-lead * lead_id 002 source www.google.com
XADD create-lead * lead_id 003 source www.google.com

XRANGE create-lead - +
XLEN create-lead

XREAD COUNT 2 BLOCK 5000 STREAMS create-lead $
XADD create-lead * lead_id 0005