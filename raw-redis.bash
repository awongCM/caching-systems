# Starting Redis Server

redis-cli

#String values
SET newkey ofsomevalue
GET newkey

SET newkey newval nx
SET newkey newval xx

# Binary values
SET pages:about "about us"
GET pages:about

# Counter values
SET counter 100
INCR counter
INCR counter
INCR counter 50

# MSET/MGET 
MSET a 10 b 20 c 30
MGET a b c

# altering querying key
set mykey hello
exists mykey
del mykey
exists mykey
type mykey

# set timeout
set key somevalue
expire key 5
get key
get key

set key 100 ex 10
ttl key

# LIST OPERATIONS
RPUSH MYLIST a
RPUSH MYLIST b
LPUSH MYLIST first
LRANGE MYLIST 0 -1

RPUSH MYLIST 1 2 3 4 5 "foo bar"
LRANGE MYLIST 0 -1

RPUSH NUMBERLIST 1 2 3
RPOP NUMBERLIST
RPOP NUMBERLIST
RPOP NUMBERLIST

# CAPPED LIST
RPUSH MYLIST 1 2 3 4 5
LTRIM MYLIST 0 2
LRANGE MYLIST 0 -1
