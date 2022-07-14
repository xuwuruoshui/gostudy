# ACL
- sub: 代表用户主体
- obj: 代表资源
- act: 代表访问权限


# RBAC
- g: 用户是某个角色,角色包含了多个角色
```csv
g, zhangsan, admin
g, lisi, developer
g, root, admin
g, wangwu,rooter
```
zhangsan的角色是admin, lisi的角色是developer, rooter的角色是admin, wangwu的角色是root

# 多RBAC
不仅用户分角色,访问对象也分角色(分组)

# Domain RBAC
全局角色,领域角色