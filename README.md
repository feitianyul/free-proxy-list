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

最后更新：2026-03-17 03:20:36 UTC（2026-03-17 11:20:36 UTC+8）

**代理总数：91**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1289ms | 否 | ✓ 655ms | ✓ 1172ms | ✓ 969ms | http |
| 147.161.210.140:8800 | ✓ 1290ms | ✓ 1379ms | ✓ 740ms | ✓ 1172ms | ✓ 1826ms | http |
| 168.235.110.63:3128 | ✓ 711ms | ✓ 1697ms | 否 | ✓ 1419ms | 否 | http |
| 35.225.22.61:80 | ✓ 757ms | ✓ 1849ms | ✓ 1017ms | ✓ 1128ms | ✓ 934ms | http |
| 12.49.24.22:8080 | ✓ 1391ms | ✓ 1273ms | ✓ 417ms | 否 | 否 | http |
| 37.187.109.70:10111 | ✓ 1140ms | ✓ 1528ms | ✓ 801ms | ✓ 1680ms | ✓ 1280ms | http |
| 137.220.150.170:6005 | ✓ 826ms | 否 | ✓ 1694ms | ✓ 1139ms | ✓ 911ms | http |
| 8.219.97.248:80 | ✓ 1181ms | 否 | ✓ 1832ms | ✓ 1592ms | 否 | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1562ms | ✓ 1709ms | ✓ 954ms | http |
| 137.220.151.110:6005 | ✓ 855ms | 否 | ✓ 930ms | ✓ 1168ms | ✓ 902ms | http |
| 45.167.124.52:8080 | ✓ 985ms | 否 | ✓ 1316ms | ✓ 1733ms | ✓ 1410ms | http |
| 47.77.193.180:1080 | 否 | 否 | ✓ 520ms | ✓ 758ms | ✓ 556ms | http |
| 1.231.81.166:3128 | ✓ 1629ms | 否 | ✓ 1459ms | ✓ 1122ms | ✓ 1167ms | http |
| 38.145.208.169:8444 | ✓ 906ms | ✓ 735ms | ✓ 754ms | 否 | ✓ 1025ms | http |
| 24.144.86.173:1080 | ✓ 1335ms | 否 | ✓ 1568ms | ✓ 794ms | ✓ 561ms | http |
| 101.43.127.100:8877 | ✓ 1827ms | ✓ 1349ms | ✓ 955ms | ✓ 1082ms | ✓ 1627ms | http |
| 8.222.175.80:6128 | ✓ 1367ms | ✓ 1677ms | ✓ 736ms | ✓ 1063ms | ✓ 834ms | http |
| 219.117.204.211:7799 | ✓ 1417ms | 否 | ✓ 1464ms | ✓ 1082ms | ✓ 1747ms | http |
| 149.50.116.240:1080 | ✓ 1138ms | 否 | ✓ 1935ms | ✓ 1865ms | ✓ 1148ms | http |
| 38.145.203.135:8453 | ✓ 809ms | ✓ 1254ms | ✓ 329ms | ✓ 1180ms | 否 | http |
| 38.145.203.32:8453 | ✓ 755ms | ✓ 773ms | ✓ 1646ms | ✓ 1519ms | ✓ 1110ms | http |
| 38.145.220.173:8444 | ✓ 364ms | ✓ 991ms | ✓ 1476ms | ✓ 1484ms | ✓ 603ms | http |
| 72.56.79.129:1080 | ✓ 655ms | 否 | ✓ 1189ms | 否 | ✓ 1036ms | http |
| 178.236.245.17:3128 | ✓ 695ms | 否 | ✓ 1042ms | ✓ 1861ms | ✓ 1792ms | http |
| 38.145.218.217:8444 | ✓ 1430ms | ✓ 1741ms | ✓ 1189ms | ✓ 1521ms | ✓ 1047ms | http |
| 38.34.179.47:8448 | ✓ 989ms | ✓ 1695ms | 否 | ✓ 728ms | ✓ 1322ms | http |
| 116.80.96.107:3172 | ✓ 1590ms | 否 | ✓ 1535ms | ✓ 1821ms | 否 | http |
| 178.236.245.59:3128 | ✓ 1286ms | 否 | ✓ 668ms | ✓ 1750ms | ✓ 1356ms | http |
| 83.219.250.8:62920 | ✓ 1266ms | 否 | ✓ 1499ms | 否 | ✓ 1530ms | http |
| 194.5.212.40:8080 | ✓ 610ms | 否 | ✓ 1364ms | 否 | ✓ 1434ms | http |
| 85.198.96.242:3128 | ✓ 776ms | 否 | ✓ 1413ms | 否 | ✓ 1433ms | http |
| 1.225.116.115:1080 | 否 | ✓ 1199ms | ✓ 1215ms | ✓ 1633ms | ✓ 1571ms | http |
| 212.192.12.90:6005 | ✓ 1187ms | 否 | ✓ 1574ms | 否 | ✓ 1203ms | http |
| 137.220.150.152:6005 | ✓ 1544ms | 否 | ✓ 1186ms | ✓ 1281ms | ✓ 1202ms | http |
| 165.227.5.10:8888 | ✓ 1711ms | 否 | 否 | ✓ 775ms | ✓ 742ms | http |
| 115.231.181.40:8128 | ✓ 941ms | 否 | ✓ 972ms | ✓ 1124ms | ✓ 963ms | http |
| 137.220.150.104:6005 | ✓ 938ms | 否 | ✓ 899ms | ✓ 1336ms | ✓ 861ms | http |
| 20.78.26.206:8561 | ✓ 1305ms | ✓ 1582ms | ✓ 535ms | ✓ 800ms | ✓ 646ms | http |
| 20.210.39.153:8561 | ✓ 1311ms | 否 | ✓ 493ms | ✓ 835ms | ✓ 791ms | http |
| 20.27.11.248:8561 | ✓ 1309ms | ✓ 1211ms | ✓ 1689ms | ✓ 803ms | ✓ 632ms | http |
| 20.27.15.111:8561 | ✓ 1309ms | ✓ 1674ms | ✓ 1226ms | ✓ 804ms | ✓ 640ms | http |
| 45.88.0.113:3128 | ✓ 1174ms | 否 | ✓ 1558ms | 否 | ✓ 1722ms | http |
| 45.88.0.117:3128 | ✓ 1168ms | 否 | ✓ 1565ms | 否 | ✓ 1384ms | http |
| 45.88.0.115:3128 | ✓ 1167ms | 否 | ✓ 1728ms | ✓ 1662ms | ✓ 1180ms | http |
| 45.149.92.147:5001 | ✓ 667ms | 否 | ✓ 683ms | ✓ 845ms | ✓ 658ms | http |
| 20.27.13.35:8561 | ✓ 666ms | ✓ 885ms | ✓ 520ms | ✓ 819ms | ✓ 1202ms | http |
| 59.46.216.131:30001 | ✓ 1990ms | 否 | ✓ 1056ms | ✓ 1452ms | ✓ 1016ms | http |
| 120.92.212.16:7890 | ✓ 1003ms | ✓ 1208ms | 否 | ✓ 1217ms | 否 | http |
| 120.55.163.237:10086 | ✓ 941ms | ✓ 1076ms | ✓ 922ms | ✓ 1109ms | ✓ 931ms | http |
| 162.240.154.26:3128 | ✓ 1452ms | ✓ 1006ms | ✓ 1124ms | ✓ 1062ms | ✓ 897ms | http |
| 103.84.95.54:7890 | ✓ 746ms | ✓ 1721ms | ✓ 989ms | ✓ 1428ms | ✓ 694ms | http |
| 3.79.194.222:852 | ✓ 1579ms | 否 | ✓ 1434ms | 否 | ✓ 1954ms | http |
| 62.60.177.204:34094 | ✓ 895ms | ✓ 1105ms | 否 | ✓ 1243ms | ✓ 805ms | http |
| 133.242.138.34:8100 | ✓ 1406ms | ✓ 1603ms | ✓ 1691ms | ✓ 1422ms | ✓ 1148ms | http |
| 34.101.184.164:3128 | ✓ 1717ms | 否 | ✓ 1197ms | ✓ 1453ms | ✓ 1038ms | http |
| 38.55.107.137:6005 | 否 | 否 | ✓ 1271ms | ✓ 1186ms | ✓ 1812ms | http |
| 45.88.0.99:3128 | 否 | 否 | ✓ 1588ms | ✓ 1857ms | ✓ 1660ms | http |
| 45.88.0.114:3128 | 否 | 否 | ✓ 1606ms | ✓ 1829ms | ✓ 1703ms | http |
| 45.88.0.116:3128 | 否 | 否 | ✓ 1607ms | ✓ 1829ms | ✓ 1669ms | http |
| 45.88.0.98:3128 | 否 | 否 | ✓ 1564ms | ✓ 1865ms | ✓ 1667ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 933ms | ✓ 1090ms | ✓ 778ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 990ms | ✓ 1236ms | ✓ 988ms | http |
| 62.113.119.14:8080 | ✓ 1872ms | 否 | ✓ 1332ms | 否 | ✓ 1201ms | http |
| 45.136.130.156:8450 | ✓ 970ms | ✓ 1468ms | ✓ 535ms | ✓ 773ms | 否 | http |
| 160.238.65.5:3128 | ✓ 1717ms | ✓ 1684ms | 否 | ✓ 1474ms | ✓ 1147ms | http |
| 160.238.65.6:3128 | ✓ 1216ms | ✓ 1857ms | 否 | ✓ 1998ms | ✓ 1631ms | http |
| 160.238.65.8:3128 | 否 | 否 | ✓ 1955ms | ✓ 1668ms | ✓ 1810ms | http |
| 106.117.208.101:7890 | ✓ 1061ms | ✓ 1614ms | ✓ 1012ms | 否 | ✓ 1144ms | http |
| 8.137.112.117:3128 | 否 | ✓ 1334ms | ✓ 1022ms | ✓ 1400ms | ✓ 1034ms | http |
| 160.238.65.4:3128 | ✓ 1011ms | 否 | ✓ 703ms | 否 | ✓ 1132ms | http |
| 86.53.183.16:1080 | ✓ 1239ms | 否 | ✓ 1471ms | 否 | ✓ 1513ms | http |
| 47.101.149.27:9010 | ✓ 1329ms | 否 | ✓ 1329ms | 否 | ✓ 1342ms | http |
| 213.220.62.62:3128 | ✓ 583ms | 否 | ✓ 573ms | ✓ 1441ms | ✓ 1148ms | http |
| 85.208.108.43:2094 | ✓ 479ms | 否 | ✓ 252ms | 否 | ✓ 1383ms | http |
| 160.238.65.7:3128 | ✓ 1206ms | 否 | ✓ 1251ms | ✓ 1432ms | ✓ 1411ms | http |
| 160.238.65.2:3128 | ✓ 1026ms | 否 | ✓ 1474ms | ✓ 1496ms | ✓ 1542ms | http |
| 159.223.42.219:3128 | ✓ 1753ms | 否 | 否 | ✓ 1657ms | ✓ 1146ms | http |
| 172.212.68.37:3128 | ✓ 384ms | 否 | ✓ 939ms | ✓ 1065ms | ✓ 1274ms | http |
| 20.120.225.109:3128 | 否 | 否 | ✓ 1826ms | ✓ 1071ms | ✓ 884ms | http |
| 45.129.141.143:3128 | ✓ 1173ms | 否 | ✓ 1634ms | 否 | ✓ 1799ms | http |
| 103.39.51.190:8080 | ✓ 1890ms | 否 | ✓ 1767ms | 否 | ✓ 1801ms | http |
| 45.88.0.111:3128 | ✓ 1650ms | 否 | ✓ 564ms | 否 | ✓ 1408ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1173ms | 否 | ✓ 1358ms | ✓ 1078ms | http |
| 193.23.200.251:10808 | 否 | 否 | ✓ 1226ms | ✓ 1897ms | ✓ 1533ms | http |
| 121.230.9.160:1080 | ✓ 996ms | ✓ 1697ms | ✓ 1572ms | 否 | 否 | http |
| 38.145.220.36:8444 | 否 | ✓ 1010ms | ✓ 789ms | 否 | ✓ 827ms | http |
| 45.136.198.40:3128 | ✓ 1126ms | 否 | ✓ 1829ms | 否 | ✓ 1654ms | http |
| 61.52.131.172:8443 | ✓ 1391ms | ✓ 1666ms | 否 | ✓ 1192ms | ✓ 942ms | http |
| 220.197.44.36:3128 | ✓ 1956ms | ✓ 1924ms | 否 | ✓ 1965ms | 否 | http |
| 38.145.203.19:8447 | ✓ 1413ms | ✓ 996ms | ✓ 1352ms | ✓ 1989ms | ✓ 613ms | http |
| 212.192.13.76:6005 | 否 | 否 | ✓ 1178ms | ✓ 1272ms | ✓ 1200ms | http |

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
