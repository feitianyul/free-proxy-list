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

最后更新：2026-03-09 06:56:59 UTC（2026-03-09 14:56:59 UTC+8）

**代理总数：79**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 79 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 101.47.73.135:3128 | ✓ 798ms | 否 | ✓ 959ms | ✓ 1555ms | ✓ 974ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1283ms | ✓ 1072ms | ✓ 1612ms | ✓ 1043ms | http |
| 120.92.212.16:7890 | ✓ 788ms | ✓ 935ms | 否 | ✓ 1002ms | ✓ 753ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1034ms | ✓ 981ms | ✓ 913ms | ✓ 1065ms | http |
| 45.140.147.155:1081 | ✓ 1088ms | 否 | ✓ 1292ms | ✓ 1755ms | ✓ 1283ms | http |
| 194.213.18.200:443 | ✓ 1159ms | 否 | ✓ 1253ms | 否 | ✓ 1628ms | http |
| 150.107.140.238:3128 | ✓ 833ms | 否 | ✓ 1043ms | 否 | ✓ 1887ms | http |
| 103.106.219.107:8081 | ✓ 1344ms | 否 | ✓ 1649ms | ✓ 1520ms | ✓ 1529ms | http |
| 165.227.5.10:8888 | ✓ 726ms | 否 | ✓ 1202ms | 否 | ✓ 482ms | http |
| 103.84.95.54:7890 | ✓ 609ms | 否 | ✓ 1142ms | 否 | ✓ 1313ms | http |
| 120.92.212.16:8890 | ✓ 1022ms | 否 | ✓ 1097ms | ✓ 956ms | ✓ 762ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 1655ms | ✓ 1268ms | ✓ 994ms | http |
| 190.9.109.207:999 | 否 | 否 | ✓ 1836ms | ✓ 1632ms | ✓ 1262ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1522ms | ✓ 1389ms | ✓ 1171ms | http |
| 45.136.198.40:3128 | ✓ 1933ms | 否 | ✓ 911ms | ✓ 1967ms | ✓ 1842ms | http |
| 39.104.201.40:7890 | ✓ 712ms | ✓ 949ms | ✓ 1287ms | ✓ 1189ms | ✓ 744ms | http |
| 115.231.181.40:8128 | ✓ 736ms | ✓ 974ms | ✓ 1492ms | ✓ 1252ms | ✓ 1065ms | http |
| 202.155.12.161:443 | ✓ 1312ms | 否 | ✓ 1771ms | ✓ 1028ms | ✓ 1201ms | http |
| 81.70.169.194:80 | ✓ 818ms | ✓ 1050ms | ✓ 1404ms | ✓ 951ms | ✓ 721ms | http |
| 101.43.255.96:80 | ✓ 1424ms | ✓ 1504ms | ✓ 752ms | ✓ 1580ms | ✓ 1381ms | http |
| 14.225.222.164:7890 | 否 | ✓ 1385ms | ✓ 1616ms | 否 | ✓ 1609ms | http |
| 222.184.48.241:22222 | ✓ 1337ms | ✓ 928ms | ✓ 1217ms | 否 | 否 | http |
| 117.159.239.51:22222 | ✓ 750ms | ✓ 980ms | ✓ 810ms | ✓ 1076ms | ✓ 773ms | http |
| 117.159.239.52:22222 | ✓ 789ms | ✓ 1015ms | ✓ 746ms | ✓ 1030ms | ✓ 805ms | http |
| 117.159.239.54:22222 | ✓ 901ms | ✓ 985ms | ✓ 760ms | ✓ 1022ms | ✓ 797ms | http |
| 120.240.35.176:22222 | ✓ 1198ms | ✓ 1156ms | ✓ 960ms | ✓ 1082ms | ✓ 804ms | http |
| 120.240.35.177:22222 | ✓ 843ms | ✓ 1108ms | ✓ 1174ms | ✓ 1132ms | ✓ 1897ms | http |
| 120.240.35.161:22222 | ✓ 897ms | 否 | ✓ 1131ms | ✓ 1117ms | ✓ 858ms | http |
| 160.250.4.245:1 | ✓ 1301ms | 否 | ✓ 1483ms | ✓ 1165ms | ✓ 959ms | http |
| 160.250.5.22:1 | ✓ 1301ms | 否 | ✓ 1334ms | ✓ 1398ms | ✓ 998ms | http |
| 183.249.5.214:22222 | ✓ 1926ms | 否 | ✓ 861ms | ✓ 1297ms | ✓ 1106ms | http |
| 120.240.35.175:22222 | ✓ 863ms | 否 | ✓ 1976ms | ✓ 1049ms | ✓ 1445ms | http |
| 113.59.32.141:22222 | ✓ 1712ms | 否 | 否 | ✓ 1973ms | ✓ 1351ms | http |
| 54.222.174.194:80 | ✓ 1159ms | 否 | ✓ 1335ms | ✓ 1432ms | 否 | http |
| 120.238.159.230:22222 | ✓ 885ms | ✓ 1181ms | ✓ 902ms | ✓ 1048ms | 否 | http |
| 120.238.159.228:22222 | ✓ 847ms | ✓ 1143ms | ✓ 981ms | ✓ 1080ms | ✓ 896ms | http |
| 120.240.35.173:22222 | ✓ 879ms | ✓ 1565ms | ✓ 911ms | ✓ 1065ms | ✓ 846ms | http |
| 83.219.250.8:62920 | ✓ 790ms | 否 | ✓ 1281ms | 否 | ✓ 1471ms | http |
| 46.39.105.157:8080 | 否 | 否 | ✓ 1861ms | ✓ 1731ms | ✓ 1518ms | http |
| 89.185.85.138:1080 | ✓ 1070ms | 否 | ✓ 1133ms | 否 | ✓ 1506ms | http |
| 35.225.22.61:80 | ✓ 1097ms | ✓ 1249ms | ✓ 564ms | ✓ 1136ms | ✓ 1050ms | http |
| 168.235.110.63:3128 | ✓ 1002ms | ✓ 1270ms | 否 | 否 | ✓ 1006ms | http |
| 138.124.93.82:1080 | ✓ 1290ms | 否 | ✓ 1414ms | ✓ 1780ms | ✓ 1632ms | http |
| 103.236.89.228:7890 | ✓ 790ms | ✓ 1425ms | ✓ 1118ms | ✓ 1037ms | ✓ 984ms | http |
| 222.184.48.251:22222 | 否 | 否 | ✓ 1203ms | ✓ 1911ms | ✓ 735ms | http |
| 190.9.109.198:999 | ✓ 1063ms | ✓ 1587ms | ✓ 1602ms | ✓ 1423ms | 否 | http |
| 67.169.98.211:443 | ✓ 597ms | ✓ 1711ms | 否 | 否 | ✓ 1060ms | http |
| 183.249.5.111:22222 | ✓ 733ms | ✓ 904ms | ✓ 829ms | ✓ 916ms | ✓ 670ms | http |
| 117.159.239.50:22222 | ✓ 781ms | ✓ 983ms | ✓ 818ms | ✓ 1136ms | ✓ 794ms | http |
| 120.240.35.178:22222 | ✓ 967ms | ✓ 1179ms | ✓ 1048ms | ✓ 1087ms | ✓ 874ms | http |
| 106.14.203.63:3333 | ✓ 1704ms | 否 | ✓ 1020ms | ✓ 1782ms | ✓ 709ms | http |
| 183.249.5.213:22222 | ✓ 1091ms | ✓ 1294ms | ✓ 916ms | ✓ 1127ms | ✓ 901ms | http |
| 120.240.35.160:22222 | ✓ 946ms | ✓ 1421ms | ✓ 1040ms | 否 | ✓ 836ms | http |
| 144.31.184.218:3128 | ✓ 1485ms | 否 | ✓ 1526ms | 否 | ✓ 1518ms | http |
| 121.230.8.181:1080 | ✓ 1218ms | 否 | ✓ 848ms | 否 | ✓ 1660ms | http |
| 222.184.48.252:22222 | ✓ 784ms | 否 | ✓ 1384ms | 否 | ✓ 1504ms | http |
| 1.55.182.125:10001 | ✓ 1756ms | 否 | ✓ 1365ms | ✓ 1792ms | 否 | http |
| 121.232.73.214:1080 | ✓ 1602ms | 否 | ✓ 1303ms | ✓ 1302ms | 否 | http |
| 183.249.5.117:22222 | ✓ 876ms | ✓ 845ms | ✓ 661ms | ✓ 864ms | ✓ 668ms | http |
| 120.232.242.119:22222 | ✓ 855ms | ✓ 1284ms | 否 | 否 | ✓ 1847ms | http |
| 58.220.95.12:12417 | ✓ 1615ms | ✓ 973ms | ✓ 889ms | ✓ 1653ms | ✓ 814ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1076ms | ✓ 1708ms | 否 | ✓ 1300ms | http |
| 113.59.32.142:22222 | ✓ 1354ms | ✓ 1681ms | 否 | ✓ 1675ms | ✓ 1303ms | http |
| 45.88.0.115:3128 | ✓ 1146ms | ✓ 1950ms | ✓ 1498ms | 否 | 否 | http |
| 45.88.0.113:3128 | ✓ 1135ms | 否 | ✓ 1154ms | ✓ 1470ms | ✓ 1179ms | http |
| 8.219.97.248:80 | ✓ 1322ms | 否 | ✓ 1393ms | ✓ 1397ms | 否 | http |
| 14.225.222.213:7890 | ✓ 1316ms | ✓ 1390ms | ✓ 1078ms | 否 | ✓ 871ms | http |
| 95.85.252.153:21064 | ✓ 1118ms | ✓ 1649ms | ✓ 1944ms | 否 | 否 | http |
| 103.139.138.194:3128 | ✓ 1075ms | 否 | ✓ 1671ms | ✓ 1474ms | ✓ 1325ms | http |
| 122.52.108.244:8082 | ✓ 1300ms | 否 | ✓ 1679ms | ✓ 1424ms | ✓ 1345ms | http |
| 103.82.23.118:5196 | ✓ 1685ms | ✓ 1898ms | ✓ 1150ms | ✓ 1296ms | ✓ 1502ms | http |
| 103.39.51.190:8080 | ✓ 1696ms | 否 | ✓ 1752ms | ✓ 1400ms | ✓ 1287ms | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1971ms | ✓ 1823ms | ✓ 1570ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1609ms | 否 | ✓ 1451ms | ✓ 994ms | http |
| 113.59.32.163:22222 | 否 | ✓ 1794ms | ✓ 1377ms | ✓ 1605ms | 否 | http |
| 103.67.46.225:3125 | ✓ 1896ms | 否 | ✓ 1702ms | ✓ 1506ms | 否 | http |
| 113.132.112.110:9000 | 否 | 否 | ✓ 1904ms | ✓ 1977ms | ✓ 1493ms | http |
| 113.59.32.161:22222 | ✓ 1118ms | ✓ 1414ms | ✓ 1077ms | ✓ 1351ms | ✓ 1125ms | http |
| 222.184.48.235:22222 | ✓ 1407ms | ✓ 1499ms | ✓ 1840ms | ✓ 1562ms | ✓ 759ms | http |

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
