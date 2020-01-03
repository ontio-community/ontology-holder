# Ontology-holder

Ontology-holder show the holder of OEP4.

## How to use?

Ontology-holder need an ontology node to sync block data, and need a mysql scheme (ontolog-holder) to save data.

Rename config-simpe.json to config.json, and set Mysql config, and set Ontology config.

There are two important item, "BlockHeight" and "Contracts". "Contracts" includes the hash of oep4 contracts which you want to get the holders. "BlockHeight" item indicates the height where program will search blocks.

Note that you need create db scheme "ontology-holder" with utf-8 charset befer setup Ontology-holder.

## API

contract must be OEP4 contract, such as b71fc841b203bcf08e81311131671885db689faf

1. Get holder list of asset

http://localhost:8080/getAssetHolder?qid=1&contract=b71fc841b203bcf08e81311131671885db689faf&from=0&count=100

from and count must larger 0, and count must smaller than 100.


2. Get asset base info

http://localhost:8080/GetAssetInfo?qid=1&contract=b71fc841b203bcf08e81311131671885db689faf


Get total holder count

```
http://localhost:8080/getAssetHolderCount?qid=1&contract=b71fc841b203bcf08e81311131671885db689faf
```


Get balance of address

```
http://localhost:8090/getBalance?address=98067c0ae9fd8f109956e06f5519a9bc0963f699&contract=b71fc841b203bcf08e81311131671885db689faf
```

contract param is option.

