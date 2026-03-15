**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-15 16:32:28 UTC（2026-03-16 00:32:28 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 137.220.150.152:6005 | ✓ 726ms | 否 | ✓ 789ms | ✓ 1210ms | ✓ 812ms | http |
| 137.220.151.110:6005 | ✓ 1458ms | 否 | ✓ 735ms | ✓ 1280ms | 否 | http |
| 205.209.118.30:3138 | ✓ 361ms | ✓ 1265ms | ✓ 1000ms | ✓ 1267ms | ✓ 969ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 913ms | ✓ 1253ms | ✓ 968ms | http |
| 137.184.1.87:3128 | ✓ 330ms | ✓ 793ms | ✓ 737ms | ✓ 674ms | ✓ 491ms | http |
| 137.184.6.37:3128 | ✓ 314ms | ✓ 1359ms | ✓ 339ms | ✓ 675ms | ✓ 491ms | http |
| 165.227.5.10:8888 | ✓ 609ms | ✓ 1675ms | ✓ 1704ms | 否 | 否 | http |
| 103.82.93.219:3128 | ✓ 839ms | 否 | ✓ 1182ms | ✓ 1177ms | ✓ 951ms | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 589ms | ✓ 642ms | ✓ 515ms | http |
| 104.129.202.127:10810 | 否 | 否 | ✓ 671ms | ✓ 673ms | ✓ 667ms | http |
| 113.160.132.26:8080 | ✓ 1572ms | ✓ 1314ms | ✓ 1109ms | ✓ 1170ms | ✓ 917ms | http |
| 165.225.113.220:11462 | ✓ 821ms | 否 | ✓ 772ms | ✓ 1126ms | 否 | http |
| 165.225.113.220:10958 | ✓ 711ms | 否 | ✓ 837ms | ✓ 1111ms | ✓ 1150ms | http |
| 45.167.124.52:8080 | ✓ 1206ms | 否 | ✓ 1414ms | 否 | ✓ 1583ms | http |
| 120.92.212.16:8890 | ✓ 1067ms | ✓ 1341ms | 否 | ✓ 1169ms | 否 | http |
| 101.43.255.96:80 | ✓ 1011ms | ✓ 1474ms | ✓ 1352ms | ✓ 1435ms | ✓ 1227ms | http |
| 101.43.127.100:8877 | ✓ 1550ms | ✓ 971ms | ✓ 1400ms | 否 | ✓ 1635ms | http |
| 165.225.113.220:10525 | ✓ 1457ms | 否 | ✓ 1615ms | ✓ 1515ms | ✓ 1200ms | http |
| 35.225.22.61:80 | ✓ 891ms | 否 | ✓ 1634ms | ✓ 1212ms | ✓ 928ms | http |
| 194.5.212.40:8080 | ✓ 1456ms | ✓ 1564ms | ✓ 1421ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1252ms | 否 | ✓ 1986ms | 否 | ✓ 1143ms | http |
| 137.220.150.104:6005 | ✓ 1781ms | 否 | 否 | ✓ 1107ms | ✓ 1784ms | http |
| 81.70.169.194:80 | 否 | ✓ 1467ms | ✓ 1469ms | ✓ 1417ms | ✓ 983ms | http |
| 34.101.184.164:3128 | ✓ 892ms | 否 | ✓ 1085ms | ✓ 1317ms | ✓ 956ms | http |
| 38.145.203.135:8443 | ✓ 848ms | ✓ 709ms | ✓ 1047ms | ✓ 644ms | ✓ 478ms | http |
| 45.136.131.39:8443 | ✓ 826ms | ✓ 1855ms | ✓ 793ms | ✓ 867ms | ✓ 1229ms | http |
| 62.60.177.204:34094 | ✓ 841ms | ✓ 1426ms | 否 | 否 | ✓ 1040ms | http |
| 150.249.255.91:3128 | ✓ 1045ms | ✓ 1859ms | 否 | 否 | ✓ 670ms | http |
| 139.159.99.242:8080 | 否 | 否 | ✓ 818ms | ✓ 1019ms | ✓ 925ms | http |
| 95.3.9.78:8080 | ✓ 1033ms | 否 | 否 | ✓ 1985ms | ✓ 1485ms | http |
| 114.231.73.23:1080 | ✓ 1047ms | ✓ 1245ms | ✓ 1928ms | ✓ 1579ms | 否 | http |
| 117.42.87.184:1080 | ✓ 924ms | ✓ 1180ms | 否 | ✓ 1298ms | ✓ 1189ms | http |
| 112.163.160.93:3128 | ✓ 774ms | ✓ 1201ms | ✓ 1275ms | 否 | 否 | http |
| 165.225.113.220:11180 | ✓ 1832ms | 否 | ✓ 1113ms | ✓ 1520ms | 否 | http |
| 165.225.113.220:11143 | ✓ 1910ms | 否 | 否 | ✓ 1510ms | ✓ 1197ms | http |
| 165.225.113.220:10086 | ✓ 1809ms | 否 | 否 | ✓ 1498ms | ✓ 1262ms | http |
| 38.55.106.208:6005 | 否 | 否 | ✓ 1413ms | ✓ 1416ms | ✓ 988ms | http |
| 47.101.149.27:9010 | 否 | 否 | ✓ 1974ms | ✓ 1940ms | ✓ 1212ms | http |
| 16.78.119.130:443 | 否 | ✓ 1819ms | 否 | ✓ 1609ms | ✓ 1534ms | http |
| 157.230.38.173:3128 | ✓ 716ms | 否 | ✓ 1370ms | ✓ 1178ms | ✓ 891ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1661ms | 否 | ✓ 1985ms | ✓ 1636ms | http |
| 178.236.245.59:3128 | ✓ 881ms | 否 | ✓ 986ms | ✓ 1866ms | ✓ 1712ms | http |
| 178.236.245.17:3128 | ✓ 880ms | 否 | ✓ 983ms | ✓ 1818ms | ✓ 1776ms | http |
| 14.225.212.37:7890 | 否 | 否 | ✓ 1089ms | ✓ 1725ms | ✓ 1093ms | http |
| 8.219.97.248:80 | ✓ 1020ms | ✓ 1812ms | 否 | ✓ 1361ms | 否 | http |
| 164.90.155.209:3128 | ✓ 1431ms | ✓ 1754ms | ✓ 932ms | ✓ 650ms | ✓ 483ms | http |
| 165.225.113.220:11191 | 否 | 否 | ✓ 1144ms | ✓ 1555ms | ✓ 1198ms | http |
| 165.225.113.220:11083 | ✓ 1096ms | 否 | ✓ 1440ms | ✓ 1522ms | 否 | http |
| 165.225.113.220:10970 | ✓ 1114ms | 否 | ✓ 1102ms | 否 | ✓ 1151ms | http |
| 45.136.198.40:3128 | ✓ 1013ms | ✓ 1910ms | ✓ 1264ms | 否 | ✓ 1789ms | http |
| 165.225.113.220:10829 | ✓ 1132ms | 否 | 否 | ✓ 1504ms | ✓ 1168ms | http |
| 165.225.113.220:11596 | ✓ 1157ms | 否 | ✓ 1445ms | ✓ 1504ms | 否 | http |
| 165.225.113.220:10017 | ✓ 1133ms | 否 | ✓ 1097ms | ✓ 1848ms | 否 | http |
| 165.225.113.220:11396 | ✓ 1141ms | 否 | ✓ 1103ms | ✓ 1463ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1375ms | 否 | 否 | ✓ 1664ms | ✓ 1389ms | http |
| 165.225.113.220:11584 | ✓ 1630ms | 否 | ✓ 1109ms | ✓ 1506ms | ✓ 1185ms | http |
| 38.145.208.138:8447 | ✓ 722ms | ✓ 733ms | ✓ 94ms | ✓ 668ms | ✓ 512ms | http |
| 45.136.130.223:8443 | ✓ 160ms | ✓ 1435ms | ✓ 81ms | ✓ 630ms | ✓ 478ms | http |
| 38.145.218.101:8447 | ✓ 1402ms | ✓ 1460ms | ✓ 84ms | ✓ 649ms | ✓ 476ms | http |
| 92.119.127.212:6005 | ✓ 1097ms | 否 | ✓ 1868ms | 否 | ✓ 1866ms | http |
| 38.145.208.234:8443 | 否 | ✓ 689ms | ✓ 1854ms | ✓ 679ms | ✓ 586ms | http |
| 149.50.116.240:1080 | ✓ 1192ms | 否 | ✓ 1305ms | 否 | ✓ 1557ms | http |
| 160.238.65.2:3128 | ✓ 1180ms | ✓ 1882ms | ✓ 1443ms | 否 | 否 | http |
| 160.238.65.7:3128 | ✓ 1185ms | ✓ 1753ms | ✓ 1565ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 1188ms | ✓ 1677ms | ✓ 1645ms | 否 | 否 | http |
| 103.139.138.194:3128 | ✓ 1167ms | 否 | ✓ 1165ms | ✓ 1615ms | ✓ 1497ms | http |
| 213.131.85.28:1976 | ✓ 1453ms | 否 | ✓ 1113ms | ✓ 1904ms | 否 | http |
| 47.77.193.180:1080 | ✓ 501ms | ✓ 828ms | ✓ 374ms | ✓ 699ms | ✓ 478ms | http |
| 165.225.113.220:10260 | ✓ 1225ms | 否 | ✓ 1483ms | ✓ 1565ms | 否 | http |
| 165.225.113.220:11807 | ✓ 1212ms | 否 | 否 | ✓ 1723ms | ✓ 1190ms | http |
| 165.225.113.220:11814 | ✓ 1126ms | 否 | ✓ 1188ms | ✓ 1532ms | ✓ 1191ms | http |
| 165.225.113.220:10155 | 否 | 否 | ✓ 1102ms | ✓ 1473ms | ✓ 1205ms | http |
| 113.176.92.71:3128 | ✓ 1829ms | ✓ 1536ms | ✓ 1172ms | ✓ 1133ms | ✓ 1145ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1260ms | 否 | ✓ 1327ms | ✓ 1739ms | http |
| 45.22.209.157:8888 | ✓ 1156ms | 否 | ✓ 1643ms | 否 | ✓ 1454ms | http |
| 172.212.68.37:3128 | ✓ 508ms | ✓ 1615ms | ✓ 905ms | 否 | ✓ 1281ms | http |
| 160.238.65.5:3128 | ✓ 1018ms | 否 | ✓ 1739ms | 否 | ✓ 1567ms | http |
| 165.225.113.220:11243 | ✓ 1108ms | 否 | ✓ 1127ms | ✓ 1545ms | ✓ 1207ms | http |
| 165.225.113.220:11481 | ✓ 1099ms | 否 | 否 | ✓ 1521ms | ✓ 1540ms | http |
| 103.183.10.169:3125 | 否 | 否 | ✓ 1230ms | ✓ 1544ms | ✓ 1459ms | http |
| 165.225.113.220:11589 | ✓ 1990ms | 否 | ✓ 1104ms | ✓ 1516ms | 否 | http |
| 165.225.113.220:11845 | ✓ 1953ms | 否 | ✓ 1107ms | ✓ 1504ms | ✓ 1203ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
