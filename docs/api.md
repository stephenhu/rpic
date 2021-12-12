# api

resource | description
---|---
systems | monitoring of system resources like cpu, memory, network, storage, perform shutdown/restart of the system, and configuration of the system.
networks | network configuration for wired, wifi
services | these are installed on top of raspbian like wireguard or custom applications
backups | backup of configuration and other persistent data

> ## systems

url | description | parameter(s)
---|---|---
`GET /api/systems/:id` | get info about all resource state (cpu, memory, network, storage) |
`PUT /api/systems/:id`| shutdown, restart, or suspend device | 0 - shutdown, 1 - restart, 2 - suspend

> ## networks

url | description | parameter(s)
---|---|---
`GET /api/networks/:id` | | 
`PUT /api/networks/:id` | toggle network on/off |


> ## services

url | parameters | description
---|---|---
`GET /api/services` | |
`POST /api/services` | name: string, imagePath: string, version: string, hash: string |
`PUT /api/services/:id` | |

> ## backups

url | parameters | description
---|---|---
`GET /api/backups` | |
`POST /api/backups` | |
