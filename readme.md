# Nila Agent 

akan bertindak seperti proxy antara odoo manajemen dengan hypervisor

## Diagram

```
                    |                                                                       
                    |                                                                       
Bisa Public / Local | Cukup local IP saja                                                   
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
 +--------+         |                   +------------+                +--------------------+
 |        |         | XML-RPC           |            |      SSH       |                    |
 |  Odoo  <---------|-------------------- Nila Agent -----------------> SmartOS / BSD Host |
 |        |         |                   |            |                |                    |
 +--------+         |                   +------------+                +--------------------+
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
```
