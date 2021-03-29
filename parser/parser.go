package parser

/*
建表语句如下：

```sql
CREATE TABLE `ppm_org_organization_out_info` (
  `id` bigint NOT NULL,
  `org_id` bigint NOT NULL DEFAULT '0',
  `out_org_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `source_channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `source_platform` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `industry` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `is_authenticated` tinyint NOT NULL DEFAULT '1',
  `auth_ticket` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `auth_level` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT '1',
  `creator` bigint NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updator` bigint NOT NULL DEFAULT '0',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `version` int NOT NULL DEFAULT '1',
  `is_delete` tinyint NOT NULL DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `index_ppm_org_organization_out_info_org_id` (`org_id`) USING BTREE,
  KEY `index_ppm_org_organization_out_info_out_org_id` (`out_org_id`) USING BTREE,
  KEY `index_ppm_org_organization_out_info_create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
```

现在我们要对其进行解析，得到其中的字段名、类型等信息。

*/

type Parser struct {
	sql string
	i int
}

// 建表语句
type CreateTableStmt struct {

}

