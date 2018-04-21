const redis = require('redis'),
      sleep = require('sleep'),
      client = redis.createClient({
        host: '127.0.0.1',
        port: 6379
      });

  // TODO: - to handle asynchronous nature of its sets core APIs 
  client.on("connect", () => {
     
    // String values
    client.set('newkey', 'ofsomevalue', redis.print);
    client.get('newkey', redis.print);

    // Binary values
    client.set('pages:about', 'about us', redis.print);
    client.get('pages:about', redis.print);

    // Counter values
    client.set('counter', 100, redis.print);
    client.incr('counter', redis.print);
    client.get('counter', redis.print);

    client.incr('counter', redis.print);
    client.get('counter', redis.print);

    client.incrby('counter', 50, redis.print);
    client.get('counter', redis.print);

    // MSET/MGET 
    client.mset('a', 10, 'b', 20, 'c', 30, redis.print);
    let array = client.mget('a','b','c', redis.print);
    console.log('MSET array: ', array);

    //  ALTERING/QUERYING KEY
    let mykeyexists;

    client.set('mykey', 'hello', redis.print);
    mykeyexists = client.exists('mykey', redis.print);
    console.log(mykeyexists);

    client.del('mykey', redis.print);
    mykeyexists = client.exists('mykey', redis.print);
    console.log('Deleted Key: ', mykeyexists);

    let keytype = client.type('mykey', redis.print);
    console.log('Keytype: ', keytype);

    //  SET TIMEOUT
    client.set('mykey', 'somevalue', redis.print);
    client.expire('mykey', 5, redis.print);
    let beforevalue = client.get('mykey', redis.print);
    console.log('before expiry: ' + beforevalue);
    sleep.sleep(6);
    let aftervalue = client.get('mykey', redis.print);
    console.log('after expiry: ' + aftervalue);
    
    //  TTL
    client.set('mykey', '100', redis.print)
    client.expire('mykey', 10, redis.print)
    let ttl_value = client.ttl('mykey', redis.print);
    console.log( 'TTL: ' + ttl_value);

    // LIST OPERATIONS

    //  remove existing values from list, if any
    client.del('mylist', redis.print); 

    client.rpush('mylist', 'a', redis.print);
    client.rpush('mylist', 'b', redis.print);
    client.lpush('mylist', 'first', redis.print);
    console.log(client.lrange('mylist', 0, -1));

    client.rpush('mylist', [ 1, 2, 3, 4, 5, "foo bar"], redis.print);
    console.log(client.lrange('mylist', 0, -1));

    client.rpush('numberlist', [1, 2, 3], redis.print);
    console.log('Popped value: ' + client.rpop('numberlist'));
    console.log('Popped value: ' + client.rpop('numberlist'));
    console.log('Popped value: ' + client.rpop('numberlist'));

    //  CAPPED LIST
    client.rpush('mycappedlist', [1, 2, 3, 4, 5], redis.print);
    client.ltrim('mycappedlist', 0, 2, redis.print);
    console.log('Capped list:.... ');
    console.log(client.lrange('mycappedlist', 0, -1));

    // close connection
    client.quit();

  });

  

  client.on("error", onError);

  function onError(error) {
    console.log("Error " + error);
  }