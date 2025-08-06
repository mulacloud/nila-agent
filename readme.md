# Nila Agent 

Akan bertindak seperti proxy antara odoo manajemen dengan Hoster.
Agent akan mendaftarkan dirinya ke odoo, lalu apabila agent sudah 
berhasil diauthentikasi, agent akan secara periodik mengirim data statistik
ke Odoo dan menarik perintah manajemen container.

## Diagram

```
                    |                                                                       
                    |                                                                       
Bisa Public / Local | Cukup local IP saja                                                   
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
 +--------+         |                   +------------+                +----------------------+
 |        |         | XML-RPC           |            |      SSH       |                      |
 |  Odoo  <---------|-------------------- Nila Agent -----------------> SmartOS / BSD Hoster |
 |        |         |                   |            |                |                      |
 +--------+         |                   +------------+                +----------------------+
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
                    |                                                                       
```
