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

最后更新：2026-05-07 10:25:40 UTC（2026-05-07 18:25:40 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 574ms | ✓ 1305ms | ✓ 950ms | ✓ 1349ms | ✓ 956ms | http |
| 113.160.132.26:8080 | ✓ 1499ms | ✓ 1558ms | ✓ 944ms | ✓ 1286ms | ✓ 1047ms | http |
| 185.191.236.162:3128 | ✓ 1163ms | ✓ 1673ms | ✓ 1841ms | 否 | ✓ 1457ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1319ms | ✓ 1560ms | ✓ 1321ms | ✓ 1197ms | http |
| 43.133.44.89:8888 | ✓ 957ms | 否 | ✓ 1231ms | ✓ 1513ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1707ms | ✓ 1463ms | ✓ 1091ms | 否 | 否 | http |
| 38.194.254.134:999 | ✓ 1734ms | 否 | 否 | ✓ 1387ms | ✓ 1226ms | http |
| 86.104.72.219:1081 | ✓ 1407ms | ✓ 1155ms | ✓ 864ms | ✓ 1230ms | ✓ 983ms | http |
| 47.79.39.142:30000 | 否 | ✓ 1744ms | ✓ 1161ms | ✓ 942ms | ✓ 1077ms | http |
| 45.125.67.37:8443 | ✓ 1073ms | 否 | ✓ 1261ms | ✓ 1280ms | ✓ 1341ms | http |
| 91.233.223.147:3128 | ✓ 1345ms | 否 | ✓ 1050ms | 否 | ✓ 1720ms | http |
| 193.160.209.58:1080 | ✓ 1060ms | 否 | ✓ 1569ms | ✓ 1950ms | ✓ 1583ms | http |
| 144.124.227.88:3128 | ✓ 695ms | 否 | ✓ 1110ms | ✓ 1845ms | ✓ 1638ms | http |
| 137.59.47.73:3128 | ✓ 1630ms | ✓ 1646ms | 否 | ✓ 1922ms | ✓ 1756ms | http |
| 77.110.119.136:3128 | ✓ 964ms | 否 | 否 | ✓ 1828ms | ✓ 872ms | http |
| 190.12.150.244:999 | 否 | 否 | ✓ 971ms | ✓ 1704ms | ✓ 1502ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1530ms | ✓ 983ms | ✓ 1306ms | ✓ 1108ms | http |
| 158.160.215.167:8125 | ✓ 897ms | ✓ 1619ms | ✓ 1702ms | 否 | 否 | http |
| 14.143.222.113:57788 | ✓ 1685ms | 否 | ✓ 1084ms | ✓ 1536ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1731ms | 否 | 否 | ✓ 1968ms | ✓ 1467ms | http |
| 43.156.132.113:3128 | ✓ 1551ms | ✓ 1451ms | ✓ 860ms | ✓ 1115ms | ✓ 1012ms | http |
| 91.242.229.129:8092 | ✓ 1304ms | ✓ 1410ms | ✓ 1323ms | 否 | ✓ 1966ms | http |
| 206.206.126.177:2412 | 否 | ✓ 1972ms | ✓ 817ms | ✓ 1262ms | ✓ 877ms | http |
| 103.157.200.126:3128 | ✓ 1714ms | 否 | ✓ 1330ms | ✓ 1871ms | 否 | http |
| 47.77.216.82:1080 | ✓ 491ms | ✓ 918ms | 否 | ✓ 793ms | ✓ 594ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1333ms | ✓ 1420ms | ✓ 1178ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1181ms | ✓ 1790ms | ✓ 1383ms | http |
| 38.180.2.107:3128 | ✓ 1047ms | ✓ 1682ms | ✓ 1846ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1229ms | ✓ 1675ms | 否 | ✓ 1560ms | ✓ 1214ms | http |
| 91.217.81.131:1080 | ✓ 978ms | 否 | 否 | ✓ 1978ms | ✓ 1409ms | http |
| 158.160.215.167:8124 | ✓ 1354ms | ✓ 1912ms | ✓ 800ms | 否 | 否 | http |
| 84.47.150.125:1080 | ✓ 808ms | ✓ 1823ms | ✓ 1825ms | 否 | ✓ 1477ms | http |
| 103.82.23.118:5185 | ✓ 1393ms | 否 | ✓ 1918ms | ✓ 1878ms | ✓ 1605ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1788ms | ✓ 768ms | ✓ 1552ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1468ms | ✓ 1884ms | ✓ 1367ms | ✓ 1274ms | 否 | http |
| 220.121.143.33:3128 | 否 | 否 | ✓ 883ms | ✓ 1101ms | ✓ 852ms | http |
| 20.78.213.56:80 | ✓ 996ms | ✓ 1406ms | 否 | ✓ 1221ms | ✓ 1075ms | http |
| 147.45.186.28:3128 | ✓ 1318ms | 否 | 否 | ✓ 1626ms | ✓ 1496ms | http |
| 121.230.8.171:1080 | ✓ 1194ms | ✓ 1522ms | ✓ 1521ms | ✓ 1606ms | ✓ 1230ms | http |
| 116.118.48.147:3128 | 否 | 否 | ✓ 1114ms | ✓ 1289ms | ✓ 1031ms | http |
| 8.154.21.175:3128 | 否 | ✓ 1252ms | ✓ 1031ms | ✓ 1163ms | ✓ 1198ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1381ms | ✓ 1359ms | ✓ 1077ms | http |
| 120.92.212.16:7890 | ✓ 1038ms | ✓ 1446ms | ✓ 1038ms | ✓ 1364ms | ✓ 1073ms | http |
| 158.160.215.167:8126 | ✓ 1358ms | ✓ 1779ms | 否 | 否 | ✓ 1875ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1164ms | ✓ 1173ms | ✓ 1204ms | http |
| 158.160.215.167:8127 | ✓ 1359ms | ✓ 1871ms | ✓ 1541ms | 否 | 否 | http |
| 103.106.218.123:8081 | 否 | 否 | ✓ 1618ms | ✓ 1684ms | ✓ 1776ms | http |
| 138.197.68.35:4857 | ✓ 361ms | 否 | ✓ 201ms | ✓ 1481ms | ✓ 992ms | http |
| 181.119.97.24:999 | ✓ 1432ms | ✓ 1956ms | ✓ 1782ms | 否 | 否 | http |
| 152.32.132.190:7890 | ✓ 1057ms | ✓ 1109ms | 否 | ✓ 1494ms | 否 | http |
| 38.211.245.34:999 | ✓ 1164ms | 否 | ✓ 884ms | 否 | ✓ 1940ms | http |
| 38.211.245.18:999 | ✓ 1305ms | 否 | ✓ 1272ms | 否 | ✓ 1966ms | http |
| 195.19.217.200:3128 | ✓ 1383ms | 否 | ✓ 1604ms | 否 | ✓ 1710ms | http |
| 178.63.155.151:8888 | ✓ 807ms | ✓ 1448ms | ✓ 1629ms | 否 | ✓ 1255ms | http |
| 45.153.231.229:8080 | ✓ 1077ms | ✓ 1959ms | 否 | 否 | ✓ 1706ms | http |
| 45.140.147.155:1081 | ✓ 1364ms | 否 | ✓ 1046ms | ✓ 1706ms | ✓ 1544ms | http |
| 157.0.142.246:10057 | ✓ 1234ms | ✓ 1415ms | ✓ 1245ms | 否 | 否 | http |
| 45.59.122.132:80 | ✓ 584ms | 否 | ✓ 1004ms | ✓ 1992ms | ✓ 1380ms | http |
| 3.101.133.120:80 | 否 | 否 | ✓ 494ms | ✓ 1116ms | ✓ 780ms | http |
| 86.104.74.110:1081 | ✓ 774ms | 否 | ✓ 1069ms | 否 | ✓ 1724ms | http |
| 86.104.74.110:1082 | 否 | 否 | ✓ 1592ms | ✓ 1895ms | ✓ 1383ms | http |
| 110.172.28.217:3128 | 否 | 否 | ✓ 1245ms | ✓ 1656ms | ✓ 1396ms | http |
| 82.114.228.67:1080 | ✓ 1909ms | ✓ 1528ms | ✓ 1115ms | ✓ 1590ms | 否 | http |
| 223.84.151.86:30005 | 否 | ✓ 1986ms | 否 | ✓ 1940ms | ✓ 1394ms | http |
| 86.104.72.219:1082 | ✓ 1285ms | 否 | ✓ 1199ms | 否 | ✓ 854ms | http |
| 222.107.27.7:8056 | 否 | ✓ 1261ms | ✓ 1219ms | ✓ 1236ms | 否 | http |
| 150.136.153.231:80 | ✓ 653ms | ✓ 1395ms | 否 | ✓ 1910ms | ✓ 1863ms | http |
| 84.247.171.137:3128 | 否 | 否 | ✓ 1531ms | ✓ 1825ms | ✓ 1241ms | http |
| 212.58.132.5:8888 | ✓ 1980ms | 否 | 否 | ✓ 1477ms | ✓ 1756ms | http |
| 116.80.96.90:3172 | ✓ 1647ms | 否 | 否 | ✓ 1989ms | ✓ 1756ms | http |
| 210.223.44.230:3128 | ✓ 1904ms | 否 | ✓ 1904ms | 否 | ✓ 1260ms | http |
| 181.78.44.63:999 | 否 | ✓ 1686ms | ✓ 1410ms | ✓ 1371ms | ✓ 1169ms | http |
| 45.140.147.155:1082 | ✓ 944ms | ✓ 1806ms | ✓ 851ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 837ms | ✓ 1527ms | ✓ 978ms | 否 | 否 | http |
| 185.230.191.240:3128 | ✓ 1077ms | ✓ 1860ms | 否 | 否 | ✓ 1591ms | http |
| 20.164.75.153:8080 | ✓ 1308ms | 否 | ✓ 1362ms | 否 | ✓ 1881ms | http |
| 175.194.173.105:3128 | 否 | ✓ 1140ms | ✓ 904ms | ✓ 1114ms | 否 | http |
| 38.180.192.119:3128 | ✓ 835ms | 否 | ✓ 1728ms | ✓ 1898ms | ✓ 1845ms | http |
| 45.236.129.64:3128 | ✓ 1888ms | 否 | ✓ 1679ms | 否 | ✓ 1887ms | http |
| 77.93.89.128:47146 | ✓ 866ms | 否 | ✓ 1152ms | ✓ 1125ms | ✓ 1308ms | http |
| 121.230.9.198:1080 | ✓ 1728ms | ✓ 1713ms | ✓ 1210ms | ✓ 1899ms | 否 | http |
| 57.128.188.167:8132 | ✓ 1773ms | 否 | ✓ 1546ms | ✓ 1871ms | ✓ 1804ms | http |
| 185.121.13.73:3128 | ✓ 867ms | ✓ 1768ms | ✓ 1712ms | ✓ 1749ms | ✓ 1657ms | http |
| 121.230.8.17:1080 | ✓ 1290ms | 否 | ✓ 1261ms | ✓ 1771ms | 否 | http |
| 175.215.147.63:3084 | 否 | 否 | ✓ 1195ms | ✓ 1803ms | ✓ 1150ms | http |
| 103.147.84.34:8080 | ✓ 1894ms | 否 | ✓ 1447ms | ✓ 1623ms | ✓ 1513ms | http |
| 160.238.65.9:3128 | ✓ 1199ms | ✓ 1887ms | 否 | ✓ 1848ms | 否 | http |
| 160.238.65.3:3128 | 否 | 否 | ✓ 1468ms | ✓ 1519ms | ✓ 1733ms | http |
| 20.127.128.70:8080 | ✓ 1518ms | 否 | ✓ 1166ms | ✓ 1603ms | ✓ 1792ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1959ms | ✓ 1917ms | ✓ 1898ms | http |

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
