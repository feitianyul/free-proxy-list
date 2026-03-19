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

最后更新：2026-03-19 03:27:17 UTC（2026-03-19 11:27:17 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1348ms | 否 | ✓ 863ms | ✓ 1200ms | ✓ 1078ms | http |
| 147.161.239.240:8800 | ✓ 1395ms | ✓ 1546ms | ✓ 1047ms | ✓ 1544ms | ✓ 1465ms | http |
| 202.155.12.161:443 | ✓ 1350ms | 否 | ✓ 1071ms | ✓ 1329ms | ✓ 1125ms | http |
| 62.113.119.14:8080 | ✓ 1168ms | 否 | ✓ 1188ms | 否 | ✓ 1899ms | http |
| 113.160.132.26:8080 | ✓ 1766ms | ✓ 1601ms | 否 | ✓ 1395ms | ✓ 1082ms | http |
| 45.167.124.52:8080 | ✓ 1910ms | 否 | 否 | ✓ 1799ms | ✓ 1445ms | http |
| 85.198.96.242:3128 | ✓ 678ms | 否 | ✓ 560ms | ✓ 1609ms | ✓ 1271ms | http |
| 45.125.67.37:443 | 否 | 否 | ✓ 1506ms | ✓ 1384ms | ✓ 1340ms | http |
| 137.220.151.110:6005 | ✓ 1555ms | 否 | ✓ 1407ms | ✓ 1334ms | ✓ 1002ms | http |
| 35.225.22.61:80 | ✓ 869ms | 否 | ✓ 718ms | ✓ 961ms | ✓ 1029ms | http |
| 180.125.216.109:8118 | ✓ 1007ms | 否 | 否 | ✓ 1613ms | ✓ 1011ms | http |
| 116.80.49.161:3172 | ✓ 1659ms | 否 | ✓ 1846ms | 否 | ✓ 1776ms | http |
| 137.220.150.152:6005 | ✓ 902ms | 否 | ✓ 855ms | ✓ 1746ms | ✓ 1169ms | http |
| 38.34.179.46:8448 | 否 | ✓ 1027ms | ✓ 284ms | ✓ 1997ms | ✓ 889ms | http |
| 168.138.175.189:7890 | ✓ 1322ms | 否 | ✓ 835ms | ✓ 1146ms | 否 | http |
| 222.184.48.241:22222 | ✓ 1039ms | 否 | ✓ 1042ms | ✓ 1301ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1047ms | 否 | ✓ 968ms | 否 | ✓ 949ms | http |
| 183.249.5.117:22222 | ✓ 908ms | ✓ 1377ms | ✓ 1140ms | 否 | ✓ 957ms | http |
| 168.235.110.63:3128 | ✓ 631ms | ✓ 1867ms | ✓ 1227ms | 否 | ✓ 1748ms | http |
| 174.138.24.77:1080 | ✓ 846ms | 否 | ✓ 1883ms | 否 | ✓ 907ms | http |
| 101.47.73.135:3128 | ✓ 1095ms | 否 | 否 | ✓ 1498ms | ✓ 1340ms | http |
| 5.102.109.41:999 | 否 | ✓ 1363ms | ✓ 1823ms | ✓ 1514ms | ✓ 1219ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1212ms | ✓ 1088ms | ✓ 1340ms | ✓ 1003ms | http |
| 115.231.181.40:8128 | ✓ 1072ms | ✓ 1958ms | ✓ 1040ms | ✓ 1483ms | ✓ 1063ms | http |
| 219.117.204.211:7799 | ✓ 1859ms | 否 | ✓ 1395ms | ✓ 1959ms | ✓ 1557ms | http |
| 47.77.193.180:1080 | ✓ 918ms | ✓ 1020ms | 否 | ✓ 863ms | ✓ 653ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1423ms | ✓ 776ms | ✓ 1003ms | ✓ 851ms | http |
| 88.80.150.82:8080 | ✓ 1226ms | 否 | ✓ 1995ms | ✓ 1988ms | ✓ 1618ms | https |
| 217.76.245.80:999 | 否 | ✓ 1355ms | ✓ 1105ms | ✓ 1404ms | ✓ 1311ms | http |
| 101.32.244.83:8080 | ✓ 1191ms | ✓ 1868ms | ✓ 1099ms | ✓ 1624ms | ✓ 1377ms | http |
| 121.43.196.213:8222 | ✓ 1061ms | ✓ 1236ms | ✓ 1020ms | ✓ 1244ms | ✓ 990ms | http |
| 121.43.196.210:8222 | ✓ 1041ms | ✓ 1218ms | ✓ 1062ms | ✓ 1268ms | ✓ 988ms | http |
| 114.55.226.123:10086 | ✓ 1164ms | 否 | ✓ 1098ms | ✓ 1466ms | ✓ 1179ms | http |
| 38.34.179.102:8448 | ✓ 334ms | ✓ 908ms | ✓ 1043ms | ✓ 1106ms | ✓ 653ms | http |
| 3.137.167.45:27628 | ✓ 929ms | 否 | ✓ 1161ms | 否 | ✓ 1771ms | http |
| 120.92.212.16:7890 | ✓ 1053ms | ✓ 1384ms | 否 | 否 | ✓ 1188ms | http |
| 138.124.53.25:7443 | ✓ 475ms | 否 | ✓ 1451ms | ✓ 1653ms | 否 | http |
| 121.237.181.137:8888 | ✓ 1090ms | ✓ 1956ms | ✓ 1015ms | ✓ 1521ms | ✓ 914ms | http |
| 3.8.4.205:1206 | ✓ 806ms | 否 | ✓ 1327ms | 否 | ✓ 1880ms | http |
| 120.92.212.16:8890 | ✓ 1065ms | ✓ 1372ms | ✓ 1567ms | ✓ 1714ms | ✓ 1107ms | http |
| 137.220.150.104:6005 | ✓ 971ms | 否 | ✓ 1607ms | 否 | ✓ 1893ms | http |
| 137.220.150.170:6005 | ✓ 866ms | 否 | ✓ 960ms | ✓ 1289ms | ✓ 1028ms | http |
| 72.56.79.129:1080 | ✓ 613ms | ✓ 1763ms | ✓ 1591ms | 否 | ✓ 1505ms | http |
| 34.101.184.164:3128 | ✓ 1787ms | 否 | ✓ 1067ms | 否 | ✓ 1100ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1435ms | ✓ 1637ms | ✓ 1682ms | http |
| 185.41.152.110:3128 | ✓ 1053ms | ✓ 1592ms | ✓ 1839ms | ✓ 1621ms | ✓ 1521ms | http |
| 212.192.12.90:6005 | 否 | 否 | ✓ 1779ms | ✓ 1987ms | ✓ 1510ms | http |
| 106.117.208.101:7890 | ✓ 1103ms | ✓ 1361ms | ✓ 1188ms | ✓ 1889ms | ✓ 1145ms | http |
| 45.186.6.104:3128 | ✓ 1210ms | ✓ 1936ms | ✓ 1832ms | 否 | 否 | http |
| 45.88.0.114:3128 | 否 | 否 | ✓ 507ms | ✓ 1334ms | ✓ 1735ms | http |
| 45.88.0.113:3128 | ✓ 1554ms | ✓ 1742ms | ✓ 1592ms | ✓ 1874ms | ✓ 1814ms | http |
| 45.88.0.98:3128 | 否 | ✓ 1504ms | ✓ 1407ms | ✓ 1874ms | ✓ 1806ms | http |
| 45.88.0.116:3128 | 否 | ✓ 1958ms | ✓ 945ms | ✓ 1873ms | ✓ 1816ms | http |
| 213.220.62.62:3128 | 否 | 否 | ✓ 901ms | ✓ 1864ms | ✓ 1816ms | http |
| 45.88.0.99:3128 | ✓ 1560ms | 否 | ✓ 1331ms | ✓ 1870ms | ✓ 1795ms | http |
| 45.88.0.115:3128 | 否 | ✓ 1317ms | ✓ 1600ms | ✓ 1872ms | ✓ 1817ms | http |
| 45.88.0.111:3128 | ✓ 1024ms | ✓ 1528ms | ✓ 491ms | 否 | ✓ 1041ms | http |
| 45.88.0.117:3128 | ✓ 494ms | 否 | ✓ 604ms | ✓ 1299ms | 否 | http |
| 103.84.95.54:7890 | ✓ 863ms | 否 | ✓ 1454ms | 否 | ✓ 1475ms | http |
| 8.222.175.80:6128 | ✓ 1224ms | ✓ 1928ms | ✓ 993ms | ✓ 1180ms | ✓ 975ms | http |
| 211.217.231.234:8080 | ✓ 778ms | ✓ 1397ms | ✓ 962ms | ✓ 1939ms | ✓ 1298ms | http |
| 38.34.183.224:8448 | 否 | ✓ 879ms | ✓ 271ms | ✓ 1221ms | ✓ 1853ms | http |
| 193.23.200.251:10808 | ✓ 1259ms | 否 | ✓ 1268ms | 否 | ✓ 1972ms | http |
| 5.35.125.77:1080 | ✓ 1539ms | 否 | ✓ 1265ms | 否 | ✓ 1412ms | http |
| 210.223.44.230:3128 | ✓ 1690ms | 否 | ✓ 1134ms | 否 | ✓ 1859ms | http |
| 45.136.198.40:3128 | ✓ 1257ms | 否 | ✓ 1990ms | ✓ 1947ms | ✓ 1679ms | http |
| 38.180.2.107:3128 | ✓ 1253ms | ✓ 1710ms | 否 | 否 | ✓ 1777ms | http |
| 45.129.141.143:3128 | ✓ 1469ms | 否 | ✓ 1975ms | ✓ 1871ms | 否 | http |
| 185.191.236.162:3128 | ✓ 591ms | ✓ 1650ms | ✓ 1057ms | 否 | ✓ 1088ms | http |
| 183.249.5.110:22222 | 否 | 否 | ✓ 1541ms | ✓ 1874ms | ✓ 1554ms | http |
| 45.140.147.155:1081 | 否 | ✓ 1168ms | ✓ 454ms | ✓ 1634ms | 否 | http |
| 45.140.147.155:1082 | ✓ 904ms | ✓ 1821ms | ✓ 827ms | 否 | 否 | http |
| 89.169.3.180:1080 | ✓ 1149ms | ✓ 1911ms | 否 | 否 | ✓ 1457ms | http |
| 38.145.208.172:8448 | ✓ 1387ms | ✓ 887ms | ✓ 297ms | ✓ 883ms | ✓ 780ms | http |
| 209.97.149.157:80 | 否 | 否 | ✓ 1486ms | ✓ 1068ms | ✓ 808ms | http |
| 183.249.5.105:22222 | 否 | 否 | ✓ 1015ms | ✓ 1557ms | ✓ 1030ms | http |
| 61.76.95.217:40088 | 否 | 否 | ✓ 1139ms | ✓ 1451ms | ✓ 1129ms | http |
| 119.93.82.171:8082 | ✓ 1729ms | 否 | ✓ 1985ms | ✓ 1464ms | ✓ 1821ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1859ms | ✓ 1692ms | ✓ 1476ms | http |
| 45.136.131.59:8450 | ✓ 305ms | ✓ 1657ms | ✓ 1305ms | ✓ 1084ms | ✓ 1515ms | http |

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
