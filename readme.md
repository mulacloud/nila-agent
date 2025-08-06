# Nila Agent 

akan bertindak seperti proxy antara odoo manajemen dengan hypervisor

## Diagram

```
+------+                                                        
|      |    XML-RPC   +------------+  SSH  +-------------------+
| Odoo +--------------> Nila Agent +------->-Smartos / BSD Box |
|      |              +------------+       +-------------------+
+------+                                                        
```
