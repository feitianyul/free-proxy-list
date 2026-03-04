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

最后更新：2026-03-04 17:43:52 UTC（2026-03-05 01:43:52 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 120.232.242.119:22222 | 否 | ✓ 1481ms | ✓ 1206ms | 否 | ✓ 1159ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1031ms | ✓ 1108ms | ✓ 1019ms | http |
| 210.223.44.230:3128 | ✓ 1352ms | 否 | 否 | ✓ 1366ms | ✓ 954ms | http |
| 120.240.35.174:22222 | ✓ 1136ms | ✓ 1423ms | ✓ 1238ms | ✓ 1430ms | ✓ 1125ms | http |
| 91.193.240.157:9877 | ✓ 1276ms | 否 | ✓ 1529ms | 否 | ✓ 1496ms | http |
| 101.43.255.96:80 | 否 | ✓ 1466ms | ✓ 1269ms | 否 | ✓ 1326ms | http |
| 160.238.65.2:3128 | 否 | 否 | ✓ 1711ms | ✓ 1639ms | ✓ 987ms | http |
| 160.238.65.9:3128 | 否 | 否 | ✓ 1698ms | ✓ 1874ms | ✓ 970ms | http |
| 192.166.82.55:1080 | ✓ 942ms | 否 | ✓ 1130ms | ✓ 1593ms | 否 | http |
| 120.240.35.176:22222 | ✓ 1163ms | ✓ 1461ms | ✓ 1229ms | 否 | ✓ 1143ms | http |
| 120.240.29.168:22222 | ✓ 1137ms | ✓ 1526ms | 否 | 否 | ✓ 1130ms | http |
| 120.240.35.160:22222 | 否 | ✓ 1502ms | ✓ 1229ms | ✓ 1443ms | ✓ 1155ms | http |
| 117.159.239.51:22222 | ✓ 912ms | ✓ 1279ms | ✓ 984ms | ✓ 1303ms | ✓ 1013ms | http |
| 120.240.35.161:22222 | 否 | ✓ 1535ms | ✓ 1168ms | ✓ 1534ms | ✓ 1118ms | http |
| 61.72.110.54:3128 | ✓ 1444ms | 否 | ✓ 1380ms | 否 | ✓ 1491ms | http |
| 38.253.85.3:999 | 否 | 否 | ✓ 1312ms | ✓ 1518ms | ✓ 1359ms | http |
| 113.59.32.162:22222 | 否 | ✓ 1583ms | ✓ 1271ms | ✓ 1511ms | ✓ 1138ms | http |
| 120.92.212.16:7890 | ✓ 1158ms | 否 | ✓ 1411ms | ✓ 1516ms | 否 | http |
| 222.184.48.248:22222 | ✓ 1172ms | ✓ 1431ms | ✓ 1105ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1209ms | ✓ 1470ms | 否 | ✓ 1787ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1525ms | ✓ 1765ms | ✓ 1779ms | 否 | 否 | http |
| 45.140.147.155:1081 | ✓ 544ms | 否 | ✓ 937ms | ✓ 1570ms | 否 | http |
| 113.59.32.148:22222 | ✓ 1280ms | 否 | ✓ 1163ms | ✓ 1470ms | 否 | http |
| 160.238.65.4:3128 | 否 | ✓ 1377ms | 否 | ✓ 1601ms | ✓ 1379ms | http |
| 160.238.65.3:3128 | 否 | 否 | ✓ 1513ms | ✓ 1507ms | ✓ 1357ms | http |
| 160.238.65.6:3128 | ✓ 1603ms | ✓ 1345ms | 否 | 否 | ✓ 1430ms | http |
| 160.238.65.8:3128 | 否 | ✓ 1250ms | 否 | ✓ 1768ms | ✓ 1359ms | http |
| 160.238.65.7:3128 | 否 | ✓ 1359ms | 否 | ✓ 1644ms | ✓ 1368ms | http |
| 103.157.79.98:3125 | 否 | 否 | ✓ 1814ms | ✓ 1939ms | ✓ 1917ms | http |
| 46.249.103.192:443 | ✓ 1324ms | 否 | ✓ 1818ms | ✓ 1839ms | 否 | http |
| 103.215.36.88:17210 | ✓ 1175ms | ✓ 1602ms | ✓ 1945ms | 否 | ✓ 1131ms | http |
| 35.234.17.221:8080 | ✓ 999ms | 否 | ✓ 1240ms | 否 | ✓ 1053ms | http |
| 45.136.198.40:3128 | ✓ 709ms | 否 | ✓ 1408ms | ✓ 1901ms | ✓ 1696ms | http |
| 205.209.118.30:3138 | ✓ 347ms | ✓ 1965ms | ✓ 720ms | 否 | 否 | http |
| 120.240.35.177:22222 | ✓ 1307ms | ✓ 1840ms | ✓ 1668ms | ✓ 1632ms | ✓ 1321ms | http |
| 104.243.46.122:3128 | ✓ 1000ms | 否 | ✓ 1521ms | ✓ 1184ms | ✓ 855ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 883ms | ✓ 1093ms | ✓ 841ms | http |
| 14.56.177.44:3128 | ✓ 1051ms | ✓ 1451ms | ✓ 1141ms | ✓ 1465ms | ✓ 1099ms | http |
| 125.128.12.144:3128 | ✓ 1049ms | ✓ 1726ms | ✓ 1249ms | ✓ 1282ms | ✓ 914ms | http |
| 121.128.121.54:3128 | ✓ 1062ms | 否 | ✓ 992ms | ✓ 1201ms | ✓ 1075ms | http |
| 125.128.12.14:3128 | ✓ 1054ms | 否 | 否 | ✓ 1255ms | ✓ 949ms | http |
| 61.72.221.94:3128 | ✓ 1062ms | 否 | ✓ 1197ms | ✓ 1343ms | ✓ 1677ms | http |
| 61.72.221.194:3128 | ✓ 1072ms | 否 | ✓ 1258ms | ✓ 1315ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1851ms | ✓ 1376ms | ✓ 1537ms | 否 | ✓ 1137ms | http |
| 183.249.5.213:22222 | ✓ 901ms | ✓ 1352ms | ✓ 925ms | ✓ 1133ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1088ms | ✓ 1672ms | ✓ 1398ms | ✓ 1282ms | ✓ 1520ms | http |
| 138.124.53.25:7443 | ✓ 1375ms | 否 | 否 | ✓ 1651ms | ✓ 1684ms | http |
| 103.126.87.203:8082 | ✓ 1991ms | 否 | 否 | ✓ 1992ms | ✓ 1885ms | http |
| 120.240.29.51:22222 | ✓ 1193ms | ✓ 1688ms | ✓ 1365ms | 否 | 否 | http |
| 14.56.107.244:3128 | ✓ 812ms | 否 | ✓ 979ms | ✓ 1420ms | ✓ 938ms | http |
| 113.59.32.141:22222 | ✓ 1281ms | ✓ 1533ms | ✓ 1389ms | 否 | 否 | http |
| 47.101.149.27:9010 | ✓ 1574ms | 否 | ✓ 1577ms | ✓ 1616ms | 否 | http |
| 61.72.221.234:3128 | ✓ 1030ms | 否 | ✓ 1399ms | 否 | ✓ 988ms | http |
| 113.59.32.142:22222 | ✓ 1220ms | 否 | ✓ 1244ms | ✓ 1479ms | 否 | http |
| 91.233.223.147:3128 | ✓ 772ms | 否 | ✓ 1148ms | ✓ 1877ms | ✓ 1414ms | http |
| 120.240.35.175:22222 | ✓ 1260ms | ✓ 1852ms | ✓ 1283ms | ✓ 1613ms | ✓ 1203ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1618ms | ✓ 1610ms | ✓ 1310ms | http |
| 81.70.169.194:80 | ✓ 1576ms | ✓ 1805ms | 否 | ✓ 1435ms | ✓ 1242ms | http |
| 103.215.36.88:16198 | ✓ 1230ms | ✓ 1647ms | ✓ 1172ms | 否 | 否 | http |
| 222.228.171.92:8080 | ✓ 980ms | 否 | 否 | ✓ 1050ms | ✓ 854ms | http |
| 162.240.154.26:3128 | ✓ 886ms | ✓ 1789ms | ✓ 1391ms | ✓ 1699ms | 否 | http |
| 117.159.239.54:22222 | ✓ 1088ms | ✓ 1288ms | ✓ 963ms | ✓ 1312ms | ✓ 1012ms | http |
| 88.80.150.82:8080 | ✓ 1868ms | ✓ 1857ms | 否 | 否 | ✓ 1650ms | https |
| 120.240.35.178:22222 | ✓ 1332ms | ✓ 1841ms | ✓ 1761ms | ✓ 1745ms | ✓ 1236ms | http |
| 195.158.8.123:3128 | ✓ 1581ms | 否 | ✓ 1898ms | 否 | ✓ 1748ms | http |
| 5.75.196.26:40000 | ✓ 1440ms | ✓ 1222ms | ✓ 512ms | ✓ 1185ms | ✓ 1501ms | http |
| 120.240.29.174:22222 | ✓ 1137ms | ✓ 1544ms | ✓ 1093ms | ✓ 1432ms | ✓ 1158ms | http |
| 45.140.147.82:1081 | ✓ 397ms | ✓ 1093ms | ✓ 991ms | 否 | ✓ 1254ms | http |
| 160.238.65.5:3128 | ✓ 898ms | 否 | 否 | ✓ 1629ms | ✓ 1534ms | http |
| 222.90.155.238:8118 | ✓ 1220ms | ✓ 1628ms | ✓ 1289ms | ✓ 1622ms | ✓ 1184ms | http |
| 172.212.68.37:3128 | 否 | ✓ 1508ms | 否 | ✓ 1682ms | ✓ 1384ms | http |
| 120.55.163.237:10086 | ✓ 1176ms | ✓ 1328ms | ✓ 1655ms | ✓ 1373ms | ✓ 1054ms | http |
| 113.59.32.160:22222 | ✓ 1223ms | ✓ 1517ms | ✓ 1198ms | ✓ 1502ms | ✓ 1168ms | http |
| 103.215.36.88:13763 | ✓ 1139ms | ✓ 1588ms | ✓ 1446ms | 否 | ✓ 1201ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1374ms | 否 | ✓ 1193ms | ✓ 831ms | http |
| 113.255.59.226:8080 | ✓ 1462ms | 否 | ✓ 1367ms | 否 | ✓ 1389ms | http |
| 113.59.32.161:22222 | ✓ 1164ms | 否 | ✓ 1328ms | ✓ 1490ms | ✓ 1198ms | http |
| 103.215.36.88:16690 | ✓ 1279ms | ✓ 1637ms | ✓ 1422ms | ✓ 1574ms | ✓ 1295ms | http |

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
