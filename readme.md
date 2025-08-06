# Nila Agent 

akan bertindak seperti proxy antara odoo manajemen dengan hypervisor
agent akan mendaftarkan dirinya ke odoo 
apabila agent sudah berhasil di authenticasi, agent akan periodik 
mengirim data statistis ke odoo dan menarik perintah manajemen container

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
