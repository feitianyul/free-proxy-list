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

最后更新：2026-03-08 23:22:58 UTC（2026-03-09 07:22:58 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 820ms | 否 | ✓ 1693ms | 否 | ✓ 1112ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1099ms | ✓ 1164ms | ✓ 1273ms | ✓ 981ms | http |
| 1.231.81.166:3128 | ✓ 1852ms | ✓ 1484ms | ✓ 994ms | ✓ 954ms | ✓ 939ms | http |
| 1.225.116.115:1080 | ✓ 958ms | ✓ 1741ms | ✓ 986ms | ✓ 1215ms | ✓ 832ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1137ms | ✓ 1771ms | ✓ 1508ms | 否 | http |
| 186.148.180.46:999 | ✓ 1380ms | 否 | ✓ 1572ms | 否 | ✓ 1645ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1974ms | ✓ 684ms | ✓ 963ms | ✓ 879ms | http |
| 152.42.213.210:8080 | ✓ 781ms | 否 | ✓ 1353ms | ✓ 1087ms | ✓ 1218ms | http |
| 194.213.18.200:443 | ✓ 1014ms | 否 | 否 | ✓ 1905ms | ✓ 1346ms | http |
| 114.4.251.26:8080 | ✓ 1325ms | 否 | ✓ 1240ms | ✓ 1851ms | ✓ 1186ms | http |
| 67.169.98.211:443 | ✓ 1928ms | ✓ 1831ms | 否 | ✓ 1150ms | 否 | http |
| 202.155.12.161:443 | ✓ 1922ms | ✓ 1367ms | 否 | ✓ 1154ms | ✓ 926ms | http |
| 14.56.107.244:3128 | ✓ 1831ms | ✓ 1797ms | ✓ 1643ms | ✓ 1977ms | ✓ 1471ms | http |
| 101.43.255.96:80 | ✓ 1000ms | ✓ 1249ms | ✓ 1046ms | ✓ 1182ms | ✓ 1262ms | http |
| 81.70.169.194:80 | ✓ 948ms | ✓ 1350ms | ✓ 995ms | ✓ 1297ms | ✓ 961ms | http |
| 59.46.216.131:30001 | ✓ 980ms | 否 | ✓ 989ms | ✓ 1374ms | 否 | http |
| 168.235.110.63:3128 | 否 | ✓ 1143ms | ✓ 500ms | ✓ 1245ms | ✓ 969ms | http |
| 190.9.109.198:999 | ✓ 1587ms | ✓ 1525ms | ✓ 1155ms | ✓ 1637ms | ✓ 1376ms | http |
| 190.9.109.207:999 | ✓ 1589ms | ✓ 1543ms | ✓ 1181ms | ✓ 1426ms | ✓ 1548ms | http |
| 116.80.82.224:3172 | 否 | 否 | ✓ 1760ms | ✓ 1858ms | ✓ 1947ms | http |
| 47.77.193.180:1080 | ✓ 719ms | ✓ 781ms | ✓ 267ms | ✓ 736ms | ✓ 554ms | http |
| 83.219.250.8:62920 | ✓ 786ms | 否 | ✓ 1229ms | 否 | ✓ 1434ms | http |
| 34.101.184.164:3128 | ✓ 814ms | 否 | ✓ 1766ms | ✓ 1286ms | ✓ 1183ms | http |
| 120.92.212.16:8890 | ✓ 1942ms | 否 | ✓ 960ms | 否 | ✓ 1687ms | http |
| 103.215.36.88:16541 | ✓ 956ms | ✓ 1388ms | ✓ 1050ms | ✓ 1276ms | ✓ 1087ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1235ms | ✓ 1464ms | ✓ 1919ms | ✓ 1851ms | http |
| 8.219.97.248:80 | ✓ 1514ms | ✓ 1877ms | 否 | ✓ 1717ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1011ms | ✓ 1919ms | ✓ 929ms | ✓ 1459ms | 否 | http |
| 178.236.245.59:3128 | ✓ 669ms | ✓ 1702ms | ✓ 1650ms | 否 | ✓ 1797ms | http |
| 178.236.245.17:3128 | ✓ 683ms | 否 | ✓ 1354ms | 否 | ✓ 1799ms | http |
| 103.39.51.190:8080 | ✓ 1466ms | 否 | 否 | ✓ 1316ms | ✓ 1789ms | http |
| 147.45.251.242:8888 | ✓ 1842ms | 否 | ✓ 1680ms | 否 | ✓ 1942ms | http |
| 91.233.223.147:3128 | ✓ 1248ms | 否 | ✓ 1329ms | 否 | ✓ 1580ms | http |
| 165.227.5.10:8888 | ✓ 779ms | 否 | ✓ 653ms | ✓ 1994ms | ✓ 749ms | http |
| 152.42.213.210:80 | ✓ 1021ms | 否 | ✓ 1614ms | ✓ 1239ms | ✓ 1106ms | http |
| 85.208.108.43:2094 | ✓ 524ms | 否 | ✓ 726ms | ✓ 1302ms | ✓ 812ms | http |
| 85.208.108.43:10808 | ✓ 517ms | 否 | ✓ 722ms | ✓ 1272ms | ✓ 846ms | http |
| 61.72.110.54:3128 | ✓ 1490ms | ✓ 870ms | ✓ 1010ms | ✓ 1014ms | ✓ 895ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1598ms | ✓ 1702ms | 否 | ✓ 1847ms | http |
| 121.230.8.109:1080 | ✓ 1039ms | ✓ 1502ms | ✓ 1180ms | ✓ 1860ms | ✓ 1131ms | http |
| 180.103.19.252:1080 | 否 | 否 | ✓ 1167ms | ✓ 1479ms | ✓ 1085ms | http |
| 180.125.216.109:8118 | ✓ 869ms | ✓ 1210ms | 否 | 否 | ✓ 973ms | http |
| 116.80.82.231:3172 | ✓ 1855ms | 否 | 否 | ✓ 1824ms | ✓ 1700ms | http |
| 159.192.133.250:8088 | ✓ 1677ms | 否 | ✓ 1280ms | ✓ 1621ms | ✓ 1258ms | http |
| 61.72.221.194:3128 | ✓ 1519ms | ✓ 1161ms | ✓ 1167ms | ✓ 1505ms | ✓ 845ms | http |
| 43.165.195.107:3128 | ✓ 1528ms | ✓ 1708ms | ✓ 1271ms | ✓ 1211ms | ✓ 960ms | http |
| 159.89.31.62:8080 | ✓ 611ms | 否 | ✓ 1829ms | ✓ 1598ms | ✓ 1657ms | http |
| 45.140.147.155:1081 | ✓ 1044ms | ✓ 1783ms | ✓ 1374ms | 否 | ✓ 1674ms | http |
| 115.231.181.40:8128 | ✓ 870ms | ✓ 1164ms | ✓ 1001ms | ✓ 1237ms | ✓ 902ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1995ms | 否 | ✓ 1614ms | ✓ 956ms | http |
| 34.96.238.40:8080 | ✓ 1292ms | 否 | ✓ 1210ms | ✓ 1223ms | ✓ 1744ms | http |
| 125.135.190.164:3076 | 否 | 否 | ✓ 1549ms | ✓ 1590ms | ✓ 1321ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 579ms | ✓ 999ms | ✓ 733ms | http |
| 129.226.155.60:3128 | ✓ 745ms | ✓ 1497ms | ✓ 1091ms | ✓ 1062ms | ✓ 1062ms | http |
| 103.215.36.88:17406 | ✓ 1135ms | ✓ 1628ms | ✓ 1093ms | 否 | ✓ 1032ms | http |
| 222.102.86.137:3040 | ✓ 1627ms | ✓ 1748ms | ✓ 1483ms | ✓ 1443ms | ✓ 1814ms | http |
| 103.183.10.169:3125 | ✓ 1297ms | 否 | 否 | ✓ 1554ms | ✓ 1388ms | http |
| 103.183.10.203:3125 | ✓ 1836ms | 否 | 否 | ✓ 1777ms | ✓ 1887ms | http |
| 20.120.225.109:3128 | ✓ 635ms | ✓ 1080ms | ✓ 1438ms | ✓ 1137ms | ✓ 830ms | http |
| 45.129.141.143:3128 | ✓ 738ms | 否 | ✓ 1623ms | 否 | ✓ 1542ms | http |
| 61.72.221.94:3128 | ✓ 1450ms | 否 | ✓ 1328ms | ✓ 1372ms | ✓ 1740ms | http |
| 103.215.36.88:19829 | 否 | 否 | ✓ 1843ms | ✓ 1456ms | ✓ 1348ms | http |
| 103.215.36.88:17039 | 否 | 否 | ✓ 1697ms | ✓ 1435ms | ✓ 1124ms | http |
| 103.215.36.88:10383 | ✓ 1110ms | ✓ 1345ms | ✓ 1136ms | ✓ 1263ms | ✓ 1027ms | http |
| 103.215.36.88:12482 | ✓ 1003ms | ✓ 1247ms | ✓ 1093ms | ✓ 1413ms | ✓ 968ms | http |
| 103.215.36.88:16241 | ✓ 929ms | ✓ 1228ms | ✓ 1307ms | ✓ 1374ms | ✓ 989ms | http |
| 103.160.202.174:3125 | ✓ 1736ms | 否 | ✓ 1824ms | ✓ 1659ms | ✓ 1559ms | http |
| 103.215.36.88:18877 | ✓ 1115ms | ✓ 1270ms | ✓ 1126ms | ✓ 1480ms | ✓ 1026ms | http |
| 103.215.36.88:17879 | ✓ 906ms | ✓ 1216ms | ✓ 947ms | ✓ 1237ms | ✓ 1020ms | http |
| 103.215.36.88:13378 | ✓ 1046ms | ✓ 1259ms | ✓ 1117ms | ✓ 1325ms | ✓ 990ms | http |
| 217.165.137.209:8181 | ✓ 1766ms | 否 | 否 | ✓ 1827ms | ✓ 1747ms | http |
| 103.215.36.88:15034 | ✓ 1018ms | ✓ 1324ms | ✓ 1007ms | ✓ 1192ms | ✓ 1017ms | http |
| 103.215.36.88:16067 | ✓ 985ms | ✓ 1213ms | ✓ 1061ms | ✓ 1330ms | ✓ 1067ms | http |
| 103.215.36.88:17516 | ✓ 1369ms | ✓ 1552ms | ✓ 1177ms | ✓ 1440ms | ✓ 1207ms | http |
| 103.215.36.88:16634 | ✓ 923ms | 否 | ✓ 1491ms | ✓ 1262ms | ✓ 980ms | http |
| 94.72.109.169:8080 | ✓ 1039ms | 否 | ✓ 1623ms | ✓ 1740ms | ✓ 1395ms | http |
| 143.244.140.119:3128 | ✓ 1897ms | 否 | ✓ 1887ms | ✓ 1982ms | ✓ 1682ms | http |
| 121.128.121.54:3128 | ✓ 1447ms | ✓ 1462ms | ✓ 1228ms | ✓ 1474ms | ✓ 1340ms | http |
| 150.107.140.238:3128 | ✓ 1582ms | 否 | ✓ 1999ms | 否 | ✓ 1938ms | http |
| 45.186.6.104:3128 | ✓ 1492ms | ✓ 1876ms | ✓ 1613ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 717ms | 否 | ✓ 1843ms | ✓ 1674ms | ✓ 1300ms | http |
| 121.230.9.160:1080 | ✓ 1537ms | ✓ 1972ms | ✓ 1154ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1013ms | ✓ 1899ms | ✓ 1768ms | 否 | 否 | http |
| 172.105.212.216:8888 | ✓ 698ms | 否 | ✓ 1495ms | ✓ 887ms | ✓ 742ms | http |
| 121.40.231.103:7890 | ✓ 1916ms | 否 | 否 | ✓ 1554ms | ✓ 902ms | http |
| 103.69.84.106:3131 | ✓ 1940ms | 否 | ✓ 1365ms | ✓ 1158ms | ✓ 954ms | http |
| 103.67.46.225:3125 | ✓ 1842ms | 否 | 否 | ✓ 1812ms | ✓ 1543ms | http |
| 103.84.95.54:7890 | ✓ 832ms | 否 | ✓ 1129ms | 否 | ✓ 1497ms | http |
| 103.215.36.88:11953 | ✓ 1054ms | ✓ 1343ms | ✓ 1138ms | ✓ 1423ms | ✓ 1044ms | http |
| 88.80.150.82:8080 | ✓ 891ms | ✓ 1697ms | ✓ 1652ms | ✓ 1770ms | ✓ 1639ms | https |

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
