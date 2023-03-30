# GoDemo

---

```shell
        quu..__
         $$$b  `---.__
          "$$b        `--.                          ___.---uuudP
           `$$b           `.__.------.__     __.---'      $$$$"              .
             "$b          -'            `-.-'            $$$"              .'|
               ".                                       d$"             _.'  |
                 `.   /                              ..."             .'     |
                   `./                           ..::-'            _.'       |
                    /                         .:::-'            .-'         .'
                   :                          ::''\          _.'            |
                  .' .-.             .-.           `.      .'               |
                  : /'$$|           .@"$\           `.   .'              _.-'
                 .'|$u$$|          |$$,$$|           |  <            _.-'
                 | `:$$:'          :$$$$$:           `.  `.       .-'
                 :                  `"--'             |    `-.     \
                :##.       ==             .###.       `.      `.    `\
                |##:                      :###:        |        >     >
                |#'     `..'`..'          `###'        x:      /     /
                 \                                   xXX|     /    ./
                  \                                xXXX'|    /   ./
                  /`-.                                  `.  /   /
                 :    `-  ...........,                   | /  .'
                 |         ``:::::::'       .            |<    `.
                 |             ```          |           x| \ `.:``.
                 |                         .'    /'   xXX|  `:`M`M':.
                 |    |                    ;    /:' xXXX'|  -'MMMMM:'
                 `.  .'                   :    /:'       |-'MMMM.-'
                  |  |                   .'   /'        .'MMM.-'
                  `'`'                   :  ,'          |MMM<
                    |                     `'            |tbap\
                     \                                  :MM.-'
                      \                 |              .''
                       \.               `.            /
                        /     .:::::::.. :           /
                       |     .:::::::::::`.         /
                       |   .:::------------\       /
                      /   .''               >::'  /
                      `',:                 :    .'
                                           `:.:'
```


## Contents

---

| package | desc          |
|---------|---------------|
| app     | application   |
| domain  | domain models |
| infra   | infra         |
| job     | crontab       |
| onchain | onchain ACL   |
| pkg     | utils         |


## Structure

---

```
      -------------------                 ------------
      |    Interface    |                 |   Job    |
      -------------------                 ------------
         ↑            ↓                        |
         |          -------------------        |
         |          |     Service     |        |
         |          -------------------        |
         |             ↑           ↓           ↓
         |             |        -------------------
         |             |        |      Domain     |
         |             |        -------------------
         |             |                 ↑
      ----------------------------------------------
      |              Infrastructure                |
      ----------------------------------------------
```

## Run Service

---

### 1. Run App Service
cmd/xxx/config.yaml
```yaml
debug: true
server:
  address: 0.0.0.0:6668
mysql:
  dns: user:password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local
redis:
  address: 127.0.0.1
  port: 6379
  DB: 0
chain:
  band:
    node_url: https://laozi1.bandchain.org/api
    block_time: 7
    irreversible_block: 10
  waxp:
    node_url: https://wax.greymass.com
    block_time: 1
    irreversible_block: 100
```

### 2. Run Job
cmd/job/config.yaml 
```yaml
debug: true
job:
  broker_dns: 127.0.0.1:6379
  backend_dns: 127.0.0.1:6379
  default_queue: default
mysql:
  dns: user:password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local
redis:
  address: 127.0.0.1
  port: 6379
  DB: 0
chain:
  band:
    node_url: https://laozi1.bandchain.org/api
    block_time: 7
    irreversible_block: 10
  waxp:
    node_url: https://wax.greymass.com
    block_time: 1
    irreversible_block: 100
```


## Undo

1. RPC Server
2. Error Code 
3. Wallet Transfer 
4. Job Schedule
5. Log